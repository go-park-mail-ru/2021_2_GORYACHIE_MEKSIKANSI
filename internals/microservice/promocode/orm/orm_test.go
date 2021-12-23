package orm

import (
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/promocode/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/promocode/orm/mocks"
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v4"
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
		case **string:
			t := r.row[i].(string)
			*dest[i].(**string) = &t
		case *float32:
			*dest[i].(*float32) = float32(r.row[i].(float64))
		case **int32:
			t := int32(r.row[i].(int))
			*dest[i].(**int32) = &t
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
				"SELECT type FROM public.promocode WHERE code = $1 AND restaurant = $2 AND end_date > NOW()",
				tt.inputQueryPromoCode, tt.inputQueryRestaurantId,
			).
			Return(&tt.outQuery)
		testUser := &Wrapper{Conn: m}
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
				"SELECT free_delivery FROM public.promocode WHERE code = $1 AND restaurant = $2",
				tt.inputQuery,
			).
			Return(&tt.outQuery)
		testUser := &Wrapper{Conn: m}
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
				"SELECT cost_for_free_dish, free_dish_id FROM public.promocode WHERE code = $1 AND restaurant = $2",
				tt.inputQuery,
			).
			Return(&tt.outQuery)
		testUser := &Wrapper{Conn: m}
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
				"SELECT cost_for_sale, sale_percent, sale_amount FROM public.promocode WHERE code = $1 AND restaurant = $2",
				tt.inputQuery,
			).
			Return(&tt.outQuery)
		testUser := &Wrapper{Conn: m}
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
	inputTime                time.Time
	out                      int
	outErr                   string
	inputQueryCode           string
	inputQueryRestaurant     int
	outQuery                 Row
	countQuery               int
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:             "Active time sale percent",
		out:                  50,
		outErr:               "",
		inputPromoCode:       "promo",
		inputAmount:          100,
		inputRestaurant:      1,
		inputTime:            time.Date(0, 0, 0, 18, 0, 0, 0, time.Local),
		inputQueryCode:       "promo",
		inputQueryRestaurant: 1,
		outQuery: Row{row: []interface{}{
			time.Date(0, 0, 0, 17, 0, 0, 0, time.Local),
			time.Date(0, 0, 0, 21, 0, 0, 0, time.Local),
			50,
			nil,
		},
		},
		countQuery:               1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
	{
		testName:             "Non active time",
		out:                  100,
		outErr:               "",
		inputPromoCode:       "promo",
		inputAmount:          100,
		inputRestaurant:      1,
		inputTime:            time.Date(0, 0, 0, 16, 0, 0, 0, time.Local),
		inputQueryCode:       "promo",
		inputQueryRestaurant: 1,
		outQuery: Row{row: []interface{}{
			time.Date(0, 0, 0, 17, 0, 0, 0, time.Local),
			time.Date(0, 0, 0, 21, 0, 0, 0, time.Local),
			nil,
			50,
		},
		},
		countQuery:               1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
	{
		testName:             "Active time sale amount",
		out:                  100,
		outErr:               "",
		inputPromoCode:       "promo",
		inputAmount:          150,
		inputTime:            time.Date(0, 0, 0, 18, 0, 0, 0, time.Local),
		inputRestaurant:      1,
		inputQueryCode:       "promo",
		inputQueryRestaurant: 1,
		outQuery: Row{
			row: []interface{}{
				time.Date(0, 0, 0, 17, 0, 0, 0, time.Local),
				time.Date(0, 0, 0, 21, 0, 0, 0, time.Local),
				nil,
				50,
			},
		},
		countQuery:               1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
	{
		testName:             "Active time sale amount more user amount",
		out:                  0,
		outErr:               "",
		inputPromoCode:       "promo",
		inputAmount:          10,
		inputTime:            time.Date(0, 0, 0, 18, 0, 0, 0, time.Local),
		inputRestaurant:      1,
		inputQueryCode:       "promo",
		inputQueryRestaurant: 1,
		outQuery: Row{
			row: []interface{}{
				time.Date(0, 0, 0, 17, 0, 0, 0, time.Local),
				time.Date(0, 0, 0, 21, 0, 0, 0, time.Local),
				nil,
				50,
			},
		},
		countQuery:               1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
	{
		testName:             "Error begin transaction",
		out:                  0,
		outErr:               errPkg.PActiveTimeForSaleTransactionNotCreate,
		inputPromoCode:       "promo",
		inputAmount:          10,
		inputTime:            time.Date(0, 0, 0, 18, 0, 0, 0, time.Local),
		inputRestaurant:      1,
		inputQueryCode:       "promo",
		inputQueryRestaurant: 1,
		outQuery: Row{
			row: []interface{}{
				time.Date(0, 0, 0, 17, 0, 0, 0, time.Local),
				time.Date(0, 0, 0, 21, 0, 0, 0, time.Local),
				nil,
				50,
			},
		},
		countQuery:               0,
		errBeginTransaction:      errors.New("text"),
		errCommitTransaction:     nil,
		countCommitTransaction:   0,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 0,
	},
	{
		testName:             "Error query",
		out:                  0,
		outErr:               errPkg.PActiveTimeForSaleRestaurantsNotSelect,
		inputPromoCode:       "promo",
		inputAmount:          10,
		inputTime:            time.Date(0, 0, 0, 18, 0, 0, 0, time.Local),
		inputRestaurant:      1,
		inputQueryCode:       "promo",
		inputQueryRestaurant: 1,
		outQuery: Row{
			errRow: errors.New("text"),
			row: []interface{}{
				time.Date(0, 0, 0, 17, 0, 0, 0, time.Local),
				time.Date(0, 0, 0, 21, 0, 0, 0, time.Local),
				nil,
				50,
			},
		},
		countQuery:               1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   0,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
	{
		testName:             "Promo code not found",
		out:                  0,
		outErr:               errPkg.PActiveTimeForSaleRestaurantsNotFound,
		inputPromoCode:       "promo",
		inputAmount:          10,
		inputTime:            time.Date(0, 0, 0, 18, 0, 0, 0, time.Local),
		inputRestaurant:      1,
		inputQueryCode:       "promo",
		inputQueryRestaurant: 1,
		outQuery: Row{
			errRow: pgx.ErrNoRows,
			row: []interface{}{
				time.Date(0, 0, 0, 17, 0, 0, 0, time.Local),
				time.Date(0, 0, 0, 21, 0, 0, 0, time.Local),
				nil,
				50,
			},
		},
		countQuery:               1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   0,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
	{
		testName:             "Error commit",
		out:                  0,
		outErr:               errPkg.PActiveTimeForSaleNotCommit,
		inputPromoCode:       "promo",
		inputAmount:          10,
		inputTime:            time.Date(0, 0, 0, 18, 0, 0, 0, time.Local),
		inputRestaurant:      1,
		inputQueryCode:       "promo",
		inputQueryRestaurant: 1,
		outQuery: Row{
			row: []interface{}{
				time.Date(0, 0, 0, 17, 0, 0, 0, time.Local),
				time.Date(0, 0, 0, 21, 0, 0, 0, time.Local),
				nil,
				50,
			},
		},
		countQuery:               1,
		errBeginTransaction:      nil,
		errCommitTransaction:     errors.New("text"),
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
				"SELECT time_for_sale_start, time_for_sale_finish, sale_in_time_percent, sale_in_time_amount FROM public.promocode WHERE code = $1 AND restaurant = $2",
				tt.inputQueryCode, tt.inputQueryRestaurant,
			).
			Return(&tt.outQuery).
			Times(tt.countQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.ActiveTimeForSale(tt.inputPromoCode, tt.inputAmount, tt.inputRestaurant, tt.inputTime)
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
	outErr                   string
	inputQueryClient         int
	inputQueryPromoCode      string
	inputQueryRestaurant     int
	inputPromoCode           string
	inputRestaurantId        int
	inputClientId            int
	errQuery                 error
	countQuery               int
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:                 "First",
		outErr:                   "",
		inputQueryClient:         1,
		inputQueryPromoCode:      "promo",
		inputQueryRestaurant:     1,
		inputPromoCode:           "promo",
		inputRestaurantId:        1,
		inputClientId:            1,
		errQuery:                 nil,
		countQuery:               1,
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

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range AddPromoCode {
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
			Exec(context.Background(),
				"INSERT INTO public.cart_user (client_id, promo_code, restaurant) VALUES ($1, $2, $3) ON CONFLICT (client_id) DO UPDATE SET promo_code = $2 WHERE cart_user.client_id =  $1",
				tt.inputQueryClient, tt.inputQueryPromoCode, tt.inputQueryRestaurant,
			).
			Return(nil, tt.errQuery).
			Times(tt.countQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := testUser.AddPromoCode(tt.inputPromoCode, tt.inputRestaurantId, tt.inputClientId)
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
	errBeginTransaction      error
	inputQuery               int
	outQuery                 Row
	countQuery               int
	errCommitTransaction     error
	countCommitTransaction   int
	countRollbackTransaction int
}{
	{
		testName:                 "First",
		input:                    1,
		out:                      "promo",
		outErr:                   "",
		errBeginTransaction:      nil,
		inputQuery:               1,
		outQuery:                 Row{row: []interface{}{"promo"}},
		countQuery:               1,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		countRollbackTransaction: 1,
	},
}

func TestGetPromoCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range GetPromoCode {
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
				"SELECT promo_code FROM public.cart_user WHERE client_id = $1",
				tt.inputQuery,
			).
			Return(&tt.outQuery).
			Times(tt.countQuery)
		testUser := &Wrapper{Conn: m}
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
