//go:generate easyjson -no_std_marshalers restaurant.go
package restaurant

import "mime/multipart"

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

type AllRestaurants struct {
	Restaurant []Restaurants `json:"restaurants_info"`
	AllTags    []Tag         `json:"tags"`
}

type AllRestaurantsPromo struct {
	Restaurant []Restaurants `json:"restaurants_info"`
	AllTags    []Tag         `json:"tags"`
	AllPromo   []PromoCode   `json:"promo_code"`
}

type PromoCode struct {
	Name         string `json:"name"`
	Description  string `json:"desc"`
	Img          string `json:"img"`
	RestaurantId int    `json:"restId"`
	Code         string `json:"code"`
}

type RestaurantIdResponse struct {
	RestaurantsGet interface{} `json:"restaurant"`
}

type RestaurantId struct {
	Id                  int     `json:"id"`
	Img                 string  `json:"img"`
	Name                string  `json:"name"`
	CostForFreeDelivery int     `json:"costFFD"`
	MinDelivery         int     `json:"minDTime"`
	MaxDelivery         int     `json:"maxDTime"`
	Rating              float32 `json:"rating"`
	Favourite           bool    `json:"favourite"`
	Tags                []Tag   `json:"tags"`
	Menu                []Menu  `json:"menu"`
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
	Id          int           `json:"id"`
	Img         string        `json:"img"`
	Title       string        `json:"name"`
	Cost        int           `json:"cost"`
	Ccal        int           `json:"ccal"`
	Description string        `json:"desc"`
	Radios      []Radios      `json:"radios,omitempty"`
	Ingredient  []Ingredients `json:"ingredients,omitempty"`
}

type Radios struct {
	Title string          `json:"name"`
	Id    int             `json:"id"`
	Rows  []ElementRadios `json:"opt"`
}

type ElementRadios struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Ingredients struct {
	Id    int    `json:"id"`
	Title string `json:"name"`
	Cost  int    `json:"cost"`
}

//easyjson:json
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
	Tags                []Tag    `json:"tags"`
	Reviews             []Review `json:"reviews,omitempty"`
	Status              bool     `json:"status_favorite"`
}

type Review struct {
	Name string `json:"name"`
	Text string `json:"text"`
	Date string `json:"date"`
	Time string `json:"time"`
	Rate int    `json:"rate"`
}

type SearchRestaurant struct {
	SearchText string `json:"searchText"`
}

//easyjson:json
type ResFavouriteNew struct {
	Id int `json:"id"`
}

type CreateDishHost struct {
	Dishes     Dishes `json:"dishes"`
	FileHeader *multipart.FileHeader
}

type UpdateDishHost struct {
	Dishes     Dishes `json:"dishes"`
	FileHeader *multipart.FileHeader
}

type DishHost struct {
	Dishes     CreateDishes `json:"dishes"`
	FileHeader *multipart.FileHeader
}

type CreateDishes struct {
	Id                 int    `json:"id"`
	Title              string `json:"name"`
	Cost               int    `json:"cost"`
	Ccal               int    `json:"ccal"`
	Description        string `json:"desc"`
	Protein            int
	Falt               int
	Carbohydrates      int
	Weight             int
	CategoryDishes     string
	CategoryRestaurant string
	Count              int
	Radios             []CreateRadios      `json:"radios,omitempty"`
	Ingredient         []CreateIngredients `json:"ingredients,omitempty"`
}

type CreateRadios struct {
	Title string                `json:"name"`
	Id    int                   `json:"id"`
	Rows  []CreateElementRadios `json:"opt"`
}

type CreateElementRadios struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Protein       int
	Falt          int
	Carbohydrates int
}

type CreateIngredients struct {
	Id            int    `json:"id"`
	Title         string `json:"name"`
	Cost          int    `json:"cost"`
	Protein       int
	Falt          int
	Carbohydrates int
	Count         int
}

type DeleteDishesHost struct {
	IdDishes int `json:"id_dishes"`
}

type ResFavouriteStatus struct {
	Status bool `json:"status"`
}
