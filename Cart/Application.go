package Cart

import (
	"2021_2_GORYACHIE_MEKSIKANSI/Utils"
)

func GetCart(db Utils.WrapperCart, id int) (Utils.Cart, error) {
	result, _ := db.GetCart(id)
	return result, nil
}

func UpdateCart(db Utils.WrapperCart, dishes Utils.Cart,  clientId int) ([]Utils.CastDishesErrs, error) {
	err := db.DeleteCart(clientId)
	if err != nil {
		return nil, err
	}
	return db.UpdateCart(dishes, clientId)
}

func DeleteCart(db Utils.WrapperCart, id int) error {
	return db.DeleteCart(id)
}
