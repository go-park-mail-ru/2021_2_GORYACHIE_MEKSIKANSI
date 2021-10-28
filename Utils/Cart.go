package Utils

type ResponseCart struct {
	Cart interface{} `json:"cart"`
}

/*type ResponseCartDefault struct {
	Restaurant RestaurantIdCastResponse `json:"restaurant"`
	Dishes     []DishesCartResponse     `json:"dishes"`
	Cost       CostCartResponse         `json:"cost"`
}*/

type ResponseCartErrors struct {
	Restaurant RestaurantIdCastResponse `json:"restaurant"`
	Dishes     []DishesCartResponse     `json:"dishes"`
	Cost       CostCartResponse         `json:"cost"`
	DishErr    []CastDishesErrs         `json:"dishesErrs,omitempty"`
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

func (c *ResponseCartErrors) Cast(a RequestCartDefault) {
	for i, dish := range a.Dishes {
		c.Dishes[i].Id = dish.Id
		c.Dishes[i].ItemNumber = dish.ItemNumber
		c.Dishes[i].Count = dish.Count

		for j, ingredient := range dish.Ingredients {
			c.Dishes[i].IngredientCart[j].Id = ingredient.Id
		}

		for j, radios := range dish.Radios {
			c.Dishes[i].RadiosCart[j].Id = radios.Id
			c.Dishes[i].RadiosCart[j].RadiosId = radios.RadiosId
		}
	}
}
