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

// Error of restaurant
const (
	RGetRestaurantRestaurantNotFound   = "restaurant not found"
	RGetRestaurantTransactionNotCreate = "transaction get restaurant not create"
	RGetRestaurantNotCommit            = "get restaurant not commit"
)

// Error of cart
const (
	CDeleteCartCartNotDelete                       = "cart not delete"
	CUpdateCartCartNotInsert                       = "cart not insert"
	CUpdateCartCartNotFound                        = "dish not found"
	CUpdateCartStructFoodStructureFoodNotInsert    = "structure food not insert"
	CUpdateCartRadiosRadiosNotInsert               = "radios not insert"
	CGetPriceDeliveryPriceNotFound                 = "delivery not found"
	CGetPriceDeliveryPriceNotScan                  = "delivery not scan"
	CUpdateCartCartNotScan                         = "cart not scan"
	CUpdateCartStructureFoodStructureFoodNotSelect = "structure dishes not select"
	CUpdateCartStructRadiosStructRadiosNotSelect   = "structure radios not select"
	CUpdateCartTransactionNotCreate                = "transaction not create"
	CUpdateCartNotCommit                           = "update cart not commit"
	CGetCartTransactionNotCreate                   = "transaction get not create"
	CGetCartNotSelect                              = "cart not select"
	CGetCartNotCommit                              = "transaction get not commit"
	CGetCartNotScan                                = "cart not scan"
	CGetCartCartNotFound                           = "cart is void"
	CDeleteCartTransactionNotCreate                = "transaction delete cart not create"
	CDeleteCartNotCommit                           = "transaction delete cart not commit"
	CGetPriceDeliveryTransactionNotCreate          = "transaction get price delivery not create"
	CGetPriceDeliveryNotCommit                     = "transaction get price delivery not commit"
	CAddPromoCodeTransactionNotCreate              = "transaction add promo code not create"
	CAddPromoCodeNotCommit                         = "transaction add promo code not commit"
	CAddPromoCodeNotUpsert                         = "promo not upsert"
	CDoPromoCodeNotSelectInfo                      = "promo code info not select"
	CDoPromoCodeNotSelectInfoDish                  = "info about free dish not select"
)
