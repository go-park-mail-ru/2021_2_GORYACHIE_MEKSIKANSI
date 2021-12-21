package orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internals/cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/cart/orm/mocks"
	cartProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart/proto"
	promoProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/promocode/proto"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var GetCart = []struct {
	testName   string
	input      int
	out        *cart.ResponseCartErrors
	outErr     string
	inputQuery *cartProto.CartId
	outQuery   *cartProto.ResponseCartErrors
	errQuery   error
}{
	{
		testName: "Get dish",
		input:    1,
		out: &cart.ResponseCartErrors{
			Restaurant: cart.RestaurantIdCastResponse{
				Id:                  1,
				Img:                 "/url/",
				Name:                "Шоколад",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              3.0,
			},
			Dishes: []cart.DishesCartResponse{{
				Id:             1,
				ItemNumber:     0,
				Img:            "/url/",
				Name:           "Орехи",
				Count:          1,
				Cost:           100,
				Kilocalorie:    500,
				Weight:         20,
				Description:    "Очень вкусный шоколад",
				RadiosCart:     nil,
				IngredientCart: nil,
			},
			},
			Cost: cart.CostCartResponse{
				DCost:   250,
				SumCost: 350,
			},
			PromoCode: cart.PromoCode{
				Name:        "Free all",
				Description: "free delivery",
				Code:        "Double Time",
			},
			DishErr: nil,
		},
		outErr: "",
		inputQuery: &cartProto.CartId{
			Id: 1,
		},
		outQuery: &cartProto.ResponseCartErrors{
			Restaurant: &cartProto.RestaurantIdCastResponse{
				Id:                  1,
				Img:                 "/url/",
				Name:                "Шоколад",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              3.0,
			},
			Dishes: []*cartProto.DishesCartResponse{{
				Id:          1,
				ItemNumber:  0,
				Img:         "/url/",
				Name:        "Орехи",
				Count:       1,
				Cost:        100,
				Ccal:        500,
				Weight:      20,
				Description: "Очень вкусный шоколад",
				Radios:      nil,
				Ingredients: nil,
			},
			},
			Cost: &cartProto.CostCartResponse{
				DeliveryCost: 250,
				SumCost:      350,
			},
			PromoCode: &cartProto.PromoCode{
				Name:        "Free all",
				Description: "free delivery",
				Code:        "Double Time",
			},
			DishesErrors: nil,
		},
		errQuery: nil,
	},
	{
		testName: "Error server",
		input:    1,
		out:      nil,
		outErr:   "text",
		inputQuery: &cartProto.CartId{
			Id: 1,
		},
		outQuery: &cartProto.ResponseCartErrors{
			Restaurant: &cartProto.RestaurantIdCastResponse{
				Id:                  1,
				Img:                 "/url/",
				Name:                "Шоколад",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              3.0,
			},
			Dishes: []*cartProto.DishesCartResponse{{
				Id:          1,
				ItemNumber:  0,
				Img:         "/url/",
				Name:        "Орехи",
				Count:       1,
				Cost:        100,
				Ccal:        500,
				Weight:      20,
				Description: "Очень вкусный шоколад",
				Radios:      nil,
				Ingredients: nil,
			},
			},
			Cost: &cartProto.CostCartResponse{
				DeliveryCost: 250,
				SumCost:      350,
			},
			PromoCode: &cartProto.PromoCode{
				Name:        "Free all",
				Description: "free delivery",
				Code:        "Double Time",
			},
			DishesErrors: nil,
			Error:        "",
		},
		errQuery: errors.New("text"),
	},
	{
		testName: "Error get cart",
		input:    1,
		out:      nil,
		outErr:   "text",
		inputQuery: &cartProto.CartId{
			Id: 1,
		},
		outQuery: &cartProto.ResponseCartErrors{
			Restaurant: &cartProto.RestaurantIdCastResponse{
				Id:                  1,
				Img:                 "/url/",
				Name:                "Шоколад",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              3.0,
			},
			Dishes: []*cartProto.DishesCartResponse{{
				Id:          1,
				ItemNumber:  0,
				Img:         "/url/",
				Name:        "Орехи",
				Count:       1,
				Cost:        100,
				Ccal:        500,
				Weight:      20,
				Description: "Очень вкусный шоколад",
				Radios:      nil,
				Ingredients: nil,
			},
			},
			Cost: &cartProto.CostCartResponse{
				DeliveryCost: 250,
				SumCost:      350,
			},
			PromoCode: &cartProto.PromoCode{
				Name:        "Free all",
				Description: "free delivery",
				Code:        "Double Time",
			},
			DishesErrors: nil,
			Error:        "text",
		},
		errQuery: nil,
	},
}

func TestGetCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectCartServiceInterface(ctrl)
	for _, tt := range GetCart {
		m.
			EXPECT().
			GetCart(gomock.Any(), tt.inputQuery).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Wrapper{ConnCart: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetCart(tt.input)
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
	testName    string
	inputDishes cart.RequestCartDefault
	inputClient int
	out         *cart.ResponseCartErrors
	outErr      string
	inputQuery  *cartProto.RequestCartDefault
	outQuery    *cartProto.ResponseCartErrors
	errQuery    error
}{
	{
		testName: "Update cart",
		inputDishes: cart.RequestCartDefault{
			Restaurant: cart.RestaurantRequest{
				Id: 1,
			},
			Dishes: []cart.DishesRequest{{
				Id:         1,
				ItemNumber: 0,
				Count:      1,
			},
			},
			PromoCode: "Free all",
		},
		inputClient: 1,
		out: &cart.ResponseCartErrors{
			Restaurant: cart.RestaurantIdCastResponse{
				Id:                  1,
				Img:                 "/url/",
				Name:                "Шоколад",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              3.0,
			},
			Dishes: []cart.DishesCartResponse{{
				Id:             1,
				ItemNumber:     0,
				Img:            "/url/",
				Name:           "Орехи",
				Count:          1,
				Cost:           100,
				Kilocalorie:    500,
				Weight:         20,
				Description:    "Очень вкусный шоколад",
				RadiosCart:     nil,
				IngredientCart: nil,
			},
			},
			Cost: cart.CostCartResponse{
				DCost:   250,
				SumCost: 350,
			},
			PromoCode: cart.PromoCode{
				Name:        "Free all",
				Description: "free delivery",
				Code:        "Double Time",
			},
			DishErr: nil,
		},
		outErr: "",
		inputQuery: &cartProto.RequestCartDefault{
			Restaurant: &cartProto.RestaurantRequest{
				Id: 1,
			},
			Dishes: []*cartProto.DishesRequest{{
				Id:         1,
				ItemNumber: 0,
				Count:      1,
			},
			},
			PromoCode: "Free all",
			ClientId:  1,
		},
		outQuery: &cartProto.ResponseCartErrors{
			Restaurant: &cartProto.RestaurantIdCastResponse{
				Id:                  1,
				Img:                 "/url/",
				Name:                "Шоколад",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              3.0,
			},
			Dishes: []*cartProto.DishesCartResponse{{
				Id:          1,
				ItemNumber:  0,
				Img:         "/url/",
				Name:        "Орехи",
				Count:       1,
				Cost:        100,
				Ccal:        500,
				Weight:      20,
				Description: "Очень вкусный шоколад",
				Radios:      nil,
				Ingredients: nil,
			},
			},
			Cost: &cartProto.CostCartResponse{
				DeliveryCost: 250,
				SumCost:      350,
			},
			PromoCode: &cartProto.PromoCode{
				Name:        "Free all",
				Description: "free delivery",
				Code:        "Double Time",
			},
			DishesErrors: nil,
			Error:        "",
		},
		errQuery: nil,
	},
	{
		testName: "Error update cart",
		inputDishes: cart.RequestCartDefault{
			Restaurant: cart.RestaurantRequest{
				Id: 1,
			},
			Dishes: []cart.DishesRequest{{
				Id:         1,
				ItemNumber: 0,
				Count:      1,
			},
			},
			PromoCode: "Free all",
		},
		inputClient: 1,
		out:         nil,
		outErr:      "text",
		inputQuery: &cartProto.RequestCartDefault{
			Restaurant: &cartProto.RestaurantRequest{
				Id: 1,
			},
			Dishes: []*cartProto.DishesRequest{{
				Id:         1,
				ItemNumber: 0,
				Count:      1,
			},
			},
			PromoCode: "Free all",
			ClientId:  1,
		},
		outQuery: &cartProto.ResponseCartErrors{
			Restaurant: &cartProto.RestaurantIdCastResponse{
				Id:                  1,
				Img:                 "/url/",
				Name:                "Шоколад",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              3.0,
			},
			Dishes: []*cartProto.DishesCartResponse{{
				Id:          1,
				ItemNumber:  0,
				Img:         "/url/",
				Name:        "Орехи",
				Count:       1,
				Cost:        100,
				Ccal:        500,
				Weight:      20,
				Description: "Очень вкусный шоколад",
				Radios:      nil,
				Ingredients: nil,
			},
			},
			Cost: &cartProto.CostCartResponse{
				DeliveryCost: 250,
				SumCost:      350,
			},
			PromoCode: &cartProto.PromoCode{
				Name:        "Free all",
				Description: "free delivery",
				Code:        "Double Time",
			},
			DishesErrors: nil,
			Error:        "text",
		},
		errQuery: nil,
	},
	{
		testName: "Error server",
		inputDishes: cart.RequestCartDefault{
			Restaurant: cart.RestaurantRequest{
				Id: 1,
			},
			Dishes: []cart.DishesRequest{{
				Id:         1,
				ItemNumber: 0,
				Count:      1,
			},
			},
			PromoCode: "Free all",
		},
		inputClient: 1,
		out:         nil,
		outErr:      "text",
		inputQuery: &cartProto.RequestCartDefault{
			Restaurant: &cartProto.RestaurantRequest{
				Id: 1,
			},
			Dishes: []*cartProto.DishesRequest{{
				Id:         1,
				ItemNumber: 0,
				Count:      1,
			},
			},
			PromoCode: "Free all",
			ClientId:  1,
		},
		outQuery: &cartProto.ResponseCartErrors{
			Restaurant: &cartProto.RestaurantIdCastResponse{
				Id:                  1,
				Img:                 "/url/",
				Name:                "Шоколад",
				CostForFreeDelivery: 250,
				MinDelivery:         15,
				MaxDelivery:         30,
				Rating:              3.0,
			},
			Dishes: []*cartProto.DishesCartResponse{
				{
					Id:          1,
					ItemNumber:  0,
					Img:         "/url/",
					Name:        "Орехи",
					Count:       1,
					Cost:        100,
					Ccal:        500,
					Weight:      20,
					Description: "Очень вкусный шоколад",
					Radios:      nil,
					Ingredients: nil,
				},
			},
			Cost: &cartProto.CostCartResponse{
				DeliveryCost: 250,
				SumCost:      350,
			},
			PromoCode: &cartProto.PromoCode{
				Name:        "Free all",
				Description: "free delivery",
				Code:        "Double Time",
			},
			DishesErrors: nil,
			Error:        "",
		},
		errQuery: errors.New("text"),
	},
	{
		testName: "Restaurant is nil",
		inputDishes: cart.RequestCartDefault{
			Restaurant: cart.RestaurantRequest{
				Id: 1,
			},
			Dishes: []cart.DishesRequest{{
				Id:         1,
				ItemNumber: 0,
				Count:      1,
			},
			},
			PromoCode: "Free all",
		},
		inputClient: 1,
		out:         nil,
		outErr:      "",
		inputQuery: &cartProto.RequestCartDefault{
			Restaurant: &cartProto.RestaurantRequest{
				Id: 1,
			},
			Dishes: []*cartProto.DishesRequest{{
				Id:         1,
				ItemNumber: 0,
				Count:      1,
			},
			},
			PromoCode: "Free all",
			ClientId:  1,
		},
		outQuery: &cartProto.ResponseCartErrors{
			Restaurant: nil,
			Dishes: []*cartProto.DishesCartResponse{{
				Id:          1,
				ItemNumber:  0,
				Img:         "/url/",
				Name:        "Орехи",
				Count:       1,
				Cost:        100,
				Ccal:        500,
				Weight:      20,
				Description: "Очень вкусный шоколад",
				Radios:      nil,
				Ingredients: nil,
			},
			},
			Cost: &cartProto.CostCartResponse{
				DeliveryCost: 250,
				SumCost:      350,
			},
			PromoCode: &cartProto.PromoCode{
				Name:        "Free all",
				Description: "free delivery",
				Code:        "Double Time",
			},
			DishesErrors: nil,
			Error:        "",
		},
		errQuery: nil,
	},
}

func TestUpdateCart(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectCartServiceInterface(ctrl)
	for _, tt := range UpdateCart {
		m.
			EXPECT().
			UpdateCart(gomock.Any(), tt.inputQuery).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Wrapper{ConnCart: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.UpdateCart(tt.inputDishes, tt.inputClient)
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
	testName        string
	inputCode       string
	inputRestaurant int
	inputClient     int
	outErr          string
	inputQuery      *promoProto.PromoCodeWithRestaurantIdAndClient
	outQuery        *promoProto.Error
	errQuery        error
}{
	{
		testName:        "Add promo code",
		inputCode:       "promo",
		inputRestaurant: 1,
		inputClient:     1,
		outErr:          "",
		inputQuery: &promoProto.PromoCodeWithRestaurantIdAndClient{
			PromoCode:  "promo",
			Restaurant: 1,
			Client:     1,
		},
		outQuery: &promoProto.Error{
			Error: "",
		},
		errQuery: nil,
	},
	{
		testName:        "Error add promo code",
		inputCode:       "promo",
		inputRestaurant: 1,
		inputClient:     1,
		outErr:          "text",
		inputQuery: &promoProto.PromoCodeWithRestaurantIdAndClient{
			PromoCode:  "promo",
			Restaurant: 1,
			Client:     1,
		},
		outQuery: &promoProto.Error{
			Error: "text",
		},
		errQuery: nil,
	},
	{
		testName:        "Error microservice",
		inputCode:       "promo",
		inputRestaurant: 1,
		inputClient:     1,
		inputQuery: &promoProto.PromoCodeWithRestaurantIdAndClient{
			PromoCode:  "promo",
			Restaurant: 1,
			Client:     1,
		},
		outErr: "text",
		outQuery: &promoProto.Error{
			Error: "text",
		},
		errQuery: errors.New("text"),
	},
}

func TestAddPromoCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectPromocodeServiceInterface(ctrl)
	for _, tt := range AddPromoCode {
		m.
			EXPECT().
			AddPromoCode(gomock.Any(), tt.inputQuery).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Wrapper{ConnPromo: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := testUser.AddPromoCode(tt.inputCode, tt.inputRestaurant, tt.inputClient)
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
