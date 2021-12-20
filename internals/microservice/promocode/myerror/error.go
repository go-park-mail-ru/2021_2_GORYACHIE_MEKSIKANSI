package myerror

type MultiLogger interface {
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Sync() error
}

type ResultError struct {
	Status  int    `json:"status"`
	Explain string `json:"explain,omitempty"`
}

type Errors struct {
	Text string
}

func (e *Errors) Error() string {
	return e.Text
}

// Error of main
const (
	MCreateDBNotConnect = "db not connect"
)

// Error of promocode
const (
	PGetTypePromoCodeTransactionNotCreate      = "transaction get type promo code not create"
	PGetTypePromoCodeNotCommit                 = "transaction get type promo code not commit"
	PGetTypePromoCodeRestaurantsNotFound       = "type not found"
	PGetTypePromoCodeRestaurantsNotSelect      = "type not select"
	PActiveFreeDeliveryTransactionNotCreate    = "transaction get cost for free delivery not create"
	PActiveFreeDeliveryNotCommit               = "transaction get cost for free delivery not commit"
	PActiveFreeDeliveryRestaurantsNotFound     = "cost for free delivery not found"
	PActiveFreeDeliveryRestaurantsNotSelect    = "cost for free delivery not select"
	PActiveCostForSaleTransactionNotCreate     = "transaction get cost for sale not create"
	PActiveCostForSaleNotCommit                = "transaction get cost for sale not commit"
	PActiveCostForSaleRestaurantsNotFound      = "cost for sale not found"
	PActiveCostForSaleRestaurantsNotSelect     = "cost for sale not select"
	PActiveTimeForSaleTransactionNotCreate     = "transaction get Time for sale not create"
	PActiveTimeForSaleNotCommit                = "transaction get Time for sale not commit"
	PActiveTimeForSaleRestaurantsNotFound      = "Time for sale not found"
	PActiveTimeForSaleRestaurantsNotSelect     = "Time for sale not select"
	PActiveCostForFreeDishTransactionNotCreate = "transaction for free dish not create"
	PActiveCostForFreeDishRestaurantsNotFound  = "free dish not found"
	PActiveCostForFreeDishRestaurantsNotSelect = "free dish not select"
	PActiveCostForFreeDishNotCommit            = "transaction for free dish not commit"
	PAddPromoCodeTransactionNotCreate          = "transaction add promo code not create"
	PAddPromoCodeNotUpsert                     = "promo not upsert"
	PAddPromoCodeNotCommit                     = "transaction add promo code not commit"
	PGetPromoCodeTransactionNotCreate          = "transaction get promo code not create"
	PGetPromoCodeNotSelect                     = "promo code not select"
	PGetPromoCodeNotCommit                     = "transaction get promo code not commit"
)
