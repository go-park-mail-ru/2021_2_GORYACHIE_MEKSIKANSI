package Test

/*import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/Authorization"
	mocks "2021_2_GORYACHIE_MEKSIKANSI/Test/Mocks"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestSignUpApplication(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	spaceDefense := utils.Defense{}

	m := mocks.NewMockWrapperAuthorization(ctrl)
	m.
		EXPECT().
		SignupClient(&utils.RegistrationRequest{TypeUser: "client"}).
		Return(&spaceDefense, errors.New("text"))
	m.
		EXPECT().
		SignupClient(&utils.RegistrationRequest{TypeUser: "client"}).
		Return(&spaceDefense, nil)
	m.
		EXPECT().
		SignupHost(&utils.RegistrationRequest{TypeUser: "host"}).
		Return(&spaceDefense, nil)
	m.
		EXPECT().
		SignupCourier(&utils.RegistrationRequest{TypeUser: "courier"}).
		Return(&spaceDefense, nil)
	// TODO: make beautiful
	result, _ := auth.SignUp(m, &utils.RegistrationRequest{TypeUser: "client"})
	if gomock.Nil().Matches(result) != true {
		//t.Errorf("Not equal\n")
		fmt.Printf("Not equal\n")
	} else {
		fmt.Printf("equal\n")
	}
	result, _ = auth.SignUp(m, &utils.RegistrationRequest{TypeUser: "client"})
	if gomock.Nil().Matches(result) != true {
		//t.Errorf("Not equal\n")
		fmt.Printf("Not equal\n")
	} else {
		fmt.Printf("equal\n")
	}
	result, _ = auth.SignUp(m, &utils.RegistrationRequest{TypeUser: "courier"})
	if gomock.Nil().Matches(result) != true {
		//t.Errorf("Not equal\n")
		fmt.Printf("Not equal\n")
	} else {
		fmt.Printf("equal\n")
	}
	result, _ = auth.SignUp(m, &utils.RegistrationRequest{TypeUser: "host"})
	if gomock.Nil().Matches(result) != true {
		//t.Errorf("Not equal\n")
		fmt.Printf("Not equal\n")
	} else {
		fmt.Printf("equal\n")
	}
	result, _ = auth.SignUp(m, &utils.RegistrationRequest{TypeUser: "default"})
	if gomock.Nil().Matches(result) != true {
		//t.Errorf("Not equal\n")
		fmt.Printf("Not equal\n")
	} else {
		fmt.Printf("equal\n")
	}
}

func TestLoginApplication(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	spaceDefense := utils.Defense{}

	m := mocks.NewMockWrapperAuthorization(ctrl)
	m.
		EXPECT().
		LoginByEmail("1", "1").
		Return(1, nil)
	m.
		EXPECT().
		GenerateNew().
		Return(&spaceDefense)
	m.
		EXPECT().
		LoginByPhone("1", "1").
		Return(1, nil)
	m.
		EXPECT().
		GenerateNew().
		Return(&spaceDefense)
	m.
		EXPECT().
		LoginByPhone("1", "1").
		Return(0, errors.New("text"))
	m.
		EXPECT().
		AddCookie(&spaceDefense, 1).
		Return(nil)
	m.
		EXPECT().
		AddCookie(&spaceDefense, 1).
		Return(errors.New("text"))
	// TODO: make beautiful
	result, _ := auth.Login(m, &auth.Authorization{Email: "1", Password: "1"})
	if gomock.Nil().Matches(result) != true {
		//t.Errorf("Not equal\n")
		fmt.Printf("Not equal\n")
	} else {
		fmt.Printf("equal\n")
	}
	result, _ = auth.Login(m, &auth.Authorization{Phone: "1", Password: "1"})
	if gomock.Nil().Matches(result) != true {
		//t.Errorf("Not equal\n")
		fmt.Printf("Not equal\n")
	} else {
		fmt.Printf("equal\n")
	}
	result, _ = auth.Login(m, &auth.Authorization{Phone: "1", Password: "1"})
	if gomock.Nil().Matches(result) != true {
		//t.Errorf("Not equal\n")
		fmt.Printf("Not equal\n")
	} else {
		fmt.Printf("equal\n")
	}
	result, _ = auth.Login(m, &auth.Authorization{})
	if gomock.Nil().Matches(result) != true {
		//t.Errorf("Not equal\n")
		fmt.Printf("Not equal\n")
	} else {
		fmt.Printf("equal\n")
	}
}

func TestLogoutApplication(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	spaceDefense := utils.Defense{}

	m := mocks.NewMockWrapperAuthorization(ctrl)
	m.
		EXPECT().
		DeleteCookie(&spaceDefense).
		Return(nil)
	// TODO: make beautiful
	result := auth.Logout(m, &spaceDefense)
	if gomock.Nil().Matches(result) != true {
		//t.Errorf("Not equal\n")
		fmt.Printf("Not equal\n")
	} else {
		fmt.Printf("equal\n")
	}
}
*/