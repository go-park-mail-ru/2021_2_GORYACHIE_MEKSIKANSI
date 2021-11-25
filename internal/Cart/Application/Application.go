package Application

import (
	Cart2 "2021_2_GORYACHIE_MEKSIKANSI/internal/Cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
)

type Cart struct {
	DB Interface.WrapperCartServer
}

func (c *Cart) CalculatePriceDelivery(id int) (int, error) {
	return c.DB.CalculatePriceDelivery(id)
}

func (c *Cart) GetCart(id int) (*Cart2.ResponseCartErrors, error) {
	return c.DB.GetCart(id)
}

func (c *Cart) UpdateCart(dishes Cart2.RequestCartDefault, clientId int) (*Cart2.ResponseCartErrors, error) {
	return c.DB.UpdateCart(dishes, clientId)
}

func (c *Cart) DeleteCart(id int) error {
	return c.DB.DeleteCart(id)
}
