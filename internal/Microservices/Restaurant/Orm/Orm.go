package Orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyError"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Restaurant"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Util"
	"context"
	"github.com/jackc/pgx/v4"
	"time"
)

type Wrapper struct {
	Conn Interface.ConnectionInterface
}

func (db *Wrapper) GetRestaurants() ([]Restaurant.Restaurants, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRestaurantsTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	row, err := tx.Query(contextTransaction,
		"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant LIMIT 50")
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRestaurantsRestaurantsNotSelect,
		}
	}

	restaurant := Restaurant.Restaurants{}
	var result []Restaurant.Restaurants
	for row.Next() {
		err := row.Scan(&restaurant.Id, &restaurant.Img, &restaurant.Name, &restaurant.CostForFreeDelivery,
			&restaurant.MinDelivery, &restaurant.MaxDelivery, &restaurant.Rating)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.RGetRestaurantsRestaurantsNotScan,
			}
		}
		result = append(result, restaurant)
	}

	if result == nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRestaurantsRestaurantsNotFound,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRestaurantsNotCommit,
		}
	}

	return result, nil
}

func (db *Wrapper) GetRestaurant(id int) (*Restaurant.RestaurantId, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRestaurantTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var restaurant Restaurant.RestaurantId
	err = tx.QueryRow(contextTransaction,
		"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant WHERE id = $1", id).Scan(
		&restaurant.Id, &restaurant.Img, &restaurant.Name, &restaurant.CostForFreeDelivery, &restaurant.MinDelivery,
		&restaurant.MaxDelivery, &restaurant.Rating)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRestaurantRestaurantNotFound,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRestaurantNotCommit,
		}
	}

	return &restaurant, nil
}

func (db *Wrapper) GetTagsRestaurant(id int) ([]Restaurant.Tag, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetTagsRestaurantTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	rowCategory, err := tx.Query(contextTransaction,
		"SELECT id, category, place FROM restaurant_category WHERE restaurant = $1", id)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetTagsRestaurantNotSelect,
		}
	}

	var tags []Restaurant.Tag
	var tag Restaurant.Tag

	place := make(map[int]Restaurant.Tag)
	for rowCategory.Next() {
		var placeCategory int
		err := rowCategory.Scan(&tag.Id, &tag.Name, &placeCategory)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.RGetTagsRestaurantRestaurantNotScan,
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
			Alias: errPkg.RGetTagsRestaurantNotCommit,
		}
	}

	return tags, nil
}

func (db *Wrapper) GetMenu(id int) ([]Restaurant.Menu, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetMenuTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var result []Restaurant.Menu

	rowDishes, err := tx.Query(contextTransaction,
		"SELECT category_restaurant, id, avatar, name, cost, kilocalorie, place, place_category FROM dishes WHERE restaurant = $1", id)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetMenuDishesCategoryNotSelect,
		}
	}

	categoryPlace := make(map[int]Restaurant.Menu)
	place := make(map[int]map[int]Restaurant.DishesMenu)

	for rowDishes.Next() {
		var menu Restaurant.Menu
		var dish Restaurant.DishesMenu
		var placeDish, placeCategory int
		err := rowDishes.Scan(&menu.Name, &dish.Id, &dish.Img, &dish.Name, &dish.Cost, &dish.Kilocalorie, &placeDish, &placeCategory)
		if err != nil {
			return nil, err
		}

		temp := place[placeCategory]
		if temp == nil {
			temp = make(map[int]Restaurant.DishesMenu)
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
			Alias: errPkg.RGetMenuDishesNotFound,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetMenuNotCommit,
		}
	}

	return result, nil
}

func (db *Wrapper) GetStructDishes(dishesId int) ([]Restaurant.Ingredients, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetStructDishesTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var ingredients []Restaurant.Ingredients
	rowDishes, err := tx.Query(contextTransaction,
		"SELECT id, name, cost, place FROM structure_dishes WHERE food = $1", dishesId)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetStructDishesStructDishesNotSelect,
		}
	}

	place := make(map[int]Restaurant.Ingredients)
	for rowDishes.Next() {
		var placeDish int
		var ingredient Restaurant.Ingredients
		err := rowDishes.Scan(&ingredient.Id, &ingredient.Title, &ingredient.Cost, &placeDish)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.RGetStructDishesStructDishesNotScan,
			}
		}
		place[placeDish] = ingredient
	}

	for i := 0; i < len(place); i++ {
		ingredients = append(ingredients, place[i])
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetStructDishesNotCommit,
		}
	}

	return ingredients, nil
}

func (db *Wrapper) GetDishes(restId int, dishesId int) (*Restaurant.Dishes, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetDishesTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	var dishes Restaurant.Dishes
	err = tx.QueryRow(contextTransaction,
		"SELECT id, avatar, name, cost, kilocalorie, description FROM dishes WHERE id = $1 AND restaurant = $2",
		dishesId, restId).Scan(
		&dishes.Id, &dishes.Img, &dishes.Title, &dishes.Cost, &dishes.Ccal, &dishes.Description)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, &errPkg.Errors{
				Alias: errPkg.RGetDishesDishesNotFound,
			}
		}
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetDishesDishesNotScan,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetDishesNotCommit,
		}
	}

	return &dishes, nil
}

func (db *Wrapper) GetRadios(dishesId int) ([]Restaurant.Radios, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRadiosTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	rowDishes, err := tx.Query(contextTransaction,
		"SELECT r.id, r.name, sr.id, sr.name, r.place, sr.place FROM radios r "+
			"LEFT JOIN structure_radios sr ON sr.radios=r.id WHERE r.food = $1", dishesId)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRadiosRadiosNotSelect,
		}
	}

	var radios []Restaurant.Radios
	place := make(map[int]map[int]Restaurant.CheckboxesRows)
	radiosInfo := make(map[int]Restaurant.Radios)

	for rowDishes.Next() {
		var rad Restaurant.Radios
		var elementRadios Restaurant.CheckboxesRows
		var placeRadios, placeElementRadios int
		err := rowDishes.Scan(&rad.Id, &rad.Title, &elementRadios.Id, &elementRadios.Name, &placeRadios, &placeElementRadios)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.RGetRadiosRadiosNotScan,
			}
		}

		temp := place[placeRadios]
		if temp == nil {
			temp = make(map[int]Restaurant.CheckboxesRows)
		}
		temp[placeElementRadios] = elementRadios
		place[placeRadios] = temp
		radiosInfo[placeRadios] = rad
	}

	for i := 0; i < len(place); i++ {
		radios = append(radios, radiosInfo[i])
		var rows []Restaurant.CheckboxesRows
		for j := 0; j < len(place[i]); j++ {
			rows = append(rows, place[i][j])
		}
		radios[i].Rows = rows
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRadiosNotCommit,
		}
	}

	return radios, nil
}

func (db *Wrapper) GetReview(id int) ([]Restaurant.Review, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetReviewTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	rowDishes, err := tx.Query(contextTransaction,
		"SELECT gn.name, r.text, r.date_create, r.rate FROM review r "+
			"LEFT JOIN general_user_info gn ON r.author=gn.id "+
			"WHERE r.restaurant = $1 ORDER BY r.date_create", id)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetReviewNotSelect,
		}
	}

	var result []Restaurant.Review
	for rowDishes.Next() {
		var review Restaurant.Review
		var date time.Time
		err := rowDishes.Scan(&review.Name, &review.Text, &date, &review.Rate)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.RGetReviewNotScan,
			}
		}
		review.Date, review.Time = Util.FormatDate(date)
		result = append(result, review)
	}

	if result == nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetReviewEmpty,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetReviewNotCommit,
		}
	}

	return result, nil
}

func (db *Wrapper) CreateReview(id int, review Restaurant.NewReview) error {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.RCreateReviewTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	_, err = tx.Exec(contextTransaction,
		"INSERT INTO review (author, restaurant, text, rate) VALUES ($1, $2, $3, $4)",
		id, review.Restaurant.Id, review.Text, review.Rate)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.RCreateReviewNotInsert,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.RCreateReviewNotCommit,
		}
	}

	return nil
}

func (db *Wrapper) SearchCategory(name string) ([]int, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RSearchCategoryTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	rows, err := tx.Query(contextTransaction,
		"SELECT id FROM restaurant_category WHERE fts @@ to_tsquery($1)",
		name)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RSearchCategoryNotSelect,
		}
	}

	var result []int
	for rows.Next() {
		var restaurantId int
		err := rows.Scan(&restaurantId)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.RSearchCategoryNotScan,
			}
		}
		result = append(result, restaurantId)
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RSearchCategoryNotCommit,
		}
	}

	return result, nil
}

func (db *Wrapper) SearchRestaurant(name string) ([]int, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RSearchRestaurantTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	rows, err := tx.Query(contextTransaction,
		"SELECT id FROM restaurant WHERE fts @@ to_tsquery($1)",
		name)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RSearchRestaurantNotSelect,
		}
	}

	var result []int
	for rows.Next() {
		var restaurantId int
		err := rows.Scan(&restaurantId)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.RSearchRestaurantNotScan,
			}
		}
		result = append(result, restaurantId)
	}

	if result == nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RSearchRestaurantEmpty,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RSearchRestaurantNotCommit,
		}
	}

	return result, nil
}

func (db *Wrapper) GetGeneralInfoRestaurant(id int) (*Restaurant.Restaurants, error) {
	contextTransaction := context.Background()
	tx, err := db.Conn.Begin(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetGeneralInfoTransactionNotCreate,
		}
	}

	defer tx.Rollback(contextTransaction)

	restaurant := Restaurant.Restaurants{}
	err = tx.QueryRow(contextTransaction,
		"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant WHERE id = $1",
		id).Scan(&restaurant.Id, &restaurant.Img, &restaurant.Name, &restaurant.CostForFreeDelivery,
		&restaurant.MinDelivery, &restaurant.MaxDelivery, &restaurant.Rating)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetGeneralInfoNotScan,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetGeneralInfoNotCommit,
		}
	}

	return &restaurant, nil
}
