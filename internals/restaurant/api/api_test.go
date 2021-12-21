package api

import (
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant/api/mocks"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
	"testing"
)

var RestaurantHandler = []struct {
	testName                     string
	inputValueReqId              interface{}
	inputValueUnmarshal          []byte
	inputGetOrderHandlerIdCtx    interface{}
	inputGetOrderHandlerId       int
	inputGetOrderHandlerIdOrdCtx interface{}
	inputGetOrderHandlerIdOrd    int
	inputUpdateName              string
	inputGetOrderHandlerDishes   int
	out                          []byte
	inputGetOrderCSRFCtx         interface{}
	inputErrorfArgs              []interface{}
	inputErrorfFormat            string
	countErrorf                  int
	inputWarnfArgs               []interface{}
	inputWarnfFormat             string
	countWarnf                   int
	outGetProfile                *restaurant.AllRestaurantsPromo
	errGetOrder                  error
	countGetOrder                int
}{
	{
		testName:                     "Successful UpdateUserName handler",
		inputValueReqId:              10,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputUpdateName:              "",
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputGetOrderCSRFCtx:         "token",
		out:                          []byte("{\"status\":200,\"body\":{\"restaurants\":null}}"),
		countErrorf:                  0,
		countWarnf:                   0,
		errGetOrder:                  nil,
		countGetOrder:                1,
	},
	{
		testName:          "Error reqId interfaceConvertInt",
		out:               []byte(errPkg.ErrNotStringAndInt),
		inputValueReqId:   nil,
		inputErrorfArgs:   []interface{}{errPkg.ErrNotStringAndInt},
		inputErrorfFormat: "%s",
		countErrorf:       1,
		countWarnf:        0,
		countGetOrder:     0,
	},
	{
		testName:                     "Error checkError-ErrCheck-404",
		out:                          []byte("{\"status\":404,\"explain\":\"restaurants not found\"}"),
		inputGetOrderCSRFCtx:         "token",
		inputValueReqId:              1,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputWarnfArgs:               []interface{}{errPkg.RGetRestaurantsRestaurantsNotFound, 1},
		inputWarnfFormat:             "%s, requestId: %d",
		countErrorf:                  0,
		countWarnf:                   1,
		errGetOrder:                  errors.New(errPkg.RGetRestaurantsRestaurantsNotFound),
		countGetOrder:                1,
	},
	{
		testName:                     "Error checkError-ErrCheck-500",
		out:                          []byte("{\"status\":500,\"explain\":\"database is not responding\"}"),
		inputGetOrderCSRFCtx:         "token",
		inputValueReqId:              1,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputErrorfArgs:              []interface{}{errPkg.RGetRestaurantsRestaurantsNotScan, 1},
		inputErrorfFormat:            "%s, requestId: %d",
		countErrorf:                  1,
		countWarnf:                   0,
		errGetOrder:                  errors.New(errPkg.RGetRestaurantsRestaurantsNotScan),
		countGetOrder:                1,
	},
}

func TestRestaurantHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrlApp := gomock.NewController(t)
	defer ctrl.Finish()
	defer ctrlApp.Finish()

	mockMultilogger := mocks.NewMockMultiLogger(ctrl)
	mockApplication := mocks.NewMockRestaurantApplicationInterface(ctrlApp)
	for _, tt := range RestaurantHandler {
		ctxIn := fasthttp.RequestCtx{}
		ctxIn.SetUserValue("reqId", tt.inputValueReqId)
		ctxIn.SetUserValue("X-Csrf-Token", tt.inputGetOrderCSRFCtx)
		ctxIn.SetUserValue("id", tt.inputGetOrderHandlerIdCtx)
		ctxIn.SetUserValue("idOrd", tt.inputGetOrderHandlerIdOrdCtx)
		ctxIn.Request.SetBody(tt.inputValueUnmarshal)
		ctxExpected := fasthttp.RequestCtx{}
		ctxExpected.Response.SetBody(tt.out)
		mockMultilogger.
			EXPECT().
			Errorf(tt.inputErrorfFormat, tt.inputErrorfArgs).
			Times(tt.countErrorf)

		mockMultilogger.
			EXPECT().
			Warnf(tt.inputWarnfFormat, tt.inputWarnfArgs).
			Times(tt.countWarnf)

		mockApplication.
			EXPECT().
			AllRestaurants().
			Return(tt.outGetProfile, tt.errGetOrder).
			Times(tt.countGetOrder)

		profileInfo := InfoRestaurant{Logger: mockMultilogger, Application: mockApplication}
		t.Run(tt.testName, func(t *testing.T) {
			profileInfo.RestaurantHandler(&ctxIn)
			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))

		})
	}

}

var RestaurantIdHandler = []struct {
	testName                     string
	inputValueReqId              interface{}
	inputValueUnmarshal          []byte
	inputGetOrderHandlerIdCtx    interface{}
	inputGetOrderHandlerId       int
	inputGetOrderHandlerIdOrdCtx interface{}
	inputGetOrderHandlerIdOrd    int
	inputUpdateName              string
	inputGetOrderHandlerDishes   int
	inputRestaurantIdIdResCtx    interface{}
	inputRestaurantIdIdRes       int
	out                          []byte
	inputGetOrderCSRFCtx         interface{}
	inputErrorfArgs              []interface{}
	inputErrorfFormat            string
	countErrorf                  int
	inputWarnfArgs               []interface{}
	inputWarnfFormat             string
	countWarnf                   int
	outGetProfile                *restaurant.RestaurantId
	errGetOrder                  error
	countGetOrder                int
}{
	{
		testName:                     "Successful UpdateUserName handler",
		inputValueReqId:              10,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputRestaurantIdIdResCtx:    1,
		inputRestaurantIdIdRes:       1,
		inputUpdateName:              "",
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputGetOrderCSRFCtx:         "token",
		out:                          []byte("{\"status\":200,\"body\":{\"restaurant\":null}}"),
		countErrorf:                  0,
		countWarnf:                   0,
		errGetOrder:                  nil,
		countGetOrder:                1,
	},
	{
		testName:          "Error reqId interfaceConvertInt",
		out:               []byte(errPkg.ErrNotStringAndInt),
		inputValueReqId:   nil,
		inputErrorfArgs:   []interface{}{errPkg.ErrNotStringAndInt},
		inputErrorfFormat: "%s",
		countErrorf:       1,
		countWarnf:        0,
		countGetOrder:     0,
	},
	{
		testName:          "Error idRes interfaceConvertInt",
		out:               []byte(errPkg.ErrNotStringAndInt),
		inputValueReqId:   1,
		inputErrorfArgs:   []interface{}{errPkg.ErrNotStringAndInt, 1},
		inputErrorfFormat: "%s, requestId: %d",
		countErrorf:       1,
		countWarnf:        0,
		countGetOrder:     0,
	},
	{
		testName:                  "Error idDish interfaceConvertInt",
		out:                       []byte(errPkg.ErrNotStringAndInt),
		inputValueReqId:           1,
		inputRestaurantIdIdResCtx: 1,
		inputErrorfArgs:           []interface{}{errPkg.ErrNotStringAndInt, 1},
		inputErrorfFormat:         "%s, requestId: %d",
		countErrorf:               1,
		countWarnf:                0,
		countGetOrder:             0,
	},
	{
		testName:                     "Error checkError-ErrCheck-500",
		out:                          []byte("{\"status\":500,\"explain\":\"database is not responding\"}"),
		inputGetOrderCSRFCtx:         "token",
		inputValueReqId:              1,
		inputRestaurantIdIdResCtx:    1,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputRestaurantIdIdRes:       1,
		inputGetOrderHandlerIdOrd:    1,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputErrorfArgs:              []interface{}{errPkg.RGetRestaurantRestaurantNotFound, 1},
		inputErrorfFormat:            "%s, requestId: %d",
		countErrorf:                  1,
		countWarnf:                   0,
		errGetOrder:                  errors.New(errPkg.RGetRestaurantRestaurantNotFound),
		countGetOrder:                1,
	},
}

func TestRestaurantIdHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrlApp := gomock.NewController(t)
	defer ctrl.Finish()
	defer ctrlApp.Finish()

	mockMultilogger := mocks.NewMockMultiLogger(ctrl)
	mockApplication := mocks.NewMockRestaurantApplicationInterface(ctrlApp)
	for _, tt := range RestaurantIdHandler {
		ctxIn := fasthttp.RequestCtx{}
		ctxIn.SetUserValue("reqId", tt.inputValueReqId)
		ctxIn.SetUserValue("X-Csrf-Token", tt.inputGetOrderCSRFCtx)
		ctxIn.SetUserValue("id", tt.inputGetOrderHandlerIdCtx)
		ctxIn.SetUserValue("idRes", tt.inputRestaurantIdIdResCtx)
		ctxIn.Request.SetBody(tt.inputValueUnmarshal)
		ctxExpected := fasthttp.RequestCtx{}
		ctxExpected.Response.SetBody(tt.out)
		mockMultilogger.
			EXPECT().
			Errorf(tt.inputErrorfFormat, tt.inputErrorfArgs).
			Times(tt.countErrorf)

		mockMultilogger.
			EXPECT().
			Warnf(tt.inputWarnfFormat, tt.inputWarnfArgs).
			Times(tt.countWarnf)

		mockApplication.
			EXPECT().
			GetRestaurant(tt.inputRestaurantIdIdRes, tt.inputGetOrderHandlerId).
			Return(tt.outGetProfile, tt.errGetOrder).
			Times(tt.countGetOrder)

		profileInfo := InfoRestaurant{Logger: mockMultilogger, Application: mockApplication}
		t.Run(tt.testName, func(t *testing.T) {
			profileInfo.RestaurantIdHandler(&ctxIn)
			println(string(ctxIn.Response.Body()))
			//println(string(ctxExpected.Response.Body()))
			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))

		})
	}

}

var RestaurantDishesHandler = []struct {
	testName                     string
	inputValueReqId              interface{}
	inputValueUnmarshal          []byte
	inputGetOrderHandlerIdCtx    interface{}
	inputGetOrderHandlerId       int
	inputGetOrderHandlerIdOrdCtx interface{}
	inputGetOrderHandlerIdOrd    int
	inputUpdateName              string
	inputGetOrderHandlerDishes   int
	inputRestaurantIdIdResCtx    interface{}
	inputRestaurantIdIdRes       int
	inputRestaurantIdidDishCtx   interface{}
	inputRestaurantIdidDish      int
	inputRestaurantidDishRes     int
	out                          []byte
	inputGetOrderCSRFCtx         interface{}
	inputErrorfArgs              []interface{}
	inputErrorfFormat            string
	countErrorf                  int
	inputWarnfArgs               []interface{}
	inputWarnfFormat             string
	countWarnf                   int
	outGetProfile                *restaurant.Dishes
	errGetOrder                  error
	countGetOrder                int
}{
	{
		testName:                     "Successful UpdateUserName handler",
		inputValueReqId:              10,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputRestaurantIdIdResCtx:    1,
		inputRestaurantIdIdRes:       1,
		inputUpdateName:              "",
		inputRestaurantIdidDishCtx:   1,
		inputRestaurantIdidDish:      1,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputGetOrderCSRFCtx:         "token",
		out:                          []byte("{\"status\":200,\"body\":{\"dishes\":null}}"),
		countErrorf:                  0,
		countWarnf:                   0,
		errGetOrder:                  nil,
		countGetOrder:                1,
	},
	{
		testName:          "Error reqId interfaceConvertInt",
		out:               []byte(errPkg.ErrNotStringAndInt),
		inputValueReqId:   nil,
		inputErrorfArgs:   []interface{}{errPkg.ErrNotStringAndInt},
		inputErrorfFormat: "%s",
		countErrorf:       1,
		countWarnf:        0,
		countGetOrder:     0,
	},
	{
		testName:          "Error idRes interfaceConvertInt",
		out:               []byte(errPkg.ErrNotStringAndInt),
		inputValueReqId:   1,
		inputErrorfArgs:   []interface{}{errPkg.ErrNotStringAndInt, 1},
		inputErrorfFormat: "%s, requestId: %d",
		countErrorf:       1,
		countWarnf:        0,
		countGetOrder:     0,
	},
	{
		testName:                   "Error idDish interfaceConvertInt",
		out:                        []byte("expected type string or int"),
		inputValueReqId:            1,
		inputRestaurantIdIdResCtx:  1,
		inputRestaurantIdidDishCtx: nil,
		inputErrorfArgs:            []interface{}{"expected type string or int", 1},
		inputErrorfFormat:          "%s, requestId: %d",
		countErrorf:                1,
		countWarnf:                 0,
		countGetOrder:              0,
	},
	{
		testName:                     "Error checkError-ErrCheck-500",
		out:                          []byte("{\"status\":500,\"explain\":\"database is not responding\"}"),
		inputGetOrderCSRFCtx:         "token",
		inputValueReqId:              1,
		inputRestaurantIdIdResCtx:    1,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputRestaurantIdIdRes:       1,
		inputGetOrderHandlerIdOrd:    1,
		inputRestaurantIdidDishCtx:   1,
		inputRestaurantIdidDish:      1,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputErrorfArgs:              []interface{}{errPkg.RGetDishesDishesNotFound, 1},
		inputErrorfFormat:            "%s, requestId: %d",
		countErrorf:                  1,
		countWarnf:                   0,
		errGetOrder:                  errors.New(errPkg.RGetDishesDishesNotFound),
		countGetOrder:                1,
	},
}

func TestRestaurantDishesHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrlApp := gomock.NewController(t)
	defer ctrl.Finish()
	defer ctrlApp.Finish()

	mockMultilogger := mocks.NewMockMultiLogger(ctrl)
	mockApplication := mocks.NewMockRestaurantApplicationInterface(ctrlApp)
	for _, tt := range RestaurantDishesHandler {
		ctxIn := fasthttp.RequestCtx{}
		ctxIn.SetUserValue("reqId", tt.inputValueReqId)
		ctxIn.SetUserValue("X-Csrf-Token", tt.inputGetOrderCSRFCtx)
		ctxIn.SetUserValue("id", tt.inputGetOrderHandlerIdCtx)
		ctxIn.SetUserValue("idRes", tt.inputRestaurantIdIdResCtx)
		ctxIn.SetUserValue("idDish", tt.inputRestaurantIdidDishCtx)
		ctxIn.Request.SetBody(tt.inputValueUnmarshal)
		ctxExpected := fasthttp.RequestCtx{}
		ctxExpected.Response.SetBody(tt.out)
		mockMultilogger.
			EXPECT().
			Errorf(tt.inputErrorfFormat, tt.inputErrorfArgs).
			Times(tt.countErrorf)

		mockMultilogger.
			EXPECT().
			Warnf(tt.inputWarnfFormat, tt.inputWarnfArgs).
			Times(tt.countWarnf)

		mockApplication.
			EXPECT().
			RestaurantDishes(tt.inputRestaurantIdIdRes, tt.inputRestaurantIdidDish).
			Return(tt.outGetProfile, tt.errGetOrder).
			Times(tt.countGetOrder)

		profileInfo := InfoRestaurant{Logger: mockMultilogger, Application: mockApplication}
		t.Run(tt.testName, func(t *testing.T) {
			profileInfo.RestaurantDishesHandler(&ctxIn)
			println(string(ctxIn.Response.Body()))
			//println(string(ctxExpected.Response.Body()))
			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))

		})
	}

}

var CreateReviewHandler = []struct {
	testName                     string
	inputValueReqId              interface{}
	inputValueUnmarshal          []byte
	inputGetOrderHandlerIdCtx    interface{}
	inputGetOrderHandlerId       int
	inputGetOrderHandlerIdOrdCtx interface{}
	inputGetOrderHandlerIdOrd    int
	inputUpdateName              string
	inputGetOrderHandlerDishes   int
	inputRestaurantIdIdResCtx    interface{}
	inputRestaurantIdIdRes       int
	inputRestaurantIdidDishCtx   interface{}
	inputRestaurantIdidDish      int
	inputRestaurantidDishRes     int
	out                          []byte
	inputGetOrderCSRFCtx         interface{}
	inputErrorfArgs              []interface{}
	inputErrorfFormat            string
	countErrorf                  int
	inputWarnfArgs               []interface{}
	inputWarnfFormat             string
	countWarnf                   int
	outGetProfile                *restaurant.Dishes
	inputGerProfile              restaurant.NewReview
	errGetOrder                  error
	countGetOrder                int
}{
	{
		testName:                     "Successful UpdateUserName handler",
		inputValueReqId:              10,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputRestaurantIdIdResCtx:    1,
		inputRestaurantIdIdRes:       1,
		inputUpdateName:              "",
		inputRestaurantIdidDishCtx:   1,
		inputRestaurantIdidDish:      1,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputGetOrderCSRFCtx:         "token",
		out:                          []byte("{\"status\":200}"),
		countErrorf:                  0,
		countWarnf:                   0,
		errGetOrder:                  nil,
		countGetOrder:                1,
	},
	{
		testName:          "Error reqId interfaceConvertInt",
		out:               []byte(errPkg.ErrNotStringAndInt),
		inputValueReqId:   nil,
		inputErrorfArgs:   []interface{}{errPkg.ErrNotStringAndInt},
		inputErrorfFormat: "%s",
		countErrorf:       1,
		countWarnf:        0,
		countGetOrder:     0,
	},
	{
		testName:            "Error unmarshal interfaceConvertInt",
		out:                 []byte(errPkg.ErrUnmarshal),
		inputValueReqId:     1,
		inputErrorfArgs:     []interface{}{errPkg.ErrUnmarshal, "parse error: expected string near offset 9 of '1'", 1},
		inputValueUnmarshal: []byte("{\"text\":1}"),
		inputErrorfFormat:   "%s, %s, requestId: %d",
		countErrorf:         1,
		countWarnf:          0,
		countGetOrder:       0,
	},
	{
		testName:            "Error id interfaceConvertInt",
		out:                 []byte(errPkg.ErrNotStringAndInt),
		inputValueUnmarshal: []byte("{\"text\":\"1\"}"),
		inputValueReqId:     1,
		inputErrorfArgs:     []interface{}{errPkg.ErrNotStringAndInt, 1},
		inputErrorfFormat:   "%s, requestId: %d",
		countErrorf:         1,
		countWarnf:          0,
		countGetOrder:       0,
	},
	{
		testName:                   "Error X-CSRF interfaceConvertInt",
		out:                        []byte("expected type string or int"),
		inputValueUnmarshal:        []byte("{\"text\":\"1\"}"),
		inputValueReqId:            1,
		inputRestaurantIdIdResCtx:  1,
		inputGetOrderHandlerId:     1,
		inputGetOrderHandlerIdCtx:  1,
		inputRestaurantIdidDishCtx: nil,
		inputErrorfArgs:            []interface{}{"expected type string or int", 1},
		inputErrorfFormat:          "%s, requestId: %d",
		countErrorf:                1,
		countWarnf:                 0,
		countGetOrder:              0,
	},
	{
		testName:                     "Error checkError-ErrCheck-500",
		out:                          []byte("{\"status\":500,\"explain\":\"database is not responding\"}"),
		inputGetOrderCSRFCtx:         "token",
		inputValueUnmarshal:          []byte("{\"text\":\"1\"}"),
		inputValueReqId:              1,
		inputRestaurantIdIdResCtx:    1,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputRestaurantIdIdRes:       1,
		inputGetOrderHandlerIdOrd:    1,
		inputRestaurantIdidDishCtx:   1,
		inputGerProfile:              restaurant.NewReview{Text: "1"},
		inputRestaurantIdidDish:      1,
		inputErrorfArgs:              []interface{}{errPkg.RGetDishesDishesNotFound, 1},
		inputErrorfFormat:            "%s, requestId: %d",
		countErrorf:                  1,
		countWarnf:                   0,
		errGetOrder:                  errors.New(errPkg.RGetDishesDishesNotFound),
		countGetOrder:                1,
	},
}

func TestCreateReviewHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrlApp := gomock.NewController(t)
	defer ctrl.Finish()
	defer ctrlApp.Finish()

	mockMultilogger := mocks.NewMockMultiLogger(ctrl)
	mockApplication := mocks.NewMockRestaurantApplicationInterface(ctrlApp)
	for _, tt := range CreateReviewHandler {
		ctxIn := fasthttp.RequestCtx{}
		ctxIn.SetUserValue("reqId", tt.inputValueReqId)
		ctxIn.SetUserValue("X-Csrf-Token", tt.inputGetOrderCSRFCtx)
		ctxIn.SetUserValue("id", tt.inputGetOrderHandlerIdCtx)
		ctxIn.SetUserValue("idRes", tt.inputRestaurantIdIdResCtx)
		ctxIn.SetUserValue("idDish", tt.inputRestaurantIdidDishCtx)
		ctxIn.Request.SetBody(tt.inputValueUnmarshal)
		ctxExpected := fasthttp.RequestCtx{}
		ctxExpected.Response.SetBody(tt.out)
		mockMultilogger.
			EXPECT().
			Errorf(tt.inputErrorfFormat, tt.inputErrorfArgs).
			Times(tt.countErrorf)

		mockMultilogger.
			EXPECT().
			Warnf(tt.inputWarnfFormat, tt.inputWarnfArgs).
			Times(tt.countWarnf)

		mockApplication.
			EXPECT().
			CreateReview(tt.inputGetOrderHandlerId, tt.inputGerProfile).
			Return(tt.errGetOrder).
			Times(tt.countGetOrder)

		profileInfo := InfoRestaurant{Logger: mockMultilogger, Application: mockApplication}
		t.Run(tt.testName, func(t *testing.T) {
			profileInfo.CreateReviewHandler(&ctxIn)
			println(string(ctxIn.Response.Body()))
			//println(string(ctxExpected.Response.Body()))
			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))

		})
	}

}

var GetReviewHandler = []struct {
	testName                     string
	inputValueReqId              interface{}
	inputValueUnmarshal          []byte
	inputGetOrderHandlerIdCtx    interface{}
	inputGetOrderHandlerId       int
	inputGetOrderHandlerIdOrdCtx interface{}
	inputGetOrderHandlerIdOrd    int
	inputUpdateName              string
	inputGetOrderHandlerDishes   int
	inputRestaurantIdIdResCtx    interface{}
	inputRestaurantIdIdRes       int
	inputRestaurantIdidDishCtx   interface{}
	inputRestaurantIdidDish      int
	inputRestaurantidDishRes     int
	out                          []byte
	inputGetOrderCSRFCtx         interface{}
	inputErrorfArgs              []interface{}
	inputErrorfFormat            string
	countErrorf                  int
	inputWarnfArgs               []interface{}
	inputWarnfFormat             string
	countWarnf                   int
	outGetProfile                *restaurant.ResReview
	inputGerProfile              restaurant.NewReview
	errGetOrder                  error
	countGetOrder                int
}{
	{
		testName:                     "Successful UpdateUserName handler",
		inputValueReqId:              10,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputGetOrderHandlerIdOrd:    1,
		inputRestaurantIdIdResCtx:    1,
		inputRestaurantIdIdRes:       1,
		inputUpdateName:              "",
		inputRestaurantIdidDishCtx:   1,
		inputRestaurantIdidDish:      1,
		inputValueUnmarshal:          []byte("{\"id\":1}"),
		inputGetOrderCSRFCtx:         "token",
		out:                          []byte("{\"status\":404,\"explain\":\"review is empty\",\"body\":{\"restaurants\":{\"id\":0,\"img\":\"\",\"name\":\"\",\"costFFD\":0,\"minDTime\":0,\"maxDTime\":0,\"rate\":0,\"tags\":null,\"status_favorite\":false}}}"),
		countErrorf:                  0,
		countWarnf:                   0,
		outGetProfile:                &restaurant.ResReview{},
		errGetOrder:                  nil,
		countGetOrder:                1,
	},
	{
		testName:          "Error reqId interfaceConvertInt",
		out:               []byte(errPkg.ErrNotStringAndInt),
		inputValueReqId:   nil,
		inputErrorfArgs:   []interface{}{errPkg.ErrNotStringAndInt},
		inputErrorfFormat: "%s",
		countErrorf:       1,
		countWarnf:        0,
		countGetOrder:     0,
	},
	{
		testName:            "Error id interfaceConvertInt",
		out:                 []byte(errPkg.ErrNotStringAndInt),
		inputValueUnmarshal: []byte("{\"text\":\"1\"}"),
		inputValueReqId:     1,
		inputErrorfArgs:     []interface{}{errPkg.ErrNotStringAndInt, 1},
		inputErrorfFormat:   "%s, requestId: %d",
		countErrorf:         1,
		countWarnf:          0,
		countGetOrder:       0,
	},
	{
		testName:                     "Error checkError-ErrCheck-500",
		out:                          []byte("{\"status\":500,\"explain\":\"database is not responding\"}"),
		inputGetOrderCSRFCtx:         "token",
		inputValueUnmarshal:          []byte("{\"text\":\"1\"}"),
		inputValueReqId:              1,
		inputRestaurantIdIdResCtx:    1,
		inputGetOrderHandlerIdCtx:    1,
		inputGetOrderHandlerId:       1,
		inputGetOrderHandlerIdOrdCtx: 1,
		inputRestaurantIdIdRes:       1,
		inputGetOrderHandlerIdOrd:    1,
		inputRestaurantIdidDishCtx:   1,
		inputGerProfile:              restaurant.NewReview{Text: "1"},
		inputRestaurantIdidDish:      1,
		inputErrorfArgs:              []interface{}{errPkg.RGetDishesDishesNotFound, 1},
		inputErrorfFormat:            "%s, requestId: %d",
		countErrorf:                  1,
		countWarnf:                   0,
		errGetOrder:                  errors.New(errPkg.RGetDishesDishesNotFound),
		countGetOrder:                1,
	},
}

func TestGetReviewHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrlApp := gomock.NewController(t)
	defer ctrl.Finish()
	defer ctrlApp.Finish()

	mockMultilogger := mocks.NewMockMultiLogger(ctrl)
	mockApplication := mocks.NewMockRestaurantApplicationInterface(ctrlApp)
	for _, tt := range GetReviewHandler {
		ctxIn := fasthttp.RequestCtx{}
		ctxIn.SetUserValue("reqId", tt.inputValueReqId)
		ctxIn.SetUserValue("X-Csrf-Token", tt.inputGetOrderCSRFCtx)
		ctxIn.SetUserValue("id", tt.inputGetOrderHandlerIdCtx)
		ctxIn.SetUserValue("idRes", tt.inputRestaurantIdIdResCtx)
		ctxIn.SetUserValue("idDish", tt.inputRestaurantIdidDishCtx)
		ctxIn.Request.SetBody(tt.inputValueUnmarshal)
		ctxExpected := fasthttp.RequestCtx{}
		ctxExpected.Response.SetBody(tt.out)
		mockMultilogger.
			EXPECT().
			Errorf(tt.inputErrorfFormat, tt.inputErrorfArgs).
			Times(tt.countErrorf)

		mockMultilogger.
			EXPECT().
			Warnf(tt.inputWarnfFormat, tt.inputWarnfArgs).
			Times(tt.countWarnf)

		mockApplication.
			EXPECT().
			GetReview(tt.inputRestaurantIdIdRes, tt.inputGetOrderHandlerId).
			Return(tt.outGetProfile, tt.errGetOrder).
			Times(tt.countGetOrder)

		profileInfo := InfoRestaurant{Logger: mockMultilogger, Application: mockApplication}
		t.Run(tt.testName, func(t *testing.T) {
			profileInfo.GetReviewHandler(&ctxIn)
			println(string(ctxIn.Response.Body()))
			//println(string(ctxExpected.Response.Body()))
			require.Equal(t, ctxExpected.Response.Body(), ctxIn.Response.Body(), fmt.Sprintf("Expected: %v\nbut got: %v", ctxExpected.Response.Body(), ctxIn.Response.Body()))

		})
	}

}
