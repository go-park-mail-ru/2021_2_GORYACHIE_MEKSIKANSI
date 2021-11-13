package Errors

import (
	"time"
)

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
	Time time.Time
}

func (e *Errors) Error() string {
	return e.Text
}

type CheckError struct {
	RequestId int
	Logger    MultiLogger
}

// Error of server
const (
	ErrDB              = "database is not responding"
	ErrEncode          = "Encode"
	ErrAtoi            = "func Atoi convert string in int"
	ErrNotStringAndInt = "expected type string or int"
	ErrMarshal         = "marshaling in json"
	ErrCheck           = "err check"
	ErrUnmarshal       = "unmarshal json"
	ErrAuth            = "Вы не авторизированы"
	IntNil             = 0
	ErrCartNull        = "Ваша корзина пустая"
	ErrNotSearchAvatar = "Поле avatar не было найдено"
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
	ASignupHostTransactionNotCreate    = "transaction host not create"
	ASignupCourierTransactionNotCreate = "transaction courier not create"
	ASignupClientTransactionNotCreate  = "transaction client not create"
	ASignUpUnknownType                 = "unknown type of user"
	ALoginVoidLogin                    = "email and password is void"
	ASignUpHostNotCommit               = "signup host not commit"
	ASignUpCourierNotCommit            = "signup courier not commit"
	ASignUpClientNotCommit             = "signup client not commit"
	ALoginByEmailTransactionNotCreate  = "transaction login by email not create"  // TODO: add in checkErrors
	ALoginByEmailNotCommit             = "login by email not commit"              // TODO: add in checkErrors
	ALoginByPhoneTransactionNotCreate  = "transaction  login by phone not create" // TODO: add in checkErrors
	ALoginByPhoneNotCommit             = "login by phone not commit"              // TODO: add in checkErrors
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

// Error of main
const (
	MCreateDBNotConnect         = "db not connect"
	MCreateDBCreateFileNotFound = "CreateTables.sql not found"
	MCreateDBDeleteFileNotFound = "DeleteTables.sql not found"
	MCreateDBFillFileNotFound   = "Fill.sql not found"
	MCreateDBNotCreateTables    = "table not create"
	MCreateDBNotDeleteTables    = "table not delete"
	MCreateDBNotFillTables      = "table not fill"
)

// Error of profile
const (
	PGetRoleByIdClientNotScan             = "check user on client not scan"
	PGetRoleByIdHostNotScan               = "check user on host not scan"
	PGetRoleByIdCourierNotScan            = "check user on courier not scan"
	PGetProfileHostHostNotScan            = "get profile host not scan"
	PGetProfileClientClientNotScan        = "get profile client not scan"
	PGetProfileCourierCourierNotScan      = "get profile courier not scan"
	PGetProfileClientBirthdayNotScan      = "birthday not scan"
	PUpdateNameNameNotUpdate              = "name not update"
	PUpdateEmailEmailNotUpdate            = "email not update"
	PUpdateEmailEmailRepeat               = "email already exist"
	PUpdatePhonePhoneNotUpdate            = "phone not update"
	PUpdatePhonePhoneRepeat               = "phone already exist"
	PUpdatePasswordSaltNotSelect          = "salt not found"
	PUpdatePasswordPasswordNotUpdate      = "password not update"
	PUpdateAvatarAvatarNotUpdate          = "avatar not update"
	PUpdateBirthdayBirthdayNotUpdate      = "birthday not update"
	PUpdateAddressAddressNotUpdate        = "address not update"
	PGetProfileUnknownRole                = "unknown role of user"
	PUpdatePhoneIncorrectPhoneFormat      = "incorrect format phone"
	PUpdateAvatarAvatarNotOpen            = "file not open"
	PUpdateAvatarAvatarNotUpload          = "avatar not send"
	PUpdateAvatarFileNameEmpty            = "file name is empty"
	PUpdateAvatarFileWithoutExtension     = "file without extension"
	PGetRoleByIdTransactionNotCreate      = "transaction role by id not create"         // TODO: add in checkErrors
	PGetRoleByIdNotCommit                 = "role by id not commit"                     // TODO: add in checkErrors
	PGetProfileClientTransactionNotCreate = "transaction get profile client not create" // TODO: add in checkErrors
	PGetProfileClientNotCommit            = "get profile client not commit"             // TODO: add in checkErrors
	PUpdatePasswordTransactionNotCreate   = "transaction update password not create"    // TODO: add in checkErrors
	PUpdatePasswordNotCommit              = "update password not commit"                // TODO: add in checkErrors
	PAddAddressAddressNotAdd              = "address not insert"                        // TODO: add in checkErrors
	PAddDeleteAddressNotDelete              = "address not delete"                        // TODO: add in checkErrors
)

// Error of restaurant
const (
	RGetRestaurantsRestaurantsNotFound    = "restaurants not found"
	RGetRestaurantsRestaurantsNotScan     = "restaurants scan error"
	RGetRestaurantsRestaurantsNotSelect   = "restaurants not select"
	RGetGeneralInfoRestaurantNotFound     = "restaurant not found"
	RGetTagsCategoryRestaurantNotScan     = "category restaurants scan error"
	RGetMenuDishesNotSelect               = "dishes not select"
	RGetDishesRestaurantDishesNotScan     = "dishes not scan"
	RGetMenuDishesNotFound                = "dishes not found"
	RGetDishesDishesNotFound              = "dishes not found"
	RGetDishesDishesNotScan               = "dishes not scan"
	RGetStructDishesStructDishesNotSelect = "dishes not select"
	RGetStructDishesStructDishesNotScan   = "dishes not scan"
	RGetStructRadiosStructRadiosNotSelect = "struct radios not select"
	RGetRadiosRadiosNotScan               = "radios not scan"
	RGetStructRadiosStructRadiosNotFound  = "radios not found"
	RGetStructRadiosStructRadiosNotScan   = "structure radios not scan"
	RGetTagsCategoryNotSelect             = "category not select"
	RGetRadiosRadiosNotSelect             = "radios not select"
	RGetTagsTagsNotFound                  = "tags not found"
	RGetMenuTransactionNotCreate          = "transaction menu dishes not create" // TODO: add in checkErrors
	RGetMenuDishesNotCommit               = "menu dishes not commit"             // TODO: add in checkErrors
	RGetRadiosNotCreate                   = "transaction get radios not create"  // TODO: add in checkErrors
	RGetRadiosNotCommit                   = "get radios not commit"              // TODO: add in checkErrors
)

// Error of Cart
const (
	CGetCartDishesNotFound                         = "dishes not found"
	CGetCartDishesNotScan                          = "dishes not scan"
	CGetStructFoodRestaurantNotSelect              = "restaurant not select"
	CGetStructFoodCheckboxNotScan                  = "checkbox not scan"
	CGetStructRadiosRadiosNotSelect                = "radios not select"
	CGetStructRadiosRadiosNotScan                  = "radios not scan"
	CGetStructRadiosStructRadiosNotFound           = "struct radios not found"
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
	CUpdateCartTransactionNotCreate                = "transaction not create"
	CUpdateCartNotCommit                           = "update cart not commit"
	CGetCartTransactionNotCreate = "transaction get not create"
	CGetCartNotSelect = "cart not select"
	CGetCartNotCommit = "transaction get not commit"
)

// Error of Order
const (
	OCreateOrderTransactionNotCreate         = "transaction create order not create"
	OCreateOrderNotCommit                    = "create order not commit"
	OGetOrdersTransactionNotCreate           = "transaction get orders not create"
	OGetOrdersDishesNotCommit                = "get orders not commit"
	OCreateOrderOrderUserNotInsert           = "not insert in order_user"
	OCreateOrderOrderRadiosListUserNotInsert = "not insert in order_radios_list"
	OCreateOrderOrderStructureListNotInsert  = "not insert in order_structure_list"
	OCreateOrderOrderListNotInsert           = "not insert in order_list"
	OCreateOrderCountNotUpdate               = "count dish not update"
	OCreateOrderCountNotCorrect              = "dishes not enough"
	OCreateOrderCartIsVoid = "cart is void"
)
