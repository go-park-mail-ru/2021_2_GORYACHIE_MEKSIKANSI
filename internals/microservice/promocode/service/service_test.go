package service

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/promocode/orm/mocks"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/promocode/proto"
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var GetTypePromoCode = []struct {
	testName             string
	input                *proto.PromoCodeWithRestaurantId
	out                  *proto.TypePromoCodeResponse
	outErr               string
	inputQueryRestaurant int
	inputQueryPromoCode  string
	outQuery             int
	errQuery             error
	countQuery           int
}{
	{
		testName: "Get type promo code",
		input: &proto.PromoCodeWithRestaurantId{
			Restaurant: 1,
			PromoCode:  "promo",
		},
		out: &proto.TypePromoCodeResponse{
			Type:  1,
			Error: "",
		},
		outErr:               "",
		inputQueryRestaurant: 1,
		inputQueryPromoCode:  "promo",
		outQuery:             1,
		errQuery:             nil,
		countQuery:           1,
	},
	{
		testName: "Error get type",
		input: &proto.PromoCodeWithRestaurantId{
			Restaurant: 1,
			PromoCode:  "promo",
		},
		out: &proto.TypePromoCodeResponse{
			Error: "text",
		},
		outErr:               "",
		inputQueryPromoCode:  "promo",
		inputQueryRestaurant: 1,
		outQuery:             0,
		errQuery:             errors.New("text"),
		countQuery:           1,
	},
}

func TestGetTypePromoCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperPromocodeInterface(ctrl)
	for _, tt := range GetTypePromoCode {
		m.
			EXPECT().
			GetTypePromoCode(tt.inputQueryPromoCode, tt.inputQueryRestaurant).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := PromocodeManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.GetTypePromoCode(context.Background(), tt.input)
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

var ActiveFreeDelivery = []struct {
	testName             string
	input                *proto.PromoCodeWithRestaurantId
	out                  *proto.FreeDeliveryResponse
	outErr               string
	inputQueryRestaurant int
	inputQueryPromoCode  string
	outQuery             bool
	errQuery             error
	countQuery           int
}{
	{
		testName: "Active free delivery",
		input: &proto.PromoCodeWithRestaurantId{
			Restaurant: 1,
			PromoCode:  "promo",
		},
		out: &proto.FreeDeliveryResponse{
			Have:  true,
			Error: "",
		},
		outErr:               "",
		inputQueryRestaurant: 1,
		inputQueryPromoCode:  "promo",
		outQuery:             true,
		errQuery:             nil,
		countQuery:           1,
	},
	{
		testName: "Error active free delivery",
		input: &proto.PromoCodeWithRestaurantId{
			Restaurant: 1,
			PromoCode:  "promo",
		},
		out: &proto.FreeDeliveryResponse{
			Error: "text",
		},
		outErr:               "",
		inputQueryPromoCode:  "promo",
		inputQueryRestaurant: 1,
		outQuery:             false,
		errQuery:             errors.New("text"),
		countQuery:           1,
	},
}

func TestActiveFreeDelivery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperPromocodeInterface(ctrl)
	for _, tt := range ActiveFreeDelivery {
		m.
			EXPECT().
			ActiveFreeDelivery(tt.inputQueryPromoCode, tt.inputQueryRestaurant).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := PromocodeManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.ActiveFreeDelivery(context.Background(), tt.input)
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

var ActiveCostForFreeDish = []struct {
	testName             string
	input                *proto.PromoCodeWithRestaurantId
	out                  *proto.FreeDishResponse
	outErr               string
	inputQueryRestaurant int
	inputQueryPromoCode  string
	outQueryCost         int
	outQueryDish         int
	errQuery             error
	countQuery           int
}{
	{
		testName: "Active free dish",
		input: &proto.PromoCodeWithRestaurantId{
			Restaurant: 1,
			PromoCode:  "promo",
		},
		out: &proto.FreeDishResponse{
			Cost:   100,
			DishId: 1,
			Error:  "",
		},
		outErr:               "",
		inputQueryRestaurant: 1,
		inputQueryPromoCode:  "promo",
		outQueryDish:         1,
		outQueryCost:         100,
		errQuery:             nil,
		countQuery:           1,
	},
	{
		testName: "Error active free dish",
		input: &proto.PromoCodeWithRestaurantId{
			Restaurant: 1,
			PromoCode:  "promo",
		},
		out: &proto.FreeDishResponse{
			Error: "text",
		},
		outErr:               "",
		inputQueryPromoCode:  "promo",
		inputQueryRestaurant: 1,
		outQueryCost:         100,
		outQueryDish:         1,
		errQuery:             errors.New("text"),
		countQuery:           1,
	},
}

func TestActiveCostForFreeDish(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperPromocodeInterface(ctrl)
	for _, tt := range ActiveCostForFreeDish {
		m.
			EXPECT().
			ActiveCostForFreeDish(tt.inputQueryPromoCode, tt.inputQueryRestaurant).
			Return(tt.outQueryCost, tt.outQueryDish, tt.errQuery).
			Times(tt.countQuery)
		test := PromocodeManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.ActiveCostForFreeDish(context.Background(), tt.input)
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
