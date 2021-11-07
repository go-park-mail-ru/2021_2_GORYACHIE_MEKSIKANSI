package Profile

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	mocks "2021_2_GORYACHIE_MEKSIKANSI/Test/Mocks"
	"2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
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
		switch dest[i].(type) {
		case *int:
			*dest[i].(*int) = r.row[i].(int)
		case *string:
			*dest[i].(*string) = r.row[i].(string)
		case *float32:
			*dest[i].(*float32) = float32(r.row[i].(float64))
		}
	}
	return nil
}

var OrmGetRoleById = []struct {
	testName                     string
	input                        int
	rowsQueryHost                Row
	inputQueryHost               int
	rowsQueryClient              Row
	inputQueryClient             int
	rowsQueryCourier             Row
	inputQueryCourier            int
	errQuery                     error
	out                          string
	inputGetDishesRestaurantName string
	inputGetDishesRestaurantId   int
	outErr                       string
	countQueryHost               int
	countQueryClient             int
	countQueryCourier            int
}{
	{
		testName:          "One",
		input:             1,
		inputQueryHost:    1,
		rowsQueryHost:     Row{row: []interface{}{0}},
		rowsQueryClient:   Row{row: []interface{}{1}},
		inputQueryClient:  1,
		rowsQueryCourier:  Row{row: []interface{}{0}},
		inputQueryCourier: 1,
		countQueryCourier: 0,
		countQueryHost:    1,
		countQueryClient:  1,
		errQuery:          nil,
		out:               "client",
		outErr:            errorsConst.RGetGeneralInfoRestaurantNotFound,
	},
}

func TestGetRoleById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmGetRoleById {
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT id FROM host WHERE client_id = $1",
				tt.inputQueryHost,
			).
			Return(&tt.rowsQueryHost).Times(tt.countQueryHost)
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT id FROM client WHERE client_id = $1",
				tt.inputQueryClient,
			).
			Return(&tt.rowsQueryClient).Times(tt.countQueryClient)
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT id FROM courier WHERE client_id = $1",
				tt.inputQueryCourier,
			).
			Return(&tt.rowsQueryCourier).Times(tt.countQueryCourier)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetRoleById(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmGetProfileHost = []struct {
	testName                     string
	input                        int
	rowsQuery                    Row
	inputQuery                   int
	errQuery                     error
	out                          *Utils.Profile
	inputGetDishesRestaurantName string
	inputGetDishesRestaurantId   int
	outErr                       string
}{
	{
		testName:   "One",
		input:      1,
		rowsQuery:  Row{row: []interface{}{"1", "1", "1", "1"}},
		inputQuery: 1,
		errQuery:   nil,
		out:        &Utils.Profile{Name: "1", Email: "1", Phone: "1", Avatar: "1", Birthday: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)},
		outErr:     errorsConst.RGetGeneralInfoRestaurantNotFound,
	},
}

func TestGetProfileHost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmGetProfileHost {
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1",
				tt.inputQuery,
			).
			Return(&tt.rowsQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetProfileHost(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmGetProfileCourier = []struct {
	testName                     string
	input                        int
	rowsQuery                    Row
	inputQuery                   int
	errQuery                     error
	out                          *Utils.Profile
	inputGetDishesRestaurantName string
	inputGetDishesRestaurantId   int
	outErr                       string
}{
	{
		testName:   "One",
		input:      1,
		rowsQuery:  Row{row: []interface{}{"1", "1", "1", "1"}},
		inputQuery: 1,
		errQuery:   nil,
		out:        &Utils.Profile{Name: "1", Email: "1", Phone: "1", Avatar: "1", Birthday: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)},
		outErr:     errorsConst.RGetGeneralInfoRestaurantNotFound,
	},
}

func TestGetProfileCourier(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmGetProfileCourier {
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1",
				tt.inputQuery,
			).
			Return(&tt.rowsQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetProfileCourier(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmGetProfileClient = []struct {
	testName                     string
	input                        int
	rowsQuery                    Row
	inputQuery                   int
	errQuery                     error
	out                          *Utils.Profile
	inputGetDishesRestaurantName string
	inputGetDishesRestaurantId   int
	outErr                       string
	rowsQueryBirthday            Row
	inputQueryBirthday           int
	countQueryBirthday           int
}{
	{
		testName:           "One",
		input:              1,
		rowsQuery:          Row{row: []interface{}{"1", "1", "1", "1"}},
		inputQuery:         1,
		errQuery:           nil,
		out:                &Utils.Profile{Name: "1", Email: "1", Phone: "1", Avatar: "1", Birthday: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)},
		outErr:             errorsConst.RGetGeneralInfoRestaurantNotFound,
		inputQueryBirthday: 1,
		rowsQueryBirthday:  Row{row: []interface{}{time.Now()}},
		countQueryBirthday: 0,
	},
}

func TestGetProfileClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmGetProfileClient {
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT email, name, avatar, phone FROM general_user_info WHERE id = $1",
				tt.inputQuery,
			).
			Return(&tt.rowsQuery)
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT date_birthday FROM client WHERE client_id = $1",
				tt.inputQueryBirthday,
			).
			Return(&tt.rowsQueryBirthday).
			Times(tt.countQueryBirthday)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetProfileClient(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmUpdateName = []struct {
	testName       string
	inputId        int
	inputName      string
	inputQueryId   int
	inputQueryName string
	errQuery       error
	outErr         string
}{
	{
		testName:       "One",
		inputQueryId:   1,
		inputQueryName: "1",
		errQuery:       nil,
		outErr:         "",
		inputId:        1,
		inputName:      "1",
	},
}

func TestUpdateName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmUpdateName {
		m.
			EXPECT().
			Exec(context.Background(),
				"UPDATE general_user_info SET name = $1 WHERE id = $2",
				tt.inputQueryName, tt.inputQueryId,
			).
			Return(nil, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := testUser.UpdateName(tt.inputId, tt.inputName)
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmUpdateEmail = []struct {
	testName        string
	inputId         int
	inputEmail      string
	inputQueryId    int
	inputQueryEmail string
	errQuery        error
	outErr          string
}{
	{
		testName:        "One",
		inputQueryId:    1,
		inputQueryEmail: "1",
		errQuery:        nil,
		outErr:          "",
		inputId:         1,
		inputEmail:      "1",
	},
}

func TestUpdateEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmUpdateEmail {
		m.
			EXPECT().
			Exec(context.Background(),
				"UPDATE general_user_info SET email = $1 WHERE id = $2",
				tt.inputQueryEmail, tt.inputQueryId,
			).
			Return(nil, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := testUser.UpdateEmail(tt.inputId, tt.inputEmail)
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmUpdatePassword = []struct {
	testName           string
	inputId            int
	inputPassword      string
	inputQueryId       int
	inputQuerySalt     int
	inputQueryPassword string
	errQuery           error
	querySalt          Row
	outErr             string
}{
	{
		testName:           "One",
		inputQueryId:       1,
		inputQueryPassword: "4fc82b26aecb47d2868c4efbe3581732a3e7cbcc6c2efb32062c08170a05eeb8",
		inputQuerySalt:     1,
		errQuery:           nil,
		querySalt:          Row{row: []interface{}{"1"}},
		outErr:             "",
		inputId:            1,
		inputPassword:      "1",
	},
}

func TestUpdatePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmUpdatePassword {
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT salt FROM general_user_info WHERE id = $1",
				tt.inputQuerySalt,
			).
			Return(&tt.querySalt)
		m.
			EXPECT().
			Exec(context.Background(),
				"UPDATE general_user_info SET password = $1 WHERE id = $2",
				tt.inputQueryPassword, tt.inputQueryId,
			).
			Return(nil, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := testUser.UpdatePassword(tt.inputId, tt.inputPassword)
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmUpdatePhone = []struct {
	testName        string
	inputId         int
	inputPhone      string
	inputQueryId    int
	inputQueryPhone string
	errQuery        error
	outErr          string
}{
	{
		testName:        "One",
		inputQueryId:    1,
		inputQueryPhone: "1",
		errQuery:        nil,
		outErr:          "",
		inputId:         1,
		inputPhone:      "1",
	},
}

func TestUpdatePhone(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmUpdatePhone {
		m.
			EXPECT().
			Exec(context.Background(),
				"UPDATE general_user_info SET phone = $1 WHERE id = $2",
				tt.inputQueryPhone, tt.inputQueryId,
			).
			Return(nil, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := testUser.UpdatePhone(tt.inputId, tt.inputPhone)
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmUpdateAvatar = []struct {
	testName         string
	inputId          int
	inputAvatar      string
	inputQueryId     int
	inputQueryAvatar string
	errQuery         error
	outErr           string
}{
	{
		testName:         "One",
		inputQueryId:     1,
		inputQueryAvatar: "1",
		errQuery:         nil,
		outErr:           "",
		inputId:          1,
		inputAvatar:      "1",
	},
}

func TestUpdateAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmUpdateAvatar {
		m.
			EXPECT().
			Exec(context.Background(),
				"UPDATE general_user_info SET avatar = $1 WHERE id = $2",
				tt.inputQueryAvatar, tt.inputQueryId,
			).
			Return(nil, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := testUser.UpdateAvatar(tt.inputId, tt.inputAvatar)
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmUpdateBirthday = []struct {
	testName           string
	inputId            int
	inputBirthday      time.Time
	inputQueryId       int
	inputQueryBirthday time.Time
	errQuery           error
	outErr             string
}{
	{
		testName:           "One",
		inputQueryId:       1,
		inputQueryBirthday: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		errQuery:           nil,
		outErr:             "",
		inputId:            1,
		inputBirthday:      time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
	},
}

func TestUpdateBirthday(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmUpdateBirthday {
		m.
			EXPECT().
			Exec(context.Background(),
				"UPDATE client SET date_birthday = $1 WHERE client_id = $2",
				tt.inputQueryBirthday, tt.inputQueryId,
			).
			Return(nil, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := testUser.UpdateBirthday(tt.inputId, tt.inputBirthday)
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var OrmUpdateAddress = []struct {
	testName          string
	inputId           int
	inputAddress      Utils.AddressCoordinates
	inputQueryId      int
	inputQueryAddress Utils.AddressCoordinates
	errQuery          error
	outErr            string
}{
	{
		testName:     "One",
		inputQueryId: 1,
		inputQueryAddress: Utils.AddressCoordinates{Alias: "1", Comment: "1", City: "1", Street: "1", House: "1",
			Floor: 1, Flat: 1, Porch: 1, Intercom: "1", Coordinates: Utils.Coordinates{Latitude: 1.0, Longitude: 1.0}},
		errQuery: nil,
		outErr:   "",
		inputId:  1,
		inputAddress: Utils.AddressCoordinates{Alias: "1", Comment: "1", City: "1", Street: "1", House: "1",
			Floor: 1, Flat: 1, Porch: 1, Intercom: "1", Coordinates: Utils.Coordinates{Latitude: 1.0, Longitude: 1.0}},
	},
}

func TestUpdateAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range OrmUpdateAddress {
		m.
			EXPECT().
			Exec(context.Background(),
				"UPDATE address_user SET alias = $1, comment = $2, city = $3, street = $4, house = $5, floor = $6,"+
					" flat = $7, porch = $8, intercom = $9, latitude = $10, longitude = $11 WHERE client_id = $12",
				tt.inputQueryAddress.Alias, tt.inputQueryAddress.Comment, tt.inputQueryAddress.City,
				tt.inputQueryAddress.Street, tt.inputQueryAddress.House, tt.inputQueryAddress.Floor,
				tt.inputQueryAddress.Flat, tt.inputQueryAddress.Porch, tt.inputQueryAddress.Intercom,
				tt.inputQueryAddress.Coordinates.Latitude, tt.inputQueryAddress.Coordinates.Longitude, tt.inputQueryId,
			).
			Return(nil, tt.errQuery)
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := testUser.UpdateAddress(tt.inputId, tt.inputAddress)
			if tt.outErr != "" && err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationGetProfile = []struct {
	testName                string
	input                   int
	out                     *Utils.Profile
	outErr                  string
	inputGetRoleById        int
	resultGetRoleById       string
	errGetRoleById          error
	inputGetProfileClient   int
	resultGetProfileClient  *Utils.Profile
	errGetProfileClient     error
	countGetProfileClient   int
	inputGetProfileCourier  int
	resultGetProfileCourier *Utils.Profile
	errGetProfileCourier    error
	countGetProfileCourier  int
	inputGetProfileHost     int
	resultGetProfileHost    *Utils.Profile
	errGetProfileHost       error
	countGetProfileHost     int
}{
	{
		testName:                "One",
		input:                   1,
		out:                     &Utils.Profile{Name: "", Email: "", Phone: "", Avatar: "", Birthday: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)},
		outErr:                  "",
		inputGetRoleById:        1,
		resultGetRoleById:       "client",
		errGetRoleById:          nil,
		inputGetProfileClient:   1,
		resultGetProfileClient:  &Utils.Profile{},
		errGetProfileClient:     nil,
		countGetProfileClient:   1,
		inputGetProfileCourier:  1,
		resultGetProfileCourier: &Utils.Profile{},
		errGetProfileCourier:    nil,
		countGetProfileCourier:  0,
		inputGetProfileHost:     0,
		resultGetProfileHost:    &Utils.Profile{},
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
		resultGetProfileClient:  &Utils.Profile{},
		errGetProfileClient:     errors.New("text"),
		countGetProfileClient:   1,
		inputGetProfileCourier:  1,
		resultGetProfileCourier: &Utils.Profile{},
		errGetProfileCourier:    nil,
		countGetProfileCourier:  0,
		inputGetProfileHost:     0,
		resultGetProfileHost:    &Utils.Profile{},
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
		resultGetProfileClient:  &Utils.Profile{},
		errGetProfileClient:     errors.New("text"),
		countGetProfileClient:   1,
		inputGetProfileCourier:  1,
		resultGetProfileCourier: &Utils.Profile{},
		errGetProfileCourier:    nil,
		countGetProfileCourier:  0,
		inputGetProfileHost:     0,
		resultGetProfileHost:    &Utils.Profile{},
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

var ApplicationUpdateAvatar = []struct {
	testName                   string
	inputId                    int
	inputNewAvatar             string
	outErr                     string
	inputUpdateAvatarId        int
	inputUpdateAvatarNewAvatar string
	errUpdateAvatar            error
}{
	{
		testName:                   "One",
		inputId:                    1,
		inputNewAvatar:             "1",
		outErr:                     "",
		inputUpdateAvatarId:        1,
		inputUpdateAvatarNewAvatar: "1",
		errUpdateAvatar:            nil,
	},
	{
		testName:                   "Two",
		inputId:                    1,
		inputNewAvatar:             "1",
		outErr:                     "text",
		inputUpdateAvatarId:        1,
		inputUpdateAvatarNewAvatar: "1",
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

var ApplicationUpdateAddress = []struct {
	testName                     string
	inputId                      int
	inputNewAddress              Utils.AddressCoordinates
	outErr                       string
	inputUpdateAddressId         int
	inputUpdateAddressNewAddress Utils.AddressCoordinates
	errUpdateAddress             error
}{
	{
		testName:                     "One",
		outErr:                       "",
		inputId:                      1,
		inputNewAddress:              Utils.AddressCoordinates{},
		inputUpdateAddressId:         1,
		inputUpdateAddressNewAddress: Utils.AddressCoordinates{},
		errUpdateAddress:             nil,
	},
	{
		testName:                     "Two",
		outErr:                       "text",
		inputId:                      1,
		inputNewAddress:              Utils.AddressCoordinates{},
		inputUpdateAddressId:         1,
		inputUpdateAddressNewAddress: Utils.AddressCoordinates{},
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
