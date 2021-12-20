package service

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart"
	cartProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart/proto"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart/service/mocks"
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var GetCart = []struct {
	testName   string
	input      *cartProto.CartId
	out        *cartProto.ResponseCartErrors
	outErr     string
	inputQuery int
	outQuery   *cart.ResponseCartErrors
	errQuery   error
	countQuery int
}{
	{
		testName: "Get cart",
		input: &cartProto.CartId{
			Id: 1,
		},
		out: &cartProto.ResponseCartErrors{
			Restaurant: &cartProto.RestaurantIdCastResponse{
				Id:                  1,
				Img:                 "/url/",
				Name:                "Шоколад",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              5.0,
			},
			Dishes: []*cartProto.DishesCartResponse{
				{
					Id:          1,
					ItemNumber:  0,
					Img:         "/url/",
					Name:        "Шоколад",
					Count:       1,
					Cost:        100,
					Ccal:        500,
					Weight:      10,
					Description: "Очень вкусно",
					Radios: []*cartProto.RadiosCartResponse{
						{
							Name:     "Вид шоколада",
							RadiosId: 1,
							Id:       1,
						},
					},
					Ingredients: []*cartProto.IngredientCartResponse{
						{
							Name: "Орехи",
							Cost: 100,
							Id:   1,
						},
					},
				},
			},
			Cost: &cartProto.CostCartResponse{
				DeliveryCost: 150,
				SumCost:      250,
			},
			DishesErrors: []*cartProto.CastDishesErrs{
				{
					ItemNumber: 0,
					NameDish:   "Пустота",
					CountAvail: 1,
				},
			},
			PromoCode: &cartProto.PromoCode{},
			Error:     "",
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
				Rating:              5.0,
			},
			Dishes: []cart.DishesCartResponse{
				{
					Id:          1,
					ItemNumber:  0,
					Img:         "/url/",
					Name:        "Шоколад",
					Count:       1,
					Cost:        100,
					Kilocalorie: 500,
					Weight:      10,
					Description: "Очень вкусно",
					RadiosCart: []cart.RadiosCartResponse{
						{
							Name:     "Вид шоколада",
							RadiosId: 1,
							Id:       1,
						},
					},
					IngredientCart: []cart.IngredientCartResponse{
						{
							Name: "Орехи",
							Cost: 100,
							Id:   1,
						},
					},
				},
			},
			Cost: cart.CostCartResponse{
				DCost:   150,
				SumCost: 250,
			},
			DishErr: []cart.CastDishesErrs{
				{
					ItemNumber: 0,
					NameDish:   "Пустота",
					CountAvail: 1,
				},
			},
		},
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName: "Error get cart",
		input: &cartProto.CartId{
			Id: 1,
		},
		out: &cartProto.ResponseCartErrors{
			Restaurant:   nil,
			Dishes:       nil,
			Cost:         nil,
			DishesErrors: nil,
			Error:        "text",
		},
		outErr:     "",
		inputQuery: 1,
		outQuery:   &cart.ResponseCartErrors{},
		errQuery:   errors.New("text"),
		countQuery: 1,
	},
}

func TestGetCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockCartInterface(ctrl)
	for _, tt := range GetCart {
		m.
			EXPECT().
			GetCart(tt.inputQuery).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := CartManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.GetCart(context.Background(), tt.input)
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
	testName           string
	input              *cartProto.RequestCartDefault
	out                *cartProto.ResponseCartErrors
	outErr             string
	inputQueryCart     cart.RequestCartDefault
	inputQueryClientId int
	outQuery           *cart.ResponseCartErrors
	errQuery           error
	countQuery         int
}{
	{
		testName: "Update cart",
		input: &cartProto.RequestCartDefault{
			Restaurant: &cartProto.RestaurantRequest{
				Id: 1,
			},
			Dishes: []*cartProto.DishesRequest{
				{
					Id:         1,
					ItemNumber: 0,
					Count:      1,
					Radios: []*cartProto.RadiosCartRequest{
						{
							RadiosId: 1,
							Id:       1,
						},
					},
					Ingredients: []*cartProto.IngredientsCartRequest{
						{
							Id: 1,
						},
					},
				},
			},
			PromoCode: "",
			ClientId:  1,
		},
		out: &cartProto.ResponseCartErrors{
			Restaurant: &cartProto.RestaurantIdCastResponse{
				Id:                  1,
				Img:                 "/url/",
				Name:                "Шоколад",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              5.0,
			},
			Dishes: []*cartProto.DishesCartResponse{
				{
					Id:          1,
					ItemNumber:  0,
					Img:         "/url/",
					Name:        "Шоколад",
					Count:       1,
					Cost:        100,
					Ccal:        500,
					Weight:      10,
					Description: "Очень вкусно",
					Radios: []*cartProto.RadiosCartResponse{
						{
							Name:     "Вид шоколада",
							RadiosId: 1,
							Id:       1,
						},
					},
					Ingredients: []*cartProto.IngredientCartResponse{
						{
							Name: "Орехи",
							Cost: 100,
							Id:   1,
						},
					},
				},
			},
			Cost: &cartProto.CostCartResponse{
				DeliveryCost: 150,
				SumCost:      250,
			},
			DishesErrors: []*cartProto.CastDishesErrs{
				{
					ItemNumber: 0,
					NameDish:   "Пустота",
					CountAvail: 1,
				},
			},
			PromoCode: &cartProto.PromoCode{},
			Error:     "",
		},
		outErr:             "",
		inputQueryClientId: 1,
		inputQueryCart: cart.RequestCartDefault{
			Restaurant: cart.RestaurantRequest{
				Id: 1,
			},
			Dishes: []cart.DishesRequest{
				{
					Id:         1,
					ItemNumber: 0,
					Count:      1,
					Radios: []cart.RadiosCartRequest{
						{
							RadiosId: 1,
							Id:       1,
						},
					},
					Ingredients: []cart.IngredientsCartRequest{
						{
							Id: 1,
						},
					},
				},
			},
			PromoCode: "",
		},
		outQuery: &cart.ResponseCartErrors{
			Restaurant: cart.RestaurantIdCastResponse{
				Id:                  1,
				Img:                 "/url/",
				Name:                "Шоколад",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              5.0,
			},
			Dishes: []cart.DishesCartResponse{
				{
					Id:          1,
					ItemNumber:  0,
					Img:         "/url/",
					Name:        "Шоколад",
					Count:       1,
					Cost:        100,
					Kilocalorie: 500,
					Weight:      10,
					Description: "Очень вкусно",
					RadiosCart: []cart.RadiosCartResponse{
						{
							Name:     "Вид шоколада",
							RadiosId: 1,
							Id:       1,
						},
					},
					IngredientCart: []cart.IngredientCartResponse{
						{
							Name: "Орехи",
							Cost: 100,
							Id:   1,
						},
					},
				},
			},
			Cost: cart.CostCartResponse{
				DCost:   150,
				SumCost: 250,
			},
			DishErr: []cart.CastDishesErrs{
				{
					ItemNumber: 0,
					NameDish:   "Пустота",
					CountAvail: 1,
				},
			},
		},
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName: "Cart is void",
		input: &cartProto.RequestCartDefault{
			Restaurant: &cartProto.RestaurantRequest{
				Id: 1,
			},
			Dishes: []*cartProto.DishesRequest{
				{
					Id:         1,
					ItemNumber: 0,
					Count:      1,
					Radios: []*cartProto.RadiosCartRequest{
						{
							RadiosId: 1,
							Id:       1,
						},
					},
					Ingredients: []*cartProto.IngredientsCartRequest{
						{
							Id: 1,
						},
					},
				},
			},
			PromoCode: "",
			ClientId:  1,
		},
		out:                &cartProto.ResponseCartErrors{},
		outErr:             "",
		inputQueryClientId: 1,
		inputQueryCart: cart.RequestCartDefault{
			Restaurant: cart.RestaurantRequest{
				Id: 1,
			},
			Dishes: []cart.DishesRequest{
				{
					Id:         1,
					ItemNumber: 0,
					Count:      1,
					Radios: []cart.RadiosCartRequest{
						{
							RadiosId: 1,
							Id:       1,
						},
					},
					Ingredients: []cart.IngredientsCartRequest{
						{
							Id: 1,
						},
					},
				},
			},
			PromoCode: "",
		},
		outQuery:   nil,
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName: "Error update cart",
		input: &cartProto.RequestCartDefault{
			Restaurant: &cartProto.RestaurantRequest{
				Id: 1,
			},
			Dishes: []*cartProto.DishesRequest{
				{
					Id:         1,
					ItemNumber: 0,
					Count:      1,
					Radios: []*cartProto.RadiosCartRequest{
						{
							RadiosId: 1,
							Id:       1,
						},
					},
					Ingredients: []*cartProto.IngredientsCartRequest{
						{
							Id: 1,
						},
					},
				},
			},
			PromoCode: "",
			ClientId:  1,
		},
		out: &cartProto.ResponseCartErrors{
			Restaurant:   nil,
			Dishes:       nil,
			Cost:         nil,
			DishesErrors: nil,
			Error:        "text",
		},
		outErr:             "",
		inputQueryClientId: 1,
		inputQueryCart: cart.RequestCartDefault{
			Restaurant: cart.RestaurantRequest{
				Id: 1,
			},
			Dishes: []cart.DishesRequest{
				{
					Id:         1,
					ItemNumber: 0,
					Count:      1,
					Radios: []cart.RadiosCartRequest{
						{
							RadiosId: 1,
							Id:       1,
						},
					},
					Ingredients: []cart.IngredientsCartRequest{
						{
							Id: 1,
						},
					},
				},
			},
			PromoCode: "",
		},
		outQuery:   &cart.ResponseCartErrors{},
		errQuery:   errors.New("text"),
		countQuery: 1,
	},
}

func TestUpdateCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockCartInterface(ctrl)
	for _, tt := range UpdateCart {
		m.
			EXPECT().
			UpdateCart(tt.inputQueryCart, tt.inputQueryClientId).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := CartManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.UpdateCart(context.Background(), tt.input)
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
