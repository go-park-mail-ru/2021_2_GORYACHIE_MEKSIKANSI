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
	var dishes []Utils.DishesCartResponse
	var radios []Utils.RadiosCartResponse
	var ingredients []Utils.IngredientCartResponse

	var restaurant Utils.RestaurantCartResponse
	restaurant.Id = id
	_ = db.Conn.QueryRow(context.Background(),
		"SELECT name FROM restaurant WHERE id = $1", id).Scan(&restaurant.Name)
	cart.Restaurant = restaurant

	rows, _ := db.Conn.Query(context.Background(),
		"SELECT food, count_food FROM cart WHERE client_id = $1", id)
	for rows.Next() {
		var dish Utils.DishesCartResponse
		_ = rows.Scan(&dish.Id, &dish.Count)
		rows, _ := db.Conn.Query(context.Background(),
			"SELECT name, cost, description, number_item, avatar FROM dishes WHERE id = $1", dish.Id)
		for rows.Next() {
			_ = rows.Scan(&dish.Name, &dish.Cost, &dish.Description, &dish.ItemNumber, &dish.Img)

			rows, _ := db.Conn.Query(context.Background(),
				"SELECT checkbox FROM cart_structure_food WHERE client_id = $1", dish.Id)
			for rows.Next() {
				var ingredient Utils.IngredientCartResponse
				_ = rows.Scan(&ingredient.Id)

				_ = db.Conn.QueryRow(context.Background(),
					"SELECT name FROM structure_dishes WHERE id = $1", ingredient.Id).Scan(&ingredient.Name)
				ingredients = append(ingredients, ingredient)
			}
			dish.IngredientCart = ingredients

			rows, _ = db.Conn.Query(context.Background(),
				"SELECT radios_id, radios FROM cart_radios_food WHERE client_id = $1", dish.Id)
			for rows.Next() {
				var radio Utils.RadiosCartResponse
				_ = rows.Scan(&radio.RadiosId, &radio.Id)

				_ = db.Conn.QueryRow(context.Background(),
					"SELECT name FROM structure_radios WHERE id = $1", radio.Id).Scan(&radio.Name)

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


func (db *Wrapper) UpdateCart(cart Utils.CartResponse, clientId int) ([]Utils.CastDishesErrs, error) {
	var dishesErrors []Utils.CastDishesErrs
	for _, dish := range cart.Dishes {
		var dishesError Utils.CastDishesErrs
		check := 0
		err := db.Conn.QueryRow(context.Background(),
			"SELECT id FROM dishes WHERE id = $1 ", dish.Id).Scan(&check)
		if err != nil {
			dishesError.ItemNumber = dish.ItemNumber
			dishesError.Explain = dish.Name
			dishesErrors = append(dishesErrors, dishesError)
			continue
		}

		_, _ = db.Conn.Exec(context.Background(),
			"INSERT INTO cart (client_id, food, count_food) VALUES ($1, $2, $3)",
			clientId, dish.Id, dish.Count)

		for _, ingredient := range dish.IngredientCart {
			check := 0
			err := db.Conn.QueryRow(context.Background(),
				"SELECT id FROM structure_dishes WHERE id = $1 ", ingredient.Id).Scan(&check)
			if err != nil {
				dishesError.ItemNumber = dish.ItemNumber
				dishesError.Explain = dish.Name
				dishesErrors = append(dishesErrors, dishesError)
				continue
			}

			_, _ = db.Conn.Exec(context.Background(),
				"INSERT INTO cart_structure_food (checkbox, client_id) VALUES ($1, $2)",
				ingredient.Id, clientId)
		}

		for _, radios := range dish.RadiosCart {
			check := 0
			err := db.Conn.QueryRow(context.Background(),
				"SELECT id FROM structure_radios WHERE id = $1", radios.Id).Scan(&check)
			if err != nil {
				dishesError.ItemNumber = dish.ItemNumber
				dishesError.Explain = dish.Name
				dishesErrors = append(dishesErrors, dishesError)
				continue
			}

			_, _ = db.Conn.Exec(context.Background(),
				"INSERT INTO cart_radios_food (radios_id, radios, client_id) VALUES ($1, $2, $3)",
				radios.RadiosId, radios.Id, clientId)
		}
	}

	return dishesErrors, nil
}
