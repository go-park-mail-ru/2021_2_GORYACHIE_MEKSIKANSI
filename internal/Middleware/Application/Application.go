package Application

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Interfaces"
	utils "2021_2_GORYACHIE_MEKSIKANSI/internal/Utils"
)

type Middleware struct {
	DB Interfaces.WrapperMiddleware
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
