package application

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internals/cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/cart/application/mocks"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var GetCart = []struct {
	testName   string
	input      int
	out        *cart.ResponseCartErrors
	outErr     string
	inputQuery int
	outQuery   *cart.ResponseCartErrors
	errQuery   error
}{
	{
		testName: "Get dish",
		input:    1,
		out: &cart.ResponseCartErrors{
			Restaurant: cart.RestaurantIdCastResponse{
				Id:                  1,
				Img:                 "/url/",
				Name:                "Шоколад",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              3.0,
			},
			Dishes: []cart.DishesCartResponse{{
				Id:             1,
				ItemNumber:     0,
				Img:            "/url/",
				Name:           "Орехи",
				Count:          1,
				Cost:           100,
				Kilocalorie:    500,
				Weight:         20,
				Description:    "Очень вкусный шоколад",
				RadiosCart:     nil,
				IngredientCart: nil,
			},
			},
			Cost: cart.CostCartResponse{
				DCost:   250,
				SumCost: 350,
			},
			PromoCode: cart.PromoCode{
				Name:        "Free all",
				Description: "free delivery",
				Code:        "Double Time",
			},
			DishErr: nil,
		},
		outErr:     "",
		inputQuery: 1,
		outQuery: &cart.ResponseCartErrors{
			Restaurant: cart.RestaurantIdCastResponse{
				Id:                  1,
				Img:                 "/url/",
				Name:                "Шоколад",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              3.0,
			},
			Dishes: []cart.DishesCartResponse{{
				Id:             1,
				ItemNumber:     0,
				Img:            "/url/",
				Name:           "Орехи",
				Count:          1,
				Cost:           100,
				Kilocalorie:    500,
				Weight:         20,
				Description:    "Очень вкусный шоколад",
				RadiosCart:     nil,
				IngredientCart: nil,
			},
			},
			Cost: cart.CostCartResponse{
				DCost:   250,
				SumCost: 350,
			},
			PromoCode: cart.PromoCode{
				Name:        "Free all",
				Description: "free delivery",
				Code:        "Double Time",
			},
			DishErr: nil,
		},
		errQuery: nil,
	},
}

func TestGetCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperCartServerInterface(ctrl)
	for _, tt := range GetCart {
		m.
			EXPECT().
			GetCart(tt.inputQuery).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Cart{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetCart(tt.input)
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

var UpdateCart = []struct {
	testName         string
	inputDishes      cart.RequestCartDefault
	inputClient      int
	out              *cart.ResponseCartErrors
	outErr           string
	inputQueryDishes cart.RequestCartDefault
	inputQueryClient int
	outQuery         *cart.ResponseCartErrors
	errQuery         error
}{
	{
		testName: "Update cart",
		inputDishes: cart.RequestCartDefault{
			Restaurant: cart.RestaurantRequest{
				Id: 1,
			},
			Dishes: []cart.DishesRequest{{
				Id:         1,
				ItemNumber: 0,
				Count:      1,
			},
			},
			PromoCode: "Free all",
		},
		inputClient: 1,
		out: &cart.ResponseCartErrors{
			Restaurant: cart.RestaurantIdCastResponse{
				Id:                  1,
				Img:                 "/url/",
				Name:                "Шоколад",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              3.0,
			},
			Dishes: []cart.DishesCartResponse{{
				Id:             1,
				ItemNumber:     0,
				Img:            "/url/",
				Name:           "Орехи",
				Count:          1,
				Cost:           100,
				Kilocalorie:    500,
				Weight:         20,
				Description:    "Очень вкусный шоколад",
				RadiosCart:     nil,
				IngredientCart: nil,
			},
			},
			Cost: cart.CostCartResponse{
				DCost:   250,
				SumCost: 350,
			},
			PromoCode: cart.PromoCode{
				Name:        "Free all",
				Description: "free delivery",
				Code:        "Double Time",
			},
			DishErr: nil,
		},
		outErr: "",
		inputQueryDishes: cart.RequestCartDefault{
			Restaurant: cart.RestaurantRequest{
				Id: 1,
			},
			Dishes: []cart.DishesRequest{{
				Id:         1,
				ItemNumber: 0,
				Count:      1,
			},
			},
			PromoCode: "Free all",
		},
		inputQueryClient: 1,
		outQuery: &cart.ResponseCartErrors{
			Restaurant: cart.RestaurantIdCastResponse{
				Id:                  1,
				Img:                 "/url/",
				Name:                "Шоколад",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              3.0,
			},
			Dishes: []cart.DishesCartResponse{{
				Id:             1,
				ItemNumber:     0,
				Img:            "/url/",
				Name:           "Орехи",
				Count:          1,
				Cost:           100,
				Kilocalorie:    500,
				Weight:         20,
				Description:    "Очень вкусный шоколад",
				RadiosCart:     nil,
				IngredientCart: nil,
			},
			},
			Cost: cart.CostCartResponse{
				DCost:   250,
				SumCost: 350,
			},
			PromoCode: cart.PromoCode{
				Name:        "Free all",
				Description: "free delivery",
				Code:        "Double Time",
			},
			DishErr: nil,
		},
		errQuery: nil,
	},
}

func TestUpdateCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperCartServerInterface(ctrl)
	for _, tt := range UpdateCart {
		m.
			EXPECT().
			UpdateCart(tt.inputQueryDishes, tt.inputQueryClient).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Cart{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.UpdateCart(tt.inputDishes, tt.inputClient)
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

var AddPromoCode = []struct {
	testName         string
	inputPromo       string
	inputClient      int
	inputRest        int
	outErr           string
	inputQueryPromo  string
	inputQueryRest   int
	inputQueryClient int
	errQuery         error
}{
	{
		testName:         "Add promo",
		inputPromo:       "promo",
		inputClient:      1,
		inputRest:        1,
		outErr:           "",
		inputQueryPromo:  "promo",
		inputQueryRest:   1,
		inputQueryClient: 1,
		errQuery:         nil,
	},
}

func TestAddPromoCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperCartServerInterface(ctrl)
	for _, tt := range AddPromoCode {
		m.
			EXPECT().
			AddPromoCode(tt.inputQueryPromo, tt.inputQueryRest, tt.inputQueryClient).
			Return(tt.errQuery)
		testUser := &Cart{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := testUser.AddPromoCode(tt.inputPromo, tt.inputRest, tt.inputClient)
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
