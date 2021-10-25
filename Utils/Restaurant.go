package Utils

type RestaurantsResponse struct {
	RestaurantsGet interface{} `json:"restaurants"`
}

type Restaurants struct {
	Id                  int     `json:"id"`
	Img                 string  `json:"img"`
	Name                string  `json:"name"`
	CostForFreeDelivery int     `json:"costFFD"`
	MinDelivery         int     `json:"minDTime"`
	MaxDelivery         int     `json:"maxDTime"`
	Rating              float32 `json:"rate"`
}

type RestaurantIdResponse struct {
	RestaurantsGet interface{} `json:"restaurant"`
}

type RestaurantId struct {
	Id                  int         `json:"id"`
	Img                 string      `json:"img"`
	Name                string      `json:"name"`
	CostForFreeDelivery int         `json:"costFFD"` // TODO(N): надо бы rename json
	MinDelivery         int         `json:"minDTime"`
	MaxDelivery         int         `json:"maxDTime"`
	Rating              float32     `json:"rating"`
	Tags                interface{} `json:"tags"`
	Menu                interface{} `json:"menu"`
}

type Tag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Menu struct {
	Name       string      `json:"name"`
	DishesMenu interface{} `json:"dishes"`
}

type DishesMenu struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Cost        int    `json:"cost"`
	Kilocalorie int    `json:"ccal"`
	Img         string `json:"img"`
}

type DishesResponse struct {
	DishesGet interface{} `json:"dishes"`
}

type Dishes struct {
	Id          int         `json:"id"`
	Img         string      `json:"img"`
	Title       string      `json:"name"`
	Cost        int         `json:"cost"`
	Ccal        int         `json:"ccal"`
	Description string      `json:"desc"`
	Radios      interface{} `json:"radios,omitempty"`
	Ingredient  interface{} `json:"ingredients,omitempty"`
}

type Radios struct {
	Title string      `json:"name"`
	Id    int         `json:"id"`
	Rows  interface{} `json:"opt"`
}

type CheckboxesRows struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Ingredients struct {
	Id    int    `json:"id"`
	Title string `json:"name"`
	Cost  int    `json:"cost"`
}
