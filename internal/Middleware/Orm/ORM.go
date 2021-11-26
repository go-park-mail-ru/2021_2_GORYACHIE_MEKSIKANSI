package Orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyError"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Util"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Util/Cast"

	"context"
)

type Wrapper struct {
	Conn Interface.ConnectionMiddleware
	Ctx  context.Context
}

func (w *Wrapper) CheckAccess(cookie *Util.Defense) (bool, error) {
	user, err := w.Conn.CheckAccessUser(w.Ctx, cast.CastDefenseToDefenseProto(cookie))
	if err != nil {
		return false, err
	}
	if user.Error != "" {
		return false, &errPkg.Errors{Alias: user.Error}
	}
	return user.CheckResult, nil
}

func (w *Wrapper) NewCSRF(cookie *Util.Defense) (string, error) {
	user, err := w.Conn.NewCSRFUser(w.Ctx, cast.CastDefenseToDefenseProto(cookie))
	if err != nil {
		return "", err
	}
	if user.Error != "" {
		return "", &errPkg.Errors{Alias: user.Error}
	}
	return user.XCsrfToken.XCsrfToken, nil
}

func (w *Wrapper) GetIdByCookie(cookie *Util.Defense) (int, error) {
	byCookie, err := w.Conn.GetIdByCookie(w.Ctx, cast.CastDefenseToDefenseProto(cookie))
	if err != nil {
		return 0, err
	}
	if byCookie.Error != "" {
		return 0, &errPkg.Errors{Alias: byCookie.Error}
	}
	return int(byCookie.IdUser), nil
}
