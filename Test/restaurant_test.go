package Test

/*import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	res "2021_2_GORYACHIE_MEKSIKANSI/Restaurant"
	mocks "2021_2_GORYACHIE_MEKSIKANSI/Test/Mocks"
	rest "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgproto3/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

type Rows struct {
	testName string
	count int
}

func (r *Rows) Close() {
}

func (r *Rows) Err() error {
	return nil
}

func (r *Rows) CommandTag() pgconn.CommandTag {
	return nil
}

func (r *Rows) FieldDescriptions() []pgproto3.FieldDescription {
	return nil
}

func (r *Rows) Values() ([]interface{}, error) {
	return nil, nil
}

func (r *Rows) RawValues() [][]byte {
	return nil
}

func (r *Rows) Scan(dest ...interface{}) error {
	switch r.testName {
	case "Two":
		return errors.New("text")
	}
	return nil
}

func (r *Rows) Next() bool {
	switch r.testName {
	case "Two":
		return true
	case "Four":
		r.count++
		if r.count == 2 {
			return false
		}
		return true
	}
	return false
}

var restaurantTests = []struct {
	testName string
	out []rest.Restaurant
	err error
	row Rows
	outErr string
}{
	{
		testName: "One",
		out: nil,
		err: errors.New("text"),
		row: Rows{},
		outErr: errorsConst.ErrRestaurantsNotSelect,
	},
	{
		testName: "Two",
		out: nil,
		err: nil,
		row: Rows{testName: "Two"},
		outErr: errorsConst.ErrRestaurantScan,
	},
	{
		testName: "Three",
		out: nil,
		err: nil,
		row: Rows{testName: "Three"},
		outErr: errorsConst.ErrRestaurantsNotFound,
	},
	{
		testName: "Four",
		out: []rest.Restaurant{rest.Restaurant{}},
		err: nil,
		row: Rows{"Four", 0},
		outErr: "",
	},
}

func TestRestaurants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	for _, tt := range restaurantTests {
		m.
			EXPECT().
			Query(context.Background(),
				"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant ORDER BY random() LIMIT 50",
			).
			Return(&tt.row, tt.err)
		testUser := &res.Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetRestaurants()
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if err != nil {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err.Error()))
			}
		})
	}

}


var restaurantApplicationTests = []struct {
	testName string
	out []rest.Restaurant
	outErr string
	err error
}{
	{
		testName: "One",
		out: []rest.Restaurant{},
		err: nil,
		outErr: "",
	},
	{
		testName: "Two",
		out: nil,
		err: errors.New("text"),
		outErr: "text",
	},
}

func TestRestaurantApplication(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurant(ctrl)
	for _, tt := range restaurantApplicationTests {
		m.
			EXPECT().
			GetRestaurants().
			Return([]rest.Restaurant{}, tt.err)
		t.Run(tt.testName, func(t *testing.T) {
			result, err := res.AllRestaurants(m)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err.Error()))
			}
		})
	}
}*/
