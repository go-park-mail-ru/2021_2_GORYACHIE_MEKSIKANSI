package orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/promocode/orm/mocks"
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type Row struct {
	row    []interface{}
	errRow error
}

func (r *Row) Scan(dest ...interface{}) error {
	if r.errRow != nil {
		return r.errRow
	}
	for i := range dest {
		if r.row[i] == nil {
			dest[i] = nil
			continue
		}
		switch dest[i].(type) {
		case *int:
			*dest[i].(*int) = r.row[i].(int)
		case *string:
			*dest[i].(*string) = r.row[i].(string)
		case *float32:
			*dest[i].(*float32) = float32(r.row[i].(float64))
		case **int32:
			t := int32(r.row[i].(int))
			*dest[i].(**int32) = &t
		case **string:
			t := r.row[i].(string)
			*dest[i].(**string) = &t
		case *time.Time:
			*dest[i].(*time.Time) = r.row[i].(time.Time)
		case *bool:
			*dest[i].(*bool) = r.row[i].(bool)
		default:
			dest[i] = nil
		}
	}
	return nil
}

var GetTypePromoCode = []struct {
	testName                 string
	inputPromoCode           string
	inputRestaurantId        int
	out                      int
	outErr                   string
	outQuery                 Row
	inputQueryPromoCode      string
	inputQueryRestaurantId   int
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		inputPromoCode:           "promo type 1",
		inputRestaurantId:        1,
		inputQueryPromoCode:      "promo type 1",
		inputQueryRestaurantId:   1,
		outQuery:                 Row{row: []interface{}{1}},
		testName:                 "First",
		outErr:                   "",
		out:                      1,
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

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range GetTypePromoCode {
		m.
			EXPECT().
			Begin(gomock.Any()).
			Return(mTx, tt.errBeginTransaction)
		mTx.
			EXPECT().
			Commit(gomock.Any()).
			Return(tt.errCommitTransaction).
			Times(tt.countCommitTransaction)
		mTx.
			EXPECT().
			Rollback(gomock.Any()).
			Return(nil).
			Times(tt.countRollbackTransaction)
		mTx.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT type FROM promocode WHERE code = $1 AND restaurant = $2 AND end_date > NOW()",
				tt.inputQueryPromoCode, tt.inputQueryRestaurantId,
			).
			Return(&tt.outQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetTypePromoCode(tt.inputPromoCode, tt.inputRestaurantId)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
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
	outQuery                 Row
	inputQuery               string
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:                 "First",
		outErr:                   "",
		out:                      true,
		inputName:                "promo",
		inputRestaurant:          1,
		inputQuery:               "promo",
		outQuery:                 Row{row: []interface{}{true}},
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

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range ActiveFreeDelivery {
		m.
			EXPECT().
			Begin(gomock.Any()).
			Return(mTx, tt.errBeginTransaction)
		mTx.
			EXPECT().
			Commit(gomock.Any()).
			Return(tt.errCommitTransaction).
			Times(tt.countCommitTransaction)
		mTx.
			EXPECT().
			Rollback(gomock.Any()).
			Return(nil).
			Times(tt.countRollbackTransaction)
		mTx.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT free_delivery FROM promocode WHERE code = $1 AND restaurant = $2",
				tt.inputQuery,
			).
			Return(&tt.outQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.ActiveFreeDelivery(tt.inputName, tt.inputRestaurant)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
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
	outQuery                 Row
	inputQuery               string
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		inputName:                "promo",
		inputRestaurant:          1,
		inputQuery:               "promo",
		outQuery:                 Row{row: []interface{}{1, 1}},
		testName:                 "First",
		outErr:                   "",
		outCost:                  1,
		outDishId:                1,
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

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range ActiveCostForFreeDish {
		m.
			EXPECT().
			Begin(gomock.Any()).
			Return(mTx, tt.errBeginTransaction)
		mTx.
			EXPECT().
			Commit(gomock.Any()).
			Return(tt.errCommitTransaction).
			Times(tt.countCommitTransaction)
		mTx.
			EXPECT().
			Rollback(gomock.Any()).
			Return(nil).
			Times(tt.countRollbackTransaction)
		mTx.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT cost_for_free_dish, free_dish_id FROM promocode WHERE code = $1 AND restaurant = $2",
				tt.inputQuery,
			).
			Return(&tt.outQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			resultFirst, resultSecond, err := testUser.ActiveCostForFreeDish(tt.inputName, tt.inputRestaurant)
			require.Equal(t, tt.outCost, resultFirst, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outCost, resultFirst))
			require.Equal(t, tt.outDishId, resultSecond, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outDishId, resultSecond))
			if tt.outErr != "" && err != nil {
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
	outQuery                 Row
	inputQuery               string
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:                 "First",
		out:                      0,
		outErr:                   "",
		inputPromoCode:           "promo",
		inputAmount:              1,
		inputRestaurant:          1,
		inputQuery:               "promo",
		outQuery:                 Row{row: []interface{}{1, nil, 50}},
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
	{
		testName:                 "Second",
		out:                      100,
		outErr:                   "",
		inputPromoCode:           "promo",
		inputAmount:              150,
		inputRestaurant:          1,
		inputQuery:               "promo",
		outQuery:                 Row{row: []interface{}{1, nil, 50}},
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
	{
		testName:                 "Third",
		out:                      75,
		outErr:                   "",
		inputPromoCode:           "promo",
		inputAmount:              150,
		inputRestaurant:          1,
		inputQuery:               "promo",
		outQuery:                 Row{row: []interface{}{1, 50, nil}},
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

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range ActiveCostForSale {
		m.
			EXPECT().
			Begin(gomock.Any()).
			Return(mTx, tt.errBeginTransaction)
		mTx.
			EXPECT().
			Commit(gomock.Any()).
			Return(tt.errCommitTransaction).
			Times(tt.countCommitTransaction)
		mTx.
			EXPECT().
			Rollback(gomock.Any()).
			Return(nil).
			Times(tt.countRollbackTransaction)
		mTx.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT cost_for_sale, sale_percent, sale_amount FROM promocode WHERE code = $1 AND restaurant = $2",
				tt.inputQuery,
			).
			Return(&tt.outQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.ActiveCostForSale(tt.inputPromoCode, tt.inputAmount, tt.inputRestaurant)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
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
	outQuery                 Row
	inputQuery               string
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:        "First",
		out:             0,
		outErr:          "",
		inputPromoCode:  "promo",
		inputAmount:     1,
		inputRestaurant: 1,
		inputQuery:      "promo",
		outQuery: Row{row: []interface{}{
			time.Date(2022, 1, 1, 1, 1, 1, 1, time.Local),
			nil,
			50,
		},
		},
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
	{
		testName:        "Second",
		out:             100,
		outErr:          "",
		inputPromoCode:  "promo",
		inputAmount:     150,
		inputRestaurant: 1,
		inputQuery:      "promo",
		outQuery: Row{row: []interface{}{
			time.Date(2022, 1, 1, 1, 1, 1, 1, time.Local),
			nil,
			50,
		},
		},
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
	{
		testName:        "Third",
		out:             75,
		outErr:          "",
		inputPromoCode:  "promo",
		inputAmount:     150,
		inputRestaurant: 1,
		inputQuery:      "promo",
		outQuery: Row{row: []interface{}{
			time.Date(2022, 1, 1, 1, 1, 1, 1, time.Local),
			50,
			nil,
		},
		},
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

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range ActiveTimeForSale {
		m.
			EXPECT().
			Begin(gomock.Any()).
			Return(mTx, tt.errBeginTransaction)
		mTx.
			EXPECT().
			Commit(gomock.Any()).
			Return(tt.errCommitTransaction).
			Times(tt.countCommitTransaction)
		mTx.
			EXPECT().
			Rollback(gomock.Any()).
			Return(nil).
			Times(tt.countRollbackTransaction)
		mTx.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT time_for_sale, sale_in_time_percent, sale_in_time_amount FROM promocode WHERE code = $1 AND restaurant = $2",
				tt.inputQuery,
			).
			Return(&tt.outQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.ActiveTimeForSale(tt.inputPromoCode, tt.inputAmount, tt.inputRestaurant)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
