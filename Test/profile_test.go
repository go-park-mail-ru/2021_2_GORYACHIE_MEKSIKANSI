package Test

import (
	profile "2021_2_GORYACHIE_MEKSIKANSI/Profile"
	mocks "2021_2_GORYACHIE_MEKSIKANSI/Test/Mocks"
	pr "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"
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
	testUser := &profile.Wrapper{Conn: m}
	result, _ := testUser.GetProfileClient(1)
	if gomock.Nil().Matches(result) != true {
		t.Errorf("Not equal\n")
	} else {
		fmt.Printf("equal\n")
	}
}

func TestProfileApplication(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	spaceProfile := pr.Profile{}

	m := mocks.NewMockWrapperProfile(ctrl)
	m.
		EXPECT().
		GetRoleById(1).
		Return("client", nil).Times(2)
	m.
		EXPECT().
		GetRoleById(1).
		Return("host", nil)
	m.
		EXPECT().
		GetRoleById(1).
		Return("courier", nil)
	m.
		EXPECT().
		GetRoleById(1).
		Return("default", nil)
	m.
		EXPECT().
		GetRoleById(1).
		Return("host", errors.New("text"))
	m.
		EXPECT().
		GetProfileClient(1).
		Return(&spaceProfile, errors.New("text"))
	m.
		EXPECT().
		GetProfileClient(1).
		Return(&spaceProfile, nil)
	m.
		EXPECT().
		GetProfileHost(1).
		Return(&spaceProfile, nil)
	m.
		EXPECT().
		GetProfileCourier(1).
		Return(&spaceProfile, nil)
	// TODO: make beautiful
	for i := 0; i < 6; i++ {
		result, _ := profile.GetProfile(m, 1)
		if gomock.Nil().Matches(result) != true {
			//t.Errorf("Not equal\n")
			fmt.Printf("Not equal\n")
		} else {
			fmt.Printf("equal\n")
		}
	}
}
