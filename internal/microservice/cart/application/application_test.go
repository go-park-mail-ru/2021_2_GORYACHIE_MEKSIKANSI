package application

import (
	cartPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/cart/application/mocks"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var CalculatePriceDelivery = []struct {
	testName            string
	input               int
	out                 int
	outErr              string
	outGetPriceDelivery int
	errGetPriceDelivery error
}{
	{
		testName:            "First",
		input:               1,
		out:                 1,
		outErr:              "",
		outGetPriceDelivery: 1,
		errGetPriceDelivery: nil,
	},
}

func TestCalculatePriceDelivery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperCartInterface(ctrl)
	for _, tt := range CalculatePriceDelivery {
		m.
			EXPECT().
			GetPriceDelivery(gomock.Any()).
			Return(tt.outGetPriceDelivery, tt.errGetPriceDelivery)
		test := Cart{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.CalculatePriceDelivery(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %d\nbut got: %d", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var CalculateCost = []struct {
	testName              string
	inputResult           *cartPkg.ResponseCartErrors
	inputRest             *cartPkg.RestaurantId
	out                   *cartPkg.CostCartResponse
	outErr                string
	inputGetPriceDelivery int
	outGetPriceDelivery   int
	errGetPriceDelivery   error
	countGetPriceDelivery int
}{
	{
		testName: "First",
		inputResult: &cartPkg.ResponseCartErrors{
			Dishes:     []cartPkg.DishesCartResponse{{Cost: 1, Count: 1}},
			Cost:       cartPkg.CostCartResponse{SumCost: 1},
			Restaurant: cartPkg.RestaurantIdCastResponse{CostForFreeDelivery: 5},
		},
		inputRest: &cartPkg.RestaurantId{
			Id:                  1,
			CostForFreeDelivery: 5,
		},
		out:                   &cartPkg.CostCartResponse{DCost: 1, SumCost: 2},
		outErr:                "",
		inputGetPriceDelivery: 1,
		outGetPriceDelivery:   1,
		errGetPriceDelivery:   nil,
		countGetPriceDelivery: 1,
	},
}

func TestCalculateCost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperCartInterface(ctrl)
	for _, tt := range CalculateCost {
		m.
			EXPECT().
			GetPriceDelivery(tt.inputGetPriceDelivery).
			Return(tt.outGetPriceDelivery, tt.errGetPriceDelivery).
			Times(tt.countGetPriceDelivery)
		test := Cart{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.CalculateCost(tt.inputResult, tt.inputRest)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var GetCart = []struct {
	testName              string
	input                 int
	out                   *cartPkg.ResponseCartErrors
	outErr                string
	inputGetCart          int
	outGetCartResult      *cartPkg.ResponseCartErrors
	outGetCartErrorDishes []cartPkg.CastDishesErrs
	errGetCart            error
	inputGetRestaurant    int
	outGetRestaurant      *cartPkg.RestaurantId
	errGetRestaurant      error
	countGetRestaurant    int
	inputGetPriceDelivery int
	outGetPriceDelivery   int
	errGetPriceDelivery   error
	countGetPriceDelivery int
}{
	{
		testName: "First",
		input:    1,
		out: &cartPkg.ResponseCartErrors{
			Restaurant: cartPkg.RestaurantIdCastResponse{
				Id:                  1,
				Img:                 "",
				Name:                "",
				CostForFreeDelivery: 0,
				MinDelivery:         0,
				MaxDelivery:         0,
				Rating:              0,
			},
			Dishes: []cartPkg.DishesCartResponse(nil), Cost: cartPkg.CostCartResponse{
				DCost:   0,
				SumCost: 0,
			},
			DishErr: []cartPkg.CastDishesErrs{},
		},
		outErr:       "",
		inputGetCart: 1,
		outGetCartResult: &cartPkg.ResponseCartErrors{
			Restaurant: cartPkg.RestaurantIdCastResponse{
				Id: 1,
			},
		},
		outGetCartErrorDishes: []cartPkg.CastDishesErrs{},
		errGetCart:            nil,
		inputGetRestaurant:    1,
		outGetRestaurant: &cartPkg.RestaurantId{
			Id: 1,
		},
		errGetRestaurant:      nil,
		countGetRestaurant:    1,
		inputGetPriceDelivery: 1,
		outGetPriceDelivery:   1,
		errGetPriceDelivery:   nil,
		countGetPriceDelivery: 0,
	},
}

func TestApplicationGetCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperCartInterface(ctrl)
	for _, tt := range GetCart {
		m.
			EXPECT().
			GetCart(tt.inputGetCart).
			Return(tt.outGetCartResult, tt.outGetCartErrorDishes, tt.errGetCart)
		m.
			EXPECT().
			GetRestaurant(tt.inputGetRestaurant).
			Return(tt.outGetRestaurant, tt.errGetRestaurant).
			Times(tt.countGetRestaurant)
		m.
			EXPECT().
			GetPriceDelivery(tt.inputGetPriceDelivery).
			Return(tt.outGetPriceDelivery, tt.errGetPriceDelivery).
			Times(tt.countGetPriceDelivery)

		test := Cart{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.GetCart(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var UpdateCart = []struct {
	testName                 string
	inputDishes              cartPkg.RequestCartDefault
	inputId                  int
	out                      *cartPkg.ResponseCartErrors
	outErr                   string
	inputDeleteCart          int
	errDeleteCart            error
	inputUpdateCartDishes    cartPkg.RequestCartDefault
	inputUpdateCartId        int
	outUpdateCartResult      *cartPkg.ResponseCartErrors
	outUpdateCartErrorDishes []cartPkg.CastDishesErrs
	errUpdateCart            error
	countUpdateCart          int
	inputGetRestaurant       int
	outGetRestaurant         *cartPkg.RestaurantId
	errGetRestaurant         error
	countGetRestaurant       int
	inputGetPriceDelivery    int
	outGetPriceDelivery      int
	errGetPriceDelivery      error
	countGetPriceDelivery    int
}{
	{
		testName: "First",
		inputDishes: cartPkg.RequestCartDefault{
			Restaurant: cartPkg.RestaurantRequest{
				Id: 1,
			},
		},
		inputId: 1,
		out: &cartPkg.ResponseCartErrors{
			Restaurant: cartPkg.RestaurantIdCastResponse{
				Id:                  1,
				Img:                 "",
				Name:                "",
				CostForFreeDelivery: 0,
				MinDelivery:         0,
				MaxDelivery:         0,
				Rating:              0,
			},
			Dishes: []cartPkg.DishesCartResponse(nil),
			Cost: cartPkg.CostCartResponse{
				DCost:   0,
				SumCost: 0,
			},
			DishErr: []cartPkg.CastDishesErrs{},
		},
		outErr:          "",
		inputDeleteCart: 1,
		errDeleteCart:   nil,
		inputUpdateCartDishes: cartPkg.RequestCartDefault{
			Restaurant: cartPkg.RestaurantRequest{
				Id: 1,
			},
		},
		inputUpdateCartId: 1,
		outUpdateCartResult: &cartPkg.ResponseCartErrors{
			Restaurant: cartPkg.RestaurantIdCastResponse{
				Id: 1,
			},
		},
		outUpdateCartErrorDishes: []cartPkg.CastDishesErrs{},
		errUpdateCart:            nil,
		inputGetRestaurant:       1,
		outGetRestaurant: &cartPkg.RestaurantId{
			Id: 1,
		},
		errGetRestaurant:      nil,
		countGetRestaurant:    1,
		countUpdateCart:       1,
		inputGetPriceDelivery: 1,
		outGetPriceDelivery:   1,
		errGetPriceDelivery:   nil,
		countGetPriceDelivery: 0,
	},
}

func TestApplicationUpdateCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperCartInterface(ctrl)
	for _, tt := range UpdateCart {
		m.
			EXPECT().
			DeleteCart(tt.inputDeleteCart).
			Return(tt.errDeleteCart)
		m.
			EXPECT().
			UpdateCart(tt.inputUpdateCartDishes, tt.inputUpdateCartId).
			Return(tt.outUpdateCartResult, tt.outUpdateCartErrorDishes, tt.errUpdateCart).
			Times(tt.countUpdateCart)
		m.
			EXPECT().
			GetRestaurant(tt.inputGetRestaurant).
			Return(tt.outGetRestaurant, tt.errGetRestaurant).
			Times(tt.countGetRestaurant)
		m.
			EXPECT().
			GetPriceDelivery(tt.inputGetPriceDelivery).
			Return(tt.outGetPriceDelivery, tt.errGetPriceDelivery).
			Times(tt.countGetPriceDelivery)

		test := Cart{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.UpdateCart(tt.inputDishes, tt.inputId)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
