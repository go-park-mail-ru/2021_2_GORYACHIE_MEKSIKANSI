package service

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/promocode/proto"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/promocode/service/mocks"
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

	m := mocks.NewMockPromocodeApplicationInterface(ctrl)
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

	m := mocks.NewMockPromocodeApplicationInterface(ctrl)
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

	m := mocks.NewMockPromocodeApplicationInterface(ctrl)
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

var ActiveCostForSale = []struct {
	testName             string
	input                *proto.PromoCodeWithAmount
	out                  *proto.NewCostResponse
	outErr               string
	inputQueryRestaurant int
	inputQueryPromoCode  string
	inputQueryAmount     int
	outQuery             int
	errQuery             error
	countQuery           int
}{
	{
		testName: "Active sale",
		input: &proto.PromoCodeWithAmount{
			Amount:     100,
			Restaurant: 1,
			PromoCode:  "promo",
		},
		out: &proto.NewCostResponse{
			Cost:  10,
			Error: "",
		},
		outErr:               "",
		inputQueryRestaurant: 1,
		inputQueryAmount:     100,
		inputQueryPromoCode:  "promo",
		outQuery:             10,
		errQuery:             nil,
		countQuery:           1,
	},
	{
		testName: "Error active sale",
		input: &proto.PromoCodeWithAmount{
			Amount:     100,
			Restaurant: 1,
			PromoCode:  "promo",
		},
		out: &proto.NewCostResponse{
			Error: "text",
		},
		outErr:               "",
		inputQueryPromoCode:  "promo",
		inputQueryRestaurant: 1,
		inputQueryAmount:     100,
		outQuery:             10,
		errQuery:             errors.New("text"),
		countQuery:           1,
	},
}

func TestActiveCostForSale(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockPromocodeApplicationInterface(ctrl)
	for _, tt := range ActiveCostForSale {
		m.
			EXPECT().
			ActiveCostForSale(tt.inputQueryPromoCode, tt.inputQueryAmount, tt.inputQueryRestaurant).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := PromocodeManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.ActiveCostForSale(context.Background(), tt.input)
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

var ActiveTimeForSale = []struct {
	testName             string
	input                *proto.PromoCodeWithAmount
	out                  *proto.NewCostResponse
	outErr               string
	inputQueryRestaurant int
	inputQueryPromoCode  string
	inputQueryAmount     int
	outQuery             int
	errQuery             error
	countQuery           int
}{
	{
		testName: "Active time sale",
		input: &proto.PromoCodeWithAmount{
			Amount:     100,
			Restaurant: 1,
			PromoCode:  "promo",
		},
		out: &proto.NewCostResponse{
			Cost:  10,
			Error: "",
		},
		outErr:               "",
		inputQueryRestaurant: 1,
		inputQueryAmount:     100,
		inputQueryPromoCode:  "promo",
		outQuery:             10,
		errQuery:             nil,
		countQuery:           1,
	},
	{
		testName: "Error active time sale",
		input: &proto.PromoCodeWithAmount{
			Amount:     100,
			Restaurant: 1,
			PromoCode:  "promo",
		},
		out: &proto.NewCostResponse{
			Error: "text",
		},
		outErr:               "",
		inputQueryPromoCode:  "promo",
		inputQueryRestaurant: 1,
		inputQueryAmount:     100,
		outQuery:             10,
		errQuery:             errors.New("text"),
		countQuery:           1,
	},
}

func TestActiveTimeForSale(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockPromocodeApplicationInterface(ctrl)
	for _, tt := range ActiveTimeForSale {
		m.
			EXPECT().
			ActiveTimeForSale(tt.inputQueryPromoCode, tt.inputQueryAmount, tt.inputQueryRestaurant).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := PromocodeManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.ActiveTimeForSale(context.Background(), tt.input)
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
	testName             string
	input                *proto.PromoCodeWithRestaurantIdAndClient
	out                  *proto.Error
	outErr               string
	inputQueryRestaurant int
	inputQueryPromoCode  string
	inputQueryClient     int
	errQuery             error
	countQuery           int
}{
	{
		testName: "Add promo code",
		input: &proto.PromoCodeWithRestaurantIdAndClient{
			Client:     1,
			Restaurant: 1,
			PromoCode:  "promo",
		},
		out: &proto.Error{
			Error: "",
		},
		outErr:               "",
		inputQueryRestaurant: 1,
		inputQueryClient:     1,
		inputQueryPromoCode:  "promo",
		errQuery:             nil,
		countQuery:           1,
	},
	{
		testName: "Error active time sale",
		input: &proto.PromoCodeWithRestaurantIdAndClient{
			Client:     1,
			Restaurant: 1,
			PromoCode:  "promo",
		},
		out: &proto.Error{
			Error: "text",
		},
		outErr:               "",
		inputQueryPromoCode:  "promo",
		inputQueryRestaurant: 1,
		inputQueryClient:     1,
		errQuery:             errors.New("text"),
		countQuery:           1,
	},
}

func TestAddPromoCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockPromocodeApplicationInterface(ctrl)
	for _, tt := range AddPromoCode {
		m.
			EXPECT().
			AddPromoCode(tt.inputQueryPromoCode, tt.inputQueryRestaurant, tt.inputQueryClient).
			Return(tt.errQuery).
			Times(tt.countQuery)
		test := PromocodeManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.AddPromoCode(context.Background(), tt.input)
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

var GetPromoCode = []struct {
	testName         string
	input            *proto.ClientId
	out              *proto.PromoCodeText
	outErr           string
	inputQueryClient int
	outQuery         string
	errQuery         error
	countQuery       int
}{
	{
		testName: "Get promo code",
		input: &proto.ClientId{
			ClientId: 1,
		},
		out: &proto.PromoCodeText{
			PromoCodeText: "promo",
			Error:         "",
		},
		outErr:           "",
		inputQueryClient: 1,
		outQuery:         "promo",
		errQuery:         nil,
		countQuery:       1,
	},
	{
		testName: "Error active time sale",
		input: &proto.ClientId{
			ClientId: 1,
		},
		out: &proto.PromoCodeText{
			Error: "text",
		},
		outErr:           "",
		inputQueryClient: 1,
		outQuery:         "",
		errQuery:         errors.New("text"),
		countQuery:       1,
	},
}

func TestGetPromoCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockPromocodeApplicationInterface(ctrl)
	for _, tt := range GetPromoCode {
		m.
			EXPECT().
			GetPromoCode(tt.inputQueryClient).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		test := PromocodeManager{Application: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.GetPromoCode(context.Background(), tt.input)
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
