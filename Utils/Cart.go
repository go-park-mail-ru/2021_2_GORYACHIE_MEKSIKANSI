package Utils

type Cart struct {
	Restaurant RestaurantCart `json:"restaurant"`
	Dishes     []DishesCart   `json:"dishes"`
}

type RestaurantCart struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DishesCart struct {
	Id             int              `json:"id"`
	ItemNumber     int              `json:"itNum"`
	Img            string           `json:"img"`
	Name           string           `json:"name"`
	Count          int              `json:"count"`
	Cost           int              `json:"cost"`
	Description    string           `json:"desc"`
	RadiosCart     []RadiosCart     `json:"radios"`
	IngredientCart []IngredientCart `json:"ingredients"`
}

type RadiosCart struct {
	Id       int    `json:"id"`
	RadiosId int    `json:"rid"`
	Name     string `json:"name"`
}

type IngredientCart struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CastErrs struct {
	CastDishesErrs []CastDishesErrs `json:"dishesErrs"`
}

type CastDishesErrs struct {
	ItemNumber int    `json:"itNum"`
	Explain    string `json:"explain"`
}
