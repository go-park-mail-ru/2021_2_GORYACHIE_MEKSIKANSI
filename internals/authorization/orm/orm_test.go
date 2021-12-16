package orm

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/internals/authorization"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/authorization/orm/mocks"
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/authorization/proto"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/util"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	timestamp "google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

var SignUp = []struct {
	testName   string
	input      *auth.RegistrationRequest
	out        *util.Defense
	outErr     string
	inputQuery *authProto.RegistrationRequest
	outQuery   *authProto.DefenseResponse
	errQuery   error
}{
	{
		testName: "Sign up",
		input: &auth.RegistrationRequest{
			TypeUser: "client",
			Name:     "Иванов Иван Иваныч",
			Email:    "root@mail.ru",
			Phone:    "89175554433",
			Password: "43iucthhnoixkh7ldkjvhifskjbgdsmvsbdhsj",
		},
		out: &util.Defense{
			SessionId: "fdsgdsgf",
			CsrfToken: "asdasd",
			DateLife:  time.Date(2006, 2, 1, 0, 0, 0, 0, time.UTC),
		},
		outErr: "",
		inputQuery: &authProto.RegistrationRequest{
			TypeUser: "client",
			Name:     "Иванов Иван Иваныч",
			Email:    "root@mail.ru",
			Phone:    "89175554433",
			Password: "43iucthhnoixkh7ldkjvhifskjbgdsmvsbdhsj",
		},
		outQuery: &authProto.DefenseResponse{
			Defense: &authProto.Defense{
				SessionId:  "fdsgdsgf",
				XCsrfToken: "asdasd",
				DateLife:   timestamp.New(time.Date(2006, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			Error: "",
		},
		errQuery: nil,
	},
	{
		testName: "error application",
		input: &auth.RegistrationRequest{
			TypeUser: "client",
			Name:     "Иванов Иван Иваныч",
			Email:    "root@mail.ru",
			Phone:    "89175554433",
			Password: "43iucthhnoixkh7ldkjvhifskjbgdsmvsbdhsj",
		},
		out:    nil,
		outErr: "text",
		inputQuery: &authProto.RegistrationRequest{
			TypeUser: "client",
			Name:     "Иванов Иван Иваныч",
			Email:    "root@mail.ru",
			Phone:    "89175554433",
			Password: "43iucthhnoixkh7ldkjvhifskjbgdsmvsbdhsj",
		},
		outQuery: &authProto.DefenseResponse{
			Defense: &authProto.Defense{
				SessionId:  "fdsgdsgf",
				XCsrfToken: "asdasd",
				DateLife:   timestamp.New(time.Date(2006, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			Error: "text",
		},
		errQuery: nil,
	},
	{
		testName: "err microserver",
		input: &auth.RegistrationRequest{
			TypeUser: "client",
			Name:     "Иванов Иван Иваныч",
			Email:    "root@mail.ru",
			Phone:    "89175554433",
			Password: "43iucthhnoixkh7ldkjvhifskjbgdsmvsbdhsj",
		},
		out:    nil,
		outErr: "text",
		inputQuery: &authProto.RegistrationRequest{
			TypeUser: "client",
			Name:     "Иванов Иван Иваныч",
			Email:    "root@mail.ru",
			Phone:    "89175554433",
			Password: "43iucthhnoixkh7ldkjvhifskjbgdsmvsbdhsj",
		},
		outQuery: &authProto.DefenseResponse{
			Defense: &authProto.Defense{
				SessionId:  "fdsgdsgf",
				XCsrfToken: "asdasd",
				DateLife:   timestamp.New(time.Date(2006, 2, 1, 0, 0, 0, 0, time.UTC)),
			},
			Error: "",
		},
		errQuery: errors.New("text"),
	},
}

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectAuthServiceInterface(ctrl)
	for _, tt := range SignUp {
		m.
			EXPECT().
			SignUp(gomock.Any(), tt.inputQuery).
			Return(tt.outQuery, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.SignUp(tt.input)
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
