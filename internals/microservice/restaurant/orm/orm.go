//go:generate mockgen -destination=mocks/orm.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant/orm WrapperRestaurantInterface,ConnectionInterface,TransactionInterface
package orm

import (
	resPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant/myerror"
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"time"
)

type WrapperRestaurantInterface interface {
	GetRestaurants() (*resPkg.AllRestaurantsPromo, error)
	GetRecommendedRestaurants() (*resPkg.AllRestaurants, error)
	GetDishes(restId int, dishesId int) (*resPkg.Dishes, error)
	GetRestaurant(id int) (*resPkg.RestaurantId, error)
	GetMenu(id int) ([]resPkg.Menu, error)
	GetTagsRestaurant(id int) ([]resPkg.Tag, error)
	GetReview(id int) ([]resPkg.Review, error)
	CreateReview(id int, review resPkg.NewReview) error
	SearchCategory(name string) ([]int, error)
	SearchRestaurant(name string) ([]int, error)
	GetGeneralInfoRestaurant(id int) (*resPkg.Restaurants, error)
	GetFavoriteRestaurants(id int) ([]resPkg.Restaurants, error)
	EditRestaurantInFavorite(idRestaurant int, idClient int) (bool, error)
	IsFavoriteRestaurant(idClient int, idRestaurant int) (bool, error)
	GetPromoCodes() ([]resPkg.Promocode, error)
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

type Wrapper struct {
	Conn ConnectionInterface
}

func (db *Wrapper) GetRestaurants() (*resPkg.AllRestaurantsPromo, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetRestaurantsTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	row, err := tx.Query(contextTransaction,
		"SELECT t.id, t.avatar, t.name, t.price_delivery, t.min_delivery_time, t.max_delivery_time,"+
			" t.rating, rc.category, rc.id "+
			"FROM (SELECT r.id, r.avatar, r.name, r.price_delivery, r.min_delivery_time, r.max_delivery_time,"+
			" r.rating FROM restaurant r ORDER BY random() LIMIT 51) t "+
			"LEFT JOIN restaurant_category rc ON rc.restaurant = t.id")
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetRestaurantsRestaurantsNotSelect,
		}
	}

	var result resPkg.AllRestaurantsPromo
	var restaurants []resPkg.Restaurants
	var tags []resPkg.Tag
	infoRestaurant := make(map[int]resPkg.Restaurants)
	namesTags := make(map[string]resPkg.Tag)
	for row.Next() {
		var restaurant resPkg.Restaurants
		var category *string
		var categoryId *int32
		err := row.Scan(&restaurant.Id, &restaurant.Img, &restaurant.Name, &restaurant.CostForFreeDelivery,
			&restaurant.MinDelivery, &restaurant.MaxDelivery, &restaurant.Rating, &category, &categoryId)
		if err != nil {
			return nil, &errPkg.Errors{
				Text: errPkg.RGetRestaurantsRestaurantsNotScan,
			}
		}

		if _, ok := namesTags[*category]; !ok {
			namesTags[*category] = resPkg.Tag{Name: *category, Id: int(*categoryId)}
			tags = append(tags, namesTags[*category])
		}

		if _, ok := infoRestaurant[restaurant.Id]; !ok {
			infoRestaurant[restaurant.Id] = restaurant
			restaurants = append(restaurants, restaurant)
		}
	}

	if restaurants == nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetRestaurantsRestaurantsNotFound,
		}
	}

	result.Restaurant = restaurants
	result.AllTags = tags

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetRestaurantsNotCommit,
		}
	}

	return &resPkg.AllRestaurantsPromo{Restaurant: restaurants, AllTags: tags}, nil
}

func (db *Wrapper) GetRecommendedRestaurants() (*resPkg.AllRestaurants, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetRecommendedRestaurantsTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	row, err := tx.Query(contextTransaction,
		"SELECT t.id, t.avatar, t.name, t.price_delivery, t.min_delivery_time, "+
			"t.max_delivery_time, t.rating, rc.category, rc.id FROM "+
			"(SELECT r.id, r.avatar, r.name, r.price_delivery, r.min_delivery_time, "+
			"r.max_delivery_time, r.rating FROM restaurant r ORDER BY rating DESC LIMIT 6) t "+
			"LEFT JOIN restaurant_category rc ON rc.restaurant = t.id")
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetRecommendedRestaurantsRestaurantsNotSelect,
		}
	}

	var result resPkg.AllRestaurants
	var restaurants []resPkg.Restaurants
	var tags []resPkg.Tag
	infoRestaurant := make(map[int]resPkg.Restaurants)
	namesTags := make(map[string]resPkg.Tag)
	for row.Next() {
		var restaurant resPkg.Restaurants
		var category *string
		var categoryId *int32
		err := row.Scan(&restaurant.Id, &restaurant.Img, &restaurant.Name, &restaurant.CostForFreeDelivery,
			&restaurant.MinDelivery, &restaurant.MaxDelivery, &restaurant.Rating, &category, &categoryId)
		if err != nil {
			return nil, &errPkg.Errors{
				Text: errPkg.RGetRecommendedRestaurantsRestaurantsNotScan,
			}
		}

		if _, ok := namesTags[*category]; !ok {
			namesTags[*category] = resPkg.Tag{Name: *category, Id: int(*categoryId)}
			tags = append(tags, namesTags[*category])
		}

		if _, ok := infoRestaurant[restaurant.Id]; !ok {
			infoRestaurant[restaurant.Id] = restaurant
			restaurants = append(restaurants, restaurant)
		}
	}

	if restaurants == nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetRecommendedRestaurantsRestaurantsNotFound,
		}
	}

	result.Restaurant = restaurants
	result.AllTags = tags

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetRecommendedRestaurantsNotCommit,
		}
	}

	return &resPkg.AllRestaurants{Restaurant: restaurants, AllTags: tags}, nil
}

func (db *Wrapper) GetRestaurant(id int) (*resPkg.RestaurantId, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetRestaurantTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var restaurant resPkg.RestaurantId
	err = tx.QueryRow(contextTransaction,
		"SELECT r.id, r.avatar, r.name, r.price_delivery, r.min_delivery_time, r.max_delivery_time, r.rating FROM restaurant r WHERE r.id = $1", id).Scan(
		&restaurant.Id, &restaurant.Img, &restaurant.Name, &restaurant.CostForFreeDelivery, &restaurant.MinDelivery,
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

func (db *Wrapper) GetTagsRestaurant(id int) ([]resPkg.Tag, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetTagsRestaurantTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	rowCategory, err := tx.Query(contextTransaction,
		"SELECT id, category, place FROM restaurant_category WHERE restaurant = $1", id)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetTagsRestaurantNotSelect,
		}
	}

	var tags []resPkg.Tag
	var tag resPkg.Tag

	place := make(map[int]resPkg.Tag)
	for rowCategory.Next() {
		var placeCategory int
		err := rowCategory.Scan(&tag.Id, &tag.Name, &placeCategory)
		if err != nil {
			return nil, &errPkg.Errors{
				Text: errPkg.RGetTagsRestaurantRestaurantNotScan,
			}
		}
		place[placeCategory] = tag
	}

	for i := 0; i < len(place); i++ {
		tags = append(tags, place[i])
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetTagsRestaurantNotCommit,
		}
	}

	return tags, nil
}

func (db *Wrapper) GetMenu(id int) ([]resPkg.Menu, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetMenuTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var result []resPkg.Menu

	rowDishes, err := tx.Query(contextTransaction,
		"SELECT category_restaurant, id, avatar, name, cost, kilocalorie, place, place_category FROM dishes WHERE restaurant = $1", id)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetMenuDishesCategoryNotSelect,
		}
	}

	categoryPlace := make(map[int]resPkg.Menu)
	place := make(map[int]map[int]resPkg.DishesMenu)

	for rowDishes.Next() {
		var menu resPkg.Menu
		var dish resPkg.DishesMenu
		var placeDish, placeCategory int
		err := rowDishes.Scan(&menu.Name, &dish.Id, &dish.Img, &dish.Name, &dish.Cost, &dish.Kilocalorie, &placeDish, &placeCategory)
		if err != nil {
			return nil, err
		}

		temp := place[placeCategory]
		if temp == nil {
			temp = make(map[int]resPkg.DishesMenu)
		}
		temp[placeDish] = dish
		place[placeCategory] = temp
		categoryPlace[placeCategory] = menu
	}

	for i := 0; i < len(place); i++ {
		result = append(result, categoryPlace[i])
		for j := 0; j < len(place[i]); j++ {
			result[i].DishesMenu = append(result[i].DishesMenu, place[i][j])
		}
	}

	if result == nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetMenuDishesNotFound,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetMenuNotCommit,
		}
	}

	return result, nil
}

func (db *Wrapper) GetDishes(restId int, dishesId int) (*resPkg.Dishes, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetDishesTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var dish resPkg.Dishes
	rows, err := tx.Query(contextTransaction,
		"SELECT d.id, d.avatar, d.name, d.cost, d.kilocalorie, d.description, r.id, r.name, sr.id, sr.name, r.place, "+
			"sr.place, sd.id, sd.name, sd.cost, sd.place "+
			"FROM dishes d"+
			" LEFT JOIN radios r ON d.id=r.food "+
			"LEFT JOIN structure_radios sr ON sr.radios=r.id "+
			"LEFT JOIN structure_dishes sd ON sd.food=d.id WHERE d.id = $1 AND restaurant = $2",
		dishesId, restId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, &errPkg.Errors{
				Text: errPkg.RGetDishesDishesNotFound,
			}
		}
		return nil, &errPkg.Errors{
			Text: errPkg.RGetDishesDishesNotScan,
		}
	}
	radiosInfo := make(map[int]resPkg.Radios)
	radios := make(map[int]map[int]resPkg.CheckboxesRows)
	ingredients := make(map[int]resPkg.Ingredients)

	for rows.Next() {
		var rad resPkg.Radios
		var ingredient resPkg.Ingredients
		var elementRadios resPkg.CheckboxesRows
		var placeRadios, placeElementRadios, placeIngredient *int32
		var radId, elementRadiosId, ingredientId, ingredientCost *int32
		var radTitle, elementRadiosName, ingredientTitle *string
		err := rows.Scan(&dish.Id, &dish.Img, &dish.Title, &dish.Cost, &dish.Ccal, &dish.Description,
			&radId, &radTitle, &elementRadiosId, &elementRadiosName, &placeRadios, &placeElementRadios,
			&ingredientId, &ingredientTitle, &ingredientCost, &placeIngredient)
		if err != nil {
			return nil, &errPkg.Errors{
				Text: errPkg.RGetRadiosRadiosNotScan,
			}
		}

		if radId != nil {
			rad.Id = int(*radId)
			rad.Title = *radTitle
			elementRadios.Id = int(*elementRadiosId)
			elementRadios.Name = *elementRadiosName
		}

		if ingredientId != nil {
			ingredient.Id = int(*ingredientId)
			ingredient.Title = *ingredientTitle
			ingredient.Cost = int(*ingredientCost)
		}

		if placeRadios != nil {
			temp := radios[int(*placeRadios)]
			if temp == nil {
				temp = make(map[int]resPkg.CheckboxesRows)
			}
			temp[int(*placeElementRadios)] = elementRadios
			radios[int(*placeRadios)] = temp
			radiosInfo[int(*placeRadios)] = rad
		}

		if placeIngredient != nil {
			ingredients[int(*placeIngredient)] = ingredient
		}

	}

	for i := 0; i < len(ingredients); i++ {
		dish.Ingredient = append(dish.Ingredient, ingredients[i])
	}

	for i := 0; i < len(radiosInfo); i++ {
		for j := 0; j < len(radios[i]); j++ {
			temp := radiosInfo[i]
			temp.Rows = append(temp.Rows, radios[i][j])
			radiosInfo[i] = temp
		}
		dish.Radios = append(dish.Radios, radiosInfo[i])
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetDishesNotCommit,
		}
	}

	return &dish, nil
}

func (db *Wrapper) GetReview(id int) ([]resPkg.Review, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetReviewTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	rowDishes, err := tx.Query(contextTransaction,
		"SELECT gn.name, r.text, r.date_create, r.rate FROM review r "+
			"LEFT JOIN general_user_info gn ON r.author = gn.id "+
			"WHERE r.restaurant = $1 ORDER BY r.date_create", id)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetReviewNotSelect,
		}
	}

	var result []resPkg.Review
	for rowDishes.Next() {
		var review resPkg.Review
		var date time.Time
		err := rowDishes.Scan(&review.Name, &review.Text, &date, &review.Rate)
		if err != nil {
			return nil, &errPkg.Errors{
				Text: errPkg.RGetReviewNotScan,
			}
		}

		review.Date, review.Time = FormatDate(date)
		result = append(result, review)
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetReviewNotCommit,
		}
	}

	return result, nil
}

func (db *Wrapper) CreateReview(id int, review resPkg.NewReview) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.RCreateReviewTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	_, err = tx.Exec(contextTransaction,
		"INSERT INTO review (author, restaurant, text, rate) VALUES ($1, $2, $3, $4)",
		id, review.Restaurant.Id, review.Text, review.Rate)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.RCreateReviewNotInsert,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Text: errPkg.RCreateReviewNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) SearchCategory(name string) ([]int, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RSearchCategoryTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	rows, err := tx.Query(contextTransaction,
		"SELECT restaurant FROM restaurant_category WHERE fts @@ to_tsquery($1)",
		name)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RSearchCategoryNotSelect,
		}
	}

	var result []int
	for rows.Next() {
		var restaurantId int
		err := rows.Scan(&restaurantId)
		if err != nil {
			return nil, &errPkg.Errors{
				Text: errPkg.RSearchCategoryNotScan,
			}
		}
		result = append(result, restaurantId)
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RSearchCategoryNotCommit,
		}
	}

	return result, nil
}

func (db *Wrapper) SearchRestaurant(name string) ([]int, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RSearchRestaurantTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	rows, err := tx.Query(contextTransaction,
		"SELECT id FROM restaurant WHERE fts @@ to_tsquery($1)",
		name)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RSearchRestaurantNotSelect,
		}
	}

	var result []int
	for rows.Next() {
		var restaurantId int
		err := rows.Scan(&restaurantId)
		if err != nil {
			return nil, &errPkg.Errors{
				Text: errPkg.RSearchRestaurantNotScan,
			}
		}
		result = append(result, restaurantId)
	}

	if result == nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RSearchRestaurantEmpty,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RSearchRestaurantNotCommit,
		}
	}

	return result, nil
}

func (db *Wrapper) GetGeneralInfoRestaurant(id int) (*resPkg.Restaurants, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetGeneralInfoTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	restaurant := resPkg.Restaurants{}
	err = tx.QueryRow(contextTransaction,
		"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant WHERE id = $1",
		id).Scan(&restaurant.Id, &restaurant.Img, &restaurant.Name, &restaurant.CostForFreeDelivery,
		&restaurant.MinDelivery, &restaurant.MaxDelivery, &restaurant.Rating)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetGeneralInfoNotScan,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetGeneralInfoNotCommit,
		}
	}

	return &restaurant, nil
}

func (db *Wrapper) GetFavoriteRestaurants(id int) ([]resPkg.Restaurants, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetFavoriteRestaurantsTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	rows, err := tx.Query(contextTransaction,
		"SELECT r.id, r.avatar, r.name, r.price_delivery, r.min_delivery_time, r.max_delivery_time, r.rating, fr.position"+
			" FROM restaurant r RIGHT JOIN favorite_restaurant fr ON fr.restaurant = r.id WHERE fr.client = $1",
		id)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetFavoriteRestaurantsRestaurantsNotSelect,
		}
	}

	var restaurants []resPkg.Restaurants
	mapRestaurants := make(map[int]resPkg.Restaurants)

	for rows.Next() {
		var position *int32
		var restaurant resPkg.Restaurants
		err := rows.Scan(&restaurant.Id, &restaurant.Img, &restaurant.Name, &restaurant.CostForFreeDelivery,
			&restaurant.MinDelivery, &restaurant.MaxDelivery, &restaurant.Rating, &position)
		if err != nil {
			return nil, &errPkg.Errors{
				Text: errPkg.RGetFavoriteRestaurantsRestaurantsNotScan,
			}
		}

		mapRestaurants[int(*position)] = restaurant
	}

	for i := 0; i < len(mapRestaurants); i++ {
		restaurants = append(restaurants, mapRestaurants[i])
	}

	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetFavoriteRestaurantsRestaurantsNotExist,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetFavoriteRestaurantsInfoNotCommit,
		}
	}

	return restaurants, nil
}

func (db *Wrapper) IsFavoriteRestaurant(idClient int, idRestaurant int) (bool, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return false, &errPkg.Errors{
			Text: errPkg.RIsFavoriteRestaurantsTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var check *int32
	err = tx.QueryRow(contextTransaction,
		"SELECT id FROM favorite_restaurant WHERE client = $1 AND restaurant = $2",
		idClient, idRestaurant).Scan(&check)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, nil
		}
		return false, &errPkg.Errors{
			Text: errPkg.RIsFavoriteRestaurantsRestaurantsNotSelect,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return false, &errPkg.Errors{
			Text: errPkg.RIsFavoriteRestaurantsInfoNotCommit,
		}
	}

	return true, nil
}

func (db *Wrapper) EditRestaurantInFavorite(idRestaurant int, idClient int) (bool, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return false, &errPkg.Errors{
			Text: errPkg.REditRestaurantInFavoriteTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var check *int32
	err = tx.QueryRow(contextTransaction,
		"DELETE FROM favorite_restaurant WHERE client = $1 AND restaurant = $2 RETURNING id",
		idClient, idRestaurant).Scan(&check)
	if err != pgx.ErrNoRows {
		err = tx.Commit(contextTransaction)
		if err != nil {
			return false, &errPkg.Errors{
				Text: errPkg.REditRestaurantInFavoriteInfoNotCommit,
			}
		}
		return false, nil
	}

	if err != nil && err != pgx.ErrNoRows {
		return false, &errPkg.Errors{
			Text: errPkg.REditRestaurantInFavoriteRestaurantsNotDelete,
		}
	}

	var positionRestaurants *int32
	err = tx.QueryRow(contextTransaction,
		"SELECT MAX(position) FROM favorite_restaurant WHERE client = $1", idClient).Scan(&positionRestaurants)
	if err != nil {
		return false, &errPkg.Errors{
			Text: errPkg.REditRestaurantInFavoriteRestaurantsNotSelect,
		}
	}

	var pos int
	if positionRestaurants == nil {
		pos = 0
	} else {
		pos = int(*positionRestaurants) + 1
	}

	_, err = tx.Exec(contextTransaction,
		"INSERT INTO favorite_restaurant (client, restaurant, position) VALUES ($1, $2, $3)",
		idClient, idRestaurant, pos)
	if err != nil {
		return false, &errPkg.Errors{
			Text: errPkg.REditRestaurantInFavoriteRestaurantsNotScan,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return false, &errPkg.Errors{
			Text: errPkg.REditRestaurantInFavoriteInfoNotCommit,
		}
	}

	return true, nil
}

func (db *Wrapper) GetPromoCodes() ([]resPkg.Promocode, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetPromoCodesTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	rows, err := tx.Query(contextTransaction,
		"SELECT name, description, avatar, restaurant, code FROM promocode ORDER BY random() LIMIT 5")
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetPromoCodesCodesNotSelect,
		}
	}
	var promoCodes []resPkg.Promocode
	for rows.Next() {
		var promoCode resPkg.Promocode
		err := rows.Scan(&promoCode.Name, &promoCode.Description, &promoCode.Img, &promoCode.RestaurantId, &promoCode.Code)
		if err != nil {
			return nil, &errPkg.Errors{
				Text: errPkg.RGetRadiosRadiosNotScan,
			}
		}
		promoCodes = append(promoCodes, promoCode)
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Text: errPkg.RGetPromoCodesNotCommit,
		}
	}

	return promoCodes, nil
}
