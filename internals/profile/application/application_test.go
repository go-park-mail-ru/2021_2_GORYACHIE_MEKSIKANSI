package application

import (
	profilePkg "2021_2_GORYACHIE_MEKSIKANSI/internals/profile"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/profile/application/mocks"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"mime/multipart"
	"testing"
)

var GetProfile = []struct {
	testName                string
	input                   int
	out                     *profilePkg.Profile
	outErr                  string
	inputGetRoleById        int
	resultGetRoleById       string
	errGetRoleById          error
	inputGetProfileClient   int
	resultGetProfileClient  *profilePkg.Profile
	errGetProfileClient     error
	countGetProfileClient   int
	inputGetProfileCourier  int
	resultGetProfileCourier *profilePkg.Profile
	errGetProfileCourier    error
	countGetProfileCourier  int
	inputGetProfileHost     int
	resultGetProfileHost    *profilePkg.Profile
	errGetProfileHost       error
	countGetProfileHost     int
}{
	{
		testName: "First",
		input:    1,
		out: &profilePkg.Profile{
			Name:     "",
			Email:    "",
			Phone:    "",
			Avatar:   "",
			Birthday: "",
		},
		outErr:                  "",
		inputGetRoleById:        1,
		resultGetRoleById:       "client",
		errGetRoleById:          nil,
		inputGetProfileClient:   1,
		resultGetProfileClient:  &profilePkg.Profile{},
		errGetProfileClient:     nil,
		countGetProfileClient:   1,
		inputGetProfileCourier:  1,
		resultGetProfileCourier: &profilePkg.Profile{},
		errGetProfileCourier:    nil,
		countGetProfileCourier:  0,
		inputGetProfileHost:     0,
		resultGetProfileHost:    &profilePkg.Profile{},
		errGetProfileHost:       nil,
		countGetProfileHost:     0,
	},
	{
		testName:                "Second",
		input:                   1,
		out:                     nil,
		outErr:                  "text",
		inputGetRoleById:        1,
		resultGetRoleById:       "client",
		errGetRoleById:          nil,
		inputGetProfileClient:   1,
		resultGetProfileClient:  &profilePkg.Profile{},
		errGetProfileClient:     errors.New("text"),
		countGetProfileClient:   1,
		inputGetProfileCourier:  1,
		resultGetProfileCourier: &profilePkg.Profile{},
		errGetProfileCourier:    nil,
		countGetProfileCourier:  0,
		inputGetProfileHost:     0,
		resultGetProfileHost:    &profilePkg.Profile{},
		errGetProfileHost:       nil,
		countGetProfileHost:     0,
	},
	{
		testName:                "Third",
		input:                   1,
		out:                     nil,
		outErr:                  "text",
		inputGetRoleById:        1,
		resultGetRoleById:       "client",
		errGetRoleById:          nil,
		inputGetProfileClient:   1,
		resultGetProfileClient:  &profilePkg.Profile{},
		errGetProfileClient:     errors.New("text"),
		countGetProfileClient:   1,
		inputGetProfileCourier:  1,
		resultGetProfileCourier: &profilePkg.Profile{},
		errGetProfileCourier:    nil,
		countGetProfileCourier:  0,
		inputGetProfileHost:     0,
		resultGetProfileHost:    &profilePkg.Profile{},
		errGetProfileHost:       nil,
		countGetProfileHost:     0,
	},
}

func TestGetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperProfileInterface(ctrl)
	for _, tt := range GetProfile {
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
		test := Profile{DB: m}
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

var UpdateName = []struct {
	testName               string
	inputId                int
	inputNewName           string
	outErr                 string
	inputUpdateNameId      int
	inputUpdateNameNewName string
	errUpdateName          error
}{
	{
		testName:               "First",
		inputId:                1,
		inputNewName:           "1",
		outErr:                 "",
		inputUpdateNameId:      1,
		inputUpdateNameNewName: "1",
		errUpdateName:          nil,
	},
	{
		testName:               "Second",
		inputId:                1,
		inputNewName:           "1",
		outErr:                 "text",
		inputUpdateNameId:      1,
		inputUpdateNameNewName: "1",
		errUpdateName:          errors.New("text"),
	},
}

func TestUpdateName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperProfileInterface(ctrl)
	for _, tt := range UpdateName {
		m.
			EXPECT().
			UpdateName(tt.inputUpdateNameId, tt.inputUpdateNameNewName).
			Return(tt.errUpdateName)
		test := Profile{DB: m}
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

var UpdateEmail = []struct {
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
		testName:                 "First",
		inputId:                  1,
		inputNewEmail:            "1",
		outErr:                   "",
		inputUpdateEmailId:       1,
		inputUpdateEmailNewEmail: "1",
		errUpdateEmail:           nil,
	},
	{
		testName:                 "Second",
		inputId:                  1,
		inputNewEmail:            "1",
		outErr:                   "text",
		inputUpdateEmailId:       1,
		inputUpdateEmailNewEmail: "1",
		errUpdateEmail:           errors.New("text"),
	},
}

func TestUpdateEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperProfileInterface(ctrl)
	for _, tt := range UpdateEmail {
		m.
			EXPECT().
			UpdateEmail(tt.inputUpdateEmailId, tt.inputUpdateEmailNewEmail).
			Return(tt.errUpdateEmail)
		test := Profile{DB: m}
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

var UpdatePassword = []struct {
	testName                       string
	inputId                        int
	inputNewPassword               string
	outErr                         string
	inputUpdatePasswordId          int
	inputUpdatePasswordNewPassword string
	errUpdatePassword              error
}{
	{
		testName:                       "First",
		inputId:                        1,
		inputNewPassword:               "1",
		outErr:                         "",
		inputUpdatePasswordId:          1,
		inputUpdatePasswordNewPassword: "1",
		errUpdatePassword:              nil,
	},
	{
		testName:                       "Second",
		inputId:                        1,
		inputNewPassword:               "1",
		outErr:                         "text",
		inputUpdatePasswordId:          1,
		inputUpdatePasswordNewPassword: "1",
		errUpdatePassword:              errors.New("text"),
	},
}

func TestUpdatePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperProfileInterface(ctrl)
	for _, tt := range UpdatePassword {
		m.
			EXPECT().
			UpdatePassword(tt.inputUpdatePasswordId, tt.inputUpdatePasswordNewPassword).
			Return(tt.errUpdatePassword)
		test := Profile{DB: m}
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

var UpdatePhone = []struct {
	testName                 string
	inputId                  int
	inputNewPhone            string
	outErr                   string
	inputUpdatePhoneId       int
	inputUpdatePhoneNewPhone string
	errUpdatePhone           error
}{
	{
		testName:                 "First",
		inputId:                  1,
		inputNewPhone:            "1",
		outErr:                   "",
		inputUpdatePhoneId:       1,
		inputUpdatePhoneNewPhone: "1",
		errUpdatePhone:           nil,
	},
	{
		testName:                 "Second",
		inputId:                  1,
		inputNewPhone:            "1",
		outErr:                   "text",
		inputUpdatePhoneId:       1,
		inputUpdatePhoneNewPhone: "1",
		errUpdatePhone:           errors.New("text"),
	},
}

func TestUpdatePhone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperProfileInterface(ctrl)
	for _, tt := range UpdatePhone {
		m.
			EXPECT().
			UpdatePhone(tt.inputUpdatePhoneId, tt.inputUpdatePhoneNewPhone).
			Return(tt.errUpdatePhone)
		test := Profile{DB: m}
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

var UpdateAvatar = []struct {
	testName            string
	inputId             int
	inputNewAvatar      *profilePkg.UpdateAvatar
	outErr              string
	inputUpdateAvatarId int
	errUpdateAvatar     error
}{
	{
		testName:            "First",
		inputId:             1,
		inputNewAvatar:      &profilePkg.UpdateAvatar{FileHeader: &multipart.FileHeader{Filename: "name_photo.jpeg"}},
		outErr:              "",
		inputUpdateAvatarId: 1,
		errUpdateAvatar:     nil,
	},
	{
		testName:            "Second",
		inputId:             1,
		inputNewAvatar:      &profilePkg.UpdateAvatar{FileHeader: &multipart.FileHeader{Filename: "name_photo.jpeg"}},
		outErr:              "text",
		inputUpdateAvatarId: 1,
		errUpdateAvatar:     errors.New("text"),
	},
}

func TestUpdateAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperProfileInterface(ctrl)
	for _, tt := range UpdateAvatar {
		m.
			EXPECT().
			UpdateAvatar(tt.inputUpdateAvatarId, gomock.Any(), gomock.Any()).
			Return(tt.errUpdateAvatar)
		test := Profile{DB: m}
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

var UpdateBirthday = []struct {
	testName                string
	inputId                 int
	inputNewBirthday        string
	outErr                  string
	inputUpdateBirthdayId   int
	inputUpdateBirthdayDate string
	errUpdateBirthday       error
}{
	{
		testName:                "First",
		inputId:                 1,
		inputNewBirthday:        "02.01.2006",
		outErr:                  "",
		inputUpdateBirthdayId:   1,
		inputUpdateBirthdayDate: "02.01.2006",
		errUpdateBirthday:       nil,
	},
	{
		testName:                "Second",
		inputId:                 1,
		inputNewBirthday:        "02.01.2006",
		outErr:                  "text",
		inputUpdateBirthdayId:   1,
		inputUpdateBirthdayDate: "02.01.2006",
		errUpdateBirthday:       errors.New("text"),
	},
}

func TestUpdateBirthday(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperProfileInterface(ctrl)
	for _, tt := range UpdateBirthday {
		m.
			EXPECT().
			UpdateBirthday(tt.inputUpdateBirthdayId, tt.inputUpdateBirthdayDate).
			Return(tt.errUpdateBirthday)
		test := Profile{DB: m}
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

var UpdateAddress = []struct {
	testName                     string
	inputId                      int
	inputNewAddress              profilePkg.AddressCoordinates
	outErr                       string
	inputUpdateAddressId         int
	inputUpdateAddressNewAddress profilePkg.AddressCoordinates
	errUpdateAddress             error
}{
	{
		testName:                     "First",
		inputId:                      1,
		inputNewAddress:              profilePkg.AddressCoordinates{},
		outErr:                       "",
		inputUpdateAddressId:         1,
		inputUpdateAddressNewAddress: profilePkg.AddressCoordinates{},
		errUpdateAddress:             nil,
	},
	{
		testName:                     "Second",
		inputId:                      1,
		outErr:                       "text",
		inputNewAddress:              profilePkg.AddressCoordinates{},
		inputUpdateAddressId:         1,
		inputUpdateAddressNewAddress: profilePkg.AddressCoordinates{},
		errUpdateAddress:             errors.New("text"),
	},
}

func TestUpdateAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperProfileInterface(ctrl)
	for _, tt := range UpdateAddress {
		m.
			EXPECT().
			UpdateAddress(tt.inputUpdateAddressId, tt.inputUpdateAddressNewAddress).
			Return(tt.errUpdateAddress)
		test := Profile{DB: m}
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
