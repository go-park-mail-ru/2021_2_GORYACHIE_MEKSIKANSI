package service

import (
	appPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/promocode/application"
)

type PromocodeManagerInterface interface {
}

type PromocodeManager struct {
	Application appPkg.PromocodeApplicationInterface
}


