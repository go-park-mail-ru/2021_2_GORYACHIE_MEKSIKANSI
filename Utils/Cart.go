package Utils

type CartResponse struct {
	Restaurant interface{}   `json:"restaurant"`
	Dishes     []interface{} `json:"dishes"`
}

type RestaurantCart struct {
	Id int `json:"id"`
}

type DishesCart struct {
	Id           int           `json:"id"`
	ItemNumber   int           `json:"itNum"`
	Name         string        `json:"name"`
	Count        int           `json:"count"`
	Cost         int           `json:"cost"`
	Description  string        `json:"desc"`
	RadiosCart   []interface{} `json:"radios"`
	CheckboxCart []interface{} `json:"checkbox"`
}

type RadiosCart struct {
	RadiosId int `json:"rid"`
	Id       int `json:"id"`
}

type CheckboxCart struct {
	Id int `json:"id"`
}
