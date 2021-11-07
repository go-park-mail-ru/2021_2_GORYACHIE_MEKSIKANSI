package Cart

import (
	mocks "2021_2_GORYACHIE_MEKSIKANSI/Test/Mocks"
	"2021_2_GORYACHIE_MEKSIKANSI/Utils"
	_ "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
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

var OrmGetStructFood = []struct {
	testName           string
	input              int
	out                []Utils.IngredientCartResponse
	outErr             string
	inputQueryCart     int
	resultQueryCart    Rows
	errQueryCart error
	inputQueryStructure int
	resultQueryStructure Row
	countQueryStructure int
}{
	{
		testName:           "One",
		input: 1,
		out:                nil,
		outErr: "",
		inputQueryCart: 1,
		resultQueryCart: Rows{},
		errQueryCart: nil,
		inputQueryStructure: 1,
		resultQueryStructure: Row{},
		countQueryStructure: 0,
	},
}

func TestOrmGetStructFood(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmGetStructFood {
		m.
			EXPECT().
			Query(context.Background(),
				"SELECT checkbox FROM cart_structure_food WHERE client_id = $1",
				tt.inputQueryCart,
			).
			Return(&tt.resultQueryCart, tt.errQueryCart)
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT name, cost FROM structure_dishes WHERE id = $1",
				tt.inputQueryStructure,
			).
			Return(&tt.resultQueryStructure).
			Times(tt.countQueryStructure)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetStructFood(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmGetStructRadios = []struct {
	testName           string
	input              int
	out                []Utils.RadiosCartResponse
	outErr             string
	inputQueryCart     int
	resultQueryCart    Rows
	errQueryCart error
	inputQueryStructure int
	resultQueryStructure Row
	countQueryStructure int
}{
	{
		testName:           "One",
		input: 1,
		out:                nil,
		outErr: "",
		inputQueryCart: 1,
		resultQueryCart: Rows{},
		errQueryCart: nil,
		inputQueryStructure: 1,
		resultQueryStructure: Row{},
		countQueryStructure: 0,
	},
}

func TestOrmGetStructRadios(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmGetStructRadios {
		m.
			EXPECT().
			Query(context.Background(),
				"SELECT radios_id, radios FROM cart_radios_food WHERE client_id = $1",
				tt.inputQueryCart,
			).
			Return(&tt.resultQueryCart, tt.errQueryCart)
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT name FROM structure_radios WHERE id = $1",
				tt.inputQueryStructure,
			).
			Return(&tt.resultQueryStructure).
			Times(tt.countQueryStructure)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetStructRadios(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmGetCart = []struct {
	testName           string
	input              int
	outOne             *Utils.ResponseCartErrors
	outTwo             []Utils.CastDishesErrs
	outErr             string
	inputQuery         int
	resultQuery        Rows
	errQuery           error
}{
	{
		testName:           "One",
		input: 1,
		outOne:                &Utils.ResponseCartErrors{Restaurant:Utils.RestaurantIdCastResponse{Id:0, Img:"", Name:"", CostForFreeDelivery:0, MinDelivery:0, MaxDelivery:0, Rating:0}, Dishes:[]Utils.DishesCartResponse(nil), Cost:Utils.CostCartResponse{DCost:0, SumCost:0}, DishErr:[]Utils.CastDishesErrs(nil)},
		outTwo:                nil,
		outErr: "",
		inputQuery: 1,
		resultQuery: Rows{},
		errQuery: nil,
	},
}

func TestOrmGetCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmGetCart {
		m.
			EXPECT().
			Query(context.Background(),
				"SELECT food, count_food, number_item, name, cost, description, avatar, restaurant_id, count, weight, kilocalorie FROM cart JOIN dishes ON cart.food = dishes.id WHERE client_id = $1",
				tt.inputQuery,
			).
			Return(&tt.resultQuery, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			resultOne, resultTwo, err := testUser.GetCart(tt.input)
			require.Equal(t, tt.outOne, resultOne, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outOne, resultOne))
			require.Equal(t, tt.outTwo, resultTwo, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outTwo, resultTwo))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmDeleteCart = []struct {
	testName           string
	input              int
	outErr             string
	inputDeleteCart         int
	inputDeleteStructure         int
	inputDeleteRadios         int
	errDeleteCart           error
	errDeleteStructure          error
	errDeleteRadios           error
	countDeleteCart             int
	countDeleteStructure           int
	countDeleteRadios           int
}{
	{
		testName:           "One",
		input: 1,
		outErr: "",
		inputDeleteCart: 1,
		inputDeleteStructure: 1,
		inputDeleteRadios: 1,
		errDeleteCart: nil,
		errDeleteStructure: nil,
		errDeleteRadios: nil,
		countDeleteCart: 1,
		countDeleteStructure: 1,
		countDeleteRadios: 1,
	},
}

func TestOrmDeleteCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmDeleteCart {
		m.
			EXPECT().
			Exec(context.Background(),
				"DELETE FROM cart WHERE client_id = $1",
				tt.inputDeleteCart,
			).
			Return(nil, tt.errDeleteCart).
			Times(tt.countDeleteCart)
		m.
			EXPECT().
			Exec(context.Background(),
			"DELETE FROM cart_structure_food WHERE client_id = $1",
				tt.inputDeleteStructure,
			).
			Return(nil, tt.errDeleteStructure).
			Times(tt.countDeleteStructure)
		m.
			EXPECT().
			Exec(context.Background(),
			"DELETE FROM cart_radios_food WHERE client_id = $1",
				tt.inputDeleteRadios,
			).
			Return(nil, tt.errDeleteRadios).
			Times(tt.countDeleteRadios)
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

var OrmUpdateCartStructFood = []struct {
	testName           string
	inputClientId      int
	inputIngredient    []Utils.IngredientsCartRequest
	out                []Utils.IngredientCartResponse
	outErr             string
	inputQuery         int
	resultQuery Row
	countQuery int
	inputInsertIngredient int
	inputInsertClient int
	errInsert error
	countInsert int
}{
	{
		testName:           "One",
		inputIngredient: []Utils.IngredientsCartRequest{{Id: 1}},
		inputClientId: 1,
		out: []Utils.IngredientCartResponse{{Name:"1", Id:1, Cost:1}},
		outErr: "",
		inputQuery: 1,
		resultQuery: Row{row: []interface{}{1, "1", 1}},
		countQuery: 1,
		inputInsertIngredient: 1,
		inputInsertClient: 1,
		errInsert: nil,
		countInsert: 1,
	},
}

func TestOrmUpdateCartStructFood(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range OrmUpdateCartStructFood {
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT id, name, cost FROM structure_dishes WHERE id = $1",
				tt.inputQuery,
			).
			Return(&tt.resultQuery).
			Times(tt.countQuery)
		mTx.
			EXPECT().
			Exec(context.Background(),
			"INSERT INTO cart_structure_food (checkbox, client_id) VALUES ($1, $2)",
				tt.inputInsertIngredient, tt.inputInsertClient,
			).
			Return(nil, tt.errInsert).
			Times(tt.countInsert)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.UpdateCartStructFood(tt.inputIngredient, tt.inputClientId, mTx)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmUpdateCartRadios = []struct {
	testName           string
	inputClientId      int
	inputRadios        []Utils.RadiosCartRequest
	out                []Utils.RadiosCartResponse
	outErr             string
	inputQuery         int
	resultQuery Row
	countQuery int
	inputInsertRadiosId int
	inputInsertRadios int
	inputInsertClient int
	errInsert error
	countInsert int
}{
	{
		testName:           "One",
		inputRadios: []Utils.RadiosCartRequest{{Id: 1, RadiosId: 1}},
		inputClientId: 1,
		out: []Utils.RadiosCartResponse{{Name:"1", RadiosId:0, Id:1}},
		outErr: "",
		inputQuery: 1,
		resultQuery: Row{row: []interface{}{1, "1"}},
		countQuery: 1,
		inputInsertRadiosId: 1,
		inputInsertRadios: 1,
		inputInsertClient: 1,
		errInsert: nil,
		countInsert: 1,
	},
}

func TestOrmUpdateCartRadios(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range OrmUpdateCartRadios {
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT id, name FROM structure_radios WHERE id = $1",
				tt.inputQuery,
			).
			Return(&tt.resultQuery).
			Times(tt.countQuery)
		mTx.
			EXPECT().
			Exec(context.Background(),
			"INSERT INTO cart_radios_food (radios_id, radios, client_id) VALUES ($1, $2, $3)",
				tt.inputInsertRadiosId, tt.inputInsertRadios, tt.inputInsertClient,
			).
			Return(nil, tt.errInsert).
			Times(tt.countInsert)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.UpdateCartRadios(tt.inputRadios, tt.inputClientId, mTx)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmUpdateCart = []struct {
	testName           string
	inputClientId      int
	inputCart          Utils.RequestCartDefault
	outOne             *Utils.ResponseCartErrors
	outTwo             []Utils.CastDishesErrs
	outErr             string
	inputInsertClientId int
	inputInsertFood int
	inputInsertCountFood int
	inputInsertRestaurantId int
	inputInsertNumberItem int
	errInsert error
	countInsert int
	inputInsertDishId int
	inputQueryRestaurantId int
	resultQuery Row
	countQuery int
	errBegin error
	inputQueryStruct         int
	resultQueryStruct Row
	countQueryStruct  int
	inputInsertStructIngredient int
	inputInsertStructClient int
	errInsertStruct error
	countInsertStruct int
	inputQueryRadios         int
	resultQueryRadios Row
	countQueryRadios int
	inputInsertRadiosId int
	inputInsertRadios int
	inputInsertClient int
	errInsertRadios error
	countInsertRadios int
	errCommit error
	errRollback error
	countCommit int
	countRollback int
}{
	{
		testName:           "One",
		inputClientId: 1,
		inputCart: Utils.RequestCartDefault{Restaurant: Utils.RestaurantRequest{Id: 1},
			Dishes: []Utils.DishesRequest{{Id: 1, ItemNumber: 1, Count: 1,
				Radios: []Utils.RadiosCartRequest{{RadiosId: 1, Id: 1}},
				Ingredients: []Utils.IngredientsCartRequest{{Id: 1}}}}},
		outOne: &Utils.ResponseCartErrors{Restaurant:Utils.RestaurantIdCastResponse{Id:0, Img:"", Name:"", CostForFreeDelivery:0, MinDelivery:0, MaxDelivery:0, Rating:0}, Dishes:[]Utils.DishesCartResponse{Utils.DishesCartResponse{Id:1, ItemNumber:0, Img:"1", Name:"1", Count:1, Cost:1, Kilocalorie:1, Weight:1, Description:"1", RadiosCart:[]Utils.RadiosCartResponse{Utils.RadiosCartResponse{Name:"1", RadiosId:0, Id:1}}, IngredientCart:[]Utils.IngredientCartResponse{Utils.IngredientCartResponse{Name:"1", Id:1, Cost:1}}}}, Cost:Utils.CostCartResponse{DCost:0, SumCost:0}, DishErr:[]Utils.CastDishesErrs(nil)},
		outTwo: []Utils.CastDishesErrs(nil),
		outErr: "",
		inputInsertClientId: 1,
		inputInsertFood: 1,
		inputInsertCountFood: 1,
		inputInsertRestaurantId: 1,
		inputInsertNumberItem: 1,
		errInsert: nil,
		countInsert: 1,
		inputInsertDishId: 1,
		inputQueryRestaurantId: 1,
		resultQuery: Row{row: []interface{}{1, "1", 1, "1", "1", 1, 1, 1}},
		countQuery: 1,
		errBegin: nil,
		inputQueryStruct: 1,
		resultQueryStruct: Row{row: []interface{}{1, "1", 1}},
		countQueryStruct: 1,
		inputInsertStructIngredient: 1,
		inputInsertStructClient: 1,
		errInsertStruct: nil,
		countInsertStruct: 1,
		inputQueryRadios: 1,
		resultQueryRadios: Row{row: []interface{}{1, "1"}},
		countQueryRadios: 1,
		inputInsertRadiosId: 1,
		inputInsertRadios: 1,
		inputInsertClient: 1,
		errInsertRadios: nil,
		countInsertRadios: 1,
		errCommit: nil,
		errRollback: nil,
		countCommit: 1,
		countRollback: 0,
	},
}

func TestOrmUpdateCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range OrmUpdateCart {
		mTx.
			EXPECT().
			Exec(context.Background(),
			"INSERT INTO cart (client_id, food, count_food, restaurant_id, number_item) VALUES ($1, $2, $3, $4, $5)",
				tt.inputInsertClientId, tt.inputInsertFood, tt.inputInsertCountFood, tt.inputInsertRestaurantId, tt.inputInsertNumberItem,
			).
			Return(nil, tt.errInsert).
			Times(tt.countInsert)
		mTx.
			EXPECT().
			Exec(context.Background(),
				"INSERT INTO cart_structure_food (checkbox, client_id) VALUES ($1, $2)",
				tt.inputInsertStructIngredient, tt.inputInsertStructClient,
			).
			Return(nil, tt.errInsertStruct).
			Times(tt.countInsertStruct)
		mTx.
			EXPECT().
			Exec(context.Background(),
				"INSERT INTO cart_radios_food (radios_id, radios, client_id) VALUES ($1, $2, $3)",
				tt.inputInsertRadiosId, tt.inputInsertRadios, tt.inputInsertClient,
			).
			Return(nil, tt.errInsertRadios).
			Times(tt.countInsertRadios)
		mTx.
			EXPECT().
			Commit(context.Background()).
			Return(tt.errCommit).
			Times(tt.countCommit)
		mTx.
			EXPECT().
			Rollback(context.Background()).
			Return(tt.errRollback).
			Times(tt.countRollback)
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT id, name FROM structure_radios WHERE id = $1",
				tt.inputQueryRadios,
			).
			Return(&tt.resultQueryRadios).
			Times(tt.countQueryRadios)
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT id, name, cost FROM structure_dishes WHERE id = $1",
				tt.inputQueryStruct,
			).
			Return(&tt.resultQueryStruct).
			Times(tt.countQueryStruct)
		m.
			EXPECT().
			QueryRow(context.Background(),
			"SELECT id, avatar, cost, name, description, count, weight, kilocalorie FROM dishes WHERE id = $1 AND restaurant = $2",
				tt.inputInsertDishId, tt.inputQueryRestaurantId,
			).
			Return(&tt.resultQuery).
			Times(tt.countQuery)
		m.
			EXPECT().
			Begin(context.Background()).
			Return(mTx, tt.errBegin)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			resultOne, resultTwo, err := testUser.UpdateCart(tt.inputCart, tt.inputClientId)
			require.Equal(t, tt.outOne, resultOne, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outOne, resultOne))
			require.Equal(t, tt.outTwo, resultTwo, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outTwo, resultTwo))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmGetPriceDelivery = []struct {
	testName    string
	out         int
	outErr      string
	input       int
	resultQuery Row
	inputQuery  int
}{
	{
		input:       1,
		inputQuery: 1,
		resultQuery: Row{row: []interface{}{1}},
		testName:    "One",
		outErr:      "",
		out:         1,
	},
}

func TestOrmGetPriceDelivery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmGetPriceDelivery {
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT price_delivery FROM restaurant WHERE id = $1",
				tt.inputQuery,
			).
			Return(&tt.resultQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetPriceDelivery(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationCalculatePriceDelivery = []struct {
	testName    string
	out         int
	outErr      string
	input       int
	inputGetPrice int
	resultGetPrice int
	errGetPrice   error
}{
	{
		input:       1,
		inputGetPrice: 1,
		resultGetPrice: 1,
		testName:    "One",
		outErr:      "",
		out:         1,
		errGetPrice:   nil,
	},
}

func TestApplicationCalculatePriceDelivery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperCart(ctrl)
	for _, tt := range ApplicationCalculatePriceDelivery {
		m.
			EXPECT().
			GetPriceDelivery(tt.inputGetPrice).
			Return(tt.resultGetPrice, tt.errGetPrice)
		test := Cart{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.CalculatePriceDelivery(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationCalculateCost = []struct {
	testName    string
	out         *Utils.CostCartResponse
	outErr      string
	inputResult     *Utils.ResponseCartErrors
	inputRest       *Utils.RestaurantId
	inputGetPrice int
	resultGetPrice int
	errGetPrice   error
	countGetPrice int
}{
	{
		inputResult: &Utils.ResponseCartErrors{Dishes: []Utils.DishesCartResponse{{Cost: 1, Count: 1}}, Cost: Utils.CostCartResponse{SumCost: 1}, Restaurant: Utils.RestaurantIdCastResponse{CostForFreeDelivery: 5}},
		inputRest:   &Utils.RestaurantId{Id: 1, CostForFreeDelivery: 5},
		inputGetPrice: 1,
		resultGetPrice: 1,
		testName:    "One",
		outErr:      "",
		out:         &Utils.CostCartResponse{DCost:1, SumCost:2},
		errGetPrice:   nil,
		countGetPrice: 1,
	},
}

func TestApplicationCalculateCost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperCart(ctrl)
	for _, tt := range ApplicationCalculateCost {
		m.
			EXPECT().
			GetPriceDelivery(tt.inputGetPrice).
			Return(tt.resultGetPrice, tt.errGetPrice).
			Times(tt.countGetPrice)
		test := Cart{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.CalculateCost(tt.inputResult, tt.inputRest)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationGetCart = []struct {
	testName    string
	input       int
	out         *Utils.ResponseCartErrors
	outErr      string
	inputGetPrice int
	resultGetPrice int
	errGetPrice   error
	countGetPrice int
	inputGetCart int
	resultGetCartResult *Utils.ResponseCartErrors
	resultGetCartErrorDishes []Utils.CastDishesErrs
	errGetCart error
	inputGeneralInfo int
	resultGeneralInfo *Utils.RestaurantId
	errGeneralInfo error
	countGeneralInfo int
}{
	{
		testName:    "One",
		input:       1,
		out:         &Utils.ResponseCartErrors{Restaurant:Utils.RestaurantIdCastResponse{Id:1, Img:"", Name:"", CostForFreeDelivery:0, MinDelivery:0, MaxDelivery:0, Rating:0}, Dishes:[]Utils.DishesCartResponse(nil), Cost:Utils.CostCartResponse{DCost:0, SumCost:0}, DishErr:[]Utils.CastDishesErrs{}},
		outErr:      "",
		inputGetPrice: 1,
		resultGetPrice: 1,
		errGetPrice: nil,
		countGetPrice: 0,
		inputGetCart: 1,
		resultGetCartResult: &Utils.ResponseCartErrors{Restaurant: Utils.RestaurantIdCastResponse{Id: 1}},
		resultGetCartErrorDishes: []Utils.CastDishesErrs{},
		errGetCart: nil,
		inputGeneralInfo: 1,
		resultGeneralInfo: &Utils.RestaurantId{Id: 1},
		errGeneralInfo: nil,
		countGeneralInfo: 1,
	},
}

func TestApplicationGetCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperCart(ctrl)
	mRest := mocks.NewMockWrapperRestaurant(ctrl)
	for _, tt := range ApplicationGetCart {
		m.
			EXPECT().
			GetCart(tt.inputGetCart).
			Return(tt.resultGetCartResult, tt.resultGetCartErrorDishes, tt.errGetCart)
		mRest.
			EXPECT().
			GetGeneralInfoRestaurant(tt.inputGeneralInfo).
			Return(tt.resultGeneralInfo, tt.errGeneralInfo).
			Times(tt.countGeneralInfo)
		m.
			EXPECT().
			GetPriceDelivery(tt.inputGetPrice).
			Return(tt.resultGetPrice, tt.errGetPrice).
			Times(tt.countGetPrice)

		test := Cart{DB: m, DBRestaurant: mRest}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.GetCart(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationUpdateCart = []struct {
	testName    string
	inputDishes Utils.RequestCartDefault
	inputId       int
	out         *Utils.ResponseCartErrors
	outErr      string
	inputGetPrice int
	resultGetPrice int
	errGetPrice   error
	countGetPrice int
	inputUpdateDishes Utils.RequestCartDefault
	inputUpdateId int
	resultUpdateResult *Utils.ResponseCartErrors
	resultUpdateErrorDishes []Utils.CastDishesErrs
	errUpdate error
	countUpdate int
	inputGeneralInfo int
	resultGeneralInfo *Utils.RestaurantId
	errGeneralInfo error
	countGeneralInfo int
	inputDelete int
	errDelete error
	countDelete int
}{
	{
		testName:    "One",
		inputDishes:   Utils.RequestCartDefault{Restaurant: Utils.RestaurantRequest{Id: 1}},
		inputId:       1,
		out:         &Utils.ResponseCartErrors{Restaurant:Utils.RestaurantIdCastResponse{Id:1, Img:"", Name:"", CostForFreeDelivery:0, MinDelivery:0, MaxDelivery:0, Rating:0}, Dishes:[]Utils.DishesCartResponse(nil), Cost:Utils.CostCartResponse{DCost:0, SumCost:0}, DishErr:[]Utils.CastDishesErrs{}},
		outErr:      "",
		inputGetPrice: 1,
		resultGetPrice: 1,
		errGetPrice: nil,
		countGetPrice: 0,
		inputUpdateDishes: Utils.RequestCartDefault{Restaurant: Utils.RestaurantRequest{Id: 1}},
		inputUpdateId: 1,
		resultUpdateResult: &Utils.ResponseCartErrors{Restaurant: Utils.RestaurantIdCastResponse{Id: 1}},
		resultUpdateErrorDishes: []Utils.CastDishesErrs{},
		errUpdate: nil,
		countUpdate: 1,
		inputGeneralInfo: 1,
		resultGeneralInfo: &Utils.RestaurantId{Id: 1},
		errGeneralInfo: nil,
		countGeneralInfo: 1,
		inputDelete: 1,
		errDelete: nil,
		countDelete: 1,
	},
}

func TestApplicationUpdateCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperCart(ctrl)
	mRest := mocks.NewMockWrapperRestaurant(ctrl)
	for _, tt := range ApplicationUpdateCart {
		m.
			EXPECT().
			DeleteCart(tt.inputDelete).
			Return(tt.errDelete).
			Times(tt.countDelete)
		m.
			EXPECT().
			UpdateCart(tt.inputUpdateDishes, tt.inputUpdateId).
			Return(tt.resultUpdateResult, tt.resultUpdateErrorDishes, tt.errUpdate).
			Times(tt.countUpdate)
		mRest.
			EXPECT().
			GetGeneralInfoRestaurant(tt.inputGeneralInfo).
			Return(tt.resultGeneralInfo, tt.errGeneralInfo).
			Times(tt.countGeneralInfo)
		m.
			EXPECT().
			GetPriceDelivery(tt.inputGetPrice).
			Return(tt.resultGetPrice, tt.errGetPrice).
			Times(tt.countGetPrice)

		test := Cart{DB: m, DBRestaurant: mRest}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.UpdateCart(tt.inputDishes, tt.inputId)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationDeleteCart = []struct {
	testName    string
	input       int
	outErr      string
	inputDelete int
	errDelete error
	countDelete int
}{
	{
		testName:    "One",
		input:       1,
		outErr:      "",
		inputDelete: 1,
		errDelete: nil,
		countDelete: 1,
	},
}

func TestApplicationDeleteCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperCart(ctrl)
	mRest := mocks.NewMockWrapperRestaurant(ctrl)
	for _, tt := range ApplicationDeleteCart {
		m.
			EXPECT().
			DeleteCart(tt.inputDelete).
			Return(tt.errDelete)

		test := Cart{DB: m, DBRestaurant: mRest}
		t.Run(tt.testName, func(t *testing.T) {
			err := test.DeleteCart(tt.input)
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
