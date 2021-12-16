package orm

import (
	rest "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant/orm/mocks"
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	"context"
	"errors"
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

func (r *Rows) Close() {
}

func (r *Rows) Err() error {
	return nil
}

func (r *Rows) CommandTag() pgconn.CommandTag {
	return nil
}

func (r *Rows) FieldDescriptions() []pgproto3.FieldDescription {
	return nil
}

func (r *Rows) Values() ([]interface{}, error) {
	return nil, nil
}

func (r *Rows) RawValues() [][]byte {
	return nil
}

func (r *Rows) Scan(dest ...interface{}) error {

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

func (r *Rows) Next() bool {
	if r.currentRow == r.rows {
		return false
	}
	r.currentRow++
	return true
}

var GetRestaurants = []struct {
	testName                 string
	out                      *rest.AllRestaurantsPromo
	outErr                   string
	outQuery                 Rows
	errQuery                 error
	countQuery               int
	errBeginTransaction      error
	countRollbackTransaction int
	errCommitTransaction     error
	countCommitTransaction   int
}{
	{
		testName: "First",
		out: &rest.AllRestaurantsPromo{
			Restaurant: []rest.Restaurants{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "restaurant",
					CostForFreeDelivery: 1,
					MinDelivery:         1,
					MaxDelivery:         1,
					Rating:              1,
				},
			},
			AllTags: []rest.Tag{
				{
					Id:   1,
					Name: "Кафе",
				},
			},
			AllPromo: nil,
		},
		errQuery: nil,
		outQuery: Rows{rows: 1,
			row:    []interface{}{1, "/url/", "restaurant", 1, 1, 1, 1.0, "Кафе", 1},
			errRow: nil,
		},
		outErr:                   "",
		errBeginTransaction:      nil,
		countQuery:               1,
		countRollbackTransaction: 1,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
	},
}

func TestGetRestaurants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range GetRestaurants {
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
				"SELECT t.id, t.avatar, t.name, t.price_delivery, t.min_delivery_time, t.max_delivery_time,"+
					" t.rating, rc.category, rc.id "+
					"FROM (SELECT r.id, r.avatar, r.name, r.price_delivery, r.min_delivery_time, r.max_delivery_time,"+
					" r.rating FROM restaurant r ORDER BY random() LIMIT 51) t "+
					"LEFT JOIN restaurant_category rc ON rc.restaurant = t.id",
			).
			Return(&tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetRestaurants()
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

var GetRecommendedRestaurants = []struct {
	testName                 string
	out                      *rest.AllRestaurants
	outErr                   string
	outQuery                 Rows
	errQuery                 error
	countQuery               int
	errBeginTransaction      error
	countRollbackTransaction int
	errCommitTransaction     error
	countCommitTransaction   int
}{
	{
		testName: "First",
		out: &rest.AllRestaurants{
			Restaurant: []rest.Restaurants{
				{
					Id:                  1,
					Img:                 "/url/",
					Name:                "restaurant",
					CostForFreeDelivery: 1,
					MinDelivery:         1,
					MaxDelivery:         1,
					Rating:              1,
				},
			},
			AllTags: []rest.Tag{
				{
					Id:   1,
					Name: "Кафе",
				},
			},
		},
		errQuery: nil,
		outQuery: Rows{rows: 1,
			row:    []interface{}{1, "/url/", "restaurant", 1, 1, 1, 1.0, "Кафе", 1},
			errRow: nil,
		},
		outErr:                   "",
		errBeginTransaction:      nil,
		countQuery:               1,
		countRollbackTransaction: 1,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
	},
}

func TestGetRecommendedRestaurants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range GetRecommendedRestaurants {
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
				"SELECT t.id, t.avatar, t.name, t.price_delivery, t.min_delivery_time, "+
					"t.max_delivery_time, t.rating, rc.category, rc.id FROM "+
					"(SELECT r.id, r.avatar, r.name, r.price_delivery, r.min_delivery_time, "+
					"r.max_delivery_time, r.rating FROM restaurant r ORDER BY rating DESC LIMIT 6) t "+
					"LEFT JOIN restaurant_category rc ON rc.restaurant = t.id",
			).
			Return(&tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetRecommendedRestaurants()
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

var GetRestaurant = []struct {
	testName                 string
	input                    int
	row                      Row
	inputQuery               int
	out                      *rest.RestaurantId
	outErr                   string
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:   "First",
		input:      1,
		inputQuery: 1,
		out: &rest.RestaurantId{Id: 1, Img: "1", Name: "1",
			CostForFreeDelivery: 1, MinDelivery: 1, MaxDelivery: 1, Rating: 1, Tags: nil, Menu: nil},
		row:                      Row{row: []interface{}{1, "1", "1", 1, 1, 1, 1.0}, errRow: nil},
		outErr:                   "",
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
	{
		testName:   "Second",
		input:      1,
		inputQuery: 1,
		out:        nil,
		row: Row{row: []interface{}{},
			errRow: errors.New(errorsConst.RGetRestaurantRestaurantNotFound)},
		outErr:                   errorsConst.RGetRestaurantRestaurantNotFound,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   0,
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
				"SELECT r.id, r.avatar, r.name, r.price_delivery, r.min_delivery_time, r.max_delivery_time, r.rating FROM restaurant r WHERE r.id = $1",
				tt.inputQuery,
			).
			Return(&tt.row)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetRestaurant(tt.input)
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

var GetTagsRestaurant = []struct {
	testName                 string
	input                    int
	out                      []rest.Tag
	outErr                   string
	inputQuery               int
	rowsQuery                Rows
	errQuery                 error
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:                 "First",
		input:                    1,
		inputQuery:               1,
		errQuery:                 nil,
		out:                      []rest.Tag{{Id: 1, Name: "1"}},
		rowsQuery:                Rows{row: []interface{}{1, "1", 0}, errRow: nil, rows: 1},
		outErr:                   "",
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestGetTagsRestaurant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range GetTagsRestaurant {
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
				"SELECT id, category, place FROM restaurant_category WHERE restaurant = $1",
				tt.inputQuery,
			).
			Return(&tt.rowsQuery, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetTagsRestaurant(tt.input)
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

var GetMenu = []struct {
	testName                 string
	inputId                  int
	out                      []rest.Menu
	outErr                   string
	inputQuery               int
	outQuery                 Rows
	errQuery                 error
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:                 "First",
		inputId:                  1,
		inputQuery:               1,
		out:                      []rest.Menu{{Name: "Сладости", DishesMenu: []rest.DishesMenu{{Id: 1, Name: "Шоколад", Cost: 100, Kilocalorie: 1, Img: "/url"}}}},
		outErr:                   "",
		outQuery:                 Rows{row: []interface{}{"Сладости", 1, "/url", "Шоколад", 100, 1, 0, 0}, errRow: nil, rows: 1},
		errQuery:                 nil,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestGetMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range GetMenu {
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
				"SELECT category_restaurant, id, avatar, name, cost, kilocalorie, place, place_category FROM dishes WHERE restaurant = $1",
				tt.inputQuery,
			).
			Return(&tt.outQuery, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetMenu(tt.inputId)
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

var GetDishes = []struct {
	testName                 string
	inputRestId              int
	inputDishesId            int
	out                      *rest.Dishes
	outErr                   string
	inputQueryDishesId       int
	inputQueryRestaurantId   int
	outQuery                 Rows
	errQuery                 error
	countQuery               int
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:      "First",
		inputRestId:   1,
		inputDishesId: 1,
		out: &rest.Dishes{Id: 1, Img: "/url/", Title: "Шоколад", Cost: 100, Ccal: 500, Description: "Очень сладкий, очень вкусный",
			Radios: nil, Ingredient: nil},
		inputQueryDishesId:     1,
		inputQueryRestaurantId: 1,
		outQuery: Rows{rows: 1, row: []interface{}{1, "/url/", "Шоколад", 100, 500,
			"Очень сладкий, очень вкусный", nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}},
		errQuery:                 nil,
		outErr:                   "",
		countQuery:               1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestGetDishes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range GetDishes {
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
				"SELECT d.id, d.avatar, d.name, d.cost, d.kilocalorie, d.description, r.id, r.name, sr.id, sr.name, r.place, "+
					"sr.place, sd.id, sd.name, sd.cost, sd.place "+
					"FROM dishes d"+
					" LEFT JOIN radios r ON d.id=r.food "+
					"LEFT JOIN structure_radios sr ON sr.radios=r.id "+
					"LEFT JOIN structure_dishes sd ON sd.food=d.id WHERE d.id = $1 AND restaurant = $2",
				tt.inputQueryDishesId, tt.inputQueryRestaurantId,
			).
			Return(&tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetDishes(tt.inputRestId, tt.inputDishesId)
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

var GetReview = []struct {
	testName                 string
	input                    int
	out                      []rest.Review
	outErr                   string
	inputQuery               int
	outQuery                 Rows
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
		input:                    1,
		out:                      []rest.Review{{Name: "1", Text: "1", Date: "11.10.2020", Time: "00:00", Rate: 5}},
		inputQuery:               1,
		outQuery:                 Rows{rows: 1, row: []interface{}{"1", "1", time.Date(2020, 10, 11, 0, 0, 0, 0, time.UTC), 5}},
		errQuery:                 nil,
		outErr:                   "",
		countQuery:               1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestGetReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range GetReview {
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
				"SELECT gn.name, r.text, r.date_create, r.rate FROM review r "+
					"LEFT JOIN general_user_info gn ON r.author = gn.id "+
					"WHERE r.restaurant = $1 ORDER BY r.date_create",
				tt.inputQuery,
			).
			Return(&tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetReview(tt.input)
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

var CreateReview = []struct {
	testName                 string
	inputClientId            int
	inputNewReview           rest.NewReview
	outErr                   string
	inputQueryClientId       int
	inputQueryRestaurantId   int
	inputQueryText           string
	inputQueryRate           int
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
		inputClientId:            1,
		inputNewReview:           rest.NewReview{Restaurant: rest.RestaurantId{Id: 1}, Text: "Very good restaurant", Rate: 5},
		outErr:                   "",
		inputQueryClientId:       1,
		inputQueryRestaurantId:   1,
		inputQueryText:           "Very good restaurant",
		inputQueryRate:           5,
		errQuery:                 nil,
		countQuery:               1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestCreateReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range CreateReview {
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
				"INSERT INTO review (author, restaurant, text, rate) VALUES ($1, $2, $3, $4)",
				tt.inputClientId, tt.inputQueryRestaurantId, tt.inputQueryText, tt.inputQueryRate,
			).
			Return(nil, tt.errQuery).
			Times(tt.countQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := testUser.CreateReview(tt.inputClientId, tt.inputNewReview)
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

var SearchCategory = []struct {
	testName                 string
	input                    string
	out                      []int
	outErr                   string
	inputQuery               string
	outQuery                 Rows
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
		input:                    "cafe",
		out:                      []int{1, 2},
		outErr:                   "",
		inputQuery:               "cafe",
		outQuery:                 Rows{row: []interface{}{1, 2}, rows: 2},
		errQuery:                 nil,
		countQuery:               1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestSearchCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range SearchCategory {
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
				"SELECT restaurant FROM restaurant_category WHERE fts @@ to_tsquery($1)",
				tt.inputQuery,
			).
			Return(&tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.SearchCategory(tt.input)
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

var GetGeneralInfoRestaurant = []struct {
	testName                 string
	input                    int
	out                      *rest.Restaurants
	outErr                   string
	inputQuery               int
	outQuery                 Row
	countQuery               int
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName: "First",
		input:    1,
		out: &rest.Restaurants{Id: 1, Img: "/url/", Name: "CafeHouse",
			CostForFreeDelivery: 250, MinDelivery: 15, MaxDelivery: 30, Rating: 3},
		outErr:                   "",
		inputQuery:               1,
		outQuery:                 Row{row: []interface{}{1, "/url/", "CafeHouse", 250, 15, 30, 3.0}},
		countQuery:               1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestGetGeneralInfoRestaurant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range GetGeneralInfoRestaurant {
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
			result, err := testUser.GetGeneralInfoRestaurant(tt.input)
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

var GetFavoriteRestaurants = []struct {
	testName                 string
	input                    int
	out                      []rest.Restaurants
	outErr                   string
	inputQuery               int
	outQuery                 Rows
	errQuery                 error
	countQuery               int
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName: "First",
		input:    1,
		out: []rest.Restaurants{{Id: 1, Img: "/url/", Name: "CafeHouse",
			CostForFreeDelivery: 250, MinDelivery: 15, MaxDelivery: 30, Rating: 3}},
		outErr:                   "",
		inputQuery:               1,
		outQuery:                 Rows{rows: 1, row: []interface{}{1, "/url/", "CafeHouse", 250, 15, 30, 3.0, 0}},
		errQuery:                 nil,
		countQuery:               1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestGetFavoriteRestaurants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range GetFavoriteRestaurants {
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
				"SELECT r.id, r.avatar, r.name, r.price_delivery, r.min_delivery_time, r.max_delivery_time, r.rating, fr.position"+
					" FROM restaurant r RIGHT JOIN favorite_restaurant fr ON fr.restaurant = r.id WHERE fr.client = $1",
				tt.inputQuery,
			).
			Return(&tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetFavoriteRestaurants(tt.input)
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

var IsFavoriteRestaurant = []struct {
	testName                 string
	inputClient              int
	inputRestaurant          int
	out                      bool
	outErr                   string
	inputQueryClientId       int
	inputQueryRestauranId    int
	outQuery                 Row
	countQuery               int
	errBeginTransaction      error
	errCommitTransaction     error
	countCommitTransaction   int
	errRollbackTransaction   error
	countRollbackTransaction int
}{
	{
		testName:                 "First",
		inputClient:              1,
		inputRestaurant:          1,
		out:                      true,
		outErr:                   "",
		inputQueryClientId:       1,
		inputQueryRestauranId:    1,
		outQuery:                 Row{row: []interface{}{1}},
		countQuery:               1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestIsFavoriteRestaurant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range IsFavoriteRestaurant {
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
				"SELECT id FROM favorite_restaurant WHERE client = $1 AND restaurant = $2",
				tt.inputQueryClientId, tt.inputQueryRestauranId,
			).
			Return(&tt.outQuery).
			Times(tt.countQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.IsFavoriteRestaurant(tt.inputClient, tt.inputRestaurant)
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

var GetPromoCodes = []struct {
	testName                 string
	input                    int
	out                      []rest.Promocode
	outErr                   string
	inputQuery               int
	outQuery                 Rows
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
		input:                    1,
		out:                      []rest.Promocode{{RestaurantId: 1, Description: "free", Img: "/url/", Name: "freedelivery"}},
		outErr:                   "",
		inputQuery:               1,
		outQuery:                 Rows{rows: 1, row: []interface{}{"freedelivery", "free", "/url/", 1}},
		errQuery:                 nil,
		countQuery:               1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,
	},
}

func TestGetPromoCodes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range GetPromoCodes {
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
				"SELECT name, description, avatar, restaurant FROM promocode ORDER BY random() LIMIT 5").
			Return(&tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetPromoCodes()
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
