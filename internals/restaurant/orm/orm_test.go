package orm

import (
	resProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant/proto"
	rest "2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant/orm/mocks"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var AllRestaurants = []struct {
	testName   string
	input      int
	out        *rest.AllRestaurantsPromo
	outErr     string
	inputQuery int
	outQuery   *resProto.RestaurantsTagsPromo
	errQuery   error
}{
	{
		testName: "Get restaurants",
		input:    1,
		out: &rest.AllRestaurantsPromo{
			Restaurant: []rest.Restaurants{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 250,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              5,
				},
			},
			AllTags: []rest.Tag{
				{
					Id:   1,
					Name: "Кофейня",
				},
			},
			AllPromo: []rest.PromoCode{
				{
					Name:         "Free all",
					Description:  "free delivery",
					Img:          "/url/",
					RestaurantId: 1,
				},
			},
		},
		outErr:     "",
		inputQuery: 1,
		outQuery: &resProto.RestaurantsTagsPromo{
			Restaurants: []*resProto.Restaurant{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 250,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              5,
				},
			},
			Tags: []*resProto.Tags{
				{
					Id:   1,
					Name: "Кофейня",
				},
			},
			Promocode: []*resProto.Promocode{
				{
					Name:   "Free all",
					Desc:   "free delivery",
					Img:    "/url/",
					RestId: 1,
				},
			},
			Error: "",
		},
		errQuery: nil,
	},
	{
		testName:   "Error get restaurant",
		input:      1,
		out:        nil,
		outErr:     "text",
		inputQuery: 1,
		outQuery: &resProto.RestaurantsTagsPromo{
			Restaurants: []*resProto.Restaurant{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 250,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              5,
				},
			},
			Tags: []*resProto.Tags{
				{
					Id:   1,
					Name: "Кофейня",
				},
			},
			Promocode: []*resProto.Promocode{
				{
					Name:   "Free all",
					Desc:   "free delivery",
					Img:    "/url/",
					RestId: 1,
				},
			},
			Error: "text",
		},
		errQuery: nil,
	},
	{
		testName:   "Error microservice",
		input:      1,
		out:        nil,
		outErr:     "text",
		inputQuery: 1,
		outQuery:   nil,
		errQuery:   errors.New("text"),
	},
}

func TestAllRestaurants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectRestaurantServiceInterface(ctrl)
	for _, tt := range AllRestaurants {
		m.
			EXPECT().
			AllRestaurants(gomock.Any(), gomock.Any()).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.AllRestaurants()
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				if err == nil {
					require.NotNil(t, err, fmt.Sprintf("Expected: %s\nbut got: nil", tt.outErr))
				}
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var RecommendedRestaurants = []struct {
	testName   string
	input      int
	out        *rest.AllRestaurants
	outErr     string
	inputQuery int
	outQuery   *resProto.RecommendedRestaurants
	errQuery   error
}{
	{
		testName: "Get recommended restaurants",
		input:    1,
		out: &rest.AllRestaurants{
			Restaurant: []rest.Restaurants{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 250,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              5,
				},
			},
			AllTags: []rest.Tag{
				{
					Id:   1,
					Name: "Кофейня",
				},
			},
		},
		outErr:     "",
		inputQuery: 1,
		outQuery: &resProto.RecommendedRestaurants{
			Restaurants: []*resProto.Restaurant{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 250,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              5,
				},
			},
			Tags: []*resProto.Tags{
				{
					Id:   1,
					Name: "Кофейня",
				},
			},
			Error: "",
		},
		errQuery: nil,
	},
	{
		testName:   "Error get recommended restaurants",
		input:      1,
		out:        nil,
		outErr:     "text",
		inputQuery: 1,
		outQuery: &resProto.RecommendedRestaurants{
			Restaurants: []*resProto.Restaurant{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 250,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              5,
				},
			},
			Tags: []*resProto.Tags{
				{
					Id:   1,
					Name: "Кофейня",
				},
			},
			Error: "text",
		},
		errQuery: nil,
	},
	{
		testName:   "Error microservice",
		input:      1,
		out:        nil,
		outErr:     "text",
		inputQuery: 1,
		outQuery: &resProto.RecommendedRestaurants{
			Restaurants: []*resProto.Restaurant{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 250,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              5,
				},
			},
			Tags: []*resProto.Tags{
				{
					Id:   1,
					Name: "Кофейня",
				},
			},
			Error: "",
		},
		errQuery: errors.New("text"),
	},
}

func TestRecommendedRestaurants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectRestaurantServiceInterface(ctrl)
	for _, tt := range RecommendedRestaurants {
		m.
			EXPECT().
			GetRecommendedRestaurants(gomock.Any(), gomock.Any()).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.RecommendedRestaurants()
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				if err == nil {
					require.NotNil(t, err, fmt.Sprintf("Expected: %s\nbut got: nil", tt.outErr))
				}
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var GetRestaurant = []struct {
	testName          string
	inputRestaurantId int
	inputClientId     int
	out               *rest.RestaurantId
	outErr            string
	inputQuery        *resProto.RestaurantId
	outQuery          *resProto.RestaurantInfo
	errQuery          error
}{
	{
		testName:          "Get restaurant",
		inputClientId:     1,
		inputRestaurantId: 1,
		out: &rest.RestaurantId{
			Id:                  1,
			Img:                 "/url/",
			Name:                "KFC",
			CostForFreeDelivery: 250,
			MinDelivery:         15,
			MaxDelivery:         30,
			Rating:              5,
			Favourite:           true,
			Tags: []rest.Tag{
				{
					Id:   1,
					Name: "Кофейня",
				},
			},
			Menu: []rest.Menu{
				{
					Name: "Напиток",
					DishesMenu: []rest.DishesMenu{
						{
							Id:          1,
							Name:        "Кофе",
							Cost:        120,
							Kilocalorie: 360,
							Img:         "/url/",
						},
					},
				},
			},
		},
		outErr:     "",
		inputQuery: &resProto.RestaurantId{Id: 1, IdClient: 1},
		outQuery: &resProto.RestaurantInfo{
			Id:                  1,
			Img:                 "/url/",
			Name:                "KFC",
			CostForFreeDelivery: 250,
			MinDelivery:         15,
			MaxDelivery:         30,
			Rating:              5,
			Favourite:           true,
			Tags: []*resProto.Tags{
				{
					Id:   1,
					Name: "Кофейня",
				},
			},
			Menu: []*resProto.Menu{
				{
					Name: "Напиток",
					Dishes: []*resProto.DishesMenu{
						{
							Id:   1,
							Name: "Кофе",
							Cost: 120,
							Ccal: 360,
							Img:  "/url/",
						},
					},
				},
			},
			Error: "",
		},
		errQuery: nil,
	},
	{
		testName:          "Error get restaurant",
		inputClientId:     1,
		inputRestaurantId: 1,
		out:               nil,
		outErr:            "text",
		inputQuery:        &resProto.RestaurantId{Id: 1, IdClient: 1},
		outQuery: &resProto.RestaurantInfo{
			Id:                  1,
			Img:                 "/url/",
			Name:                "KFC",
			CostForFreeDelivery: 250,
			MinDelivery:         15,
			MaxDelivery:         30,
			Rating:              5,
			Favourite:           true,
			Tags: []*resProto.Tags{
				{
					Id:   1,
					Name: "Кофейня",
				},
			},
			Menu: []*resProto.Menu{
				{
					Name: "Напиток",
					Dishes: []*resProto.DishesMenu{
						{
							Id:   1,
							Name: "Кофе",
							Cost: 120,
							Ccal: 360,
							Img:  "/url/",
						},
					},
				},
			},
			Error: "text",
		},
		errQuery: nil,
	},
	{
		testName:          "Error microservice",
		inputClientId:     1,
		inputRestaurantId: 1,
		out:               nil,
		outErr:            "text",
		inputQuery:        &resProto.RestaurantId{Id: 1, IdClient: 1},
		outQuery: &resProto.RestaurantInfo{
			Id:                  1,
			Img:                 "/url/",
			Name:                "KFC",
			CostForFreeDelivery: 250,
			MinDelivery:         15,
			MaxDelivery:         30,
			Rating:              5,
			Favourite:           true,
			Tags: []*resProto.Tags{
				{
					Id:   1,
					Name: "Кофейня",
				},
			},
			Menu: []*resProto.Menu{
				{
					Name: "Напиток",
					Dishes: []*resProto.DishesMenu{
						{
							Id:   1,
							Name: "Кофе",
							Cost: 120,
							Ccal: 360,
							Img:  "/url/",
						},
					},
				},
			},
			Error: "",
		},
		errQuery: errors.New("text"),
	},
}

func TestGetRestaurant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectRestaurantServiceInterface(ctrl)
	for _, tt := range GetRestaurant {
		m.
			EXPECT().
			GetRestaurant(gomock.Any(), tt.inputQuery).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetRestaurant(tt.inputRestaurantId, tt.inputClientId)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				if err == nil {
					require.NotNil(t, err, fmt.Sprintf("Expected: %s\nbut got: nil", tt.outErr))
				}
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var RestaurantDishes = []struct {
	testName          string
	inputRestaurantId int
	inputClientId     int
	out               *rest.Dishes
	outErr            string
	inputQuery        *resProto.DishInfo
	outQuery          *resProto.Dishes
	errQuery          error
}{
	{
		testName:          "Restaurant dishes",
		inputClientId:     1,
		inputRestaurantId: 1,
		out: &rest.Dishes{
			Id:          1,
			Img:         "/url",
			Title:       "Шоколад",
			Cost:        30,
			Ccal:        500,
			Description: "Очень вкусно",
			Radios: []rest.Radios{
				{
					Title: "Тип шоколада",
					Id:    1,
					Rows: []rest.CheckboxesRows{
						{
							Id:   1,
							Name: "Белый",
						},
						{
							Id:   2,
							Name: "Чёрный",
						},
					},
				},
			},
			Ingredient: []rest.Ingredients{
				{
					Id:    1,
					Cost:  20,
					Title: "Орехи",
				},
			},
		},
		outErr:     "",
		inputQuery: &resProto.DishInfo{DishId: 1, RestaurantId: 1},
		outQuery: &resProto.Dishes{
			Id:          1,
			Img:         "/url",
			Name:        "Шоколад",
			Cost:        30,
			Ccal:        500,
			Description: "Очень вкусно",
			Radios: []*resProto.Radios{
				{
					Name: "Тип шоколада",
					Id:   1,
					Rows: []*resProto.CheckboxesRows{
						{
							Id:   1,
							Name: "Белый",
						},
						{
							Id:   2,
							Name: "Чёрный",
						},
					},
				},
			},
			Ingredients: []*resProto.Ingredients{
				{
					Id:   1,
					Cost: 20,
					Name: "Орехи",
				},
			},
			Error: "",
		},
		errQuery: nil,
	},
	{
		testName:          "Error get dish",
		inputClientId:     1,
		inputRestaurantId: 1,
		out:               nil,
		outErr:            "text",
		inputQuery:        &resProto.DishInfo{DishId: 1, RestaurantId: 1},
		outQuery: &resProto.Dishes{
			Id:          1,
			Img:         "/url",
			Name:        "Шоколад",
			Cost:        30,
			Ccal:        500,
			Description: "Очень вкусно",
			Radios: []*resProto.Radios{
				{
					Name: "Тип шоколада",
					Id:   1,
					Rows: []*resProto.CheckboxesRows{
						{
							Id:   1,
							Name: "Белый",
						},
						{
							Id:   2,
							Name: "Чёрный",
						},
					},
				},
			},
			Ingredients: []*resProto.Ingredients{
				{
					Id:   1,
					Cost: 20,
					Name: "Орехи",
				},
			},
			Error: "text",
		},
		errQuery: nil,
	},
	{
		testName:          "Error microservice",
		inputClientId:     1,
		inputRestaurantId: 1,
		out:               nil,
		outErr:            "text",
		inputQuery:        &resProto.DishInfo{DishId: 1, RestaurantId: 1},
		outQuery: &resProto.Dishes{
			Id:          1,
			Img:         "/url",
			Name:        "Шоколад",
			Cost:        30,
			Ccal:        500,
			Description: "Очень вкусно",
			Radios: []*resProto.Radios{
				{
					Name: "Тип шоколада",
					Id:   1,
					Rows: []*resProto.CheckboxesRows{
						{
							Id:   1,
							Name: "Белый",
						},
						{
							Id:   2,
							Name: "Чёрный",
						},
					},
				},
			},
			Ingredients: []*resProto.Ingredients{
				{
					Id:   1,
					Cost: 20,
					Name: "Орехи",
				},
			},
			Error: "",
		},
		errQuery: errors.New("text"),
	},
}

func TestRestaurantDishes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectRestaurantServiceInterface(ctrl)
	for _, tt := range RestaurantDishes {
		m.
			EXPECT().
			RestaurantDishes(gomock.Any(), tt.inputQuery).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.RestaurantDishes(tt.inputRestaurantId, tt.inputClientId)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				if err == nil {
					require.NotNil(t, err, fmt.Sprintf("Expected: %s\nbut got: nil", tt.outErr))
				}
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var CreateReview = []struct {
	testName       string
	inputNewReview rest.NewReview
	inputClientId  int
	outErr         string
	inputQuery     *resProto.NewReview
	outQuery       *resProto.Error
	errQuery       error
}{
	{
		testName:      "Create review",
		inputClientId: 1,
		inputNewReview: rest.NewReview{
			Restaurant: rest.RestaurantId{
				Id:                  1,
				Img:                 "/url/",
				Name:                "KFC",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              5,
				Favourite:           true,
				Tags: []rest.Tag{
					{
						Id:   1,
						Name: "Кофейня",
					},
				},
				Menu: []rest.Menu{
					{
						Name: "Напиток",
						DishesMenu: []rest.DishesMenu{
							{
								Id:          1,
								Name:        "Кофе",
								Cost:        120,
								Kilocalorie: 360,
								Img:         "/url/",
							},
						},
					},
				},
			},
			Text: "Very cool dishes",
			Rate: 5,
		},
		outErr: "",
		inputQuery: &resProto.NewReview{Id: 1,
			Restaurant: &resProto.RestaurantInfo{
				Id:                  1,
				Img:                 "/url/",
				Name:                "KFC",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              5,
				Favourite:           true,
				Tags: []*resProto.Tags{
					{
						Id:   1,
						Name: "Кофейня",
					},
				},
				Menu: []*resProto.Menu{
					{
						Name: "Напиток",
						Dishes: []*resProto.DishesMenu{
							{
								Id:   1,
								Name: "Кофе",
								Cost: 120,
								Ccal: 360,
								Img:  "/url/",
							},
						},
					},
				},
			},
			Text: "Very cool dishes",
			Rate: 5,
		},
		outQuery: &resProto.Error{},
		errQuery: nil,
	},
	{
		testName:      "Error create review",
		inputClientId: 1,
		inputNewReview: rest.NewReview{
			Restaurant: rest.RestaurantId{
				Id:                  1,
				Img:                 "/url/",
				Name:                "KFC",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              5,
				Favourite:           true,
				Tags: []rest.Tag{
					{
						Id:   1,
						Name: "Кофейня",
					},
				},
				Menu: []rest.Menu{
					{
						Name: "Напиток",
						DishesMenu: []rest.DishesMenu{
							{
								Id:          1,
								Name:        "Кофе",
								Cost:        120,
								Kilocalorie: 360,
								Img:         "/url/",
							},
						},
					},
				},
			},
			Text: "Very cool dishes",
			Rate: 5,
		},
		outErr: "text",
		inputQuery: &resProto.NewReview{Id: 1,
			Restaurant: &resProto.RestaurantInfo{
				Id:                  1,
				Img:                 "/url/",
				Name:                "KFC",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              5,
				Favourite:           true,
				Tags: []*resProto.Tags{
					{
						Id:   1,
						Name: "Кофейня",
					},
				},
				Menu: []*resProto.Menu{
					{
						Name: "Напиток",
						Dishes: []*resProto.DishesMenu{
							{
								Id:   1,
								Name: "Кофе",
								Cost: 120,
								Ccal: 360,
								Img:  "/url/",
							},
						},
					},
				},
			},
			Text: "Very cool dishes",
			Rate: 5,
		},
		outQuery: &resProto.Error{Error: "text"},
		errQuery: nil,
	},
	{
		testName:      "Error microservice",
		inputClientId: 1,
		inputNewReview: rest.NewReview{
			Restaurant: rest.RestaurantId{
				Id:                  1,
				Img:                 "/url/",
				Name:                "KFC",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              5,
				Favourite:           true,
				Tags: []rest.Tag{
					{
						Id:   1,
						Name: "Кофейня",
					},
				},
				Menu: []rest.Menu{
					{
						Name: "Напиток",
						DishesMenu: []rest.DishesMenu{
							{
								Id:          1,
								Name:        "Кофе",
								Cost:        120,
								Kilocalorie: 360,
								Img:         "/url/",
							},
						},
					},
				},
			},
			Text: "Very cool dishes",
			Rate: 5,
		},
		outErr: "text",
		inputQuery: &resProto.NewReview{Id: 1,
			Restaurant: &resProto.RestaurantInfo{
				Id:                  1,
				Img:                 "/url/",
				Name:                "KFC",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              5,
				Favourite:           true,
				Tags: []*resProto.Tags{
					{
						Id:   1,
						Name: "Кофейня",
					},
				},
				Menu: []*resProto.Menu{
					{
						Name: "Напиток",
						Dishes: []*resProto.DishesMenu{
							{
								Id:   1,
								Name: "Кофе",
								Cost: 120,
								Ccal: 360,
								Img:  "/url/",
							},
						},
					},
				},
			},
			Text: "Very cool dishes",
			Rate: 5,
		},
		outQuery: &resProto.Error{},
		errQuery: errors.New("text"),
	},
}

func TestCreateReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectRestaurantServiceInterface(ctrl)
	for _, tt := range CreateReview {
		m.
			EXPECT().
			CreateReview(gomock.Any(), tt.inputQuery).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := testUser.CreateReview(tt.inputClientId, tt.inputNewReview)
			if tt.outErr != "" {
				if err == nil {
					require.NotNil(t, err, fmt.Sprintf("Expected: %s\nbut got: nil", tt.outErr))
				}
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var GetReview = []struct {
	testName          string
	inputRestaurantId int
	inputClientId     int
	out               *rest.ResReview
	outErr            string
	inputQuery        *resProto.RestaurantClientId
	outQuery          *resProto.ResReview
	errQuery          error
}{
	{
		testName:          "Get reviews",
		inputClientId:     1,
		inputRestaurantId: 1,
		out: &rest.ResReview{
			Id:                  1,
			Img:                 "/url/",
			Name:                "KFC",
			CostForFreeDelivery: 250,
			MinDelivery:         15,
			MaxDelivery:         30,
			Rating:              5,
			Status:              true,
			Tags: []rest.Tag{
				{
					Id:   1,
					Name: "Кофейня",
				},
			},
			Reviews: []rest.Review{
				{
					Name: "Иванов Иван",
					Text: "Very good",
					Date: "11.11.2011",
					Time: "11:11",
					Rate: 1,
				},
			},
		},
		outErr:     "",
		inputQuery: &resProto.RestaurantClientId{IdRestaurant: 1, IdClient: 1},
		outQuery: &resProto.ResReview{
			Id:                  1,
			Img:                 "/url/",
			Name:                "KFC",
			CostForFreeDelivery: 250,
			MinDelivery:         15,
			MaxDelivery:         30,
			Rating:              5,
			Status:              true,
			Tags: []*resProto.Tags{
				{
					Id:   1,
					Name: "Кофейня",
				},
			},
			Review: []*resProto.Review{
				{
					Name: "Иванов Иван",
					Text: "Very good",
					Date: "11.11.2011",
					Time: "11:11",
					Rate: 1,
				},
			},
			Error: "",
		},
		errQuery: nil,
	},
	{
		testName:          "Error get reviews",
		inputClientId:     1,
		inputRestaurantId: 1,
		out:               nil,
		outErr:            "text",
		inputQuery:        &resProto.RestaurantClientId{IdRestaurant: 1, IdClient: 1},
		outQuery: &resProto.ResReview{
			Id:                  1,
			Img:                 "/url/",
			Name:                "KFC",
			CostForFreeDelivery: 250,
			MinDelivery:         15,
			MaxDelivery:         30,
			Rating:              5,
			Status:              true,
			Tags: []*resProto.Tags{
				{
					Id:   1,
					Name: "Кофейня",
				},
			},
			Review: []*resProto.Review{
				{
					Name: "Иванов Иван",
					Text: "Very good",
					Date: "11.11.2011",
					Time: "11:11",
					Rate: 1,
				},
			},
			Error: "text",
		},
		errQuery: nil,
	},
	{
		testName:          "Error microservice",
		inputClientId:     1,
		inputRestaurantId: 1,
		out:               nil,
		outErr:            "text",
		inputQuery:        &resProto.RestaurantClientId{IdRestaurant: 1, IdClient: 1},
		outQuery: &resProto.ResReview{
			Id:                  1,
			Img:                 "/url/",
			Name:                "KFC",
			CostForFreeDelivery: 250,
			MinDelivery:         15,
			MaxDelivery:         30,
			Rating:              5,
			Status:              true,
			Tags: []*resProto.Tags{
				{
					Id:   1,
					Name: "Кофейня",
				},
			},
			Review: []*resProto.Review{
				{
					Name: "Иванов Иван",
					Text: "Very good",
					Date: "11.11.2011",
					Time: "11:11",
					Rate: 1,
				},
			},
			Error: "",
		},
		errQuery: errors.New("text"),
	},
}

func TestGetReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectRestaurantServiceInterface(ctrl)
	for _, tt := range GetReview {
		m.
			EXPECT().
			GetReview(gomock.Any(), tt.inputQuery).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetReview(tt.inputRestaurantId, tt.inputClientId)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				if err == nil {
					require.NotNil(t, err, fmt.Sprintf("Expected: %s\nbut got: nil", tt.outErr))
				}
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var SearchRestaurant = []struct {
	testName   string
	input      string
	out        []rest.Restaurants
	outErr     string
	inputQuery *resProto.SearchRestaurantText
	outQuery   *resProto.Restaurants
	errQuery   error
}{
	{
		testName: "Search restaurants",
		input:    "KFC",
		out: []rest.Restaurants{
			{
				Id:                  1,
				Img:                 "/url/",
				Name:                "KFC",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              5.0,
			},
		},
		outErr:     "",
		inputQuery: &resProto.SearchRestaurantText{Text: "KFC"},
		outQuery: &resProto.Restaurants{
			Restaurants: []*resProto.Restaurant{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 250,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              5.0,
				},
			},
			Error: "",
		},
		errQuery: nil,
	},
	{
		testName:   "Error search restaurants",
		input:      "KFC",
		out:        nil,
		outErr:     "text",
		inputQuery: &resProto.SearchRestaurantText{Text: "KFC"},
		outQuery: &resProto.Restaurants{
			Restaurants: []*resProto.Restaurant{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 250,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              5.0,
				},
			},
			Error: "text",
		},
		errQuery: nil,
	},
	{
		testName:   "Error microservice",
		input:      "KFC",
		out:        nil,
		outErr:     "text",
		inputQuery: &resProto.SearchRestaurantText{Text: "KFC"},
		outQuery: &resProto.Restaurants{
			Restaurants: []*resProto.Restaurant{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 250,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              5.0,
				},
			},
			Error: "",
		},
		errQuery: errors.New("text"),
	},
}

func TestSearchRestaurant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectRestaurantServiceInterface(ctrl)
	for _, tt := range SearchRestaurant {
		m.
			EXPECT().
			SearchRestaurant(gomock.Any(), tt.inputQuery).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.SearchRestaurant(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				if err == nil {
					require.NotNil(t, err, fmt.Sprintf("Expected: %s\nbut got: nil", tt.outErr))
				}
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var GetFavoriteRestaurants = []struct {
	testName   string
	input      int
	out        []rest.Restaurants
	outErr     string
	inputQuery *resProto.UserId
	outQuery   *resProto.Restaurants
	errQuery   error
}{
	{
		testName: "Get favorite restaurants",
		input:    1,
		out: []rest.Restaurants{
			{
				Id:                  1,
				Img:                 "/url/",
				Name:                "KFC",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              5.0,
			},
		},
		outErr:     "",
		inputQuery: &resProto.UserId{Id: 1},
		outQuery: &resProto.Restaurants{
			Restaurants: []*resProto.Restaurant{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 250,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              5.0,
				},
			},
			Error: "",
		},
		errQuery: nil,
	},
	{
		testName:   "Error get favorite restaurants",
		input:      1,
		out:        nil,
		outErr:     "text",
		inputQuery: &resProto.UserId{Id: 1},
		outQuery: &resProto.Restaurants{
			Restaurants: []*resProto.Restaurant{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 250,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              5.0,
				},
			},
			Error: "text",
		},
		errQuery: nil,
	},
	{
		testName:   "Error microservice",
		input:      1,
		out:        nil,
		outErr:     "text",
		inputQuery: &resProto.UserId{Id: 1},
		outQuery: &resProto.Restaurants{
			Restaurants: []*resProto.Restaurant{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 250,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              5.0,
				},
			},
			Error: "",
		},
		errQuery: errors.New("text"),
	},
}

func TestGetFavoriteRestaurants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectRestaurantServiceInterface(ctrl)
	for _, tt := range GetFavoriteRestaurants {
		m.
			EXPECT().
			GetFavoriteRestaurants(gomock.Any(), tt.inputQuery).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetFavoriteRestaurants(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				if err == nil {
					require.NotNil(t, err, fmt.Sprintf("Expected: %s\nbut got: nil", tt.outErr))
				}
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var EditRestaurantInFavorite = []struct {
	testName          string
	inputRestaurantId int
	inputClientId     int
	out               bool
	outErr            string
	inputQuery        *resProto.EditRestaurantInFavoriteRequest
	outQuery          *resProto.ResponseEditRestaurantInFavorite
	errQuery          error
}{
	{
		testName:          "Edit favorite status restaurant",
		inputRestaurantId: 1,
		inputClientId:     1,
		out:               true,
		outErr:            "",
		inputQuery: &resProto.EditRestaurantInFavoriteRequest{
			IdClient:     1,
			IdRestaurant: 1,
		},
		outQuery: &resProto.ResponseEditRestaurantInFavorite{
			Status: true,
			Error:  "",
		},
		errQuery: nil,
	},
	{
		testName:          "Error edit favorite status restaurant",
		inputRestaurantId: 1,
		inputClientId:     1,
		out:               false,
		outErr:            "text",
		inputQuery: &resProto.EditRestaurantInFavoriteRequest{
			IdClient:     1,
			IdRestaurant: 1,
		},
		outQuery: &resProto.ResponseEditRestaurantInFavorite{
			Status: true,
			Error:  "text",
		},
		errQuery: nil,
	},
	{
		testName:          "Error microservice",
		inputRestaurantId: 1,
		inputClientId:     1,
		out:               false,
		outErr:            "text",
		inputQuery: &resProto.EditRestaurantInFavoriteRequest{
			IdClient:     1,
			IdRestaurant: 1,
		},
		outQuery: &resProto.ResponseEditRestaurantInFavorite{
			Status: true,
			Error:  "",
		},
		errQuery: errors.New("text"),
	},
}

func TestEditRestaurantInFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectRestaurantServiceInterface(ctrl)
	for _, tt := range EditRestaurantInFavorite {
		m.
			EXPECT().
			EditRestaurantInFavorite(gomock.Any(), tt.inputQuery).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.EditRestaurantInFavorite(tt.inputRestaurantId, tt.inputClientId)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				if err == nil {
					require.NotNil(t, err, fmt.Sprintf("Expected: %s\nbut got: nil", tt.outErr))
				}
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
