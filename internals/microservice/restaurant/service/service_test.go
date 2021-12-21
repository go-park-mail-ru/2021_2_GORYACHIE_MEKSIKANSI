package service

import (
	res "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant"
	resProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant/proto"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant/service/mocks"
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var AllRestaurants = []struct {
	testName   string
	input      *resProto.Empty
	out        *resProto.RestaurantsTagsPromo
	outErr     string
	outQuery   *res.AllRestaurantsPromo
	errQuery   error
	countQuery int
}{
	{
		testName: "Get restaurants",
		input:    &resProto.Empty{},
		out: &resProto.RestaurantsTagsPromo{
			Restaurants: []*resProto.Restaurant{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 150,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              5.0,
				},
			},
			Tags: []*resProto.Tags{
				{
					Name: "Cafe",
					Id:   1,
				},
			},
			Promocode: []*resProto.Promocode{
				{
					Name:   "promo",
					Desc:   "free delivery",
					Img:    "/url/",
					Code:   "promo",
					RestId: 1,
				},
			},
			Error: "",
		},
		outErr: "",
		outQuery: &res.AllRestaurantsPromo{
			Restaurant: []res.Restaurants{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 150,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              5.0,
				},
			},
			AllTags: []res.Tag{
				{
					Name: "Cafe",
					Id:   1,
				},
			},
			AllPromo: []res.Promocode{
				{
					Name:         "promo",
					Description:  "free delivery",
					Img:          "/url/",
					Code:         "promo",
					RestaurantId: 1,
				},
			},
		},
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName: "Error get restaurants",
		input:    &resProto.Empty{},
		out: &resProto.RestaurantsTagsPromo{
			Error: "text",
		},
		outErr:     "",
		outQuery:   &res.AllRestaurantsPromo{},
		errQuery:   errors.New("text"),
		countQuery: 1,
	},
}

func TestAllRestaurants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockRestaurantApplicationInterface(ctrl)
	for _, tt := range AllRestaurants {
		m.
			EXPECT().
			AllRestaurantsPromo().
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := RestaurantManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.AllRestaurants(context.Background(), tt.input)
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

var GetRecommendedRestaurants = []struct {
	testName   string
	input      *resProto.Empty
	out        *resProto.RecommendedRestaurants
	outErr     string
	outQuery   *res.AllRestaurants
	errQuery   error
	countQuery int
}{
	{
		testName: "Get recommended restaurants",
		input:    &resProto.Empty{},
		out: &resProto.RecommendedRestaurants{
			Restaurants: []*resProto.Restaurant{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 150,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              5.0,
				},
			},
			Tags: []*resProto.Tags{
				{
					Name: "Cafe",
					Id:   1,
				},
			},
			Error: "",
		},
		outErr: "",
		outQuery: &res.AllRestaurants{
			Restaurant: []res.Restaurants{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 150,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              5.0,
				},
			},
			AllTags: []res.Tag{
				{
					Name: "Cafe",
					Id:   1,
				},
			},
		},
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName: "Error get recommended restaurants",
		input:    &resProto.Empty{},
		out: &resProto.RecommendedRestaurants{
			Error: "text",
		},
		outErr:     "",
		outQuery:   &res.AllRestaurants{},
		errQuery:   errors.New("text"),
		countQuery: 1,
	},
}

func TestGetRecommendedRestaurants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockRestaurantApplicationInterface(ctrl)
	for _, tt := range GetRecommendedRestaurants {
		m.
			EXPECT().
			RecommendedRestaurants().
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := RestaurantManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.GetRecommendedRestaurants(context.Background(), tt.input)
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
	testName             string
	input                *resProto.RestaurantId
	out                  *resProto.RestaurantInfo
	outErr               string
	inputQueryRestaurant int
	inputQueryClientId   int
	outQuery             *res.RestaurantId
	errQuery             error
	countQuery           int
}{
	{
		testName: "Get restaurant",
		input: &resProto.RestaurantId{
			Id:       1,
			IdClient: 1,
		},
		out: &resProto.RestaurantInfo{
			Id:                  1,
			Img:                 "/url/",
			Name:                "KFC",
			CostForFreeDelivery: 250,
			MinDelivery:         15,
			MaxDelivery:         30,
			Rating:              3,
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
		outErr:               "",
		inputQueryRestaurant: 1,
		inputQueryClientId:   1,
		outQuery: &res.RestaurantId{
			Id:                  1,
			Img:                 "/url/",
			Name:                "KFC",
			CostForFreeDelivery: 250,
			MinDelivery:         15,
			MaxDelivery:         30,
			Rating:              3,
			Favourite:           true,
			Tags: []res.Tag{
				{
					Id:   1,
					Name: "Кофейня",
				},
			},
			Menu: []res.Menu{
				{
					Name: "Напиток",
					DishesMenu: []res.DishesMenu{
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
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName: "Error get restaurant",
		input: &resProto.RestaurantId{
			Id:       1,
			IdClient: 1,
		},
		out: &resProto.RestaurantInfo{
			Error: "text",
		},
		outErr:               "",
		inputQueryRestaurant: 1,
		inputQueryClientId:   1,
		outQuery:             &res.RestaurantId{},
		errQuery:             errors.New("text"),
		countQuery:           1,
	},
}

func TestGetRestaurant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockRestaurantApplicationInterface(ctrl)
	for _, tt := range GetRestaurant {
		m.
			EXPECT().
			GetRestaurant(tt.inputQueryRestaurant, tt.inputQueryClientId).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := RestaurantManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.GetRestaurant(context.Background(), tt.input)
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
	testName             string
	input                *resProto.DishInfo
	out                  *resProto.Dishes
	outErr               string
	inputQueryRestaurant int
	inputQueryClientId   int
	outQuery             *res.Dishes
	errQuery             error
	countQuery           int
}{
	{
		testName: "Get dish",
		input: &resProto.DishInfo{
			RestaurantId: 1,
			DishId:       1,
		},
		out: &resProto.Dishes{
			Id:          1,
			Img:         "/url/",
			Name:        "Шоколад",
			Cost:        100,
			Ccal:        500,
			Description: "Вкусно",
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
					Name: "Орехи",
					Cost: 20,
				},
			},
			Error: "",
		},
		outErr:               "",
		inputQueryRestaurant: 1,
		inputQueryClientId:   1,
		outQuery: &res.Dishes{
			Id:          1,
			Img:         "/url/",
			Title:       "Шоколад",
			Cost:        100,
			Ccal:        500,
			Description: "Вкусно",
			Radios: []res.Radios{
				{
					Title: "Тип шоколада",
					Id:    1,
					Rows: []res.CheckboxesRows{
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
			Ingredient: []res.Ingredients{
				{
					Id:    1,
					Title: "Орехи",
					Cost:  20,
				},
			},
		},
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName: "Error get dish",
		input: &resProto.DishInfo{
			RestaurantId: 1,
			DishId:       1,
		},
		out: &resProto.Dishes{
			Error: "text",
		},
		outErr:               "",
		inputQueryRestaurant: 1,
		inputQueryClientId:   1,
		outQuery:             &res.Dishes{},
		errQuery:             errors.New("text"),
		countQuery:           1,
	},
}

func TestRestaurantDishes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockRestaurantApplicationInterface(ctrl)
	for _, tt := range RestaurantDishes {
		m.
			EXPECT().
			RestaurantDishes(tt.inputQueryRestaurant, tt.inputQueryClientId).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := RestaurantManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.RestaurantDishes(context.Background(), tt.input)
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
	testName             string
	input                *resProto.NewReview
	out                  *resProto.Error
	outErr               string
	inputQueryRestaurant int
	inputQueryNewReview  res.NewReview
	errQuery             error
	countQuery           int
}{
	{
		testName: "Create review",
		input: &resProto.NewReview{
			Restaurant: &resProto.RestaurantInfo{
				Id:                  1,
				Img:                 "/url/",
				Name:                "KFC",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              3,
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
			Text: "Very good",
			Rate: 5.0,
			Id:   1,
		},
		out: &resProto.Error{
			Error: "",
		},
		outErr:               "",
		inputQueryRestaurant: 1,
		inputQueryNewReview: res.NewReview{
			Restaurant: res.RestaurantId{
				Id: 1,
			},
			Text: "Very good",
			Rate: 5.0,
		},
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName: "Error create review",
		input: &resProto.NewReview{
			Restaurant: &resProto.RestaurantInfo{
				Id:                  1,
				Img:                 "/url/",
				Name:                "KFC",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              3,
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
			Text: "Very good",
			Rate: 5.0,
			Id:   1,
		},
		out: &resProto.Error{
			Error: "text",
		},
		outErr:               "",
		inputQueryRestaurant: 1,
		inputQueryNewReview: res.NewReview{
			Restaurant: res.RestaurantId{
				Id: 1,
			},
			Text: "Very good",
			Rate: 5.0,
		},
		errQuery:   errors.New("text"),
		countQuery: 1,
	},
}

func TestCreateReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockRestaurantApplicationInterface(ctrl)
	for _, tt := range CreateReview {
		m.
			EXPECT().
			CreateReview(tt.inputQueryRestaurant, tt.inputQueryNewReview).
			Return(tt.errQuery).
			Times(tt.countQuery)
		test := RestaurantManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.CreateReview(context.Background(), tt.input)
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

var GetReview = []struct {
	testName             string
	input                *resProto.RestaurantClientId
	out                  *resProto.ResReview
	outErr               string
	inputQueryRestaurant int
	inputQueryClientId   int
	outQuery             *res.ResReview
	errQuery             error
	countQuery           int
}{
	{
		testName: "Get reviews",
		input: &resProto.RestaurantClientId{
			IdRestaurant: 1,
			IdClient:     1,
		},
		out: &resProto.ResReview{
			Id:                  1,
			Img:                 "/url/",
			Name:                "KFC",
			CostForFreeDelivery: 150,
			MinDelivery:         15,
			MaxDelivery:         30,
			Rating:              3.0,
			Tags: []*resProto.Tags{
				{
					Id:   1,
					Name: "Кафе",
				},
			},
			Review: []*resProto.Review{
				{
					Name: "Good user",
					Text: "Restaurant so good",
					Date: "11.11.2011",
					Time: "11:11",
					Rate: 3,
				},
			},
			Status: true,
			Error:  "",
		},
		outErr:               "",
		inputQueryRestaurant: 1,
		inputQueryClientId:   1,
		outQuery: &res.ResReview{
			Id:                  1,
			Img:                 "/url/",
			Name:                "KFC",
			CostForFreeDelivery: 150,
			MinDelivery:         15,
			MaxDelivery:         30,
			Rating:              3.0,
			Tags: []res.Tag{
				{
					Id:   1,
					Name: "Кафе",
				},
			},
			Reviews: []res.Review{
				{
					Name: "Good user",
					Text: "Restaurant so good",
					Date: "11.11.2011",
					Time: "11:11",
					Rate: 3,
				},
			},
			Status: true,
		},
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName: "Error get dish",
		input: &resProto.RestaurantClientId{
			IdRestaurant: 1,
			IdClient:     1,
		},
		out: &resProto.ResReview{
			Error: "text",
		},
		outErr:               "",
		inputQueryRestaurant: 1,
		inputQueryClientId:   1,
		outQuery:             &res.ResReview{},
		errQuery:             errors.New("text"),
		countQuery:           1,
	},
}

func TestGetReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockRestaurantApplicationInterface(ctrl)
	for _, tt := range GetReview {
		m.
			EXPECT().
			GetReview(tt.inputQueryRestaurant, tt.inputQueryClientId).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := RestaurantManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.GetReview(context.Background(), tt.input)
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
	input      *resProto.SearchRestaurantText
	out        *resProto.Restaurants
	outErr     string
	inputQuery string
	outQuery   []res.Restaurants
	errQuery   error
	countQuery int
}{
	{
		testName: "Search restaurants",
		input: &resProto.SearchRestaurantText{
			Text: "Cafe",
		},
		out: &resProto.Restaurants{
			Restaurants: []*resProto.Restaurant{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 150,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              3.0,
				},
			},
			Error: "",
		},
		outErr:     "",
		inputQuery: "Cafe",
		outQuery: []res.Restaurants{
			{
				Id:                  1,
				Img:                 "/url/",
				Name:                "KFC",
				CostForFreeDelivery: 150,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              3.0,
			},
		},
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName: "Error search restaurant",
		input: &resProto.SearchRestaurantText{
			Text: "Cafe",
		},
		out: &resProto.Restaurants{
			Error: "text",
		},
		outErr:     "",
		inputQuery: "Cafe",
		outQuery:   []res.Restaurants{},
		errQuery:   errors.New("text"),
		countQuery: 1,
	},
}

func TestSearchRestaurant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockRestaurantApplicationInterface(ctrl)
	for _, tt := range SearchRestaurant {
		m.
			EXPECT().
			SearchRestaurant(tt.inputQuery).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := RestaurantManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.SearchRestaurant(context.Background(), tt.input)
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
	input      *resProto.UserId
	out        *resProto.Restaurants
	outErr     string
	inputQuery int
	outQuery   []res.Restaurants
	errQuery   error
	countQuery int
}{
	{
		testName: "Get favorite restaurant",
		input: &resProto.UserId{
			Id: 1,
		},
		out: &resProto.Restaurants{
			Restaurants: []*resProto.Restaurant{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 150,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              3.0,
				},
			},
			Error: "",
		},
		outErr:     "",
		inputQuery: 1,
		outQuery: []res.Restaurants{
			{
				Id:                  1,
				Img:                 "/url/",
				Name:                "KFC",
				CostForFreeDelivery: 150,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              3.0,
			},
		},
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName: "Error get favorite restaurant",
		input: &resProto.UserId{
			Id: 1,
		},
		out: &resProto.Restaurants{
			Error: "text",
		},
		outErr:     "",
		inputQuery: 1,
		outQuery:   []res.Restaurants{},
		errQuery:   errors.New("text"),
		countQuery: 1,
	},
}

func TestGetFavoriteRestaurants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockRestaurantApplicationInterface(ctrl)
	for _, tt := range GetFavoriteRestaurants {
		m.
			EXPECT().
			GetFavoriteRestaurants(tt.inputQuery).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := RestaurantManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.GetFavoriteRestaurants(context.Background(), tt.input)
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
	testName               string
	input                  *resProto.EditRestaurantInFavoriteRequest
	out                    *resProto.ResponseEditRestaurantInFavorite
	outErr                 string
	inputQueryRestaurantId int
	inputQueryClientId     int
	outQuery               bool
	errQuery               error
	countQuery             int
}{
	{
		testName: "Edit favorite restaurant",
		input: &resProto.EditRestaurantInFavoriteRequest{
			IdRestaurant: 1,
			IdClient:     1,
		},
		out: &resProto.ResponseEditRestaurantInFavorite{
			Status: true,
			Error:  "",
		},
		outErr:                 "",
		inputQueryRestaurantId: 1,
		inputQueryClientId:     1,
		outQuery:               true,
		errQuery:               nil,
		countQuery:             1,
	},
	{
		testName: "Error edit favorite restaurant",
		input: &resProto.EditRestaurantInFavoriteRequest{
			IdRestaurant: 1,
			IdClient:     1,
		},
		out: &resProto.ResponseEditRestaurantInFavorite{
			Status: false,
			Error:  "text",
		},
		outErr:                 "",
		inputQueryRestaurantId: 1,
		inputQueryClientId:     1,
		outQuery:               false,
		errQuery:               errors.New("text"),
		countQuery:             1,
	},
}

func TestEditRestaurantInFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockRestaurantApplicationInterface(ctrl)
	for _, tt := range EditRestaurantInFavorite {
		m.
			EXPECT().
			EditRestaurantInFavorite(tt.inputQueryRestaurantId, tt.inputQueryClientId).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := RestaurantManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.EditRestaurantInFavorite(context.Background(), tt.input)
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
