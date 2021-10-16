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

type DishesResponse struct {
	DishesGet interface{} `json:"dishes"`
}

type Dishes struct {
	Id				int				`json:"id"`
	Img				string			`json:"img"`
	Title			string			`json:"title"`
	Cost			int				`json:"cost"`
	Description		string			`json:"description"`
	Radios			[]interface{}	`json:"radios"`
	Ingredient		[]interface{}	`json:"CheckboxesRows"`  // TODO: завтра подумать над итогом json названием
}

type Radios struct {
	Title	string			`json:"title"`
	Id		int				`json:"id"`
	Rows	[]interface{}	`json:"rows"`
}

type Rows struct {
	Id		int		`json:"id"`
	Name	string	`json:"name"`
}

type Ingredient struct {
	Id		int		`json:"id"`
	Title	string	`json:"title"`
	Cost	int		`json:"cost"`
}
