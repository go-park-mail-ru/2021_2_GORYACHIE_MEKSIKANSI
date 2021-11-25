package Orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Authorization"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	authProto "2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Authorization/proto"
	Utils2 "2021_2_GORYACHIE_MEKSIKANSI/internal/Util"
	"context"
)

type Wrapper struct {
	Conn Interface.ConnectAuthService
	Ctx context.Context
}

func (w *Wrapper) SignUp(signup *Authorization.RegistrationRequest) (*Utils2.Defense, error) {
	// TODO: add convert func
	var a authProto.RegistrationRequest
	a.TypeUser = signup.TypeUser
	a.Name = signup.Name
	a.Phone = signup.Phone
	a.Email = signup.Email
	a.Password = signup.Password
	result, err := w.Conn.SignUp(w.Ctx, &a)
	if err != nil {
		return nil, err
	}

	var res Utils2.Defense
	res.SessionId = result.Defense.SessionId
	res.CsrfToken = result.Defense.XCsrfToken
	//res.DateLife = result.Defense.DateLife
	return &res, nil
}

func (w *Wrapper) Login(login *Authorization.Authorization) (*Utils2.Defense, error) {
	// TODO: add convert func
	var a authProto.Authorization
	a.Phone = login.Phone
	a.Email = login.Email
	a.Password = login.Password
	response, err := w.Conn.Login(w.Ctx, &a)
	if err != nil {
		return nil, err
	}

	var cookie Utils2.Defense
	cookie.SessionId = response.Defense.SessionId
	cookie.CsrfToken = response.Defense.XCsrfToken
	//cookie.DateLife = response.Defense.DateLife
	return &cookie, nil
}

func (w *Wrapper) Logout(CSRF string) (string, error) {
	// TODO: add convert func
	var csrfToken authProto.CSRF
	csrfToken.XCsrfToken = CSRF
	logout, err := w.Conn.Logout(w.Ctx, &csrfToken)
	if err != nil {
		return "", err
	}
	return logout.XCsrfToken.XCsrfToken, nil
}

func (w *Wrapper) CheckAccess(cookie *Utils2.Defense) (bool, error) {
	// TODO: add convert func
	var send authProto.Defense
	//send.DateLife = cookie.DateLife
	send.XCsrfToken = cookie.CsrfToken
	send.SessionId = cookie.SessionId
	user, err := w.Conn.CheckAccessUser(w.Ctx, &send)
	if err != nil {
		return false, err
	}
	return user.CheckResult, nil
}

func (w *Wrapper) NewCSRF(cookie *Utils2.Defense) (string, error){
	// TODO: add convert func
	var send authProto.Defense
	//send.DateLife = cookie.DateLife
	send.XCsrfToken = cookie.CsrfToken
	send.SessionId = cookie.SessionId
	user, err := w.Conn.NewCSRFUser(w.Ctx, &send)
	if err != nil {
		return "", err
	}
	return user.XCsrfToken.XCsrfToken, nil
}

func (w *Wrapper) GetIdByCookie(cookie *Utils2.Defense) (int, error){
	// TODO: add convert func
	var send authProto.Defense
	//send.DateLife = cookie.DateLife
	send.XCsrfToken = cookie.CsrfToken
	send.SessionId = cookie.SessionId
	byCookie, err := w.Conn.GetIdByCookie(w.Ctx, &send)
	if err != nil {
		return 0, err
	}
	return int(byCookie.IdUser), nil
}
