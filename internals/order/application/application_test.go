package application

import (
	authPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/authorization"
	orderPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/order"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/order/application/mocks"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var GetOrders = []struct {
	testName       string
	input          int
	out            *orderPkg.HistoryOrderArray
	outErr         string
	inputGetOrders int
	outGetOrders   *orderPkg.HistoryOrderArray
	errGetOrders   error
}{
	{
		testName: "First",
		input:    1,
		out: &orderPkg.HistoryOrderArray{
			Orders: []orderPkg.HistoryOrder(nil),
		},
		outErr:         "",
		inputGetOrders: 1,
		outGetOrders: &orderPkg.HistoryOrderArray{
			Orders: []orderPkg.HistoryOrder(nil),
		},
		errGetOrders: nil,
	},
	{
		testName:       "Second",
		outErr:         "text",
		input:          1,
		inputGetOrders: 1,
		outGetOrders:   nil,
		errGetOrders:   errors.New("text"),
	},
}

func TestGetOrders(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperOrderInterface(ctrl)
	for _, tt := range GetOrders {
		m.
			EXPECT().
			GetOrders(tt.inputGetOrders).
			Return(tt.outGetOrders, tt.errGetOrders)
		test := Order{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.GetOrders(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var GetActiveOrder = []struct {
	testName                    string
	inputOrderIdClient          int
	inputOrderIdOrder           int
	out                         *orderPkg.ActiveOrder
	outErr                      string
	inputGetActiveOrderIdClient int
	inputGetActiveOrderIdOrder  int
	outGetActiveOrder           *orderPkg.ActiveOrder
	errGetActiveOrder           error
}{
	{
		testName:                    "First",
		inputOrderIdClient:          1,
		inputOrderIdOrder:           1,
		out:                         &orderPkg.ActiveOrder{},
		outErr:                      "",
		inputGetActiveOrderIdClient: 1,
		inputGetActiveOrderIdOrder:  1,
		outGetActiveOrder:           &orderPkg.ActiveOrder{},
		errGetActiveOrder:           nil,
	},
	{
		testName:                    "Second",
		inputOrderIdClient:          1,
		inputOrderIdOrder:           1,
		outErr:                      "text",
		inputGetActiveOrderIdClient: 1,
		inputGetActiveOrderIdOrder:  1,
		outGetActiveOrder:           nil,
		errGetActiveOrder:           errors.New("text"),
	},
}

func TestGetActiveOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperOrderInterface(ctrl)
	for _, tt := range GetActiveOrder {
		m.
			EXPECT().
			GetOrder(tt.inputGetActiveOrderIdClient, tt.inputGetActiveOrderIdOrder).
			Return(tt.outGetActiveOrder, tt.errGetActiveOrder)
		test := Order{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.GetActiveOrder(tt.inputOrderIdClient, tt.inputOrderIdOrder)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var UpdateStatusOrder = []struct {
	testName                       string
	inputOrderIdOrder              int
	inputOrderStatus               int
	outErr                         string
	outChanId                      int
	outChanStatus                  int
	inputUpdateStatusOrderIdClient int
	inputUpdateStatusOrderIdOrder  int
	errUpdateStatusOrder           error
}{
	{
		testName:                       "First",
		inputOrderIdOrder:              1,
		inputOrderStatus:               1,
		outErr:                         "",
		outChanId:                      1,
		outChanStatus:                  1,
		inputUpdateStatusOrderIdClient: 1,
		inputUpdateStatusOrderIdOrder:  1,
		errUpdateStatusOrder:           nil,
	},
	{
		testName:                       "Second",
		inputOrderIdOrder:              1,
		inputOrderStatus:               1,
		outErr:                         "text",
		outChanId:                      1,
		outChanStatus:                  1,
		inputUpdateStatusOrderIdClient: 1,
		inputUpdateStatusOrderIdOrder:  1,
		errUpdateStatusOrder:           errors.New("text"),
	},
}

func TestUpdateStatusOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperOrderInterface(ctrl)
	for _, tt := range UpdateStatusOrder {
		//m.
		//	EXPECT().
		//	UpdateStatusOrder(tt.inputUpdateStatusOrderIdClient, tt.inputUpdateStatusOrderIdOrder).
		//	Return(tt.errUpdateStatusOrder)
		m.
			EXPECT().
			UpdateStatusOrder(tt.inputUpdateStatusOrderIdClient, gomock.Any()).
			Return(tt.errUpdateStatusOrder).Times(5)
		test := Order{DB: m, IntCh: make(chan authPkg.WebSocketOrder, 10)}
		t.Run(tt.testName, func(t *testing.T) {
			err := test.UpdateStatusOrder(tt.inputOrderIdOrder, tt.inputOrderStatus)
			result := <-test.IntCh
			require.Equal(t, tt.outChanId, result.Id, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outChanId, result.Id))
			require.Equal(t, tt.outChanStatus, result.Status, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outChanStatus, result.Status))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
