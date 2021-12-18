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
		testName: "Get restaurant",
		out: &resPkg.RestaurantId{
			Id:                  1,
			Img:                 "/url/",
			Name:                "KFC",
			CostForFreeDelivery: 250,
			MinDelivery:         15,
			MaxDelivery:         30,
			Rating:              3,
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
		outErr: "",
		outQuery: &resPkg.RestaurantId{
			Id:                  1,
			Img:                 "/url/",
			Name:                "KFC",
			CostForFreeDelivery: 250,
			MinDelivery:         15,
			MaxDelivery:         30,
			Rating:              3,
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
