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
	Alias string
	Text  string
}

func (e *Errors) Error() string {
	return e.Alias
}

// Error of main
const (
	MCreateDBNotConnect = "db not connect"
)

// Error of promocode
const (
	PGetTypePromoCodeTransactionNotCreate          = "transaction get type promo code not create"
	PGetTypePromoCodeNotCommit                     = "transaction get type promo code not commit"
	PGetTypePromoCodeRestaurantsNotFound           = "type not found"
	PGetTypePromoCodeRestaurantsNotSelect          = "type not select"
	PActiveCostForFreeDeliveryTransactionNotCreate = "transaction get cost for free delivery not create"
	PActiveCostForFreeDeliveryNotCommit            = "transaction get cost for free delivery not commit"
	PActiveCostForFreeDeliveryRestaurantsNotFound  = "cost for free delivery not found"
	PActiveCostForFreeDeliveryRestaurantsNotSelect = "cost for free delivery not select"
	PActiveCostForSaleTransactionNotCreate         = "transaction get cost for sale not create"
	PActiveCostForSaleNotCommit                    = "transaction get cost for sale not commit"
	PActiveCostForSaleRestaurantsNotFound          = "cost for sale not found"
	PActiveCostForSaleRestaurantsNotSelect         = "cost for sale not select"
	PActiveTimeForSaleTransactionNotCreate         = "transaction get Time for sale not create"
	PActiveTimeForSaleNotCommit                    = "transaction get Time for sale not commit"
	PActiveTimeForSaleRestaurantsNotFound          = "Time for sale not found"
	PActiveTimeForSaleRestaurantsNotSelect         = "Time for sale not select"
	PActiveCostForFreeDishTransactionNotCreate     = "transaction for free dish not create"
	PActiveCostForFreeDishRestaurantsNotFound      = "free dish not found"
	PActiveCostForFreeDishRestaurantsNotSelect     = "free dish not select"
	PActiveCostForFreeDishNotCommit                = "transaction for free dish not commit"
)
