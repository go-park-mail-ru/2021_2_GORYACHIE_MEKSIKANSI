package orm

import (
	cartPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart/orm/mocks"
	promoProtoPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/promocode/proto"
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

var GetCart = []struct {
	testName                 string
	input                    int
	outOne                   *cartPkg.ResponseCartErrors
	outTwo                   []cartPkg.CastDishesErrs
	outErr                   string
	errBeginTransaction      error
	countRollbackTransaction int
	inputQuery               int
	outQuery                 Rows
	errQuery                 error
	countQuery               int
	errCommitTransaction     error
	countCommitTransaction   int
}{
	{
		testName:                 "First",
		input:                    1,
		outOne:                   nil,
		outTwo:                   nil,
		outErr:                   errPkg.CGetCartCartNotFound,
		errBeginTransaction:      nil,
		countRollbackTransaction: 1,
		inputQuery:               1,
		outQuery:                 Rows{},
		errQuery:                 nil,
		countQuery:               1,
		errCommitTransaction:     nil,
		countCommitTransaction:   0,
	},
	{
		testName: "Second",
		input:    1,
		outOne: &cartPkg.ResponseCartErrors{
			Restaurant: cartPkg.RestaurantIdCastResponse{
				Id:                  1,
				Img:                 "",
				Name:                "",
				CostForFreeDelivery: 0,
				MinDelivery:         0,
				MaxDelivery:         0,
				Rating:              0,
			},
			Dishes: []cartPkg.DishesCartResponse{
				{
					Id:             1,
					ItemNumber:     0,
					Img:            "/address/",
					Name:           "Яблоко",
					Count:          5,
					Cost:           50,
					Kilocalorie:    300,
					Weight:         750,
					Description:    "Очень вкусно",
					RadiosCart:     []cartPkg.RadiosCartResponse(nil),
					IngredientCart: []cartPkg.IngredientCartResponse(nil),
				},
			},
			Cost: cartPkg.CostCartResponse{
				DCost:   0,
				SumCost: 0,
			},
			DishErr: []cartPkg.CastDishesErrs(nil),
		},
		outTwo:                   nil,
		outErr:                   "",
		errBeginTransaction:      nil,
		countRollbackTransaction: 1,
		inputQuery:               1,
		outQuery: Rows{
			row: []interface{}{
				1, 1, 0, "/address/", "Яблоко", 5, 50, 60, 150,
				"Очень вкусно", nil, nil, nil, nil, nil, nil, 1, 1000,
				nil, nil, 0, nil, nil},
			rows: 1,
		},
		errQuery:               nil,
		countQuery:             1,
		errCommitTransaction:   nil,
		countCommitTransaction: 1,
	},
	{
		testName: "Second",
		input:    1,
		outOne: &cartPkg.ResponseCartErrors{
			Restaurant: cartPkg.RestaurantIdCastResponse{
				Id:                  1,
				Img:                 "",
				Name:                "",
				CostForFreeDelivery: 0,
				MinDelivery:         0,
				MaxDelivery:         0,
				Rating:              0,
			},
			Dishes: []cartPkg.DishesCartResponse{
				{
					Id:             1,
					ItemNumber:     0,
					Img:            "/address/",
					Name:           "Яблоко",
					Count:          5,
					Cost:           50,
					Kilocalorie:    300,
					Weight:         750,
					Description:    "Очень вкусно",
					RadiosCart:     []cartPkg.RadiosCartResponse(nil),
					IngredientCart: []cartPkg.IngredientCartResponse(nil),
				},
				{
					Id:          2,
					ItemNumber:  0,
					Img:         "/address/",
					Name:        "Яблоко с ингредиентом",
					Count:       5,
					Cost:        50,
					Kilocalorie: 425,
					Weight:      750,
					Description: "Очень вкусно",
					RadiosCart:  []cartPkg.RadiosCartResponse(nil),
					IngredientCart: []cartPkg.IngredientCartResponse{
						{
							Id:   1,
							Name: "Червяк",
							Cost: 5,
						},
					},
				},
			},
			Cost: cartPkg.CostCartResponse{
				DCost:   0,
				SumCost: 0,
			},
			DishErr: []cartPkg.CastDishesErrs(nil),
		},
		outTwo:                   nil,
		outErr:                   "",
		errBeginTransaction:      nil,
		countRollbackTransaction: 1,
		inputQuery:               1,
		outQuery: Rows{
			row: []interface{}{
				1, 1, 0, "/address/", "Яблоко", 5, 50, 60, 150,
				"Очень вкусно", nil, nil, nil, nil, nil, nil, 1, 1000,
				nil, nil, 0, nil, nil,

				1, 2, 0, "/address/", "Яблоко с ингредиентом", 5, 50, 60, 150,
				"Очень вкусно", nil, nil, nil, "Червяк", 1, 5, 1, 1000,
				nil, 25, 1, nil, 0,
			},
			rows: 2,
		},
		errQuery:               nil,
		countQuery:             1,
		errCommitTransaction:   nil,
		countCommitTransaction: 1,
	},
	{
		testName: "Third",
		input:    1,
		outOne: &cartPkg.ResponseCartErrors{
			Restaurant: cartPkg.RestaurantIdCastResponse{
				Id:                  1,
				Img:                 "",
				Name:                "",
				CostForFreeDelivery: 0,
				MinDelivery:         0,
				MaxDelivery:         0,
				Rating:              0,
			},
			Dishes: []cartPkg.DishesCartResponse{
				{
					Id:             1,
					ItemNumber:     0,
					Img:            "/address/",
					Name:           "Яблоко",
					Count:          5,
					Cost:           50,
					Kilocalorie:    300,
					Weight:         750,
					Description:    "Очень вкусно",
					RadiosCart:     []cartPkg.RadiosCartResponse(nil),
					IngredientCart: []cartPkg.IngredientCartResponse(nil),
				},
				{
					Id:          2,
					ItemNumber:  0,
					Img:         "/address/",
					Name:        "Яблоко с радиусом",
					Count:       5,
					Cost:        50,
					Kilocalorie: 1050,
					Weight:      750,
					Description: "Очень вкусно",
					RadiosCart: []cartPkg.RadiosCartResponse{
						{
							Id:       1,
							Name:     "Червяк",
							RadiosId: 1,
						},
						{
							Id:       2,
							Name:     "Листок",
							RadiosId: 1,
						},
					},
					IngredientCart: []cartPkg.IngredientCartResponse(nil),
				},
			},
			Cost: cartPkg.CostCartResponse{
				DCost:   0,
				SumCost: 0,
			},
			DishErr: []cartPkg.CastDishesErrs(nil),
		},
		outTwo:                   nil,
		outErr:                   "",
		errBeginTransaction:      nil,
		countRollbackTransaction: 1,
		inputQuery:               1,
		outQuery: Rows{
			row: []interface{}{
				1, 1, 0, "/address/", "Яблоко", 5, 50, 60, 150,
				"Очень вкусно", nil, nil, nil, nil, nil, nil, 1, 1000,
				25, nil, 0, nil, nil,

				1, 2, 0, "/address/", "Яблоко с радиусом", 5, 50, 60, 150,
				"Очень вкусно", "Листок", 2, 1, nil, nil, nil, 1, 1000,
				150, nil, 1, 1, nil,

				1, 2, 0, "/address/", "Яблоко с радиусом", 5, 50, 60, 150,
				"Очень вкусно", "Червяк", 1, 1, nil, nil, nil, 1, 1000,
				150, nil, 1, 0, nil,
			},
			rows: 3,
		},
		errQuery:               nil,
		countQuery:             1,
		errCommitTransaction:   nil,
		countCommitTransaction: 1,
	},
	{
		testName: "Fourth",
		input:    1,
		outOne: &cartPkg.ResponseCartErrors{
			Restaurant: cartPkg.RestaurantIdCastResponse{
				Id:                  1,
				Img:                 "",
				Name:                "",
				CostForFreeDelivery: 0,
				MinDelivery:         0,
				MaxDelivery:         0,
				Rating:              0,
			},
			Dishes: []cartPkg.DishesCartResponse{
				{
					Id:             1,
					ItemNumber:     0,
					Img:            "/address/",
					Name:           "Яблоко",
					Count:          5,
					Cost:           50,
					Kilocalorie:    300,
					Weight:         750,
					Description:    "Очень вкусно",
					RadiosCart:     []cartPkg.RadiosCartResponse(nil),
					IngredientCart: []cartPkg.IngredientCartResponse(nil),
				},
				{
					Id:          2,
					ItemNumber:  0,
					Img:         "/address/",
					Name:        "Яблоко с радиусом",
					Count:       5,
					Cost:        50,
					Kilocalorie: 1050,
					Weight:      750,
					Description: "Очень вкусно",
					RadiosCart: []cartPkg.RadiosCartResponse{
						{
							Id:       1,
							Name:     "Червяк",
							RadiosId: 1,
						},
						{
							Id:       2,
							Name:     "Листок",
							RadiosId: 1,
						},
					},
					IngredientCart: []cartPkg.IngredientCartResponse(nil),
				},
				{
					Id:          3,
					ItemNumber:  0,
					Img:         "/address/",
					Name:        "Яблоко с ингредиентом",
					Count:       5,
					Cost:        50,
					Kilocalorie: 425,
					Weight:      750,
					Description: "Очень вкусно",
					RadiosCart:  []cartPkg.RadiosCartResponse(nil),
					IngredientCart: []cartPkg.IngredientCartResponse{
						{
							Id:   1,
							Name: "Червяк",
							Cost: 5,
						},
					},
				},
			},
			Cost: cartPkg.CostCartResponse{
				DCost:   0,
				SumCost: 0,
			},
			DishErr: []cartPkg.CastDishesErrs(nil),
		},
		outTwo:                   nil,
		outErr:                   "",
		errBeginTransaction:      nil,
		countRollbackTransaction: 1,
		inputQuery:               1,
		outQuery: Rows{
			row: []interface{}{
				1, 1, 0, "/address/", "Яблоко", 5, 50, 60, 150,
				"Очень вкусно", nil, nil, nil, nil, nil, nil, 1, 1000,
				25, nil, 0, nil, nil,

				1, 2, 0, "/address/", "Яблоко с радиусом", 5, 50, 60, 150,
				"Очень вкусно", "Листок", 2, 1, nil, nil, nil, 1, 1000,
				150, nil, 1, 1, nil,

				1, 2, 0, "/address/", "Яблоко с радиусом", 5, 50, 60, 150,
				"Очень вкусно", "Червяк", 1, 1, nil, nil, nil, 1, 1000,
				150, nil, 1, 0, nil,

				1, 3, 0, "/address/", "Яблоко с ингредиентом", 5, 50, 60, 150,
				"Очень вкусно", nil, nil, nil, "Червяк", 1, 5, 1, 1000,
				nil, 25, 2, nil, 0,
			},
			rows: 4,
		},
		errQuery:               nil,
		countQuery:             1,
		errCommitTransaction:   nil,
		countCommitTransaction: 1,
	},
	{
		testName:                 "Fifth",
		input:                    1,
		outOne:                   nil,
		outTwo:                   nil,
		outErr:                   errPkg.CGetCartTransactionNotCreate,
		errBeginTransaction:      errors.New("text"),
		countRollbackTransaction: 0,
		inputQuery:               1,
		outQuery:                 Rows{},
		errQuery:                 nil,
		countQuery:               0,
		errCommitTransaction:     nil,
		countCommitTransaction:   0,
	},
	{
		testName:                 "Sixth",
		input:                    1,
		outOne:                   nil,
		outTwo:                   nil,
		outErr:                   errPkg.CGetCartNotSelect,
		errBeginTransaction:      nil,
		countRollbackTransaction: 1,
		inputQuery:               1,
		outQuery:                 Rows{},
		errQuery:                 errors.New("text"),
		countQuery:               1,
		errCommitTransaction:     nil,
		countCommitTransaction:   0,
	},
	{
		testName:                 "Seventh",
		input:                    1,
		outOne:                   nil,
		outTwo:                   nil,
		outErr:                   errPkg.CGetCartNotCommit,
		errBeginTransaction:      nil,
		countRollbackTransaction: 1,
		inputQuery:               1,
		outQuery: Rows{
			row: []interface{}{
				1, 1, 0, "/address/", "Яблоко", 5, 50, 60, 150,
				"Очень вкусно", nil, nil, nil, nil, nil, nil, 1, 1000,
				nil, nil, 0, nil, nil},
			rows: 1,
		},
		errQuery:               nil,
		countQuery:             1,
		errCommitTransaction:   errors.New("text"),
		countCommitTransaction: 1,
	},
}

func TestGetCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range GetCart {
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
				"SELECT cart_food.id, cart_food.food, cart_food.number_item, d.avatar, d.name, cart_food.count_food, d.cost, d.kilocalorie, d.weight,"+
					" d.description, sr.name, sr.id, sr.radios, sd.name, sd.id, sd.cost, d.restaurant, d.count, sr.kilocalorie, sd.kilocalorie,"+
					" cart_food.place, crf.place, csf.place "+
					"FROM public.cart_food "+
					"LEFT JOIN public.dishes d ON d.id = cart_food.food "+
					"LEFT JOIN public.cart_structure_food csf ON csf.client_id = cart_food.client_id and d.id=csf.food and cart_food.id=csf.cart_id "+
					"LEFT JOIN public.structure_dishes sd ON sd.id = csf.checkbox and sd.food=cart_food.food "+
					"LEFT JOIN public.cart_radios_food crf ON crf.client_id = cart_food.client_id and cart_food.id=crf.cart_id "+
					"LEFT JOIN public.structure_radios sr ON sr.id = crf.radios "+
					"WHERE public.cart_food.client_id = $1",
				tt.inputQuery,
			).
			Return(&tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			resultOne, resultTwo, err := testUser.GetCart(tt.input)
			require.Equal(t, tt.outOne, resultOne, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outOne, resultOne))
			require.Equal(t, tt.outTwo, resultTwo, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outTwo, resultTwo))
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

var DeleteCart = []struct {
	testName                 string
	input                    int
	outErr                   string
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
		outErr:                   "",
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
				"DELETE FROM public.cart_food CASCADE WHERE client_id = $1",
				tt.inputDelete,
			).
			Return(nil, tt.errDelete)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := testUser.DeleteCart(tt.input)
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
	testName   string
	input      int
	out        string
	outErr     string
	inputQuery *promoProtoPkg.ClientId
	outQuery   *promoProtoPkg.PromoCodeText
	errQuery   error
	countQuery int
}{
	{
		testName: "Get promo code",
		input:    1,
		out:      "promo",
		outErr:   "",
		inputQuery: &promoProtoPkg.ClientId{
			ClientId: 1,
		},
		outQuery: &promoProtoPkg.PromoCodeText{
			PromoCodeText: "promo",
		},
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName: "Error get promo code",
		input:    1,
		out:      "",
		outErr:   "text",
		inputQuery: &promoProtoPkg.ClientId{
			ClientId: 1,
		},
		outQuery: &promoProtoPkg.PromoCodeText{
			PromoCodeText: "promo",
			Error:         "text",
		},
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName: "Error microservice",
		input:    1,
		out:      "",
		outErr:   "text",
		inputQuery: &promoProtoPkg.ClientId{
			ClientId: 1,
		},
		outQuery: &promoProtoPkg.PromoCodeText{
			PromoCodeText: "promo",
		},
		errQuery:   errors.New("text"),
		countQuery: 1,
	},
}

func TestGetPromoCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectPromoCodeServiceInterface(ctrl)
	for _, tt := range GetPromoCode {
		m.
			EXPECT().
			GetPromoCode(gomock.Any(), tt.inputQuery).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		testUser := &Wrapper{ConnPromoService: m}
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

var UpdateCartStructFood = []struct {
	testName              string
	inputIngredient       []cartPkg.IngredientsCartRequest
	inputClientId         int
	inputCartId           int
	out                   []cartPkg.IngredientCartResponse
	outErr                string
	inputQuery            int
	outQuery              Row
	countQuery            int
	inputInsertIngredient int
	inputInsertClient     int
	errInsert             error
	countInsert           int
	inputInsertFood       int
	inputInsertCartId     int
}{
	{
		testName: "First",
		inputIngredient: []cartPkg.IngredientsCartRequest{
			{
				Id: 1,
			},
		},
		inputClientId: 1,
		inputCartId:   1,
		out: []cartPkg.IngredientCartResponse{
			{
				Name: "1",
				Id:   1,
				Cost: 1,
			},
		},
		outErr:                "",
		inputQuery:            1,
		outQuery:              Row{row: []interface{}{1, "1", 1, 1}},
		countQuery:            1,
		inputInsertIngredient: 1,
		inputInsertClient:     1,
		inputInsertFood:       1,
		inputInsertCartId:     1,
		errInsert:             nil,
		countInsert:           1,
	},
}

func TestUpdateCartStructFood(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range UpdateCartStructFood {
		mTx.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT id, name, cost, food FROM public.structure_dishes WHERE id = $1",
				tt.inputQuery,
			).
			Return(&tt.outQuery).
			Times(tt.countQuery)
		for i := 0; i < tt.countInsert; i++ {
			mTx.
				EXPECT().
				Exec(context.Background(),
					"INSERT INTO public.cart_structure_food (checkbox, client_id, food, cart_id, place) VALUES ($1, $2, $3, $4, $5)",
					tt.inputInsertIngredient, tt.inputInsertClient, tt.inputInsertFood, tt.inputInsertCartId, i,
				).
				Return(nil, tt.errInsert).
				Times(1)
		}
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.updateCartStructFood(tt.inputIngredient, tt.inputClientId, tt.inputCartId, mTx, context.Background())
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

var UpdateCartRadios = []struct {
	testName            string
	inputClientId       int
	inputRadios         []cartPkg.RadiosCartRequest
	inputCartId         int
	out                 []cartPkg.RadiosCartResponse
	outErr              string
	inputQuery          int
	outQuery            Row
	countQuery          int
	inputInsertRadiosId int
	inputInsertRadios   int
	inputInsertClient   int
	inputInsertCartId   int
	errInsert           error
	countInsert         int
}{
	{
		testName:            "First",
		inputRadios:         []cartPkg.RadiosCartRequest{{Id: 1, RadiosId: 1}},
		inputClientId:       1,
		inputCartId:         1,
		out:                 []cartPkg.RadiosCartResponse{{Name: "1", RadiosId: 0, Id: 1}},
		outErr:              "",
		inputQuery:          1,
		outQuery:            Row{row: []interface{}{1, "1"}},
		countQuery:          1,
		inputInsertRadiosId: 1,
		inputInsertRadios:   1,
		inputInsertClient:   1,
		inputInsertCartId:   1,
		errInsert:           nil,
		countInsert:         1,
	},
}

func TestUpdateCartRadios(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range UpdateCartRadios {
		mTx.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT id, name FROM public.structure_radios WHERE id = $1",
				tt.inputQuery,
			).
			Return(&tt.outQuery).
			Times(tt.countQuery)
		for i := 0; i < tt.countInsert; i++ {
			mTx.
				EXPECT().
				Exec(context.Background(),
					"INSERT INTO public.cart_radios_food (radios_id, radios, client_id, cart_id, place) VALUES ($1, $2, $3, $4, $5)",
					tt.inputInsertRadiosId, tt.inputInsertRadios, tt.inputInsertClient, tt.inputInsertCartId, i,
				).
				Return(nil, tt.errInsert).
				Times(1)
		}
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.updateCartRadios(tt.inputRadios, tt.inputClientId, tt.inputCartId, mTx, context.Background())
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
	testName                      string
	inputClientId                 int
	inputCart                     cartPkg.RequestCartDefault
	outOne                        *cartPkg.ResponseCartErrors
	outTwo                        []cartPkg.CastDishesErrs
	outErr                        string
	errBeginTransaction           error
	errCommitTransaction          error
	countCommitTransaction        int
	errRollbackTransaction        error
	countRollbackTransaction      int
	inputSelectDishesIdDish       int
	inputSelectDishesIdRestaurant int
	outSelectDishes               Row
	errSelectDishes               error
	countSelectDishes             int
	inputInsertDishClientId       int
	inputInsertDishFood           int
	inputInsertDishCountFood      int
	inputInsertDishRestaurantId   int
	inputInsertDishNumberItem     int
	inputInsertDishPlace          int
	outInsertDish                 Row
	countInsertDish               int
	inputSelectRadios             int
	outSelectRadios               Row
	errSelectRadios               error
	countSelectRadios             int
	inputInsertRadiosId           int
	inputInsertRadiosRadios       int
	inputInsertRadiosClientId     int
	inputInsertRadiosCartId       int
	inputInsertRadiosPlace        int
	errInsertRadios               error
	countInsertRadios             int
	inputSelectIngredient         int
	outSelectIngredient           Row
	errSelectIngredient           error
	countSelectIngredient         int
	inputInsertIngredientCheckbox int
	inputInsertIngredientClient   int
	inputInsertIngredientFood     int
	inputInsertIngredientCartId   int
	inputInsertIngredientPlace    int
	errInsertIngredient           error
	countInsertIngredient         int
}{
	{
		testName:      "First",
		inputClientId: 1,
		inputCart: cartPkg.RequestCartDefault{
			Restaurant: cartPkg.RestaurantRequest{Id: 1},
			Dishes: []cartPkg.DishesRequest{
				{
					Id:         1,
					ItemNumber: 1,
					Count:      1,
					Radios: []cartPkg.RadiosCartRequest{
						{
							RadiosId: 1,
							Id:       1,
						},
					},
					Ingredients: []cartPkg.IngredientsCartRequest{
						{
							Id: 1,
						},
					},
				},
			},
		},
		outOne: &cartPkg.ResponseCartErrors{
			Restaurant: cartPkg.RestaurantIdCastResponse{
				Id:                  0,
				Img:                 "",
				Name:                "",
				CostForFreeDelivery: 0,
				MinDelivery:         0,
				MaxDelivery:         0,
				Rating:              0,
			},
			Dishes: []cartPkg.DishesCartResponse{
				{
					Id:          1,
					ItemNumber:  0,
					Img:         "1",
					Name:        "1",
					Count:       1,
					Cost:        1,
					Kilocalorie: 1,
					Weight:      1,
					Description: "1",
					RadiosCart: []cartPkg.RadiosCartResponse{
						{
							Name:     "1",
							RadiosId: 0,
							Id:       1,
						},
					},
					IngredientCart: []cartPkg.IngredientCartResponse{
						{
							Name: "1",
							Id:   1,
							Cost: 1,
						},
					},
				},
			},
			Cost: cartPkg.CostCartResponse{
				DCost:   0,
				SumCost: 0,
			},
			DishErr: []cartPkg.CastDishesErrs(nil),
		},
		outTwo:                        []cartPkg.CastDishesErrs(nil),
		outErr:                        "",
		errBeginTransaction:           nil,
		errCommitTransaction:          nil,
		countCommitTransaction:        1,
		errRollbackTransaction:        nil,
		countRollbackTransaction:      1,
		inputSelectDishesIdDish:       1,
		inputSelectDishesIdRestaurant: 1,
		outSelectDishes:               Row{row: []interface{}{1, "1", 1, "1", "1", 1, 1, 1}},
		errSelectDishes:               nil,
		countSelectDishes:             1,
		inputInsertDishClientId:       1,
		inputInsertDishFood:           1,
		inputInsertDishCountFood:      1,
		inputInsertDishRestaurantId:   1,
		inputInsertDishNumberItem:     1,
		inputInsertDishPlace:          0,
		outInsertDish:                 Row{row: []interface{}{1}},
		countInsertDish:               1,
		inputSelectRadios:             1,
		outSelectRadios:               Row{row: []interface{}{1, "1"}},
		errSelectRadios:               nil,
		countSelectRadios:             1,
		inputInsertRadiosId:           1,
		inputInsertRadiosRadios:       1,
		inputInsertRadiosClientId:     1,
		inputInsertRadiosCartId:       1,
		inputInsertRadiosPlace:        0,
		errInsertRadios:               nil,
		countInsertRadios:             1,
		inputSelectIngredient:         1,
		outSelectIngredient:           Row{row: []interface{}{1, "1", 1, 1}},
		errSelectIngredient:           nil,
		countSelectIngredient:         1,
		inputInsertIngredientCheckbox: 1,
		inputInsertIngredientClient:   1,
		inputInsertIngredientFood:     1,
		inputInsertIngredientCartId:   1,
		inputInsertIngredientPlace:    0,
		errInsertIngredient:           nil,
		countInsertIngredient:         1,
	},
}

func TestUpdateCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range UpdateCart {
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
				"SELECT id, avatar, cost, name, description, count, weight, kilocalorie FROM public.dishes WHERE id = $1 AND restaurant = $2",
				tt.inputSelectDishesIdDish, tt.inputSelectDishesIdRestaurant,
			).
			Return(&tt.outSelectDishes).
			Times(1)
		mTx.
			EXPECT().
			QueryRow(context.Background(),
				"INSERT INTO public.cart_food (client_id, food, count_food, restaurant_id, number_item, place) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
				tt.inputInsertDishClientId, tt.inputInsertDishFood, tt.inputInsertDishCountFood,
				tt.inputInsertDishRestaurantId, tt.inputInsertDishNumberItem, tt.inputInsertDishPlace,
			).
			Return(&tt.outInsertDish).
			Times(tt.countInsertDish)
		mTx.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT id, name FROM public.structure_radios WHERE id = $1",
				tt.inputSelectRadios,
			).
			Return(&tt.outSelectRadios).
			Times(tt.countSelectRadios)
		mTx.
			EXPECT().
			Exec(context.Background(),
				"INSERT INTO public.cart_radios_food (radios_id, radios, client_id, cart_id, place) VALUES ($1, $2, $3, $4, $5)",
				tt.inputInsertRadiosId, tt.inputInsertRadiosRadios,
				tt.inputInsertRadiosClientId, tt.inputInsertRadiosCartId, tt.inputInsertRadiosPlace,
			).
			Return(nil, tt.errInsertRadios).
			Times(tt.countInsertRadios)
		mTx.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT id, name, cost, food FROM public.structure_dishes WHERE id = $1",
				tt.inputSelectIngredient,
			).
			Return(&tt.outSelectIngredient).
			Times(tt.countSelectIngredient)
		mTx.
			EXPECT().
			Exec(context.Background(),
				"INSERT INTO public.cart_structure_food (checkbox, client_id, food, cart_id, place) VALUES ($1, $2, $3, $4, $5)",
				tt.inputInsertIngredientCheckbox, tt.inputInsertIngredientClient, tt.inputInsertIngredientFood,
				tt.inputInsertIngredientCartId, tt.inputInsertIngredientPlace,
			).
			Return(nil, tt.errInsertIngredient).
			Times(tt.countSelectIngredient)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			resultOne, resultTwo, err := testUser.UpdateCart(tt.inputCart, tt.inputClientId)
			require.Equal(t, tt.outOne, resultOne, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outOne, resultOne))
			require.Equal(t, tt.outTwo, resultTwo, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outTwo, resultTwo))
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

var GetPriceDelivery = []struct {
	testName                 string
	out                      int
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
		input:                    1,
		inputQuery:               1,
		outQuery:                 Row{row: []interface{}{1}},
		countQuery:               1,
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

func TestGetPriceDelivery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	for _, tt := range GetPriceDelivery {
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
				"SELECT price_delivery FROM public.restaurant WHERE id = $1",
				tt.inputQuery,
			).
			Return(&tt.outQuery).
			Times(tt.countQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetPriceDelivery(tt.input)
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
	out                      *cartPkg.RestaurantId
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
		out: &cartPkg.RestaurantId{
			Id:                  1,
			Img:                 "1",
			Name:                "1",
			CostForFreeDelivery: 1,
			MinDelivery:         1,
			MaxDelivery:         1,
			Rating:              1,
			Tags:                []cartPkg.Tag(nil),
			Menu:                []cartPkg.Menu(nil),
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
				"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM public.restaurant WHERE id = $1",
				tt.inputQuery,
			).
			Return(&tt.outQuery).
			Times(tt.countQuery)
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

var AddPromoCode = []struct {
	testName          string
	outErr            string
	inputQuery        *promoProtoPkg.PromoCodeWithRestaurantIdAndClient
	inputPromoCode    string
	inputRestaurantId int
	inputClientId     int
	outQuery          *promoProtoPkg.Error
	errQuery          error
	countQuery        int
}{
	{
		testName:          "Add promo code",
		inputRestaurantId: 1,
		inputPromoCode:    "promo",
		inputClientId:     1,
		outErr:            "",
		inputQuery: &promoProtoPkg.PromoCodeWithRestaurantIdAndClient{
			Restaurant: 1,
			Client:     1,
			PromoCode:  "promo",
		},
		outQuery:   &promoProtoPkg.Error{},
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName:          "Error add promo code",
		inputRestaurantId: 1,
		inputPromoCode:    "promo",
		inputClientId:     1,
		outErr:            "text",
		inputQuery: &promoProtoPkg.PromoCodeWithRestaurantIdAndClient{
			Restaurant: 1,
			Client:     1,
			PromoCode:  "promo",
		},
		outQuery: &promoProtoPkg.Error{
			Error: "text",
		},
		errQuery:   nil,
		countQuery: 1,
	},
	{
		testName:          "Error microservice",
		inputRestaurantId: 1,
		inputPromoCode:    "promo",
		inputClientId:     1,
		outErr:            "text",
		inputQuery: &promoProtoPkg.PromoCodeWithRestaurantIdAndClient{
			Restaurant: 1,
			Client:     1,
			PromoCode:  "promo",
		},
		outQuery: &promoProtoPkg.Error{
			Error: "",
		},
		errQuery:   errors.New("text"),
		countQuery: 1,
	},
}

func TestAddPromoCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectPromoCodeServiceInterface(ctrl)
	for _, tt := range AddPromoCode {
		m.
			EXPECT().
			AddPromoCode(gomock.Any(), tt.inputQuery).
			Return(tt.outQuery, tt.errQuery).
			Times(tt.countQuery)
		testUser := &Wrapper{ConnPromoService: m}
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

var DoPromoCode = []struct {
	testName                             string
	inputCart                            *cartPkg.ResponseCartErrors
	inputPromoCode                       string
	inputRestaurantId                    int
	out                                  *cartPkg.ResponseCartErrors
	outErr                               string
	inputSelectPromoCodeInfoPromoCode    string
	inputSelectPromoCodeInfoRestaurantId int
	outSelectPromoCodeInfo               Row
	countSelectPromoCodeInfo             int
	inputGetTypePromoCode                *promoProtoPkg.PromoCodeWithRestaurantId
	outGetTypePromoCode                  *promoProtoPkg.TypePromoCodeResponse
	errGetTypePromoCode                  error
	countGetTypePromoCode                int
	errBeginTransaction                  error
	errCommitTransaction                 error
	countCommitTransaction               int
	errRollbackTransaction               error
	countRollbackTransaction             int

	inputActiveFreeDelivery *promoProtoPkg.PromoCodeWithRestaurantId
	outActiveFreeDelivery   *promoProtoPkg.FreeDeliveryResponse
	errActiveFreeDelivery   error
	countActiveFreeDelivery int

	inputActiveCostForSale *promoProtoPkg.PromoCodeWithAmount
	outActiveCostForSale   *promoProtoPkg.NewCostResponse
	errActiveCostForSale   error
	countActiveCostForSale int

	inputActiveTimeForSale *promoProtoPkg.PromoCodeWithAmount
	outActiveTimeForSale   *promoProtoPkg.NewCostResponse
	errActiveTimeForSale   error
	countActiveTimeForSale int

	inputActiveCostForFreeDish *promoProtoPkg.PromoCodeWithRestaurantId
	outActiveCostForFreeDish   *promoProtoPkg.FreeDishResponse
	errActiveCostForFreeDish   error
	countActiveCostForFreeDish int

	inputSelectInfoDish int
	outSelectInfoDish   Row
	countSelectInfoDish int
}{
	{
		testName:          "Promo code on free delivery",
		inputPromoCode:    "promo",
		inputRestaurantId: 1,
		inputCart: &cartPkg.ResponseCartErrors{
			Restaurant: cartPkg.RestaurantIdCastResponse{
				Id:                  0,
				Img:                 "",
				Name:                "",
				CostForFreeDelivery: 0,
				MinDelivery:         0,
				MaxDelivery:         0,
				Rating:              0,
			},
			Dishes: []cartPkg.DishesCartResponse{
				{
					Id:          1,
					ItemNumber:  0,
					Img:         "1",
					Name:        "1",
					Count:       1,
					Cost:        1,
					Kilocalorie: 1,
					Weight:      1,
					Description: "1",
					RadiosCart: []cartPkg.RadiosCartResponse{
						{
							Name:     "1",
							RadiosId: 0,
							Id:       1,
						},
					},
					IngredientCart: []cartPkg.IngredientCartResponse{
						{
							Name: "1",
							Id:   1,
							Cost: 1,
						},
					},
				},
			},
			Cost: cartPkg.CostCartResponse{
				DCost:   100,
				SumCost: 500,
			},
			DishErr: []cartPkg.CastDishesErrs(nil),
		},
		out: &cartPkg.ResponseCartErrors{
			Restaurant: cartPkg.RestaurantIdCastResponse{
				Id:                  0,
				Img:                 "",
				Name:                "",
				CostForFreeDelivery: 0,
				MinDelivery:         0,
				MaxDelivery:         0,
				Rating:              0,
			},
			Dishes: []cartPkg.DishesCartResponse{
				{
					Id:          1,
					ItemNumber:  0,
					Img:         "1",
					Name:        "1",
					Count:       1,
					Cost:        1,
					Kilocalorie: 1,
					Weight:      1,
					Description: "1",
					RadiosCart: []cartPkg.RadiosCartResponse{
						{
							Name:     "1",
							RadiosId: 0,
							Id:       1,
						},
					},
					IngredientCart: []cartPkg.IngredientCartResponse{
						{
							Name: "1",
							Id:   1,
							Cost: 1,
						},
					},
				},
			},
			Cost: cartPkg.CostCartResponse{
				DCost:   0,
				SumCost: 400,
			},
			PromoCode: cartPkg.PromoCode{
				Name:        "Double Time",
				Description: "Description",
				Code:        "promo",
			},
			DishErr: []cartPkg.CastDishesErrs(nil),
		},
		outErr:                               "",
		inputSelectPromoCodeInfoPromoCode:    "promo",
		inputSelectPromoCodeInfoRestaurantId: 1,
		outSelectPromoCodeInfo:               Row{row: []interface{}{"Double Time", "Description"}},
		countSelectPromoCodeInfo:             1,
		inputGetTypePromoCode: &promoProtoPkg.PromoCodeWithRestaurantId{
			PromoCode:  "promo",
			Restaurant: 1,
		},
		outGetTypePromoCode:      &promoProtoPkg.TypePromoCodeResponse{Type: PromoCodeFreeDelivery},
		errGetTypePromoCode:      nil,
		countGetTypePromoCode:    1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,

		inputActiveFreeDelivery: &promoProtoPkg.PromoCodeWithRestaurantId{
			Restaurant: 1,
			PromoCode:  "promo",
		},
		outActiveFreeDelivery: &promoProtoPkg.FreeDeliveryResponse{
			Have: true,
		},
		errActiveFreeDelivery:   nil,
		countActiveFreeDelivery: 1,

		inputActiveCostForSale: &promoProtoPkg.PromoCodeWithAmount{
			Amount:     500,
			PromoCode:  "promo",
			Restaurant: 1,
		},
		outActiveCostForSale:   &promoProtoPkg.NewCostResponse{Cost: 100},
		errActiveCostForSale:   nil,
		countActiveCostForSale: 0,

		inputActiveTimeForSale: &promoProtoPkg.PromoCodeWithAmount{
			Amount:     500,
			PromoCode:  "promo",
			Restaurant: 1,
		},
		outActiveTimeForSale: &promoProtoPkg.NewCostResponse{
			Cost: 20,
		},
		errActiveTimeForSale:   nil,
		countActiveTimeForSale: 0,

		inputActiveCostForFreeDish: &promoProtoPkg.PromoCodeWithRestaurantId{
			PromoCode:  "promo",
			Restaurant: 1,
		},
		outActiveCostForFreeDish: &promoProtoPkg.FreeDishResponse{
			DishId: 1,
		},
		errActiveCostForFreeDish:   nil,
		countActiveCostForFreeDish: 0,

		inputSelectInfoDish: 1,
		outSelectInfoDish:   Row{row: []interface{}{"/url/url/", "Бесплатное кофе", 100, 500, "Очень вкусный и очень бесплатный кофе"}},
		countSelectInfoDish: 0,
	},
	{
		testName:          "Promo code for sale over cost",
		inputPromoCode:    "promo",
		inputRestaurantId: 1,
		inputCart: &cartPkg.ResponseCartErrors{
			Restaurant: cartPkg.RestaurantIdCastResponse{
				Id:                  0,
				Img:                 "",
				Name:                "",
				CostForFreeDelivery: 0,
				MinDelivery:         0,
				MaxDelivery:         0,
				Rating:              0,
			},
			Dishes: []cartPkg.DishesCartResponse{
				{
					Id:          1,
					ItemNumber:  0,
					Img:         "1",
					Name:        "1",
					Count:       1,
					Cost:        1,
					Kilocalorie: 1,
					Weight:      1,
					Description: "1",
					RadiosCart: []cartPkg.RadiosCartResponse{
						{
							Name:     "1",
							RadiosId: 0,
							Id:       1,
						},
					},
					IngredientCart: []cartPkg.IngredientCartResponse{
						{
							Name: "1",
							Id:   1,
							Cost: 1,
						},
					},
				},
			},
			Cost: cartPkg.CostCartResponse{
				DCost:   100,
				SumCost: 500,
			},
			DishErr: []cartPkg.CastDishesErrs(nil),
		},
		out: &cartPkg.ResponseCartErrors{
			Restaurant: cartPkg.RestaurantIdCastResponse{
				Id:                  0,
				Img:                 "",
				Name:                "",
				CostForFreeDelivery: 0,
				MinDelivery:         0,
				MaxDelivery:         0,
				Rating:              0,
			},
			Dishes: []cartPkg.DishesCartResponse{
				{
					Id:          1,
					ItemNumber:  0,
					Img:         "1",
					Name:        "1",
					Count:       1,
					Cost:        1,
					Kilocalorie: 1,
					Weight:      1,
					Description: "1",
					RadiosCart: []cartPkg.RadiosCartResponse{
						{
							Name:     "1",
							RadiosId: 0,
							Id:       1,
						},
					},
					IngredientCart: []cartPkg.IngredientCartResponse{
						{
							Name: "1",
							Id:   1,
							Cost: 1,
						},
					},
				},
			},
			Cost: cartPkg.CostCartResponse{
				DCost:   100,
				SumCost: 100,
			},
			PromoCode: cartPkg.PromoCode{
				Name:        "Double Time",
				Description: "Description",
				Code:        "promo",
			},
			DishErr: []cartPkg.CastDishesErrs(nil),
		},
		outErr:                               "",
		inputSelectPromoCodeInfoPromoCode:    "promo",
		inputSelectPromoCodeInfoRestaurantId: 1,
		outSelectPromoCodeInfo:               Row{row: []interface{}{"Double Time", "Description"}},
		countSelectPromoCodeInfo:             1,
		inputGetTypePromoCode: &promoProtoPkg.PromoCodeWithRestaurantId{
			PromoCode:  "promo",
			Restaurant: 1,
		},
		outGetTypePromoCode:      &promoProtoPkg.TypePromoCodeResponse{Type: PromoCodeSaleOverCost},
		errGetTypePromoCode:      nil,
		countGetTypePromoCode:    1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,

		inputActiveFreeDelivery: &promoProtoPkg.PromoCodeWithRestaurantId{
			Restaurant: 1,
			PromoCode:  "promo",
		},
		outActiveFreeDelivery: &promoProtoPkg.FreeDeliveryResponse{
			Have: false,
		},
		errActiveFreeDelivery:   nil,
		countActiveFreeDelivery: 0,

		inputActiveCostForSale: &promoProtoPkg.PromoCodeWithAmount{
			Amount:     500,
			PromoCode:  "promo",
			Restaurant: 1,
		},
		outActiveCostForSale:   &promoProtoPkg.NewCostResponse{Cost: 100},
		errActiveCostForSale:   nil,
		countActiveCostForSale: 1,

		inputActiveTimeForSale: &promoProtoPkg.PromoCodeWithAmount{
			Amount:     500,
			PromoCode:  "promo",
			Restaurant: 1,
		},
		outActiveTimeForSale: &promoProtoPkg.NewCostResponse{
			Cost: 20,
		},
		errActiveTimeForSale:   nil,
		countActiveTimeForSale: 0,

		inputActiveCostForFreeDish: &promoProtoPkg.PromoCodeWithRestaurantId{
			PromoCode:  "promo",
			Restaurant: 1,
		},
		outActiveCostForFreeDish: &promoProtoPkg.FreeDishResponse{
			DishId: 1,
		},
		errActiveCostForFreeDish:   nil,
		countActiveCostForFreeDish: 0,

		inputSelectInfoDish: 1,
		outSelectInfoDish:   Row{row: []interface{}{"/url/url/", "Бесплатное кофе", 100, 500, "Очень вкусный и очень бесплатный кофе"}},
		countSelectInfoDish: 0,
	},
	{
		testName:          "Promo code on sale over time",
		inputPromoCode:    "promo",
		inputRestaurantId: 1,
		inputCart: &cartPkg.ResponseCartErrors{
			Restaurant: cartPkg.RestaurantIdCastResponse{
				Id:                  0,
				Img:                 "",
				Name:                "",
				CostForFreeDelivery: 0,
				MinDelivery:         0,
				MaxDelivery:         0,
				Rating:              0,
			},
			Dishes: []cartPkg.DishesCartResponse{
				{
					Id:          1,
					ItemNumber:  0,
					Img:         "1",
					Name:        "1",
					Count:       1,
					Cost:        1,
					Kilocalorie: 1,
					Weight:      1,
					Description: "1",
					RadiosCart: []cartPkg.RadiosCartResponse{
						{
							Name:     "1",
							RadiosId: 0,
							Id:       1,
						},
					},
					IngredientCart: []cartPkg.IngredientCartResponse{
						{
							Name: "1",
							Id:   1,
							Cost: 1,
						},
					},
				},
			},
			Cost: cartPkg.CostCartResponse{
				DCost:   100,
				SumCost: 500,
			},
			DishErr: []cartPkg.CastDishesErrs(nil),
		},
		out: &cartPkg.ResponseCartErrors{
			Restaurant: cartPkg.RestaurantIdCastResponse{
				Id:                  0,
				Img:                 "",
				Name:                "",
				CostForFreeDelivery: 0,
				MinDelivery:         0,
				MaxDelivery:         0,
				Rating:              0,
			},
			Dishes: []cartPkg.DishesCartResponse{
				{
					Id:          1,
					ItemNumber:  0,
					Img:         "1",
					Name:        "1",
					Count:       1,
					Cost:        1,
					Kilocalorie: 1,
					Weight:      1,
					Description: "1",
					RadiosCart: []cartPkg.RadiosCartResponse{
						{
							Name:     "1",
							RadiosId: 0,
							Id:       1,
						},
					},
					IngredientCart: []cartPkg.IngredientCartResponse{
						{
							Name: "1",
							Id:   1,
							Cost: 1,
						},
					},
				},
			},
			Cost: cartPkg.CostCartResponse{
				DCost:   100,
				SumCost: 20,
			},
			PromoCode: cartPkg.PromoCode{
				Name:        "Double Time",
				Description: "Description",
				Code:        "promo",
			},
			DishErr: []cartPkg.CastDishesErrs(nil),
		},
		outErr:                               "",
		inputSelectPromoCodeInfoPromoCode:    "promo",
		inputSelectPromoCodeInfoRestaurantId: 1,
		outSelectPromoCodeInfo:               Row{row: []interface{}{"Double Time", "Description"}},
		countSelectPromoCodeInfo:             1,
		inputGetTypePromoCode: &promoProtoPkg.PromoCodeWithRestaurantId{
			PromoCode:  "promo",
			Restaurant: 1,
		},
		outGetTypePromoCode:      &promoProtoPkg.TypePromoCodeResponse{Type: PromoCodeSaleOverTime},
		errGetTypePromoCode:      nil,
		countGetTypePromoCode:    1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,

		inputActiveFreeDelivery: &promoProtoPkg.PromoCodeWithRestaurantId{
			Restaurant: 1,
			PromoCode:  "promo",
		},
		outActiveFreeDelivery: &promoProtoPkg.FreeDeliveryResponse{
			Have: false,
		},
		errActiveFreeDelivery:   nil,
		countActiveFreeDelivery: 0,

		inputActiveCostForSale: &promoProtoPkg.PromoCodeWithAmount{
			Amount:     500,
			PromoCode:  "promo",
			Restaurant: 1,
		},
		outActiveCostForSale:   &promoProtoPkg.NewCostResponse{Cost: 100},
		errActiveCostForSale:   nil,
		countActiveCostForSale: 0,

		inputActiveTimeForSale: &promoProtoPkg.PromoCodeWithAmount{
			Amount:     500,
			PromoCode:  "promo",
			Restaurant: 1,
		},
		outActiveTimeForSale: &promoProtoPkg.NewCostResponse{
			Cost: 20,
		},
		errActiveTimeForSale:   nil,
		countActiveTimeForSale: 1,

		inputActiveCostForFreeDish: &promoProtoPkg.PromoCodeWithRestaurantId{
			PromoCode:  "promo",
			Restaurant: 1,
		},
		outActiveCostForFreeDish: &promoProtoPkg.FreeDishResponse{
			DishId: 1,
		},
		errActiveCostForFreeDish:   nil,
		countActiveCostForFreeDish: 0,

		inputSelectInfoDish: 1,
		outSelectInfoDish:   Row{row: []interface{}{"/url/url/", "Бесплатное кофе", 100, 500, "Очень вкусный и очень бесплатный кофе"}},
		countSelectInfoDish: 0,
	},
	{
		testName:          "Promo code on free dishes",
		inputPromoCode:    "promo",
		inputRestaurantId: 1,
		inputCart: &cartPkg.ResponseCartErrors{
			Restaurant: cartPkg.RestaurantIdCastResponse{
				Id:                  0,
				Img:                 "",
				Name:                "",
				CostForFreeDelivery: 0,
				MinDelivery:         0,
				MaxDelivery:         0,
				Rating:              0,
			},
			Dishes: []cartPkg.DishesCartResponse{
				{
					Id:          1,
					ItemNumber:  0,
					Img:         "1",
					Name:        "1",
					Count:       1,
					Cost:        1,
					Kilocalorie: 1,
					Weight:      1,
					Description: "1",
					RadiosCart: []cartPkg.RadiosCartResponse{
						{
							Name:     "1",
							RadiosId: 0,
							Id:       1,
						},
					},
					IngredientCart: []cartPkg.IngredientCartResponse{
						{
							Name: "1",
							Id:   1,
							Cost: 1,
						},
					},
				},
			},
			Cost: cartPkg.CostCartResponse{
				DCost:   100,
				SumCost: 500,
			},
			DishErr: []cartPkg.CastDishesErrs(nil),
		},
		out: &cartPkg.ResponseCartErrors{
			Restaurant: cartPkg.RestaurantIdCastResponse{
				Id:                  0,
				Img:                 "",
				Name:                "",
				CostForFreeDelivery: 0,
				MinDelivery:         0,
				MaxDelivery:         0,
				Rating:              0,
			},
			Dishes: []cartPkg.DishesCartResponse{
				{
					Id:          1,
					ItemNumber:  0,
					Img:         "1",
					Name:        "1",
					Count:       1,
					Cost:        1,
					Kilocalorie: 1,
					Weight:      1,
					Description: "1",
					RadiosCart: []cartPkg.RadiosCartResponse{
						{
							Name:     "1",
							RadiosId: 0,
							Id:       1,
						},
					},
					IngredientCart: []cartPkg.IngredientCartResponse{
						{
							Name: "1",
							Id:   1,
							Cost: 1,
						},
					},
				},
				{
					Id:             1,
					ItemNumber:     0,
					Img:            "/url/url/",
					Name:           "Бесплатное кофе",
					Count:          1,
					Cost:           0,
					Kilocalorie:    100,
					Weight:         500,
					Description:    "Очень вкусный и очень бесплатный кофе",
					RadiosCart:     []cartPkg.RadiosCartResponse{},
					IngredientCart: []cartPkg.IngredientCartResponse{},
				},
			},
			Cost: cartPkg.CostCartResponse{
				DCost:   100,
				SumCost: 500,
			},
			PromoCode: cartPkg.PromoCode{
				Name:        "Double Time",
				Description: "Description",
				Code:        "promo",
			},
			DishErr: []cartPkg.CastDishesErrs(nil),
		},
		outErr:                               "",
		inputSelectPromoCodeInfoPromoCode:    "promo",
		inputSelectPromoCodeInfoRestaurantId: 1,
		outSelectPromoCodeInfo:               Row{row: []interface{}{"Double Time", "Description"}},
		countSelectPromoCodeInfo:             1,
		inputGetTypePromoCode: &promoProtoPkg.PromoCodeWithRestaurantId{
			PromoCode:  "promo",
			Restaurant: 1,
		},
		outGetTypePromoCode:      &promoProtoPkg.TypePromoCodeResponse{Type: PromoCodeFreeDishes},
		errGetTypePromoCode:      nil,
		countGetTypePromoCode:    1,
		errBeginTransaction:      nil,
		errCommitTransaction:     nil,
		countCommitTransaction:   1,
		errRollbackTransaction:   nil,
		countRollbackTransaction: 1,

		inputActiveFreeDelivery: &promoProtoPkg.PromoCodeWithRestaurantId{
			Restaurant: 1,
			PromoCode:  "promo",
		},
		outActiveFreeDelivery: &promoProtoPkg.FreeDeliveryResponse{
			Have: false,
		},
		errActiveFreeDelivery:   nil,
		countActiveFreeDelivery: 0,

		inputActiveCostForSale: &promoProtoPkg.PromoCodeWithAmount{
			Amount:     400,
			PromoCode:  "promo",
			Restaurant: 1,
		},
		outActiveCostForSale:   &promoProtoPkg.NewCostResponse{Cost: 100},
		errActiveCostForSale:   nil,
		countActiveCostForSale: 0,

		inputActiveTimeForSale: &promoProtoPkg.PromoCodeWithAmount{
			Amount:     500,
			PromoCode:  "promo",
			Restaurant: 1,
		},
		outActiveTimeForSale: &promoProtoPkg.NewCostResponse{
			Cost: 20,
		},
		errActiveTimeForSale:   nil,
		countActiveTimeForSale: 0,

		inputActiveCostForFreeDish: &promoProtoPkg.PromoCodeWithRestaurantId{
			PromoCode:  "promo",
			Restaurant: 1,
		},
		outActiveCostForFreeDish: &promoProtoPkg.FreeDishResponse{
			DishId: 1,
		},
		errActiveCostForFreeDish:   nil,
		countActiveCostForFreeDish: 1,

		inputSelectInfoDish: 1,
		outSelectInfoDish:   Row{row: []interface{}{"/url/url/", "Бесплатное кофе", 100, 500, "Очень вкусный и очень бесплатный кофе"}},
		countSelectInfoDish: 1,
	},
}

func TestDoPromoCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mTx := mocks.NewMockTransactionInterface(ctrl)
	mPromo := mocks.NewMockConnectPromoCodeServiceInterface(ctrl)
	for _, tt := range DoPromoCode {
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
				"SELECT name, description FROM public.promocode WHERE code = $1 AND restaurant = $2",
				tt.inputSelectPromoCodeInfoPromoCode, tt.inputSelectPromoCodeInfoRestaurantId,
			).
			Return(&tt.outSelectPromoCodeInfo).
			Times(tt.countSelectPromoCodeInfo)
		mPromo.
			EXPECT().
			GetTypePromoCode(gomock.Any(), tt.inputGetTypePromoCode).
			Return(tt.outGetTypePromoCode, tt.errGetTypePromoCode).
			Times(tt.countGetTypePromoCode)

		mPromo.
			EXPECT().
			ActiveFreeDelivery(gomock.Any(), tt.inputActiveFreeDelivery).
			Return(tt.outActiveFreeDelivery, tt.errActiveFreeDelivery).
			Times(tt.countActiveFreeDelivery)

		mPromo.
			EXPECT().
			ActiveCostForSale(gomock.Any(), tt.inputActiveCostForSale).
			Return(tt.outActiveCostForSale, tt.errActiveCostForSale).
			Times(tt.countActiveCostForSale)

		mPromo.
			EXPECT().
			ActiveTimeForSale(gomock.Any(), tt.inputActiveTimeForSale).
			Return(tt.outActiveTimeForSale, tt.errActiveTimeForSale).
			Times(tt.countActiveTimeForSale)

		mPromo.
			EXPECT().
			ActiveCostForFreeDish(gomock.Any(), tt.inputActiveCostForFreeDish).
			Return(tt.outActiveCostForFreeDish, tt.errActiveCostForFreeDish).
			Times(tt.countActiveCostForFreeDish)
		mTx.
			EXPECT().
			QueryRow(gomock.Any(),
				"SELECT avatar, name, kilocalorie, weight, description FROM public.dishes WHERE id = $1 AND count > 1",
				tt.inputSelectInfoDish).
			Return(&tt.outSelectInfoDish).
			Times(tt.countSelectInfoDish)
		testUser := &Wrapper{Conn: m, Ctx: context.Background(), ConnPromoService: mPromo}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.DoPromoCode(tt.inputPromoCode, tt.inputRestaurantId, tt.inputCart)
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
