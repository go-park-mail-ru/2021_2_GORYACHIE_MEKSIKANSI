package orm

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/profile"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/profile/orm/mocks"
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"mime/multipart"
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

var GetRoleById = []struct {
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
		outErr:            errorsConst.RGetRestaurantRestaurantNotFound,
	},
}

func TestGetRoleById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range GetRoleById {
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

var GetProfileHost = []struct {
	testName                     string
	input                        int
	rowsQuery                    Row
	inputQuery                   int
	errQuery                     error
	out                          *profile.Profile
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
		out: &profile.Profile{Name: "1", Email: "1", Phone: "1", Avatar: "1",
			Birthday: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)},
		outErr: errorsConst.RGetRestaurantRestaurantNotFound,
	},
}

func TestGetProfileHost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range GetProfileHost {
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

var GetProfileCourier = []struct {
	testName                     string
	input                        int
	rowsQuery                    Row
	inputQuery                   int
	errQuery                     error
	out                          *profile.Profile
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
		out: &profile.Profile{Name: "1", Email: "1", Phone: "1", Avatar: "1",
			Birthday: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)},
		outErr: errorsConst.RGetRestaurantRestaurantNotFound,
	},
}

func TestGetProfileCourier(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range GetProfileCourier {
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

var GetProfileClient = []struct {
	testName                     string
	input                        int
	rowsQuery                    Row
	inputQuery                   int
	errQuery                     error
	out                          *profile.Profile
	inputGetDishesRestaurantName string
	inputGetDishesRestaurantId   int
	outErr                       string
	rowsQueryBirthday            Row
	inputQueryBirthday           int
	countQueryBirthday           int
}{
	{
		testName:   "One",
		input:      1,
		rowsQuery:  Row{row: []interface{}{"1", "1", "1", "1"}},
		inputQuery: 1,
		errQuery:   nil,
		out: &profile.Profile{Name: "1", Email: "1", Phone: "1", Avatar: "1",
			Birthday: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)},
		outErr:             errorsConst.RGetRestaurantRestaurantNotFound,
		inputQueryBirthday: 1,
		rowsQueryBirthday:  Row{row: []interface{}{time.Now()}},
		countQueryBirthday: 0,
	},
}

func TestGetProfileClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range GetProfileClient {
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

var UpdateName = []struct {
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
	for _, tt := range UpdateName {
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

var UpdateEmail = []struct {
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
	for _, tt := range UpdateEmail {
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

var UpdatePassword = []struct {
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
	for _, tt := range UpdatePassword {
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

var UpdatePhone = []struct {
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
	for _, tt := range UpdatePhone {
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

var UpdateAvatar = []struct {
	testName         string
	inputId          int
	inputAvatar      *profile.UpdateAvatar
	inputQueryId     int
	countQuery       int
	inputQueryAvatar string
	errQuery         error
	errUpload        error
	countUpload      int
	outErr           string
}{
	{
		testName:         "One",
		inputQueryId:     1,
		inputQueryAvatar: "1",
		errQuery:         nil,
		outErr:           errorsConst.PUpdateAvatarAvatarNotOpen,
		inputId:          1,
		inputAvatar:      &profile.UpdateAvatar{FileHeader: &multipart.FileHeader{Filename: "name.txt"}}, //TODO: make fill
		countQuery:       0,
		errUpload:        nil,
		countUpload:      0,
	},
}

func TestUpdateAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	mUploader := mocks.NewMockUploaderInterface(ctrl)
	for _, tt := range UpdateAvatar {
		m.
			EXPECT().
			Exec(context.Background(),
				"UPDATE general_user_info SET avatar = $1 WHERE id = $2",
				tt.inputQueryAvatar, tt.inputQueryId,
			).
			Return(nil, tt.errQuery).
			Times(tt.countQuery)
		mUploader.
			EXPECT().
			Upload(gomock.Any()).
			Return(nil, tt.errUpload).
			Times(tt.countUpload)
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

var UpdateBirthday = []struct {
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
	for _, tt := range UpdateBirthday {
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

var UpdateAddress = []struct {
	testName          string
	inputId           int
	inputAddress      profile.AddressCoordinates
	inputQueryId      int
	inputQueryAddress profile.AddressCoordinates
	errQuery          error
	outErr            string
}{
	{
		testName:     "One",
		inputQueryId: 1,
		inputQueryAddress: profile.AddressCoordinates{Alias: "1", Comment: "1", City: "1", Street: "1", House: "1",
			Floor: 1, Flat: "1", Porch: 1, Intercom: "1", Coordinates: profile.Coordinates{Latitude: 1.0, Longitude: 1.0}},
		errQuery: nil,
		outErr:   "",
		inputId:  1,
		inputAddress: profile.AddressCoordinates{Alias: "1", Comment: "1", City: "1", Street: "1", House: "1",
			Floor: 1, Flat: "1", Porch: 1, Intercom: "1", Coordinates: profile.Coordinates{Latitude: 1.0, Longitude: 1.0}},
	},
}

func TestUpdateAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range UpdateAddress {
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
