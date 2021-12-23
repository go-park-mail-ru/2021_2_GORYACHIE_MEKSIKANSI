//go:generate mockgen -destination=mocks/orm.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart/orm WrapperCartInterface,ConnectionInterface,TransactionInterface,ConnectPromoCodeServiceInterface
package orm

import (
	cartPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart/myerror"
	promoProtoPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/promocode/proto"
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
)

type WrapperCartInterface interface {
	GetCart(id int) (*cartPkg.ResponseCartErrors, []cartPkg.CastDishesErrs, error)
	UpdateCart(dishes cartPkg.RequestCartDefault, clientId int) (*cartPkg.ResponseCartErrors, []cartPkg.CastDishesErrs, error)
	DeleteCart(id int) error
	GetPriceDelivery(id int) (int, error)
	GetRestaurant(id int) (*cartPkg.RestaurantId, error)
	AddPromoCode(promoCode string, restaurantId int, clientId int) error
	DoPromoCode(promoCode string, restaurantId int, cart *cartPkg.ResponseCartErrors) (*cartPkg.ResponseCartErrors, error)
	GetPromoCode(id int) (string, error)
}

type ConnectionInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

type TransactionInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	LargeObjects() pgx.LargeObjects
	Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error)
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
	Conn() *pgx.Conn
}

type ConnectPromoCodeServiceInterface interface {
	GetTypePromoCode(ctx context.Context, in *promoProtoPkg.PromoCodeWithRestaurantId, opts ...grpc.CallOption) (*promoProtoPkg.TypePromoCodeResponse, error)
	ActiveFreeDelivery(ctx context.Context, in *promoProtoPkg.PromoCodeWithRestaurantId, opts ...grpc.CallOption) (*promoProtoPkg.FreeDeliveryResponse, error)
	ActiveCostForFreeDish(ctx context.Context, in *promoProtoPkg.PromoCodeWithRestaurantId, opts ...grpc.CallOption) (*promoProtoPkg.FreeDishResponse, error)
	ActiveCostForSale(ctx context.Context, in *promoProtoPkg.PromoCodeWithAmount, opts ...grpc.CallOption) (*promoProtoPkg.NewCostResponse, error)
	ActiveTimeForSale(ctx context.Context, in *promoProtoPkg.PromoCodeWithAmount, opts ...grpc.CallOption) (*promoProtoPkg.NewCostResponse, error)
	AddPromoCode(ctx context.Context, in *promoProtoPkg.PromoCodeWithRestaurantIdAndClient, opts ...grpc.CallOption) (*promoProtoPkg.Error, error)
	GetPromoCode(ctx context.Context, in *promoProtoPkg.ClientId, opts ...grpc.CallOption) (*promoProtoPkg.PromoCodeText, error)
}

type Wrapper struct {
	Conn             ConnectionInterface
	ConnPromoService ConnectPromoCodeServiceInterface
	Ctx              context.Context
}

func (db *Wrapper) GetCart(id int) (*cartPkg.ResponseCartErrors, []cartPkg.CastDishesErrs, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, nil, &errPkg.Errors{
			Text: errPkg.CGetCartTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var result cartPkg.ResponseCartErrors
	row, err := tx.Query(contextTransaction,
		"SELECT cart_food.id, cart_food.food, cart_food.number_item, d.avatar, d.name, cart_food.count_food, d.cost, d.kilocalorie, d.weight,"+
			" d.description, sr.name, sr.id, sr.radios, sd.name, sd.id, sd.cost, d.restaurant, d.count, sr.kilocalorie, sd.kilocalorie,"+
			" cart_food.place, crf.place, csf.place "+
			"FROM public.cart_food "+
			"LEFT JOIN public.dishes d ON d.id = cart_food.food "+
			"LEFT JOIN public.cart_structure_food csf ON csf.client_id = cart_food.client_id and d.id=csf.food and cart_food.id=csf.cart_id "+
			"LEFT JOIN public.structure_dishes sd ON sd.id = csf.checkbox and sd.food=cart_food.food "+
			"LEFT JOIN public.cart_radios_food crf ON crf.client_id = cart_food.client_id and cart_food.id=crf.cart_id "+
			"LEFT JOIN public.structure_radios sr ON sr.id = crf.radios "+
			"WHERE public.cart_food.client_id = $1",
		id)
	if err != nil {
		return nil, nil, &errPkg.Errors{
			Text: errPkg.CGetCartNotSelect,
		}
	}

	placeIngredients := make(map[int]map[int]cartPkg.IngredientCartResponse)
	placeRadios := make(map[int]map[int]cartPkg.RadiosCartResponse)
	infoDishes := make(map[int]cartPkg.DishesCartResponse)
	var restaurant cartPkg.RestaurantIdCastResponse

	for row.Next() {
		var dish cartPkg.DishesCartResponse
		var count, cartId int

		var getPlaceDishes, getPlaceRadios, getPlaceIngredient *int32
		var ingredientKilocalorie, radiosKilocalorie *int32
		var radiosId, radiosRadiosId, ingredientId, ingredientCost *int32
		var radiosName, ingredientName *string
		err := row.Scan(&cartId, &dish.Id, &dish.ItemNumber, &dish.Img, &dish.Name, &dish.Count, &dish.Cost, &dish.Kilocalorie,
			&dish.Weight, &dish.Description, &radiosName, &radiosId, &radiosRadiosId, &ingredientName,
			&ingredientId, &ingredientCost, &restaurant.Id, &count, &radiosKilocalorie, &ingredientKilocalorie,
			&getPlaceDishes, &getPlaceRadios, &getPlaceIngredient)

		if err != nil {
			return nil, nil, &errPkg.Errors{
				Text: errPkg.CGetCartNotScan,
			}
		}

		placeDishes := ConvertInt32ToInt(getPlaceDishes)
		placeRadio := ConvertInt32ToInt(getPlaceRadios)
		placeIngredient := ConvertInt32ToInt(getPlaceIngredient)

		var radios cartPkg.RadiosCartResponse
		if radiosName != nil {
			radios.Name = *radiosName
			radios.Id = int(*radiosId)
			radios.RadiosId = int(*radiosRadiosId)
			dish.Kilocalorie += int(*radiosKilocalorie)
		}

		var ingredient cartPkg.IngredientCartResponse
		if ingredientName != nil {
			ingredient.Name = *ingredientName
			ingredient.Id = int(*ingredientId)
			ingredient.Cost = int(*ingredientCost)
			dish.Kilocalorie += int(*ingredientKilocalorie)
		}

		if dish.Count > count && count != -1 {
			var dishesErrors []cartPkg.CastDishesErrs
			var dishesError cartPkg.CastDishesErrs
			dishesError.ItemNumber = dish.ItemNumber
			dishesError.NameDish = dish.Name
			dishesError.CountAvail = count
			dishesErrors = append(dishesErrors, dishesError)
		}
		dish.Weight = dish.Weight * dish.Count
		dish.Kilocalorie = dish.Kilocalorie * dish.Count

		if placeIngredient != -1 {
			temp := placeIngredients[placeDishes]
			if temp == nil {
				temp = make(map[int]cartPkg.IngredientCartResponse)
			}
			temp[placeIngredient] = ingredient
			placeIngredients[placeDishes] = temp
		}

		if placeRadio != -1 {
			temp := placeRadios[placeDishes]
			if temp == nil {
				temp = make(map[int]cartPkg.RadiosCartResponse)
			}
			temp[placeRadio] = radios
			placeRadios[placeDishes] = temp
		}

		infoDishes[placeDishes] = dish
	}

	for i := 0; i < len(infoDishes); i++ {
		dish := infoDishes[i]
		for j := 0; j < len(placeIngredients[i]); j++ {
			dish.IngredientCart = append(dish.IngredientCart, placeIngredients[i][j])
		}

		for j := 0; j < len(placeRadios[i]); j++ {
			dish.RadiosCart = append(dish.RadiosCart, placeRadios[i][j])
		}

		result.Dishes = append(result.Dishes, dish)
	}

	if len(infoDishes) == 0 {
		return nil, nil, &errPkg.Errors{
			Text: errPkg.CGetCartCartNotFound,
		}
	}

	result.Restaurant = restaurant

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, nil, &errPkg.Errors{
			Text: errPkg.CGetCartNotCommit,
		}
	}

	return &result, nil, nil
}

func (db *Wrapper) DeleteCart(id int) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.CDeleteCartTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	_, err = tx.Exec(contextTransaction,
		"DELETE FROM public.cart_food CASCADE WHERE client_id = $1", id)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.CDeleteCartCartNotDelete,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.CDeleteCartNotCommit,
		}
	}
	return nil
}

func (db *Wrapper) updateCartStructFood(ingredients []cartPkg.IngredientsCartRequest, clientId int, cartId int, tx TransactionInterface, contextTransaction context.Context) ([]cartPkg.IngredientCartResponse, error) {
	var result []cartPkg.IngredientCartResponse
	place := 0
	for _, ingredient := range ingredients {
		var checkedIngredient cartPkg.IngredientCartResponse
		var dishId int
		err := tx.QueryRow(contextTransaction,
			"SELECT id, name, cost, food FROM public.structure_dishes WHERE id = $1", ingredient.Id).Scan(
			&checkedIngredient.Id, &checkedIngredient.Name, &checkedIngredient.Cost, &dishId)
		if err != nil {
			return nil, &errPkg.Errors{
				Text: errPkg.CUpdateCartStructureFoodStructureFoodNotSelect,
			}
		}
		result = append(result, checkedIngredient)

		_, err = tx.Exec(contextTransaction,
			"INSERT INTO public.cart_structure_food (checkbox, client_id, food, cart_id, place) VALUES ($1, $2, $3, $4, $5)",
			ingredient.Id, clientId, dishId, cartId, place)
		if err != nil {
			return nil, &errPkg.Errors{
				Text: errPkg.CUpdateCartStructFoodStructureFoodNotInsert,
			}
		}
		place++
	}
	return result, nil
}

func (db *Wrapper) updateCartRadios(radios []cartPkg.RadiosCartRequest, clientId int, cartId int, tx TransactionInterface, contextTransaction context.Context) ([]cartPkg.RadiosCartResponse, error) {
	var result []cartPkg.RadiosCartResponse
	radiosPlace := 0
	for _, radio := range radios {
		var checkedRadios cartPkg.RadiosCartResponse
		err := tx.QueryRow(contextTransaction,
			"SELECT id, name FROM public.structure_radios WHERE id = $1", radio.Id).Scan(
			&checkedRadios.Id, &checkedRadios.Name)
		if err != nil {
			return nil, &errPkg.Errors{
				Text: errPkg.CUpdateCartStructRadiosStructRadiosNotSelect,
			}
		}
		result = append(result, checkedRadios)

		_, err = tx.Exec(contextTransaction,
			"INSERT INTO public.cart_radios_food (radios_id, radios, client_id, cart_id, place) VALUES ($1, $2, $3, $4, $5)",
			radio.RadiosId, radio.Id, clientId, cartId, radiosPlace)
		if err != nil {
			return nil, &errPkg.Errors{
				Text: errPkg.CUpdateCartRadiosRadiosNotInsert,
			}
		}
		radiosPlace++
	}
	return result, nil
}

func (db *Wrapper) UpdateCart(newCart cartPkg.RequestCartDefault, clientId int) (*cartPkg.ResponseCartErrors, []cartPkg.CastDishesErrs, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, nil, &errPkg.Errors{
			Text: errPkg.CUpdateCartTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var dishesErrors []cartPkg.CastDishesErrs
	var cart cartPkg.ResponseCartErrors

	for i, dish := range newCart.Dishes {
		var dishes cartPkg.DishesCartResponse
		count := 0
		err := tx.QueryRow(contextTransaction,
			"SELECT id, avatar, cost, name, description, count, weight, kilocalorie FROM public.dishes WHERE id = $1 AND restaurant = $2",
			dish.Id, newCart.Restaurant.Id).Scan(
			&dishes.Id, &dishes.Img, &dishes.Cost, &dishes.Name, &dishes.Description, &count, &dishes.Weight, &dishes.Kilocalorie)
		if err != nil {
			if err == pgx.ErrNoRows {
				return nil, nil, &errPkg.Errors{
					Text: errPkg.CUpdateCartCartNotFound,
				}
			}
			return nil, nil, &errPkg.Errors{
				Text: errPkg.CUpdateCartCartNotScan,
			}
		}

		dishes.Count = dish.Count

		if dish.Count > count && count != UnlimitedCount {
			var dishesError cartPkg.CastDishesErrs
			dishesError.ItemNumber = dish.ItemNumber
			dishesError.NameDish = dishes.Name
			dishesError.CountAvail = count
			dishesErrors = append(dishesErrors, dishesError)

			dishes.Count = count
			dish.Count = count
		}
		dishes.Weight = dishes.Weight * dishes.Count
		dishes.Kilocalorie = dishes.Kilocalorie * dishes.Count

		cart.Dishes = append(cart.Dishes, dishes)

		var idCart int
		err = tx.QueryRow(contextTransaction,
			"INSERT INTO public.cart_food (client_id, food, count_food, restaurant_id, number_item, place) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
			clientId, dish.Id, dish.Count, newCart.Restaurant.Id, newCart.Dishes[i].ItemNumber, i).Scan(&idCart)
		if err != nil {
			return nil, nil, &errPkg.Errors{
				Text: errPkg.CUpdateCartCartNotInsert,
			}
		}
		cart.Dishes[i].RadiosCart, err = db.updateCartRadios(dish.Radios, clientId, idCart, tx, contextTransaction)
		if err != nil {
			return nil, nil, err
		}

		cart.Dishes[i].IngredientCart, err = db.updateCartStructFood(dish.Ingredients, clientId, idCart, tx, contextTransaction)
		if err != nil {
			return nil, nil, err
		}
	}
	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, nil, &errPkg.Errors{
			Text: errPkg.CUpdateCartNotCommit,
		}
	}
	return &cart, dishesErrors, nil
}

func (db *Wrapper) GetPriceDelivery(id int) (int, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Text: errPkg.CGetPriceDeliveryTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var price int
	err = tx.QueryRow(contextTransaction,
		"SELECT price_delivery FROM public.restaurant WHERE id = $1", id).Scan(&price)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, &errPkg.Errors{
				Text: errPkg.CGetPriceDeliveryPriceNotFound,
			}
		}
		return 0, &errPkg.Errors{
			Text: errPkg.CGetPriceDeliveryPriceNotScan,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return 0, &errPkg.Errors{
			Text: errPkg.CGetPriceDeliveryNotCommit,
		}
	}

	return price, nil
}

func (db *Wrapper) GetRestaurant(id int) (*cartPkg.RestaurantId, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetRestaurantTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var restaurant cartPkg.RestaurantId
	err = tx.QueryRow(contextTransaction,
		"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM public.restaurant WHERE id = $1",
		id).Scan(&restaurant.Id, &restaurant.Img, &restaurant.Name, &restaurant.CostForFreeDelivery, &restaurant.MinDelivery,
		&restaurant.MaxDelivery, &restaurant.Rating)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetRestaurantRestaurantNotFound,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetRestaurantNotCommit,
		}
	}

	return &restaurant, nil
}

func (db *Wrapper) AddPromoCode(promoCode string, restaurantId int, clientId int) error {
	errorMicroservice, err := db.ConnPromoService.AddPromoCode(db.Ctx, &promoProtoPkg.PromoCodeWithRestaurantIdAndClient{
		PromoCode:  promoCode,
		Restaurant: int64(restaurantId),
		Client:     int64(clientId),
	})
	if err != nil {
		return err
	}
	if errorMicroservice.Error != "" {
		return &errPkg.Errors{
			Text: errorMicroservice.Error,
		}
	}
	return nil
}

func (db *Wrapper) GetPromoCode(id int) (string, error) {
	promoCodeText, err := db.ConnPromoService.GetPromoCode(db.Ctx, &promoProtoPkg.ClientId{ClientId: int64(id)})
	if err != nil {
		return "", err
	}
	if promoCodeText.Error != "" {
		return "", &errPkg.Errors{
			Text: promoCodeText.Error,
		}
	}
	return promoCodeText.PromoCodeText, nil
}

func (db *Wrapper) DoPromoCode(promoCode string, restaurantId int, cart *cartPkg.ResponseCartErrors) (*cartPkg.ResponseCartErrors, error) {
	typePromoCode, err := db.ConnPromoService.GetTypePromoCode(db.Ctx, &promoProtoPkg.PromoCodeWithRestaurantId{
		PromoCode:  promoCode,
		Restaurant: int64(restaurantId),
	},
	)

	if err != nil {
		return nil, err
	}
	if typePromoCode.Error != "" {
		return nil, &errPkg.Errors{
			Text: typePromoCode.Error,
		}
	}

	transactionPromoCode := context.Background()
	tx, err := db.Conn.Begin(transactionPromoCode)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.CDoPromoCodeNotSelectInfo,
		}
	}

	defer tx.Rollback(transactionPromoCode)

	switch typePromoCode.Type {
	case PromoCodeFreeDelivery:
		{
			freeDelivery, err := db.ConnPromoService.ActiveFreeDelivery(db.Ctx, &promoProtoPkg.PromoCodeWithRestaurantId{
				PromoCode:  promoCode,
				Restaurant: int64(restaurantId),
			},
			)

			if err != nil {
				return nil, err
			}
			if freeDelivery.Error != "" {
				return nil, &errPkg.Errors{
					Text: freeDelivery.Error,
				}
			}

			if freeDelivery.Have {
				cart.Cost.SumCost -= cart.Cost.DCost
				cart.Cost.DCost = 0
			}
		}
	case PromoCodeSaleOverCost:
		{
			newCost, err := db.ConnPromoService.ActiveCostForSale(db.Ctx, &promoProtoPkg.PromoCodeWithAmount{
				PromoCode:  promoCode,
				Amount:     int64(cart.Cost.SumCost),
				Restaurant: int64(restaurantId),
			},
			)

			if err != nil {
				return nil, err
			}
			if newCost.Error != "" {
				return nil, &errPkg.Errors{
					Text: newCost.Error,
				}
			}

			cart.Cost.SumCost = int(newCost.Cost)
		}
	case PromoCodeSaleOverTime:
		{
			newCost, err := db.ConnPromoService.ActiveTimeForSale(db.Ctx, &promoProtoPkg.PromoCodeWithAmount{
				PromoCode:  promoCode,
				Amount:     int64(cart.Cost.SumCost),
				Restaurant: int64(restaurantId),
			},
			)

			if err != nil {
				return nil, err
			}
			if newCost.Error != "" {
				return nil, &errPkg.Errors{
					Text: newCost.Error,
				}
			}
			cart.Cost.SumCost = int(newCost.Cost)
		}
	case PromoCodeFreeDishes:
		{
			freeDishId, err := db.ConnPromoService.ActiveCostForFreeDish(db.Ctx, &promoProtoPkg.PromoCodeWithRestaurantId{
				PromoCode:  promoCode,
				Restaurant: int64(restaurantId),
			},
			)

			if err != nil {
				return nil, err
			}
			if freeDishId.Error != "" {
				return nil, &errPkg.Errors{
					Text: freeDishId.Error,
				}
			}

			if int(freeDishId.Cost) < cart.Cost.SumCost {
				var newDish cartPkg.DishesCartResponse
				err = tx.QueryRow(transactionPromoCode,
					"SELECT avatar, name, kilocalorie, weight, description FROM public.dishes WHERE id = $1 AND count > 1",
					int(freeDishId.DishId)).Scan(&newDish.Img, &newDish.Name, &newDish.Kilocalorie,
					&newDish.Weight, &newDish.Description)
				if err != nil {
					return nil, &errPkg.Errors{
						Text: errPkg.CDoPromoCodeNotSelectInfoDish,
					}
				}

				newDish.Count = 1
				newDish.Cost = 0
				newDish.ItemNumber = 0
				newDish.Id = int(freeDishId.DishId)
				newDish.RadiosCart = []cartPkg.RadiosCartResponse{}
				newDish.IngredientCart = []cartPkg.IngredientCartResponse{}

				cart.Dishes = append(cart.Dishes, newDish)
			}
		}
	}

	err = tx.QueryRow(transactionPromoCode,
		"SELECT name, description FROM public.promocode WHERE code = $1 AND restaurant = $2",
		promoCode, restaurantId).Scan(&cart.PromoCode.Name, &cart.PromoCode.Description)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.CDoPromoCodeNotSelectInfo,
		}
	}
	err = tx.Commit(transactionPromoCode)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.CDoPromoCodeNotSelectInfo,
		}
	}
	cart.PromoCode.Code = promoCode
	return cart, nil
}
