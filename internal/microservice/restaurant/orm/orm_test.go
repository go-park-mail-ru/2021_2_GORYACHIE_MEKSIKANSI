package orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/restaurant/orm/mocks"
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
	rest "2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant"
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgproto3/v2"
	"github.com/stretchr/testify/require"
	"testing"
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
		switch dest[i].(type) {
		case *int:
			*dest[i].(*int) = r.row[i].(int)
		case *string:
			*dest[i].(*string) = r.row[i].(string)
		case *float32:
			*dest[i].(*float32) = float32(r.row[i].(float64))
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
		switch dest[i].(type) {
		case *int:
			*dest[i].(*int) = r.row[i].(int)
		case *string:
			*dest[i].(*string) = r.row[i].(string)
		case *float32:
			*dest[i].(*float32) = float32(r.row[i].(float64))
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
	testName string
	row      Rows
	errQuery error
	out      []rest.Restaurants
	outErr   string
}{
	{
		testName: "One",
		out:      []rest.Restaurants{{Id: 1, Img: "1", Name: "1", CostForFreeDelivery: 1, MinDelivery: 1, MaxDelivery: 1, Rating: 1.0}},
		errQuery: nil,
		row:      Rows{rows: 1, row: []interface{}{1, "1", "1", 1, 1, 1, 1.0}, errRow: nil},
		outErr:   errorsConst.RGetRestaurantsRestaurantsNotSelect,
	},
}

func TestGetRestaurants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range GetRestaurants {
		m.
			EXPECT().
			Query(context.Background(),
				"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating"+
					" FROM restaurant ORDER BY random() LIMIT 50",
			).
			Return(&tt.row, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetRestaurants()
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var GetGeneralInfoRestaurant = []struct {
	testName   string
	input      int
	row        Row
	inputQuery int
	out        *rest.RestaurantId
	outErr     string
}{
	{
		testName:   "One",
		input:      1,
		inputQuery: 1,
		out: &rest.RestaurantId{Id: 1, Img: "1", Name: "1",
			CostForFreeDelivery: 1, MinDelivery: 1, MaxDelivery: 1, Rating: 1, Tags: interface{}(nil), Menu: interface{}(nil)},
		row:    Row{row: []interface{}{1, "1", "1", 1, 1, 1, 1.0}, errRow: nil},
		outErr: errorsConst.RGetRestaurantRestaurantNotFound,
	},
	{
		testName:   "Two",
		input:      1,
		inputQuery: 1,
		out:        nil,
		row: Row{row: []interface{}{},
			errRow: errors.New(errorsConst.RGetRestaurantRestaurantNotFound)},
		outErr: errorsConst.RGetRestaurantRestaurantNotFound,
	},
}

func TestGetGeneralInfoRestaurant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range GetGeneralInfoRestaurant {
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant WHERE id = $1",
				tt.inputQuery,
			).
			Return(&tt.row)
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

var GetTagsRestaurant = []struct {
	testName   string
	input      int
	rowsQuery  Rows
	inputQuery int
	errQuery   error
	out        []rest.Tag
	outErr     string
}{
	{
		testName:   "One",
		input:      1,
		inputQuery: 1,
		errQuery:   nil,
		out:        []rest.Tag{{Id: 1, Name: "1"}},
		rowsQuery:  Rows{row: []interface{}{1, "1"}, errRow: nil, rows: 1},
		outErr:     errorsConst.RGetRestaurantRestaurantNotFound,
	},
}

func TestGetTagsRestaurant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range GetTagsRestaurant {
		m.
			EXPECT().
			Query(context.Background(),
				"SELECT id, category FROM restaurant_category WHERE restaurant = $1",
				tt.inputQuery,
			).
			Return(&tt.rowsQuery, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetTagsRestaurant(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var GetMenu = []struct {
	testName                     string
	inputId                      int
	rowsQuery                    Rows
	inputQueryId                 int
	errQuery                     error
	out                          []rest.Menu
	inputGetDishesRestaurantName string
	inputGetDishesRestaurantId   int
	outGetDishesRestaurant       []rest.Menu
	errGetDishesRestaurant       error
	outErr                       string
	inputQueryDishesName         string
	inputQueryDishesId           int
	rowsQueryDishes              Rows
	errQueryDishes               error
}{
	{
		testName:               "One",
		inputId:                1,
		inputQueryId:           1,
		errQuery:               nil,
		out:                    []rest.Menu{{Name: "1", DishesMenu: []rest.DishesMenu{{Id: 1, Name: "1", Cost: 1, Kilocalorie: 1, Img: "1"}}}},
		outGetDishesRestaurant: []rest.Menu{{}},
		errGetDishesRestaurant: nil,
		rowsQuery:              Rows{row: []interface{}{"1"}, errRow: nil, rows: 1},
		outErr:                 errorsConst.RGetRestaurantRestaurantNotFound,
		inputQueryDishesName:   "1",
		inputQueryDishesId:     1,
		rowsQueryDishes:        Rows{row: []interface{}{1, "1", "1", 1, 1}, errRow: nil, rows: 1},
		errQueryDishes:         nil,
	},
}

func TestGetMenu(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range GetMenu {
		m.
			EXPECT().
			Query(context.Background(),
				"SELECT id, avatar, name, cost, kilocalorie FROM dishes WHERE category_restaurant = $1 AND restaurant = $2",
				tt.inputQueryDishesName, tt.inputQueryDishesId,
			).
			Return(&tt.rowsQueryDishes, tt.errQueryDishes)
		m.
			EXPECT().
			Query(context.Background(),
				"SELECT DISTINCT category_restaurant FROM dishes WHERE restaurant = $1",
				tt.inputQueryId,
			).
			Return(&tt.rowsQuery, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetMenu(tt.inputId)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var GetDishes = []struct {
	testName                     string
	inputRestId                  int
	inputDishesId                int
	rowsQuery                    Rows
	inputQueryId                 int
	errQuery                     error
	out                          *rest.Dishes
	inputGetDishesRestaurantName string
	inputGetDishesRestaurantId   int
	outErr                       string
}{
	{
		testName:      "One",
		inputRestId:   1,
		inputDishesId: 1,
		inputQueryId:  1,
		errQuery:      nil,
		out: &rest.Dishes{Id: 1, Img: "1", Title: "1", Cost: 1, Ccal: 1, Description: "1",
			Radios: interface{}(nil), Ingredient: interface{}(nil)},
		rowsQuery: Rows{rows: 1, row: []interface{}{1, "1", "1", 1, 1, "1"}},
		outErr:    errorsConst.RGetRestaurantRestaurantNotFound,
	},
}

func TestGetDishes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range GetDishes {
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT id, avatar, name, cost, kilocalorie, description FROM dishes WHERE id = $1 AND restaurant = $2",
				tt.inputQueryId,
			).
			Return(&tt.rowsQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetDishes(tt.inputRestId, tt.inputDishesId)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
