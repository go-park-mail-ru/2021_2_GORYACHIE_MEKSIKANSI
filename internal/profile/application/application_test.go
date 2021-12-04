package application

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var ApplicationGetProfile = []struct {
	testName                string
	input                   int
	out                     *profile.Profile
	outErr                  string
	inputGetRoleById        int
	resultGetRoleById       string
	errGetRoleById          error
	inputGetProfileClient   int
	resultGetProfileClient  *profile.Profile
	errGetProfileClient     error
	countGetProfileClient   int
	inputGetProfileCourier  int
	resultGetProfileCourier *profile.Profile
	errGetProfileCourier    error
	countGetProfileCourier  int
	inputGetProfileHost     int
	resultGetProfileHost    *profile.Profile
	errGetProfileHost       error
	countGetProfileHost     int
}{
	{
		testName: "One",
		input:    1,
		out: &profile.Profile{Name: "", Email: "", Phone: "", Avatar: "",
			Birthday: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)},
		outErr:                  "",
		inputGetRoleById:        1,
		resultGetRoleById:       "client",
		errGetRoleById:          nil,
		inputGetProfileClient:   1,
		resultGetProfileClient:  &profile.Profile{},
		errGetProfileClient:     nil,
		countGetProfileClient:   1,
		inputGetProfileCourier:  1,
		resultGetProfileCourier: &profile.Profile{},
		errGetProfileCourier:    nil,
		countGetProfileCourier:  0,
		inputGetProfileHost:     0,
		resultGetProfileHost:    &profile.Profile{},
		errGetProfileHost:       nil,
		countGetProfileHost:     0,
	},
	{
		testName:                "Two",
		input:                   1,
		out:                     nil,
		outErr:                  "text",
		inputGetRoleById:        1,
		resultGetRoleById:       "client",
		errGetRoleById:          nil,
		inputGetProfileClient:   1,
		resultGetProfileClient:  &profile.Profile{},
		errGetProfileClient:     errors.New("text"),
		countGetProfileClient:   1,
		inputGetProfileCourier:  1,
		resultGetProfileCourier: &profile.Profile{},
		errGetProfileCourier:    nil,
		countGetProfileCourier:  0,
		inputGetProfileHost:     0,
		resultGetProfileHost:    &profile.Profile{},
		errGetProfileHost:       nil,
		countGetProfileHost:     0,
	},
	{
		testName:                "Three",
		input:                   1,
		out:                     nil,
		outErr:                  "text",
		inputGetRoleById:        1,
		resultGetRoleById:       "client",
		errGetRoleById:          nil,
		inputGetProfileClient:   1,
		resultGetProfileClient:  &profile.Profile{},
		errGetProfileClient:     errors.New("text"),
		countGetProfileClient:   1,
		inputGetProfileCourier:  1,
		resultGetProfileCourier: &profile.Profile{},
		errGetProfileCourier:    nil,
		countGetProfileCourier:  0,
		inputGetProfileHost:     0,
		resultGetProfileHost:    &profile.Profile{},
		errGetProfileHost:       nil,
		countGetProfileHost:     0,
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
			Return(tt.resultGetRoleById, tt.errGetRoleById)
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
		test := application.Profile{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.GetProfile(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationUpdateName = []struct {
	testName               string
	inputId                int
	inputNewName           string
	outErr                 string
	inputUpdateNameId      int
	inputUpdateNameNewName string
	errUpdateName          error
}{
	{
		testName:               "One",
		inputId:                1,
		inputNewName:           "1",
		outErr:                 "",
		inputUpdateNameId:      1,
		inputUpdateNameNewName: "1",
		errUpdateName:          nil,
	},
	{
		testName:               "Two",
		inputId:                1,
		inputNewName:           "1",
		outErr:                 "text",
		inputUpdateNameId:      1,
		inputUpdateNameNewName: "1",
		errUpdateName:          errors.New("text"),
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
			Return(tt.errUpdateName)
		test := application.Profile{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := test.UpdateName(tt.inputId, tt.inputNewName)
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationUpdateEmail = []struct {
	testName                 string
	inputId                  int
	inputNewEmail            string
	outErr                   string
	inputUpdateEmailId       int
	inputUpdateEmailNewEmail string
	errUpdateEmail           error
	countUpdateEmail         int
}{
	{
		testName:                 "One",
		inputId:                  1,
		inputNewEmail:            "1",
		outErr:                   "",
		inputUpdateEmailId:       1,
		inputUpdateEmailNewEmail: "1",
		errUpdateEmail:           nil,
	},
	{
		testName:                 "Two",
		inputId:                  1,
		inputNewEmail:            "1",
		outErr:                   "text",
		inputUpdateEmailId:       1,
		inputUpdateEmailNewEmail: "1",
		errUpdateEmail:           errors.New("text"),
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
			Return(tt.errUpdateEmail)
		test := application.Profile{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := test.UpdateEmail(tt.inputId, tt.inputNewEmail)
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationUpdatePassword = []struct {
	testName                       string
	inputId                        int
	inputNewPassword               string
	outErr                         string
	inputUpdatePasswordId          int
	inputUpdatePasswordNewPassword string
	errUpdatePassword              error
}{
	{
		testName:                       "One",
		inputId:                        1,
		inputNewPassword:               "1",
		outErr:                         "",
		inputUpdatePasswordId:          1,
		inputUpdatePasswordNewPassword: "1",
		errUpdatePassword:              nil,
	},
	{
		testName:                       "Two",
		inputId:                        1,
		inputNewPassword:               "1",
		outErr:                         "text",
		inputUpdatePasswordId:          1,
		inputUpdatePasswordNewPassword: "1",
		errUpdatePassword:              errors.New("text"),
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
			Return(tt.errUpdatePassword)
		test := application.Profile{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := test.UpdatePassword(tt.inputId, tt.inputNewPassword)
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationUpdatePhone = []struct {
	testName                 string
	inputId                  int
	inputNewPhone            string
	outErr                   string
	inputUpdatePhoneId       int
	inputUpdatePhoneNewPhone string
	errUpdatePhone           error
}{
	{
		testName:                 "One",
		inputId:                  1,
		inputNewPhone:            "1",
		outErr:                   "",
		inputUpdatePhoneId:       1,
		inputUpdatePhoneNewPhone: "1",
		errUpdatePhone:           nil,
	},
	{
		testName:                 "Two",
		inputId:                  1,
		inputNewPhone:            "1",
		outErr:                   "text",
		inputUpdatePhoneId:       1,
		inputUpdatePhoneNewPhone: "1",
		errUpdatePhone:           errors.New("text"),
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
			Return(tt.errUpdatePhone)
		test := application.Profile{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := test.UpdatePhone(tt.inputId, tt.inputNewPhone)
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationUpdateAvatar = []struct {
	testName                   string
	inputId                    int
	inputNewAvatar             *profile.UpdateAvatar
	outErr                     string
	inputUpdateAvatarId        int
	inputUpdateAvatarNewAvatar *profile.UpdateAvatar
	errUpdateAvatar            error
}{
	{
		testName:                   "One",
		inputId:                    1,
		inputNewAvatar:             &profile.UpdateAvatar{},
		outErr:                     "",
		inputUpdateAvatarId:        1,
		inputUpdateAvatarNewAvatar: &profile.UpdateAvatar{},
		errUpdateAvatar:            nil,
	},
	{
		testName:                   "Two",
		inputId:                    1,
		inputNewAvatar:             &profile.UpdateAvatar{},
		outErr:                     "text",
		inputUpdateAvatarId:        1,
		inputUpdateAvatarNewAvatar: &profile.UpdateAvatar{},
		errUpdateAvatar:            errors.New("text"),
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
			Return(tt.errUpdateAvatar)
		test := application.Profile{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := test.UpdateAvatar(tt.inputId, tt.inputNewAvatar)
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationUpdateBirthday = []struct {
	testName                string
	inputId                 int
	inputNewBirthday        time.Time
	outErr                  string
	inputUpdateBirthdayId   int
	inputUpdateBirthdayDate time.Time
	errUpdateBirthday       error
}{
	{
		testName:                "One",
		inputId:                 1,
		inputNewBirthday:        time.Time{},
		outErr:                  "",
		inputUpdateBirthdayId:   1,
		inputUpdateBirthdayDate: time.Time{},
		errUpdateBirthday:       nil,
	},
	{
		testName:                "Two",
		inputId:                 1,
		inputNewBirthday:        time.Time{},
		outErr:                  "text",
		inputUpdateBirthdayId:   1,
		inputUpdateBirthdayDate: time.Time{},
		errUpdateBirthday:       errors.New("text"),
	},
}

func TestApplicationUpdateBirthday(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperProfile(ctrl)
	for _, tt := range ApplicationUpdateBirthday {
		m.
			EXPECT().
			UpdateBirthday(tt.inputUpdateBirthdayId, tt.inputUpdateBirthdayDate).
			Return(tt.errUpdateBirthday)
		test := application.Profile{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := test.UpdateBirthday(tt.inputId, tt.inputNewBirthday)
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationUpdateAddress = []struct {
	testName                     string
	inputId                      int
	inputNewAddress              profile.AddressCoordinates
	outErr                       string
	inputUpdateAddressId         int
	inputUpdateAddressNewAddress profile.AddressCoordinates
	errUpdateAddress             error
}{
	{
		testName:                     "One",
		outErr:                       "",
		inputId:                      1,
		inputNewAddress:              profile.AddressCoordinates{},
		inputUpdateAddressId:         1,
		inputUpdateAddressNewAddress: profile.AddressCoordinates{},
		errUpdateAddress:             nil,
	},
	{
		testName:                     "Two",
		outErr:                       "text",
		inputId:                      1,
		inputNewAddress:              profile.AddressCoordinates{},
		inputUpdateAddressId:         1,
		inputUpdateAddressNewAddress: profile.AddressCoordinates{},
		errUpdateAddress:             errors.New("text"),
	},
}

func TestApplicationUpdateAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperProfile(ctrl)
	for _, tt := range ApplicationUpdateAddress {
		m.
			EXPECT().
			UpdateAddress(tt.inputUpdateAddressId, tt.inputUpdateAddressNewAddress).
			Return(tt.errUpdateAddress)
		test := application.Profile{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := test.UpdateAddress(tt.inputId, tt.inputNewAddress)
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
