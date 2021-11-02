package Authorization

import (
	mocks "2021_2_GORYACHIE_MEKSIKANSI/Test/Mocks"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var ApplicationSignUp = []struct {
	testName                string
	out                     *utils.Defense
	outErr                  string
	input                   *utils.RegistrationRequest
	inputSignupClient       *utils.RegistrationRequest
	resultSignupClient      *utils.Defense
	errSignupClient         error
	countSignupClient       int
	inputSignupCourier  int
	resultSignupCourier *utils.Defense
	errSignupCourier    error
	countSignupCourier  int
	inputSignupHost         int
	resultSignupHost        *utils.Defense
	errSignupHost           error
	countSignupHost         int
}{
	{
		input:    &utils.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "client"},
		testName: "One",
		outErr:   "",
		resultSignupClient: &utils.Defense{},
		inputSignupClient:  &utils.RegistrationRequest{Email: "", Phone: "", Password: "", TypeUser: "client"},
		out: &utils.Defense{},
		errSignupClient: nil,
		countSignupClient: 1,
	},
}

func TestApplicationSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperAuthorization(ctrl)
	for _, tt := range ApplicationSignUp {
		m.
			EXPECT().
			SignupClient(tt.inputSignupClient).
			Return(tt.resultSignupClient, tt.errSignupClient).
			Times(tt.countSignupClient)
		m.
			EXPECT().
			SignupCourier(tt.inputSignupCourier).
			Return(tt.resultSignupCourier, tt.errSignupCourier).
			Times(tt.countSignupCourier)
		m.
			EXPECT().
			SignupHost(tt.inputSignupHost).
			Return(tt.resultSignupHost, tt.errSignupHost).
			Times(tt.countSignupHost)
		t.Run(tt.testName, func(t *testing.T) {
			result, err := SignUp(m, tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
