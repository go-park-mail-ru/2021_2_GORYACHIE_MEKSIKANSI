package Cart

import (
	"2021_2_GORYACHIE_MEKSIKANSI/Utils"
)

func GetCart(db Utils.WrapperCart, id int) (Utils.CartResponse, error) {
	result, _ := db.GetCart(id)
	return result, nil
}

func UpdateCart(db Utils.WrapperCart, dishes []Utils.DishesCart, restaurantId int, clientId int) error {
	return db.UpdateCart(dishes, restaurantId, clientId)
}

func DeleteCart(db Utils.WrapperCart, id int) error {
	return db.DeleteCart(id)
}
