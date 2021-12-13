package application

import (
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/promocode/orm"
)

type PromocodeApplicationInterface interface {
}

type Promocode struct {
	DB ormPkg.WrapperPromocodeInterface
}
