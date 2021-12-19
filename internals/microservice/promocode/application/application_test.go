package application

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/promocode/application/mocks"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var GetTypePromoCode = []struct {
	testName                 string
	inputPromoCode           string
	inputRestaurantId        int
	out                      int
	outErr                   string
	inputQueryPromoCode      string
	inputQueryRestaurantId   int
	outQuery                 int
	errQuery                 error
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:                 "First",
		outErr:                   "",
		out:                      1,
		inputPromoCode:           "promo free delivery",
		inputRestaurantId:        1,
		inputQueryPromoCode:      "promo free delivery",
		inputQueryRestaurantId:   1,
		outQuery:                 1,
		errQuery:                 nil,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestGetTypePromoCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperPromocodeInterface(ctrl)
	for _, tt := range GetTypePromoCode {
		m.
			EXPECT().
			GetTypePromoCode(tt.inputQueryPromoCode, tt.inputQueryRestaurantId).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Promocode{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetTypePromoCode(tt.inputPromoCode, tt.inputRestaurantId)
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
	testName                 string
	inputName                string
	inputRestaurant          int
	out                      bool
	outErr                   string
	inputQueryName           string
	inputQueryRestaurant     int
	outQuery                 bool
	errQuery                 error
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:                 "First",
		out:                      true,
		outErr:                   "",
		inputName:                "promo",
		inputRestaurant:          1,
		inputQueryName:           "promo",
		inputQueryRestaurant:     1,
		outQuery:                 true,
		errQuery:                 nil,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestActiveFreeDelivery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperPromocodeInterface(ctrl)
	for _, tt := range ActiveFreeDelivery {
		m.
			EXPECT().
			ActiveFreeDelivery(tt.inputQueryName, tt.inputQueryRestaurant).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Promocode{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.ActiveFreeDelivery(tt.inputName, tt.inputRestaurant)
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
	testName                 string
	inputName                string
	inputRestaurant          int
	outCost                  int
	outDishId                int
	outErr                   string
	inputQueryName           string
	inputQueryRestaurant     int
	outQueryCost             int
	outQueryDishId           int
	errQuery                 error
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:                 "First",
		outCost:                  1,
		outDishId:                1,
		outErr:                   "",
		inputName:                "promo",
		inputRestaurant:          1,
		inputQueryName:           "promo",
		inputQueryRestaurant:     1,
		outQueryDishId:           1,
		outQueryCost:             1,
		errQuery:                 nil,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestActiveCostForFreeDish(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperPromocodeInterface(ctrl)
	for _, tt := range ActiveCostForFreeDish {
		m.
			EXPECT().
			ActiveCostForFreeDish(tt.inputQueryName, tt.inputQueryRestaurant).
			Return(tt.outQueryCost, tt.outQueryDishId, tt.errQuery)
		testUser := &Promocode{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			resultFirst, resultSecond, err := testUser.ActiveCostForFreeDish(tt.inputName, tt.inputRestaurant)
			require.Equal(t, tt.outCost, resultFirst, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outCost, resultFirst))
			require.Equal(t, tt.outDishId, resultSecond, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outDishId, resultSecond))
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
	testName                 string
	inputPromoCode           string
	inputAmount              int
	inputRestaurant          int
	out                      int
	outErr                   string
	inputQueryPromoCode      string
	inputQueryAmount         int
	inputQueryRestaurant     int
	outQuery                 int
	errQuery                 error
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:                 "First",
		out:                      1,
		outErr:                   "",
		inputPromoCode:           "promo",
		inputAmount:              1,
		inputRestaurant:          1,
		inputQueryPromoCode:      "promo",
		inputQueryAmount:         1,
		inputQueryRestaurant:     1,
		outQuery:                 1,
		errQuery:                 nil,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestActiveCostForSale(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperPromocodeInterface(ctrl)
	for _, tt := range ActiveCostForSale {
		m.
			EXPECT().
			ActiveCostForSale(tt.inputQueryPromoCode, tt.inputQueryAmount, tt.inputQueryRestaurant).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Promocode{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.ActiveCostForSale(tt.inputPromoCode, tt.inputAmount, tt.inputRestaurant)
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
	testName                 string
	inputPromoCode           string
	inputAmount              int
	inputRestaurant          int
	out                      int
	outErr                   string
	inputQueryPromoCode      string
	inputQueryAmount         int
	inputQueryRestaurant     int
	outQuery                 int
	errQuery                 error
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:                 "First",
		out:                      1,
		outErr:                   "",
		inputPromoCode:           "promo",
		inputAmount:              1,
		inputRestaurant:          1,
		inputQueryPromoCode:      "promo",
		inputQueryAmount:         1,
		inputQueryRestaurant:     1,
		outQuery:                 1,
		errQuery:                 nil,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestActiveTimeForSale(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperPromocodeInterface(ctrl)
	for _, tt := range ActiveTimeForSale {
		m.
			EXPECT().
			ActiveTimeForSale(tt.inputQueryPromoCode, tt.inputQueryAmount, tt.inputQueryRestaurant).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Promocode{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.ActiveTimeForSale(tt.inputPromoCode, tt.inputAmount, tt.inputRestaurant)
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
	testName                 string
	inputPromoCode           string
	inputRestaurantId        int
	inputClient              int
	out                      int
	outErr                   string
	inputQueryPromoCode      string
	inputQueryRestaurantId   int
	inputQueryClient         int
	errQuery                 error
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:                 "Add promo code",
		inputPromoCode:           "promo free delivery",
		inputRestaurantId:        1,
		inputClient:              1,
		outErr:                   "",
		out:                      1,
		inputQueryPromoCode:      "promo free delivery",
		inputQueryRestaurantId:   1,
		inputQueryClient:         1,
		errQuery:                 nil,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestAddPromoCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperPromocodeInterface(ctrl)
	for _, tt := range AddPromoCode {
		m.
			EXPECT().
			AddPromoCode(tt.inputQueryPromoCode, tt.inputQueryRestaurantId, tt.inputQueryClient).
			Return(tt.errQuery)
		testUser := &Promocode{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := testUser.AddPromoCode(tt.inputPromoCode, tt.inputRestaurantId, tt.inputClient)
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
	testName                 string
	input                    int
	out                      string
	outErr                   string
	inputQuery               int
	outQuery                 string
	errQuery                 error
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:                 "Get promo code",
		outErr:                   "",
		out:                      "promo",
		input:                    1,
		inputQuery:               1,
		outQuery:                 "promo",
		errQuery:                 nil,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestGetPromoCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperPromocodeInterface(ctrl)
	for _, tt := range GetPromoCode {
		m.
			EXPECT().
			GetPromoCode(tt.inputQuery).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Promocode{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetPromoCode(tt.input)
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
