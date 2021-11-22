package Restaurant

import "time"

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
	CostForFreeDelivery int         `json:"costFFD"`
	MinDelivery         int         `json:"minDTime"`
	MaxDelivery         int         `json:"maxDTime"`
	Rating              float32     `json:"rating"`
	Tags                interface{} `json:"tags"`
	Menu                []Menu      `json:"menu"`
}

type Tag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Menu struct {
	Name       string       `json:"name"`
	DishesMenu []DishesMenu `json:"dishes"`
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

type NewReview struct {
	Restaurant RestaurantId `json:"restaurant"`
	Text       string       `json:"text"`
	Rate       int          `json:"rate"`
}

type ResNewReview struct {
	Id int `json:"id"`
}

type ResReview struct {
	Id                  int      `json:"id"`
	Img                 string   `json:"img"`
	Name                string   `json:"name"`
	CostForFreeDelivery int      `json:"costFFD"`
	MinDelivery         int      `json:"minDTime"`
	MaxDelivery         int      `json:"maxDTime"`
	Rating              float32  `json:"rate"`
	Reviews             []Review `json:"reviews"`
}

type Review struct {
	Name string    `json:"name"`
	Text string    `json:"text"`
	Date time.Time `json:"date"`
	Time time.Time `json:"time"`
	Rate int       `json:"rate"`
}
