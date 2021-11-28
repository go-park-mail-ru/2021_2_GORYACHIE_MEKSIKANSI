package orm

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interface"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/util"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/util/cast"

	"context"
)

type Wrapper struct {
	Conn Interface.ConnectionMiddleware
	Ctx  context.Context
}

func (w *Wrapper) CheckAccess(cookie *util.Defense) (bool, error) {
	user, err := w.Conn.CheckAccessUser(w.Ctx, cast.CastDefenseToDefenseProto(cookie))
	if err != nil {
		return false, err
	}
	if user.Error != "" {
		return false, &errPkg.Errors{Alias: user.Error}
	}
	return user.CheckResult, nil
}

func (w *Wrapper) NewCSRF(cookie *util.Defense) (string, error) {
	user, err := w.Conn.NewCSRFUser(w.Ctx, cast.CastDefenseToDefenseProto(cookie))
	if err != nil {
		return "", err
	}
	if user.Error != "" {
		return "", &errPkg.Errors{Alias: user.Error}
	}
	return user.XCsrfToken.XCsrfToken, nil
}

func (w *Wrapper) GetIdByCookie(cookie *util.Defense) (int, error) {
	byCookie, err := w.Conn.GetIdByCookie(w.Ctx, cast.CastDefenseToDefenseProto(cookie))
	if err != nil {
		return 0, err
	}
	if byCookie.Error != "" {
		return 0, &errPkg.Errors{Alias: byCookie.Error}
	}
	return int(byCookie.IdUser), nil
}
