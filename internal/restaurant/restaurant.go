package restaurant

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
	Title string           `json:"name"`
	Id    int              `json:"id"`
	Rows  []CheckboxesRows `json:"opt"`
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
	Tags                []Tag    `json:"tags"`
	Reviews             []Review `json:"reviews,omitempty"`
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

type ResFavouriteNew struct {
	Id int `json:"id"`
}

type ResFavouriteStatus struct {
	Status bool `json:"status"`
}

func (r *ResReview) CastFromRestaurantId(rest RestaurantId) {
	r.Id = rest.Id
	r.Img = rest.Img
	r.Name = rest.Name
	r.CostForFreeDelivery = rest.CostForFreeDelivery
	r.MinDelivery = rest.MinDelivery
	r.MaxDelivery = rest.MaxDelivery
	r.Rating = rest.Rating
	r.Tags = rest.Tags
}