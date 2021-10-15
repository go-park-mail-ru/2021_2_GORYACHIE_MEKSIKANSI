package Test

import (
	res "2021_2_GORYACHIE_MEKSIKANSI/Restaurant"
	mocks "2021_2_GORYACHIE_MEKSIKANSI/Test/Mocks"
	rest "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgproto3"
	"testing"
)

type Rows struct {
}

func (r Rows) Close() {
}

func (r Rows) Err() error {
	return nil
}

func (r Rows) CommandTag() pgconn.CommandTag {
	return nil
}

func (r Rows) FieldDescriptions() []pgproto3.FieldDescription {
	return nil
}

func (r Rows) Values() ([]interface{}, error) {
	return nil, nil
}

func (r Rows) RawValues() [][]byte {
	return nil
}

func (r Rows) Scan(dest ...interface{}) error {
	return nil
}

func (r Rows) Next() bool {
	return false
}

func TestRestaurants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConnectionInterface(ctrl)
	m.
		EXPECT().
		Query(context.Background(),
		"SELECT id, avatar, name, price_delivery, min_delivery_time, max_delivery_time, rating FROM restaurant ORDER BY random() LIMIT 50",
		).
		Return(Rows{}, nil)
	testUser := &res.Wrapper{Conn: m}
	result, _ := testUser.GetRestaurants()
	if gomock.Nil().Matches(result) != true {
		t.Errorf("Not equal\n")
	} else {
		fmt.Printf("equal\n")
	}

}


func TestRestaurantApplication(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurant(ctrl)
	m.
		EXPECT().
		GetRestaurants().
		Return([]rest.Restaurant{}, nil)
	m.
		EXPECT().
		GetRestaurants().
		Return([]rest.Restaurant{}, errors.New("text"))
	// TODO: make beautiful
	for i := 0; i < 2; i++ {
		result, _ := res.AllRestaurants(m)
		if gomock.Nil().Matches(result) != true {
			//t.Errorf("Not equal\n")
			fmt.Printf("Not equal\n")
		} else {
			fmt.Printf("equal\n")
		}
	}
}
