package application

import (
	resPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant/application/mocks"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var AllRestaurants = []struct {
	testName string
	out      *resPkg.AllRestaurantsPromo
	outErr   string
	outQuery *resPkg.AllRestaurantsPromo
	errQuery error
}{
	{
		testName: "Get restaurants",
		out: &resPkg.AllRestaurantsPromo{
			Restaurant: []resPkg.Restaurants{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 250,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              3,
				},
			},
			AllTags: []resPkg.Tag{{
				Id:   1,
				Name: "Кофейня",
			},
			},
			AllPromo: []resPkg.PromoCode{
				{
					Name:         "Free all",
					Description:  "free delivery",
					Img:          "/url/",
					RestaurantId: 1,
				},
			},
		},
		outErr: "",
		outQuery: &resPkg.AllRestaurantsPromo{
			Restaurant: []resPkg.Restaurants{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 250,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              3,
				},
			},
			AllTags: []resPkg.Tag{{
				Id:   1,
				Name: "Кофейня",
			},
			},
			AllPromo: []resPkg.PromoCode{
				{
					Name:         "Free all",
					Description:  "free delivery",
					Img:          "/url/",
					RestaurantId: 1,
				},
			},
		},
		errQuery: nil,
	},
}

func TestAllRestaurants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurantServerInterface(ctrl)
	for _, tt := range AllRestaurants {
		m.
			EXPECT().
			AllRestaurants().
			Return(tt.outQuery, tt.errQuery)
		testUser := &Restaurant{DB: m}
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
	testName string
	out      *resPkg.AllRestaurants
	outErr   string
	outQuery *resPkg.AllRestaurants
	errQuery error
}{
	{
		testName: "Get recommended restaurant",
		out: &resPkg.AllRestaurants{
			Restaurant: []resPkg.Restaurants{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 250,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              3,
				},
			},
			AllTags: []resPkg.Tag{
				{
					Id:   1,
					Name: "Кофейня",
				},
			},
		},
		outErr: "",
		outQuery: &resPkg.AllRestaurants{
			Restaurant: []resPkg.Restaurants{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "KFC",
					CostForFreeDelivery: 250,
					MinDelivery:         15,
					MaxDelivery:         30,
					Rating:              3,
				},
			},
			AllTags: []resPkg.Tag{
				{
					Id:   1,
					Name: "Кофейня",
				},
			},
		},
		errQuery: nil,
	},
}

func TestRecommendedRestaurants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurantServerInterface(ctrl)
	for _, tt := range RecommendedRestaurants {
		m.
			EXPECT().
			RecommendedRestaurants().
			Return(tt.outQuery, tt.errQuery)
		testUser := &Restaurant{DB: m}
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
	testName             string
	inputRestaurant      int
	inputClientId        int
	out                  *resPkg.RestaurantId
	outErr               string
	inputQueryRestaurant int
	inputQueryClientId   int
	outQuery             *resPkg.RestaurantId
	errQuery             error
}{
	{
		testName:        "Get restaurant",
		inputRestaurant: 1,
		inputClientId:   1,
		out: &resPkg.RestaurantId{
			Id:                  1,
			Img:                 "/url/",
			Name:                "KFC",
			CostForFreeDelivery: 250,
			MinDelivery:         15,
			MaxDelivery:         30,
			Rating:              3,
			Favourite:           true,
			Tags: []resPkg.Tag{
				{
					Id:   1,
					Name: "Кофейня",
				},
			},
			Menu: []resPkg.Menu{
				{
					Name: "Напиток",
					DishesMenu: []resPkg.DishesMenu{
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
		outErr:               "",
		inputQueryClientId:   1,
		inputQueryRestaurant: 1,
		outQuery: &resPkg.RestaurantId{
			Id:                  1,
			Img:                 "/url/",
			Name:                "KFC",
			CostForFreeDelivery: 250,
			MinDelivery:         15,
			MaxDelivery:         30,
			Rating:              3,
			Favourite:           true,
			Tags: []resPkg.Tag{
				{
					Id:   1,
					Name: "Кофейня",
				},
			},
			Menu: []resPkg.Menu{
				{
					Name: "Напиток",
					DishesMenu: []resPkg.DishesMenu{
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
		errQuery: nil,
	},
}

func TestGetRestaurant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurantServerInterface(ctrl)
	for _, tt := range GetRestaurant {
		m.
			EXPECT().
			GetRestaurant(tt.inputQueryRestaurant, tt.inputQueryClientId).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Restaurant{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetRestaurant(tt.inputRestaurant, tt.inputClientId)
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
	inputRestaurant      int
	inputDishId          int
	out                  *resPkg.Dishes
	outErr               string
	inputQueryRestaurant int
	inputQueryDishId     int
	outQuery             *resPkg.Dishes
	errQuery             error
}{
	{
		testName:        "Get dish",
		inputRestaurant: 1,
		inputDishId:     1,
		out: &resPkg.Dishes{
			Id:          1,
			Img:         "/url/",
			Title:       "Шоколад",
			Cost:        100,
			Ccal:        500,
			Description: "Вкусно",
			Radios:      nil,
			Ingredient:  nil,
		},
		outErr:               "",
		inputQueryRestaurant: 1,
		inputQueryDishId:     1,
		outQuery: &resPkg.Dishes{
			Id:          1,
			Img:         "/url/",
			Title:       "Шоколад",
			Cost:        100,
			Ccal:        500,
			Description: "Вкусно",
			Radios:      nil,
			Ingredient:  nil,
		},
		errQuery: nil,
	},
}

func TestRestaurantDishes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurantServerInterface(ctrl)
	for _, tt := range RestaurantDishes {
		m.
			EXPECT().
			RestaurantDishes(tt.inputQueryRestaurant, tt.inputQueryDishId).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Restaurant{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.RestaurantDishes(tt.inputRestaurant, tt.inputDishId)
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
	testName         string
	inputId          int
	inputReview      resPkg.NewReview
	outErr           string
	inputQueryId     int
	inputQueryReview resPkg.NewReview
	errQuery         error
}{
	{
		testName: "Create review",
		inputId:  1,
		inputReview: resPkg.NewReview{
			Restaurant: resPkg.RestaurantId{
				Id:                  1,
				Img:                 "/url/",
				Name:                "KFC",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              3,
				Favourite:           true,
				Tags: []resPkg.Tag{
					{
						Id:   1,
						Name: "Кофейня",
					},
				},
				Menu: []resPkg.Menu{
					{
						Name: "Напиток",
						DishesMenu: []resPkg.DishesMenu{
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
			Text: "Very cool restaurant",
			Rate: 5,
		},
		outErr:       "",
		inputQueryId: 1,
		inputQueryReview: resPkg.NewReview{
			Restaurant: resPkg.RestaurantId{
				Id:                  1,
				Img:                 "/url/",
				Name:                "KFC",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              3,
				Favourite:           true,
				Tags: []resPkg.Tag{
					{
						Id:   1,
						Name: "Кофейня",
					},
				},
				Menu: []resPkg.Menu{
					{
						Name: "Напиток",
						DishesMenu: []resPkg.DishesMenu{
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
			Text: "Very cool restaurant",
			Rate: 5,
		},
		errQuery: nil,
	},
}

func TestCreateReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurantServerInterface(ctrl)
	for _, tt := range CreateReview {
		m.
			EXPECT().
			CreateReview(tt.inputQueryId, tt.inputQueryReview).
			Return(tt.errQuery)
		testUser := &Restaurant{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := testUser.CreateReview(tt.inputId, tt.inputReview)
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
	inputRestaurant      int
	inputClientId        int
	out                  *resPkg.ResReview
	outErr               string
	inputQueryRestaurant int
	inputQueryClientId   int
	outQuery             *resPkg.ResReview
	errQuery             error
}{
	{
		testName:        "Get review",
		inputRestaurant: 1,
		inputClientId:   1,
		out: &resPkg.ResReview{
			Id:                  1,
			Img:                 "/url/",
			Name:                "KFC",
			CostForFreeDelivery: 150,
			MinDelivery:         15,
			MaxDelivery:         30,
			Rating:              3.0,
			Tags:                nil,
			Reviews: []resPkg.Review{
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
		outErr:               "",
		inputQueryRestaurant: 1,
		inputQueryClientId:   1,
		outQuery: &resPkg.ResReview{
			Id:                  1,
			Img:                 "/url/",
			Name:                "KFC",
			CostForFreeDelivery: 150,
			MinDelivery:         15,
			MaxDelivery:         30,
			Rating:              3.0,
			Tags:                nil,
			Reviews: []resPkg.Review{
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
		errQuery: nil,
	},
}

func TestGetReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurantServerInterface(ctrl)
	for _, tt := range GetReview {
		m.
			EXPECT().
			GetReview(tt.inputQueryRestaurant, tt.inputQueryClientId).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Restaurant{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetReview(tt.inputRestaurant, tt.inputClientId)
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
	testName         string
	inputSearch      string
	out              []resPkg.Restaurants
	outErr           string
	inputQuerySearch string
	outQuery         []resPkg.Restaurants
	errQuery         error
}{
	{
		testName:    "Search restaurant",
		inputSearch: "cafe",
		out: []resPkg.Restaurants{
			{
				Id:                  1,
				Img:                 "/url/",
				Name:                "Cafe",
				CostForFreeDelivery: 125,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              5,
			},
		},
		outErr:           "",
		inputQuerySearch: "cafe",
		outQuery: []resPkg.Restaurants{
			{
				Id:                  1,
				Img:                 "/url/",
				Name:                "Cafe",
				CostForFreeDelivery: 125,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              5,
			},
		},
		errQuery: nil,
	},
}

func TestSearchRestaurant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurantServerInterface(ctrl)
	for _, tt := range SearchRestaurant {
		m.
			EXPECT().
			SearchRestaurant(tt.inputQuerySearch).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Restaurant{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.SearchRestaurant(tt.inputSearch)
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
	testName     string
	inputId      int
	out          []resPkg.Restaurants
	outErr       string
	inputQueryId int
	outQuery     []resPkg.Restaurants
	errQuery     error
}{
	{
		testName: "Get favorite restaurants",
		inputId:  1,
		out: []resPkg.Restaurants{
			{
				Id:                  1,
				Img:                 "/url/",
				Name:                "Cafe",
				CostForFreeDelivery: 125,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              5,
			},
		},
		outErr:       "",
		inputQueryId: 1,
		outQuery: []resPkg.Restaurants{
			{
				Id:                  1,
				Img:                 "/url/",
				Name:                "Cafe",
				CostForFreeDelivery: 125,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              5,
			},
		},
		errQuery: nil,
	},
}

func TestGetFavoriteRestaurants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurantServerInterface(ctrl)
	for _, tt := range GetFavoriteRestaurants {
		m.
			EXPECT().
			GetFavoriteRestaurants(tt.inputQueryId).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Restaurant{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetFavoriteRestaurants(tt.inputId)
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
	testName             string
	inputRestaurant      int
	inputClientId        int
	out                  bool
	outErr               string
	inputQueryRestaurant int
	inputQueryClientId   int
	outQuery             bool
	errQuery             error
}{
	{
		testName:             "Edit restaurant",
		inputRestaurant:      1,
		inputClientId:        1,
		out:                  true,
		outErr:               "",
		inputQueryRestaurant: 1,
		inputQueryClientId:   1,
		outQuery:             true,
		errQuery:             nil,
	},
}

func TestEditRestaurantInFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurantServerInterface(ctrl)
	for _, tt := range EditRestaurantInFavorite {
		m.
			EXPECT().
			EditRestaurantInFavorite(tt.inputQueryRestaurant, tt.inputQueryClientId).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Restaurant{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.EditRestaurantInFavorite(tt.inputRestaurant, tt.inputClientId)
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
