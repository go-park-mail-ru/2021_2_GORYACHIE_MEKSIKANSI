package cart

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant"
)

type ResponseCart struct {
	Cart interface{} `json:"cart"`
}

type ResponseCartErrors struct {
	Restaurant RestaurantIdCastResponse `json:"restaurant"`
	Dishes     []DishesCartResponse     `json:"dishes"`
	Cost       CostCartResponse         `json:"cost"`
	PromoCode  PromoCode                `json:"promo_code"`
	DishErr    []CastDishesErrs         `json:"dishesErrs,omitempty"`
}

type PromoCode struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
	Code        string `json:"code"`
}

type CostCartResponse struct {
	DCost   int `json:"dCost"`
	SumCost int `json:"sumCost"`
}

type DishesCartResponse struct {
	Id             int                      `json:"id"`
	ItemNumber     int                      `json:"itNum"`
	Img            string                   `json:"img"`
	Name           string                   `json:"name"`
	Count          int                      `json:"count"`
	Cost           int                      `json:"cost"`
	Kilocalorie    int                      `json:"ccal"`
	Weight         int                      `json:"weight"`
	Description    string                   `json:"desc"`
	RadiosCart     []RadiosCartResponse     `json:"radios,omitempty"`
	IngredientCart []IngredientCartResponse `json:"ingredients,omitempty"`
}

type RadiosCartResponse struct {
	Name     string `json:"name"`
	RadiosId int    `json:"rId"`
	Id       int    `json:"id"`
}

type IngredientCartResponse struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
	Cost int    `json:"cost"`
}

type RestaurantIdCastResponse struct {
	Id                  int     `json:"id"`
	Img                 string  `json:"img"`
	Name                string  `json:"name"`
	CostForFreeDelivery int     `json:"costFFD"`
	MinDelivery         int     `json:"minDTime"`
	MaxDelivery         int     `json:"maxDTime"`
	Rating              float32 `json:"rating"`
}

type CastErrs struct {
	CastDishesErrs []CastDishesErrs `json:"dishesErrs"`
}

type CastDishesErrs struct {
	ItemNumber int    `json:"itNum"`
	NameDish   string `json:"nameDish"`
	CountAvail int    `json:"countAvail"`
}

type CartRequest struct {
	Cart RequestCartDefault `json:"cart"`
}

type RequestCartDefault struct {
	Restaurant RestaurantRequest `json:"restaurant"`
	Dishes     []DishesRequest   `json:"dishes"`
	PromoCode  string            `json:"promo_code"`
}

type RestaurantRequest struct {
	Id int `json:"id"`
}

type DishesRequest struct {
	Id          int                      `json:"id"`
	ItemNumber  int                      `json:"itNum"`
	Count       int                      `json:"count"`
	Radios      []RadiosCartRequest      `json:"radios"`
	Ingredients []IngredientsCartRequest `json:"ingredients"`
}

type RadiosCartRequest struct {
	RadiosId int `json:"rId"`
	Id       int `json:"id"`
}

type IngredientsCartRequest struct {
	Id int `json:"id"`
}

func (c *ResponseCartErrors) CastFromRequestCartDefault(a RequestCartDefault) {
	for i, dish := range a.Dishes {
		c.Dishes[i].Id = dish.Id
		c.Dishes[i].ItemNumber = dish.ItemNumber

		for j, ingredient := range dish.Ingredients {
			c.Dishes[i].IngredientCart[j].Id = ingredient.Id
		}

		for j, radios := range dish.Radios {
			c.Dishes[i].RadiosCart[j].Id = radios.Id
			c.Dishes[i].RadiosCart[j].RadiosId = radios.RadiosId
		}
	}
}

func (c *ResponseCartErrors) CastFromRestaurantId(a restaurant.RestaurantId) {
	c.Restaurant.Id = a.Id
	c.Restaurant.Img = a.Img
	c.Restaurant.Rating = a.Rating
	c.Restaurant.CostForFreeDelivery = a.CostForFreeDelivery
	c.Restaurant.Name = a.Name
	c.Restaurant.MaxDelivery = a.MaxDelivery
	c.Restaurant.MinDelivery = a.MinDelivery
}
