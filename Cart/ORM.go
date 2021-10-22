package Cart

import (
	"2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
)

type Wrapper struct {
	Conn Utils.ConnectionInterface
}

func (db *Wrapper) GetCart(id int) (Utils.CartResponse, error) {
	var cart Utils.CartResponse
	var dishes []Utils.DishesCart
	var radios []Utils.RadiosCart
	var ingredients []Utils.CheckboxCart
	i := 0

	var restaurant Utils.RestaurantCart
	restaurant.Id = id
	cart.Restaurant = restaurant

	rows, _ := db.Conn.Query(context.Background(),
		"SELECT food, count_food FROM cart WHERE client_id = $1", id)
	for rows.Next() {
		var dish Utils.DishesCart
		_ = rows.Scan(&dish.Id, &dish.Count)
		rows, _ := db.Conn.Query(context.Background(),
			"SELECT name, cost, description FROM dishes WHERE id = $1", dish.Id)
		for rows.Next() {
			i++
			dish.ItemNumber = i

			_ = rows.Scan(&dish.Name, &dish.Cost, &dish.Description)

			rows, _ := db.Conn.Query(context.Background(),
				"SELECT checkbox FROM cart_structure_food WHERE client_id = $1", dish.Id)
			for rows.Next() {
				var ingredient Utils.CheckboxCart
				_ = rows.Scan(&ingredient.Id)
				ingredients = append(ingredients, ingredient)
			}
			dish.CheckboxCart = ingredients

			rows, _ = db.Conn.Query(context.Background(),
				"SELECT radios_id, radios FROM cart_radios_food WHERE client_id = $1", dish.Id)
			for rows.Next() {
				var radio Utils.RadiosCart
				_ = rows.Scan(&radio.RadiosId, &radio.Id)
				radios = append(radios, radio)
			}
			dish.RadiosCart = radios
			dishes = append(dishes, dish)
		}
	}
	cart.Dishes = dishes
	return cart, nil
}


func (db *Wrapper) DeleteCart(id int) error {
	_, _ = db.Conn.Exec(context.Background(),
		"DELETE FROM cart WHERE client_id = $1", id)
	_, _ = db.Conn.Exec(context.Background(),
		"DELETE FROM cart_structure_food WHERE client_id = $1", id)
	_, _ = db.Conn.Exec(context.Background(),
		"DELETE FROM cart_radios_food WHERE client_id = $1", id)
	return nil
}


func (db *Wrapper) UpdateCart(dishes []Utils.DishesCart, restaurantId int, clientId int) error {
	for _, dish := range dishes {
		check := 0
		_ = db.Conn.QueryRow(context.Background(),
			"SELECT id FROM dishes WHERE restaurant = $1 ", restaurantId).Scan(&check)

		if dish.Count == 0 {
			_, _ = db.Conn.Exec(context.Background(),
				"DELETE FROM cart WHERE food = $1 AND client_id = $2",
				dish.Id, clientId)
		}
		_, _ = db.Conn.Exec(context.Background(),
			"INSERT INTO cart (client_id, food, count_food) VALUES ($1, $2, $3)" +
			" ON CONFLICT (client_id, food) DO UPDATE SET count_food = $3",
			clientId, dish.Id, dish.Count)
	}

	return nil
}
