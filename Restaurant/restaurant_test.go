package Restaurant

import (
	errorsConst "2021_2_GORYACHIE_MEKSIKANSI/Errors"
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
	count    int
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
	out      []rest.Restaurants
	err      error
	row      Rows
	outErr   string
}{
	{
		testName: "One",
		out:      nil,
		err:      errors.New("text"),
		row:      Rows{},
		outErr:   errorsConst.ErrRestaurantsNotSelect,
	},
	{
		testName: "Two",
		out:      nil,
		err:      nil,
		row:      Rows{testName: "Two"},
		outErr:   errorsConst.ErrRestaurantsScan,
	},
	{
		testName: "Three",
		out:      nil,
		err:      nil,
		row:      Rows{testName: "Three"},
		outErr:   errorsConst.ErrRestaurantsNotFound,
	},
	{
		testName: "Four",
		out:      []rest.Restaurants{rest.Restaurants{}},
		err:      nil,
		row:      Rows{"Four", 0},
		outErr:   "",
	},
}

func TestRestaurantsOrm(t *testing.T) {
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
		testUser := &Wrapper{Conn: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := testUser.GetRestaurants()
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}

}

var ApplicationAllRestaurants = []struct {
	testName string
	out      []rest.Restaurants
	outErr   string
	err      error
}{
	{
		testName: "One",
		out:      []rest.Restaurants{},
		err:      nil,
		outErr:   "",
	},
	{
		testName: "Two",
		out:      nil,
		err:      errors.New("text"),
		outErr:   "text",
	},
}

func TestApplicationAllRestaurants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurant(ctrl)
	for _, tt := range ApplicationAllRestaurants {
		m.
			EXPECT().
			GetRestaurants().
			Return([]rest.Restaurants{}, tt.err)
		t.Run(tt.testName, func(t *testing.T) {
			result, err := AllRestaurants(m)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationGetRestaurant = []struct {
	testName string
	out      *rest.RestaurantId
	outErr   string
	input    int
	inputGetGeneralInfoRestaurant int
	resultGetGeneralInfoRestaurant *rest.RestaurantId
	errGetGeneralInfoRestaurant error
	countGetGeneralInfoRestaurant int
	inputGetTagsRestaurant int
	resultGetTagsRestaurant []rest.Tag
	errGetTagsRestaurant error
	countGetTagsRestaurant int
	inputGetMenu int
	resultGetMenu []rest.Menu
	errGetMenu error
	countGetMenu int
}{
	{
		testName: "One",
		out:      &rest.RestaurantId{},
		outErr:   "",
	},
}

func TestApplicationGetRestaurant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurant(ctrl)
	for _, tt := range ApplicationGetRestaurant {
		m.
			EXPECT().
			GetGeneralInfoRestaurant(tt.inputGetGeneralInfoRestaurant).
			Return(tt.resultGetGeneralInfoRestaurant, tt.errGetGeneralInfoRestaurant).
			Times(tt.countGetGeneralInfoRestaurant)
		m.
			EXPECT().
			GetTagsRestaurant(tt.inputGetTagsRestaurant).
			Return(tt.resultGetTagsRestaurant, tt.errGetTagsRestaurant).
			Times(tt.countGetTagsRestaurant)
		m.
			EXPECT().
			GetMenu(tt.inputGetMenu).
			Return(tt.resultGetMenu, tt.errGetMenu).
			Times(tt.countGetMenu)
		t.Run(tt.testName, func(t *testing.T) {
			result, err := GetRestaurant(m, tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}


var ApplicationRestaurantDishes= []struct {
	testName string
	inputRestId    int
	inputDishId    int
	out      *rest.Dishes
	outErr   string
	inputGetDishesRestId int
	inputGetDishesDishId int
	resultGetDishes *rest.RestaurantId
	errGetDishes error
	countGetDishes int
	inputGetStructureDishes int
	resultGetStructureDishes []rest.Tag
	errGetStructureDishes error
	countGetStructureDishes int
	inputGetRadios int
	resultGetRadios []rest.Menu
	errGetRadios error
	countGetRadios int
}{
	{
		testName: "One",
		out:      &rest.Dishes{},
		outErr:   "",
	},
}

func TestApplicationRestaurantDishes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurant(ctrl)
	for _, tt := range ApplicationRestaurantDishes {
		m.
			EXPECT().
			GetDishes(tt.inputGetDishesRestId, tt.inputGetDishesDishId).
			Return(tt.resultGetDishes, tt.errGetDishes).
			Times(tt.countGetDishes)
		m.
			EXPECT().
			GetTagsRestaurant(tt.inputGetStructureDishes).
			Return(tt.resultGetStructureDishes, tt.errGetStructureDishes).
			Times(tt.countGetStructureDishes)
		m.
			EXPECT().
			GetMenu(tt.inputGetRadios).
			Return(tt.resultGetRadios, tt.errGetRadios).
			Times(tt.countGetRadios)
		t.Run(tt.testName, func(t *testing.T) {
			result, err := RestaurantDishes(m, tt.inputRestId, tt.inputDishId)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
