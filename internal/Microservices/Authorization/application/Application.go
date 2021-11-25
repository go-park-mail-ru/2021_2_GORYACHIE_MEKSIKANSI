package Application

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Authorization/Interface"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Authorization/proto"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyError"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Util"
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type AuthorizationManager struct {
	DB Interface.WrapperAuthorization
}

func (am *AuthorizationManager) CheckAccessUser(ctx context.Context, cookie *proto.Defense) (*proto.CheckAccess, error) {
	// TODO: add convert func
	status, err := am.DB.CheckAccess(cookie)
	if err != nil {
		return nil, err
	}
	return &proto.CheckAccess{CheckResult: status}, nil
}

func (am *AuthorizationManager) NewCSRFUser(ctx context.Context, cookie *proto.Defense) (*proto.CSRFResponse, error) {
	// TODO: add convert func
	csrf, err := am.DB.NewCSRF(cookie)
	if err != nil {
		return nil, err
	}
	return &proto.CSRFResponse{XCsrfToken: &proto.CSRF{XCsrfToken: csrf}}, nil
}

func (am *AuthorizationManager) GetIdByCookie(ctx context.Context, cookie *proto.Defense) (*proto.IdClientResponse, error) {
	// TODO: add convert func
	byCookie, err := am.DB.GetIdByCookie(cookie)
	if err != nil {
		return nil, err
	}
	return &proto.IdClientResponse{IdUser: int64(byCookie)}, nil
}

func (am *AuthorizationManager) SignUp(ctx context.Context, signup *proto.RegistrationRequest) (*proto.DefenseResponse, error) {
	// TODO: add convert func
	var cookie *Util.Defense
	var err error
	newCookie := am.DB.NewDefense()
	switch signup.TypeUser {
	case "client":
		cookie, err = am.DB.SignupClient(signup, newCookie)
	case "courier":
		cookie, err = am.DB.SignupCourier(signup, newCookie)
	case "host":
		cookie, err = am.DB.SignupHost(signup, newCookie)
	default:
		return nil, &errPkg.Errors{
			Alias: errPkg.ASignUpUnknownType,
		}
	}

	if err != nil {
		return nil, err
	}

	return &proto.DefenseResponse{Defense: &proto.Defense{
		DateLife:   &timestamp.Timestamp{},
		SessionId:  cookie.SessionId,
		XCsrfToken: cookie.CsrfToken,
	}}, nil
}

func (am *AuthorizationManager) Login(ctx context.Context, login *proto.Authorization) (*proto.DefenseResponse, error) {
	// TODO: add convert func
	var userId int
	var err error
	switch {
	case login.Email != "":
		userId, err = am.DB.LoginByEmail(login.Email, login.Password)

	case login.Phone != "":
		userId, err = am.DB.LoginByPhone(login.Phone, login.Password)
	default:
		return nil, &errPkg.Errors{
			Alias: errPkg.ALoginVoidLogin,
		}
	}

	if err != nil {
		return nil, err
	}

	cookie := am.DB.NewDefense()
	err = am.DB.AddCookie(cookie, userId)

	if err != nil {
		return nil, err
	}
	return &proto.DefenseResponse{Defense: &proto.Defense{
		DateLife:   &timestamp.Timestamp{},
		SessionId:  cookie.SessionId,
		XCsrfToken: cookie.CsrfToken,
	}}, nil
}

func (am *AuthorizationManager) Logout(ctx context.Context, CSRF *proto.CSRF) (*proto.CSRFResponse, error) {
	// TODO: add convert func
	cookie, err := am.DB.DeleteCookie(CSRF.XCsrfToken)
	if err != nil {
		return nil, err
	}
	return &proto.CSRFResponse{XCsrfToken: &proto.CSRF{XCsrfToken: cookie}}, nil
}
