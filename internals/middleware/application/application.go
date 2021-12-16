//go:generate mockgen -destination=mocks/application.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/middleware/orm WrapperMiddlewareInterface
package application

import (
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/middleware/orm"
	utils "2021_2_GORYACHIE_MEKSIKANSI/internals/util"
)

type MiddlewareApplicationInterface interface {
	CheckAccess(cookie *utils.Defense) (bool, error)
	NewCSRF(cookie *utils.Defense) (string, error)
	GetIdByCookie(cookie *utils.Defense) (int, error)
	CheckAccessWebsocket(cookie string) (bool, error)
}

type Middleware struct {
	DB ormPkg.WrapperMiddlewareInterface
}

func (m *Middleware) CheckAccess(cookie *utils.Defense) (bool, error) {
	return m.DB.CheckAccess(cookie)
}

func (m *Middleware) NewCSRF(cookie *utils.Defense) (string, error) {
	return m.DB.NewCSRF(cookie)
}

func (m *Middleware) GetIdByCookie(cookie *utils.Defense) (int, error) {
	return m.DB.GetIdByCookie(cookie)
}

func (m *Middleware) CheckAccessWebsocket(cookie string) (bool, error) {
	return m.DB.CheckAccessWebsocket(cookie)
}
