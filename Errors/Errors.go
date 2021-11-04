package Errors

import (
	"go.uber.org/zap"
	"time"
)

type ResultError struct {
	Status  int    `json:"status"`
	Explain string `json:"explain,omitempty"`
}

type Errors struct {
	Text string
	Time time.Time
}

func (e *Errors) Error() string {
	return e.Text
}

type CheckError struct {
	RequestId     *int
	LoggerErrWarn *zap.SugaredLogger
	LoggerInfo    *zap.SugaredLogger
	LoggerTest    *zap.SugaredLogger
}

// Error of server
const (
	ErrDB              = "ERROR: database is not responding"
	ErrEncode          = "ERROR: Encode"
	ErrAtoi            = "ERROR: func Atoi convert string in int"
	ErrNotStringAndInt = "ERROR: expected type string or int"
	ErrMarshal         = "ERROR: marshaling in json"
	ErrCheck           = "ERROR: err check"
	ErrUnmarshal       = "ERROR: unmarshal json"
	ErrAuth            = "Вы не авторизированы"
	IntNil             = 0
	ErrCartNull        = "Ваша корзина пустая"
)

// Error of Authorization
const (
	ASaltNotSelect                     = "salt not scan"
	ALoginOrPasswordIncorrect          = "Неправильный логин или пароль"
	AGeneralSignUpNotInsert            = "general_user_info not scan"
	ASignUpHostHostNotInsert           = "host not insert"
	ASignUpCourierCourierNotInsert     = "courier not insert"
	ASignUpClientClientNotInsert       = "client not insert"
	AAddTransactionCookieNotInsert     = "cookie with transaction not insert"
	ADeleteCookieCookieNotDelete       = "cookie not delete"
	AAddCookieCookieNotInsert          = "cookie not insert"
	AGeneralSignUpLoginNotUnique       = "Телефон или Email уже зарегистрирован"
	AGeneralSignUpIncorrectPhoneFormat = "Неверный формат телефона"
	ALoginNotFound                     = "user not found"
	ASignupHostTransactionNotCreate    = "transaction not create"     // TODO: add handler
	ASignupCourierTransactionNotCreate = "transaction not create"     // TODO: add handler
	ASignupClientTransactionNotCreate  = "transaction not create"     // TODO: add handler
	ASignUpUnknownType                 = "unknown type of user"       // TODO: add handler
	ALoginVoidLogin                    = "email and password is void" // TODO: add handler
	ASignUpHostNotCommit               = "signup host not commit"     // TODO: add handler
	ASignUpCourierNotCommit            = "signup courier not commit"  // TODO: add handler
	ASignUpClientNotCommit             = "signup client not commit"   // TODO: add handler
)

// Error of Middleware
const (
	MCheckAccessCookieNotScan    = "cookie not scan"
	MGetIdByCookieCookieNotScan  = "cookie not scan"
	MCheckAccessCookieNotFound   = "cookie not found"
	MNewCSRFCSRFNotUpdate        = "csrf not updated"
	MGetIdByCookieCookieExpired  = "cookie expired"
	MGetIdByCookieCookieNotFound = "cookie not found"
)

// Error of Utils
const (
	UCreateDBNotConnect         = "db not connect"             // TODO: add handler
	UCreateDBCreateFileNotFound = "CreateTables.sql not found" // TODO: add handler
	UCreateDBDeleteFileNotFound = "DeleteTables.sql not found" // TODO: add handler
	UCreateDBFillFileNotFound   = "Fill.sql not found"         // TODO: add handler
	UCreateDBNotCreateTables    = "table not create"           // TODO: add handler
	UCreateDBNotDeleteTables    = "table not delete"           // TODO: add handler
	UCreateDBNotFillTables      = "table not fill"             // TODO: add handler
)

// Error of profile
const (
	PGetRoleByIdClientNotScan        = "check user on client not scan"
	PGetRoleByIdHostNotScan          = "check user on host not scan"
	PGetRoleByIdCourierNotScan       = "check user on courier not scan"
	PGetProfileHostHostNotScan       = "get profile host not scan"
	PGetProfileClientClientNotScan   = "get profile client not scan"
	PGetProfileCourierCourierNotScan = "get profile courier not scan"
	PGetProfileClientBirthdayNotScan = "birthday not scan"
	PUpdateNameNameNotUpdate         = "name not update"
	PUpdateEmailEmailNotUpdate       = "email not update"
	PUpdateEmailEmailRepeat          = "email already exist"
	PUpdatePhonePhoneNotUpdate       = "phone not update"
	PUpdatePhonePhoneRepeat          = "phone already exist"
	PUpdatePasswordSaltNotSelect     = "salt not found"
	PUpdatePasswordPasswordNotUpdate = "password not update"
	PUpdateAvatarAvatarNotUpdate     = "avatar not update"
	PUpdateBirthdayBirthdayNotUpdate = "birthday not update"
	PUpdateAddressAddressNotUpdate   = "address not update"
	PGetProfileUnknownRole           = "unknown role of user" // TODO: add handler
)

// Error of restaurant
const (
	RGetRestaurantsRestaurantsNotFound    = "restaurants not found"
	RGetRestaurantsRestaurantsNotScan     = "restaurants scan error"
	RGetRestaurantsRestaurantsNotSelect   = "restaurants not select" // TODO: add handler
	RGetGeneralInfoRestaurantNotFound     = "restaurant not found"
	RGetTagsCategoryRestaurantNotScan     = "category restaurants scan error"
	RGetMenuDishesNotSelect               = "dishes not select"
	RGetDishesRestaurantDishesNotScan     = "dishes not scan"
	RGetMenuDishesNotFound                = "dishes not found"
	RGetDishesDishesNotFound              = "dishes not found"    // TODO: add handler
	RGetDishesDishesNotScan               = "dishes not scan"     // TODO: add handler
	RGetStructDishesStructDishesNotSelect = "dishes not select"   // TODO: add handler
	RGetStructDishesStructDishesNotScan   = "dishes not scan"     // TODO: add handler
	RGetStructRadiosStructRadiosNotSelect = "radios not select"   // TODO: add handler
	RGetRadiosRadiosNotScan               = "radios not scan"     // TODO: add handler
	RGetStructRadiosStructRadiosNotFound  = "radios not found"    // TODO: add handler
	RGetStructRadiosStructRadiosNotScan   = "radios not scan"     // TODO: add handler
	RGetTagsCategoryNotSelect             = "category not select" // TODO: add handler
	RGetRadiosRadiosNotSelect             = "radios not select"   // TODO: add handler
	RGetTagsTagsNotFound                  = "tags not found"
)

// Error of Cart
const (
	CGetCartDishesNotFound                         = "dishes not found"
	CGetCartDishesNotScan                          = "dishes not scan"
	CGetStructFoodRestaurantNotSelect              = "restaurant not select"
	CGetStructFoodCheckboxNotScan                  = "checkbox not scan"
	CGetStructRadiosRadiosNotSelect                = "radios not select"
	CGetStructRadiosRadiosNotScan                  = "radios not scan"
	CGetStructRadiosStructRadiosNotFound           = "struct radios not found" // TODO: add handler
	CGetStructRadiosStructRadiosNotScan            = "struct radios not scan"
	CDeleteCartCartNotDelete                       = "cart not delete"
	CDeleteCartStructureFoodNotDelete              = "food not delete"
	CDeleteCartRadiosFoodNotDelete                 = "radios not delete"
	CUpdateCartCartNotInsert                       = "cart not insert"
	CUpdateCartCartNotFound                        = "dish not found"
	CUpdateCartStructFoodStructureFoodNotInsert    = "structure food not insert"
	CUpdateCartRadiosRadiosNotInsert               = "radios not insert"
	CGetPriceDeliveryPriceNotFound                 = "delivery not found"
	CGetPriceDeliveryPriceNotScan                  = "delivery not scan"
	CUpdateCartCartNotScan                         = "cart not scan"
	CUpdateCartStructureFoodStructureFoodNotSelect = "structure dishes not select"
	CUpdateCartStructRadiosStructRadiosNotSelect   = "structure radios not select"
	CUpdateCartTransactionNotCreate                = "transaction not create" // TODO: add handler
	CUpdateCartNotCommit                           = "update cart not commit" // TODO: add handler
)
