package Middleware

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
		case *time.Time:
			*dest[i].(*time.Time) = r.row[i].(time.Time)
		}
	}
	return nil
}

var ApplicationCheckAccess = []struct {
	testName            string
	input               *Utils.Defense
	out                 bool
	outErr              string
	inputQuerySessionId string
	inputQueryCSRFToken string
	outQuery            Row
}{
	{
		testName:            "One",
		outErr:              "",
		inputQuerySessionId: "1",
		inputQueryCSRFToken: "1",
		outQuery:            Row{row: []interface{}{1, time.Now()}},
		input:               &Utils.Defense{SessionId: "1", CsrfToken: "1"},
	},
}

func TestApplicationCheckAccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range ApplicationCheckAccess {
		m.
			EXPECT().
			QueryRow(context.Background(),
				"SELECT client_id, date_life FROM cookie WHERE session_id = $1 AND csrf_token = $2",
				tt.inputQuerySessionId, tt.inputQueryCSRFToken).
			Return(&tt.outQuery)
		t.Run(tt.testName, func(t *testing.T) {
			result, err := CheckAccess(m, tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
