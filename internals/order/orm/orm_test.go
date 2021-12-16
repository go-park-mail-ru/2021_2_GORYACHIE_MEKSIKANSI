package orm

import (
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	orderPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/order"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/order/orm/mocks"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant"
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgproto3/v2"
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

type Rows struct {
	row        []interface{}
	rows       int
	currentRow int
	errRow     error
}

func (r Rows) Close() {
}

func (r Rows) Err() error {
	return nil
}

func (r Rows) CommandTag() pgconn.CommandTag {
	return nil
}

func (r Rows) FieldDescriptions() []pgproto3.FieldDescription {
	return nil
}

func (r Rows) Values() ([]interface{}, error) {
	return nil, nil
}

func (r Rows) RawValues() [][]byte {
	return nil
}

func (r Rows) Scan(dest ...interface{}) error {
	for i := range dest {
		j := i + len(dest)*(r.currentRow-1)
		if r.row[j] == nil {
			dest[i] = nil
			continue
		}
		switch dest[i].(type) {
		case *int:
			*dest[i].(*int) = r.row[j].(int)
		case *string:
			*dest[i].(*string) = r.row[j].(string)
		case **string:
			t := r.row[j].(string)
			*dest[i].(**string) = &t
		case *float32:
			*dest[i].(*float32) = float32(r.row[j].(float64))
		case **int32:
			t := int32(r.row[j].(int))
			*dest[i].(**int32) = &t
		case *time.Time:
			*dest[i].(*time.Time) = r.row[j].(time.Time)
		case *bool:
			*dest[i].(*bool) = r.row[j].(bool)
		default:
			dest[i] = nil
		}
	}
	return r.errRow
}

func (r Rows) Next() bool {
	if r.currentRow == r.rows {
		return false
	}
	r.currentRow++
	return true
}

var GetOrder = []struct {
	testName                 string
	inputClientId            int
	inputOrderId             int
	out                      *orderPkg.ActiveOrder
	outErr                   string
	inputQueryClientId       int
	inputQueryOrderId        int
	outQuery                 Rows
	errQuery                 error
	countQuery               int
	errBeginTransaction      error
	inputDelete              int
	errDelete                error
	countDelete              int
	errCommitTransaction     error
	countCommitTransaction   int
	countRollbackTransaction int
}{
	{
		testName:                 "First",
		inputClientId:            1,
		inputOrderId:             1,
		out:                      nil,
		outErr:                   errPkg.OGetOrderNotExist,
		inputQueryClientId:       1,
		inputQueryOrderId:        1,
		outQuery:                 Rows{},
		errQuery:                 nil,
		countQuery:               1,
		errBeginTransaction:      nil,
		inputDelete:              1,
		errDelete:                nil,
		countDelete:              1,
		errCommitTransaction:     nil,
		countCommitTransaction:   0,
		countRollbackTransaction: 1,
	},
}

func TestGetOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range GetOrder {
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
			Query(context.Background(),
				"SELECT order_user.id, ol.item_number, date_order, status, au.alias, au.city, au.street, au.house,"+
					" au.flat, au.porch, au.floor, au.intercom, au.comment, au.latitude,"+
					" au.longitude, d.id, d.avatar, d.name, ol.count_dishes, "+
					"d.cost, d.kilocalorie, d.weight, d.description, sr.name, "+
					"sr.radios, sr.id, sd.name, sd.id, sd.cost, restaurant_id, r.name, r.avatar, r.city, r.street,"+
					" r.house, r.floor, r.latitude, r.longitude, dCost, sumCost, ol.place, orl.place, osl.place, r.max_delivery_time "+
					"FROM order_user"+
					" LEFT JOIN address_user au ON au.id = order_user.address_id"+
					" LEFT JOIN order_list ol ON ol.order_id = order_user.id"+
					" LEFT JOIN dishes d ON d.id = ol.food"+
					" LEFT JOIN order_structure_list osl ON osl.order_id = order_user.id and d.id=osl.food and ol.id=osl.list_id"+
					" LEFT JOIN order_radios_list orl ON orl.order_id = order_user.id and ol.food=orl.food and ol.id=orl.list_id"+
					" LEFT JOIN structure_radios sr ON sr.id = orl.radios"+
					" LEFT JOIN structure_dishes sd ON sd.id = osl.structure_food"+
					" LEFT JOIN restaurant r ON r.id = order_user.restaurant_id WHERE order_user.client_id = $1 AND order_user.id = $2",
				tt.inputQueryClientId, tt.inputQueryOrderId,
			).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetOrder(tt.inputClientId, tt.inputOrderId)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var GetOrders = []struct {
	testName                 string
	inputClientId            int
	inputOrderId             int
	out                      *orderPkg.HistoryOrderArray
	outErr                   string
	inputQueryClientId       int
	outQuery                 Rows
	errQuery                 error
	countQuery               int
	errBeginTransaction      error
	inputDelete              int
	errDelete                error
	countDelete              int
	errCommitTransaction     error
	countCommitTransaction   int
	countRollbackTransaction int
}{
	{
		testName:                 "First",
		inputClientId:            1,
		inputOrderId:             1,
		out:                      nil,
		outErr:                   errPkg.OGetOrdersOrdersIsVoid,
		inputQueryClientId:       1,
		outQuery:                 Rows{},
		errQuery:                 nil,
		countQuery:               1,
		errBeginTransaction:      nil,
		inputDelete:              1,
		errDelete:                nil,
		countDelete:              1,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		countRollbackTransaction: 1,
	},
}

func TestGetOrders(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range GetOrders {
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
			Query(context.Background(),

				"SELECT order_user.id, ol.item_number, date_order, status, au.alias, au.city, au.street, au.house,"+
					" au.flat, au.porch, au.floor, au.intercom, au.comment, au.latitude,"+
					" au.longitude, d.id, d.avatar, d.name, ol.count_dishes, "+
					"d.cost, d.kilocalorie, d.weight, d.description, sr.name, "+
					"sr.radios, sr.id, sd.name, sd.id, sd.cost, restaurant_id, r.name, r.avatar, r.city, r.street,"+
					" r.house, r.floor, r.latitude, r.longitude, dCost, sumCost, ol.place, orl.place, osl.place "+
					"FROM order_user"+
					" LEFT JOIN address_user au ON au.id = order_user.address_id"+
					" LEFT JOIN order_list ol ON ol.order_id = order_user.id"+
					" LEFT JOIN dishes d ON d.id = ol.food"+
					" LEFT JOIN order_structure_list osl ON osl.order_id = order_user.id and d.id=osl.food and ol.id=osl.list_id"+
					" LEFT JOIN order_radios_list orl ON orl.order_id = order_user.id and ol.food=orl.food and ol.id=orl.list_id"+
					" LEFT JOIN structure_radios sr ON sr.id = orl.radios"+
					" LEFT JOIN structure_dishes sd ON sd.id = osl.structure_food"+
					" LEFT JOIN restaurant r ON r.id = order_user.restaurant_id WHERE order_user.client_id = $1 ORDER BY date_order",
				tt.inputQueryClientId,
			).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetOrders(tt.inputClientId)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var UpdateStatusOrder = []struct {
	testName                 string
	inputClientId            int
	inputStatus              int
	outErr                   string
	inputQueryClientId       int
	inputQueryStatus         int
	errQuery                 error
	countQuery               int
	errBeginTransaction      error
	inputDelete              int
	errDelete                error
	countDelete              int
	errCommitTransaction     error
	countCommitTransaction   int
	countRollbackTransaction int
}{
	{
		testName:                 "First",
		inputClientId:            1,
		inputStatus:              1,
		outErr:                   errPkg.OGetOrdersOrdersIsVoid,
		inputQueryClientId:       1,
		inputQueryStatus:         1,
		errQuery:                 nil,
		countQuery:               1,
		errBeginTransaction:      nil,
		inputDelete:              1,
		errDelete:                nil,
		countDelete:              1,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		countRollbackTransaction: 1,
	},
}

func TestUpdateStatusOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range UpdateStatusOrder {
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
				"UPDATE order_user SET status = $1 WHERE id = $2",
				tt.inputQueryClientId, tt.inputQueryStatus,
			).
			Return(nil, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := testUser.UpdateStatusOrder(tt.inputClientId, tt.inputStatus)
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var DeleteCart = []struct {
	testName                 string
	input                    int
	outErr                   string
	inputQuery               int
	errQuery                 error
	countQuery               int
	errBeginTransaction      error
	inputDelete              int
	errDelete                error
	countDelete              int
	errCommitTransaction     error
	countCommitTransaction   int
	countRollbackTransaction int
}{
	{
		testName:                 "First",
		input:                    1,
		outErr:                   errPkg.OGetOrdersOrdersIsVoid,
		inputQuery:               1,
		errQuery:                 nil,
		countQuery:               1,
		errBeginTransaction:      nil,
		inputDelete:              1,
		errDelete:                nil,
		countDelete:              1,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		countRollbackTransaction: 1,
	},
}

func TestDeleteCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range DeleteCart {
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
				"DELETE FROM cart_food CASCADE WHERE client_id = $1",
				tt.inputQuery,
			).
			Return(nil, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := testUser.DeleteCart(tt.input)
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var GetRestaurant = []struct {
	testName                 string
	out                      *restaurant.RestaurantId
	outErr                   string
	input                    int
	outQuery                 Row
	countQuery               int
	inputQuery               int
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		input:      1,
		inputQuery: 1,
		outQuery:   Row{row: []interface{}{1, "1", "1", 1, 1, 1, 1.0}},
		countQuery: 1,
		testName:   "First",
		outErr:     "",
		out: &restaurant.RestaurantId{
			Id:                  1,
			Img:                 "1",
			Name:                "1",
			CostForFreeDelivery: 1,
			MinDelivery:         1,
			MaxDelivery:         1,
			Rating:              1,
			Tags:                []restaurant.Tag(nil),
			Menu:                []restaurant.Menu(nil),
		},
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestGetRestaurant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range GetRestaurant {
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
				"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant WHERE id = $1",
				tt.inputQuery,
			).
			Return(&tt.outQuery).
			Times(tt.countQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetRestaurant(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
