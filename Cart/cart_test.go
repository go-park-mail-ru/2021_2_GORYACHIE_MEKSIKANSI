package Cart

import (
	mocks "2021_2_GORYACHIE_MEKSIKANSI/Test/Mocks"
	_ "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var ApplicationCalculatePriceDelivery = []struct {
	testName    string
	out         int
	outErr      string
	input       int
	inputGetPrice int
	resultGetPrice int
	errDelete   error
}{
	{
		input:       1,
		inputGetPrice: 1,
		resultGetPrice: 1,
		testName:    "One",
		outErr:      "",
		out:         1,
		errDelete:   nil,
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
			Return(tt.resultGetPrice, tt.errDelete)
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
