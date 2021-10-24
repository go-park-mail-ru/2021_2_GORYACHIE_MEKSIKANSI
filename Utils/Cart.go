package Utils

type CartResponse struct {
	Restaurant RestaurantId         `json:"restaurant"`
	Dishes []DishesCartResponse     `json:"dishes"`
	Cost   CostCartResponse         `json:"cost"`
}

/*type RestaurantCartResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}*/

type CostCartResponse struct {
	DCost int `json:"dCost"`
	SumCost int `json:"sumCost"`
}

type DishesCartResponse struct {
	Id             int              `json:"id"`
	ItemNumber     int              `json:"itNum"`
	Img            string           `json:"img"`
	Name           string           `json:"name"`
	Count          int              `json:"count"`
	Cost           int              `json:"cost"`
	Description    string               `json:"desc"`
	RadiosCart     []RadiosCartResponse     `json:"radios"`
	IngredientCart []IngredientCartResponse `json:"ingredients"`
}

type RadiosCartResponse struct {
	Name     string `json:"name"`
	RadiosId int    `json:"rid"`
	Id       int    `json:"id"`
}

type IngredientCartResponse struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
	Cost int    `json:"cost"`
}

type CastErrs struct {
	CastDishesErrs []CastDishesErrs `json:"dishesErrs"`
}

type CastDishesErrs struct {
	ItemNumber int    `json:"itNum"`
	Explain    string `json:"explain"`
}

type CartRequest struct {
	Restaurant RestaurantRequest `json:"restaurant"`
	Dishes     []DishesRequest   `json:"dishes"`
}

type RestaurantRequest struct {
	Id int `json:"id"`
}

type DishesRequest struct {
	Id int `json:"id"`
	Count int `json:"count"`
	Radios []RadiosCartRequest `json:"radios"`
	Ingredients []IngredientsCartRequest `json:"ingredients"`
}

type RadiosCartRequest struct {
	RadiosId int `json:"rId"`
	Id int `json:"id"`
}

type IngredientsCartRequest struct {
	Id int `json:"id"`
}
