package Profile

import (
	mocks "2021_2_GORYACHIE_MEKSIKANSI/Test/Mocks"
	"2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type Row struct {
}

func (r *Row) Scan(dest ...interface{}) error {
	return nil
}

func TestProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	m.
		EXPECT().
		QueryRow(
			context.Background(),
			"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1", 1,
		).
		Return(&Row{})
	m.
		EXPECT().
		QueryRow(
			context.Background(),
			"SELECT date_birthday FROM client WHERE client_id = $1", 1,
		).
		Return(&Row{})
	testUser := &Wrapper{Conn: m}
	result, _ := testUser.GetProfileClient(1)
	if gomock.Nil().Matches(result) != true {
		//t.Errorf("Not equal\n")
		fmt.Printf("Not equal\n")
	} else {
		fmt.Printf("equal\n")
	}
}

var ApplicationGetProfile= []struct {
	testName string
	input    int
	out      *Utils.Profile
	outErr   string
	inputGetRoleById int
	resultGetRoleById string
	errGetRoleById error
	countGetRoleById int
	inputGetProfileClient int
	resultGetProfileClient *Utils.Profile
	errGetProfileClient error
	countGetProfileClient int
	inputGetProfileCourier int
	resultGetProfileCourier *Utils.Profile
	errGetProfileCourier error
	countGetProfileCourier int
	inputGetProfileHost int
	resultGetProfileHost *Utils.Profile
	errGetProfileHost error
	countGetProfileHost int
}{
	{
		testName: "One",
		out:      nil,
		outErr:   "",
	},
}

func TestApplicationGetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperProfile(ctrl)
	for _, tt := range ApplicationGetProfile {
		m.
			EXPECT().
			GetRoleById(tt.inputGetRoleById).
			Return(tt.resultGetRoleById, tt.errGetRoleById).
			Times(tt.countGetRoleById)
		m.
			EXPECT().
			GetProfileClient(tt.inputGetProfileClient).
			Return(tt.resultGetProfileClient, tt.errGetProfileClient).
			Times(tt.countGetProfileClient)
		m.
			EXPECT().
			GetProfileCourier(tt.inputGetProfileCourier).
			Return(tt.resultGetProfileCourier, tt.errGetProfileCourier).
			Times(tt.countGetProfileCourier)
		m.
			EXPECT().
			GetProfileHost(tt.inputGetProfileHost).
			Return(tt.resultGetProfileHost, tt.errGetProfileHost).
			Times(tt.countGetProfileHost)
		t.Run(tt.testName, func(t *testing.T) {
			result, err := GetProfile(m, tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationUpdateName= []struct {
	testName string
	inputId    int
	inputNewName    string
	outErr   string
	inputUpdateNameId int
	inputUpdateNameNewName string
	errUpdateName error
	countUpdateName int
}{
	{
		testName: "One",
		outErr:   "",
	},
}

func TestApplicationUpdateName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperProfile(ctrl)
	for _, tt := range ApplicationUpdateName {
		m.
			EXPECT().
			UpdateName(tt.inputUpdateNameId, tt.inputUpdateNameNewName).
			Return(tt.errUpdateName).
			Times(tt.countUpdateName)
		t.Run(tt.testName, func(t *testing.T) {
			err := UpdateName(m, tt.inputId, tt.inputNewName)
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationUpdateEmail= []struct {
	testName string
	inputId    int
	inputNewEmail    string
	outErr   string
	inputUpdateEmailId int
	inputUpdateEmailNewEmail string
	errUpdateEmail error
	countUpdateEmail int
}{
	{
		testName: "One",
		outErr:   "",
	},
}

func TestApplicationUpdateEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperProfile(ctrl)
	for _, tt := range ApplicationUpdateEmail {
		m.
			EXPECT().
			UpdateEmail(tt.inputUpdateEmailId, tt.inputUpdateEmailNewEmail).
			Return(tt.errUpdateEmail).
			Times(tt.countUpdateEmail)
		t.Run(tt.testName, func(t *testing.T) {
			err := UpdateEmail(m, tt.inputId, tt.inputNewEmail)
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationUpdatePassword= []struct {
	testName string
	inputId    int
	inputNewPassword    string
	outErr   string
	inputUpdatePasswordId int
	inputUpdatePasswordNewPassword string
	errUpdatePassword error
	countUpdatePassword int
}{
	{
		testName: "One",
		outErr:   "",
	},
}

func TestApplicationUpdatePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperProfile(ctrl)
	for _, tt := range ApplicationUpdatePassword {
		m.
			EXPECT().
			UpdatePassword(tt.inputUpdatePasswordId, tt.inputUpdatePasswordNewPassword).
			Return(tt.errUpdatePassword).
			Times(tt.countUpdatePassword)
		t.Run(tt.testName, func(t *testing.T) {
			err := UpdatePassword(m, tt.inputId, tt.inputNewPassword)
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationUpdatePhone= []struct {
	testName string
	inputId    int
	inputNewPhone   string
	outErr   string
	inputUpdatePhoneId int
	inputUpdatePhoneNewPhone string
	errUpdatePhone error
	countUpdatePhone int
}{
	{
		testName: "One",
		outErr:   "",
	},
}

func TestApplicationUpdatePhone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperProfile(ctrl)
	for _, tt := range ApplicationUpdatePhone {
		m.
			EXPECT().
			UpdatePhone(tt.inputUpdatePhoneId, tt.inputUpdatePhoneNewPhone).
			Return(tt.errUpdatePhone).
			Times(tt.countUpdatePhone)
		t.Run(tt.testName, func(t *testing.T) {
			err := UpdatePhone(m, tt.inputId, tt.inputNewPhone)
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationUpdateAvatar= []struct {
	testName string
	inputId    int
	inputNewAvatar    string
	outErr   string
	inputUpdateAvatarId int
	inputUpdateAvatarNewAvatar string
	errUpdateAvatar error
	countUpdateAvatar int
}{
	{
		testName: "One",
		outErr:   "",
	},
}

func TestApplicationUpdateAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperProfile(ctrl)
	for _, tt := range ApplicationUpdateAvatar {
		m.
			EXPECT().
			UpdateAvatar(tt.inputUpdateAvatarId, tt.inputUpdateAvatarNewAvatar).
			Return(tt.errUpdateAvatar).
			Times(tt.countUpdateAvatar)
		t.Run(tt.testName, func(t *testing.T) {
			err := UpdateAvatar(m, tt.inputId, tt.inputNewAvatar)
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationUpdateBirthday= []struct {
	testName string
	inputId    int
	inputNewBirthday    time.Time
	outErr   string
	inputUpdateBirthdayId int
	inputUpdateBirthdayName string
	errUpdateBirthday error
	countUpdateBirthday int
}{
	{
		testName: "One",
		outErr:   "",
	},
}

func TestApplicationUpdateBirthday(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperProfile(ctrl)
	for _, tt := range ApplicationUpdateBirthday {
		m.
			EXPECT().
			UpdateBirthday(tt.inputUpdateBirthdayId, tt.inputNewBirthday).
			Return(tt.errUpdateBirthday).
			Times(tt.countUpdateBirthday)
		t.Run(tt.testName, func(t *testing.T) {
			err := UpdateBirthday(m, tt.inputId, tt.inputNewBirthday)
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationUpdateAddress= []struct {
	testName string
	inputId    int
	inputNewName    Utils.AddressCoordinates
	outErr   string
	inputUpdateAddressId int
	inputUpdateAddressNewName string
	errUpdateAddress error
	countUpdateAddress int
}{
	{
		testName: "One",
		outErr:   "",
	},
}

func TestApplicationUpdateAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperProfile(ctrl)
	for _, tt := range ApplicationUpdateAddress {
		m.
			EXPECT().
			UpdateAddress(tt.inputUpdateAddressId, tt.inputUpdateAddressNewName).
			Return(tt.errUpdateAddress).
			Times(tt.countUpdateAddress)
		t.Run(tt.testName, func(t *testing.T) {
			err := UpdateAddress(m, tt.inputId, tt.inputNewName)
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
