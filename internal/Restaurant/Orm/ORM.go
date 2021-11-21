package Orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyError"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Restaurant"
	"context"
	"github.com/jackc/pgx/v4"
)

type Wrapper struct {
	Conn Interface.ConnectionInterface
}

func (db *Wrapper) GetRestaurants() ([]Restaurant.Restaurants, error) {
	row, err := db.Conn.Query(context.Background(),
		"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant ORDER BY random() LIMIT 50")
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

	return result, nil
}

func (db *Wrapper) GetGeneralInfoRestaurant(id int) (*Restaurant.RestaurantId, error) {
	var restaurant Restaurant.RestaurantId
	err := db.Conn.QueryRow(context.Background(),
		"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant WHERE id = $1", id).Scan(
		&restaurant.Id, &restaurant.Img, &restaurant.Name, &restaurant.CostForFreeDelivery, &restaurant.MinDelivery,
		&restaurant.MaxDelivery, &restaurant.Rating)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetGeneralInfoRestaurantNotFound,
		}
	}
	return &restaurant, nil
}

func (db *Wrapper) GetTagsRestaurant(id int) ([]Restaurant.Tag, error) {
	rowCategory, err := db.Conn.Query(context.Background(),
		"SELECT id, category FROM restaurant_category WHERE restaurant = $1", id)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetTagsCategoryNotSelect,
		}
	}
	var tags []Restaurant.Tag
	tag := Restaurant.Tag{}
	for rowCategory.Next() {
		err := rowCategory.Scan(&tag.Id, &tag.Name)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.RGetTagsCategoryRestaurantNotScan,
			}
		}
		tags = append(tags, tag)
	}
	if tags == nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetTagsTagsNotFound,
		}
	}
	return tags, nil
}

func getDishesRestaurant(db *Wrapper, name string, id int) ([]Restaurant.DishesMenu, error) {
	var dishes []Restaurant.DishesMenu
	dish := Restaurant.DishesMenu{}
	rowDishes, err := db.Conn.Query(context.Background(),
		"SELECT id, avatar, name, cost, kilocalorie FROM dishes WHERE category_restaurant = $1 AND restaurant = $2",
		name, id)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetMenuDishesNotSelect,
		}
	}

	for rowDishes.Next() {
		err := rowDishes.Scan(&dish.Id, &dish.Img, &dish.Name, &dish.Cost, &dish.Kilocalorie)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.RGetDishesRestaurantDishesNotScan,
			}
		}
		dishes = append(dishes, dish)
	}
	return dishes, nil
}

func (db *Wrapper) GetMenu(id int) ([]Restaurant.Menu, error) {
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetMenuTransactionNotCreate,
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	var result []Restaurant.Menu
	rowDishes, err := tx.Query(context.Background(),
		"SELECT DISTINCT category_restaurant FROM dishes WHERE restaurant = $1", id)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetMenuDishesNotSelect,
		}
	}

	for rowDishes.Next() {
		var menu Restaurant.Menu
		err := rowDishes.Scan(&menu.Name)
		if err != nil {
			return nil, err
		}

		dishes, err := getDishesRestaurant(db, menu.Name, id)
		if err != nil {
			return nil, err
		}

		menu.DishesMenu = dishes
		result = append(result, menu)
	}

	if result == nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetMenuDishesNotFound,
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetMenuDishesNotCommit,
		}
	}

	return result, nil
}

func (db *Wrapper) GetStructDishes(dishesId int) ([]Restaurant.Ingredients, error) {
	var ingredients []Restaurant.Ingredients
	rowDishes, err := db.Conn.Query(context.Background(),
		"SELECT id, name, cost FROM structure_dishes WHERE food = $1", dishesId)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetStructDishesStructDishesNotSelect,
		}
	}

	for rowDishes.Next() {
		var ingredient Restaurant.Ingredients
		err := rowDishes.Scan(&ingredient.Id, &ingredient.Title, &ingredient.Cost)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.RGetStructDishesStructDishesNotScan,
			}
		}
		ingredients = append(ingredients, ingredient)
	}
	return ingredients, nil
}

func (db *Wrapper) GetDishes(restId int, dishesId int) (*Restaurant.Dishes, error) {
	var dishes Restaurant.Dishes
	err := db.Conn.QueryRow(context.Background(),
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
	return &dishes, nil
}

func (db *Wrapper) GetRadios(dishesId int) ([]Restaurant.Radios, error) {
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRadiosNotCreate,
		}
	}

	defer func(tx pgx.Tx) {
		tx.Rollback(context.Background())
	}(tx)

	var radios []Restaurant.Radios
	m := make(map[int][]Restaurant.CheckboxesRows)

	rowDishes, err := tx.Query(context.Background(),
		"SELECT r.id, r.name, sr.id, sr.name FROM radios r "+
			"LEFT JOIN structure_radios sr ON sr.radios=r.id WHERE r.food = $1", dishesId)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRadiosRadiosNotSelect,
		}
	}

	for rowDishes.Next() {
		var rad Restaurant.Radios
		var elementRadios Restaurant.CheckboxesRows
		err := rowDishes.Scan(&rad.Id, &rad.Title, &elementRadios.Id, &elementRadios.Name)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.RGetRadiosRadiosNotScan,
			}
		}

		if _, ok := m[rad.Id]; !ok {
			radios = append(radios, rad)
		}
		m[rad.Id] = append(m[rad.Id], elementRadios)
	}

	for i, rad := range radios {
		radios[i].Rows = m[rad.Id]
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.RGetRadiosNotCommit,
		}
	}

	return radios, nil
}
