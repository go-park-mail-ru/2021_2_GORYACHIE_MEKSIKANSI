package Utils

type Restaurant struct {
	Id                  int     `json:"id"`
	Img                 string  `json:"img"`
	Name                string  `json:"name"`
	CostForFreeDelivery int     `json:"costForFreeDelivery"`
	MinDelivery         int     `json:"minDeliveryTime"`
	MaxDelivery         int     `json:"maxDeliveryTime"`
	Rating              float32 `json:"rating"`
}

type RestaurantResponse struct {
	RestaurantsGet	interface{}	`json:"restaurants"`
}
