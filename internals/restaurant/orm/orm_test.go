package orm

import (
	resProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant/proto"
	rest "2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant/orm/mocks"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var GetPromoCodes = []struct {
	testName   string
	input      int
	out        *rest.AllRestaurantsPromo
	outErr     string
	inputQuery int
	outQuery   *resProto.RestaurantsTagsPromo
	errQuery   error
}{
	{
		testName:   "First",
		input:      1,
		out:        &rest.AllRestaurantsPromo{},
		outErr:     "",
		inputQuery: 1,
		outQuery:   &resProto.RestaurantsTagsPromo{},
		errQuery:   nil,
	},
}

func TestGetPromoCodes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectRestaurantServiceInterface(ctrl)
	for _, tt := range GetPromoCodes {
		m.
			EXPECT().
			AllRestaurants(gomock.Any(), gomock.Any()).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.AllRestaurants()
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
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
		testName:   "First",
		input:      1,
		out:        &rest.AllRestaurants{},
		outErr:     "",
		inputQuery: 1,
		outQuery:   &resProto.RecommendedRestaurants{},
		errQuery:   nil,
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
			if tt.outErr != "" && err != nil {
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
		testName:          "First",
		inputClientId:     1,
		inputRestaurantId: 1,
		out:               &rest.RestaurantId{},
		outErr:            "",
		inputQuery:        &resProto.RestaurantId{Id: 1, IdClient: 1},
		outQuery:          &resProto.RestaurantInfo{},
		errQuery:          nil,
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
			if tt.outErr != "" && err != nil {
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
		testName:          "First",
		inputClientId:     1,
		inputRestaurantId: 1,
		out:               &rest.Dishes{},
		outErr:            "",
		inputQuery:        &resProto.DishInfo{DishId: 1, RestaurantId: 1},
		outQuery:          &resProto.Dishes{},
		errQuery:          nil,
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
			if tt.outErr != "" && err != nil {
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
		testName:       "First",
		inputClientId:  1,
		inputNewReview: rest.NewReview{},
		outErr:         "",
		inputQuery:     &resProto.NewReview{Id: 1, Restaurant: &resProto.RestaurantInfo{}},
		outQuery:       &resProto.Error{},
		errQuery:       nil,
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
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
