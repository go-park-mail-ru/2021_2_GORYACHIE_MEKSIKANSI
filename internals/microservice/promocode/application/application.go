//go:generate mockgen -destination=mocks/orm.go -package=mocks 2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/promocode/orm WrapperPromocodeInterface

package application

import (
	ormPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/promocode/orm"
	"time"
)

type PromocodeApplicationInterface interface {
	GetTypePromoCode(promoCode string, restaurantId int) (int, error)
	ActiveFreeDelivery(promoCode string, restaurantId int) (bool, error)
	ActiveCostForFreeDish(promoCode string, restaurantId int) (int, int, error)
	ActiveCostForSale(promoCode string, amount int, restaurantId int) (int, error)
	ActiveTimeForSale(promoCode string, amount int, restaurantId int) (int, error)
	AddPromoCode(promoCode string, restaurantId int, clientId int) error
	GetPromoCode(id int) (string, error)
}

type Promocode struct {
	DB ormPkg.WrapperPromocodeInterface
}

func (db *Promocode) GetTypePromoCode(promoCode string, restaurantId int) (int, error) {
	return db.DB.GetTypePromoCode(promoCode, restaurantId)
}

func (db *Promocode) ActiveFreeDelivery(promoCode string, restaurantId int) (bool, error) {
	return db.DB.ActiveFreeDelivery(promoCode, restaurantId)
}

func (db *Promocode) ActiveCostForFreeDish(promoCode string, restaurantId int) (int, int, error) {
	return db.DB.ActiveCostForFreeDish(promoCode, restaurantId)
}

func (db *Promocode) ActiveCostForSale(promoCode string, amount int, restaurantId int) (int, error) {
	return db.DB.ActiveCostForSale(promoCode, amount, restaurantId)
}

func (db *Promocode) ActiveTimeForSale(promoCode string, amount int, restaurantId int) (int, error) {
	return db.DB.ActiveTimeForSale(promoCode, amount, restaurantId, time.Now())
}

func (db *Promocode) AddPromoCode(promoCode string, restaurantId int, clientId int) error {
	return db.DB.AddPromoCode(promoCode, restaurantId, clientId)
}

func (db *Promocode) GetPromoCode(id int) (string, error) {
	return db.DB.GetPromoCode(id)
}
