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
	HttpNil            = 0
	ErrCartNull        = "Ваша корзина пустая"
)

// Error of Authorization
const (
	ErrSelectSaltInLogin                 = "salt in login not scan"
	ErrLoginOrPasswordIncorrect          = "Неправильный логин или пароль"
	ErrGeneralInfoScan                   = "general_user_info not scan"
	ErrInsertHost                        = "host not insert"
	ErrInsertCourier                     = "courier not insert"
	ErrInsertClient                      = "client not insert"
	ErrInsertTransactionCookie           = "cookie with transaction not insert"
	ErrDeleteCookie                      = "cookie not delete"
	ErrInsertCookie                      = "cookie not insert"
	ErrGeneralInfoUnique                 = "Телефон или Email уже зарегистрирован"
	ErrPhoneFormat                       = "Неверный формат телефона"
	ErrUserNotFoundLogin                 = "user not found"
	ErrSignupHostTransactionNotCreate    = "transaction not create in SignupHost"    // TODO: add handler
	ErrSignupCourierTransactionNotCreate = "transaction not create in SignupCourier" // TODO: add handler
	ErrSignupClientTransactionNotCreate  = "transaction not create in SignupClient"  // TODO: add handler
)

// Error of Middleware
const (
	ErrNotConnect                = "db not connect"
	ErrCookieNotScan             = "cookie not scan"
	ErrCookieScan                = "cookie not scan"
	ErrCheckAccessCookieNotFound = "cookie not found in CheckAccess"
	ErrUpdateCSRF                = "csrf not updated"
	ErrCookieExpired             = "cookie expired"
	ErrCookieNotFound            = "cookie not found"
	ErrFileNotFound              = "CreateTables.sql not found" // TODO: add handler
	ErrDeleteFileNotFound        = "DeleteTables.sql not found" // TODO: add handler
	ErrFillFileNotFound          = "Fill.sql not found"         // TODO: add handler
	ErrNotCreateTables           = "table not create"           // TODO: add handler
	ErrNotDeleteTables           = "table not delete"           // TODO: add handler
	ErrNotFillTables             = "table not fill"             // TODO: add handler
)

// Error of profile
const (
	ErrClientScan            = "check user on client not scan"
	ErrHostScan              = "check user on host not scan"
	ErrCourierScan           = "check user on courier not scan"
	ErrGetProfileHostScan    = "get profile host not scan"
	ErrGetProfileClientScan  = "get profile client not scan"
	ErrGetProfileCourierScan = "get profile courier not scan"
	ErrGetBirthdayScan       = "birthday not scan"
	ErrUpdateName            = "name not update"
	ErrUpdateEmail           = "email not update"
	ErrUpdateEmailRepeat     = "email already exist"
	ErrUpdatePhone           = "phone not update"
	ErrUpdatePhoneRepeat     = "phone already exist"
	ErrSelectSaltInUpdate    = "salt not found in update"
	ErrUpdatePassword        = "password not update"
	ErrUpdateAvatar          = "avatar not update"
	ErrUpdateBirthday        = "birthday not update"
	ErrUpdateAddress         = "address not update"
)

// Error of restaurant
const (
	ErrRestaurantsNotFound        = "restaurants not found"
	ErrRestaurantsScan            = "restaurants scan error"
	ErrRestaurantsNotSelect       = "restaurants not select"
	ErrRestaurantNotFound         = "restaurant not found"
	ErrCategoryRestaurantScan     = "category restaurants scan error"
	ErrRestaurantsDishesNotSelect = "dishes in restaurant not select"
	ErrRestaurantDishesScan       = "dishes in restaurant not scan"
	ErrRestaurantDishesNotFound   = "dishes in restaurant not found"
	DishesDishesNotFound          = "dishes not found"
	DishesDishesNotScan           = "dishes not scan"
	DishesStructDishesNotSelect   = "dishes not select"
	DishesStructDishesNotScan     = "dishes not scan"
	DishesStructRadiosNotSelect   = "radios not select"
	DishesRadiosNotScan           = "radios not scan"
	DishesStructRadiosNotFound    = "radios not found"
	DishesStructRadiosNotScan     = "radios not scan"
	ErrTagNotFound                = "tag not found"
)

// Error of Cart
const (
	GetCartRestaurantNotFound         = "restaurant not found"
	GetCartRestaurantNotScan          = "restaurant not scan"
	GetCartCartNotFound               = "cart not found"
	GetCartCartNotScan                = "cart not scan"
	GetCartDishesNotFound             = "dishes not found"
	GetCartDishesNotScan              = "dishes not scan"
	GetCartRestaurantNotSelect        = "restaurant not select"
	GetCartCheckboxNotScan            = "checkbox not scan"
	GetCartRadiosNotSelect            = "radios not select"
	GetCartRadiosNotScan              = "radios not scan"
	GetCartStructRadiosNotFound       = "struct radios not found"
	GetCartStructRadiosNowScan        = "struct radios not scan"
	CartNotDelete                     = "cat not delete"
	StructureFoodNotDelete            = "food not delete"
	CartRadiosFoodNotDelete           = "radios not delete"
	UpdateCartCartNotInsert           = "cart not insert"
	UpdateCartCartNotFound            = "dish not found"
	UpdateCartStructureFoodNotInsert  = "structure food not insert"
	UpdateCartRadiosNotInsert         = "radios not insert"
	GetPriceDeliveryNotFound          = "delivery not found"
	GetPriceDeliveryNotScan           = "delivery not scan"
	UpdateCartCartNotScan             = "cart not scan"
	UpdateCartStructureNotSelect      = "structure dishes not select"
	UpdateCartStructRadiosNotSelect   = "structure radios not select"
	ErrUpdateCartTransactionNotCreate = "transaction not create in UpdateCart" // TODO: add handler
)

// TODO: make TODO
