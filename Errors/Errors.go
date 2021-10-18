package Errors

import "time"

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

// Error of server
const (
	ErrDB        = "ERROR: database is not responding"
	ErrEncode    = "ERROR: Encode"
	ErrMarshal   = "ERROR: marshaling in json"
	ErrCheck     = "ERROR: err check"
	ErrUnmarshal = "ERROR: unmarshal json"
	ErrAuth      = "Вы не авторизированы"
	HttpNil      = 0
)

// Error of Authorization
const (
	ErrSelectSaltInLogin        = "ERROR: salt in login not scan"
	ErrLoginOrPasswordIncorrect = "Неправильный логин или пароль"
	ErrGeneralInfoScan          = "ERROR: general_user_info not scan"
	ErrInsertHost               = "ERROR: host not insert"
	ErrInsertCourier            = "ERROR: courier not insert"
	ErrInsertClient             = "ERROR: client not insert"
	ErrInsertTransactionCookie  = "ERROR: cookie with transaction not insert"
	ErrDeleteCookie             = "ERROR: cookie not delete"
	ErrInsertCookie             = "ERROR: cookie not insert"
	ErrGeneralInfoUnique        = "Телефон или Email уже зарегистрирован"
	ErrPhoneFormat              = "Неверный формат телефона"
	ErrUserNotFoundLogin        = "ERROR: user not found"
)

// Error of Middleware
const (
	ErrNotConnect                = "ERROR: db not connect"
	ErrCookieNotScan             = "ERROR: cookie not scan"
	ErrCookieScan                = "ERROR: cookie not scan"
	ErrCheckAccessCookieNotFound = "ERROR: cookie not found in CheckAccess"
	ErrUpdateCSRF                = "ERROR: csrf not updated"
	ErrCookieExpired             = "ERROR: cookie expired"
	ErrCookieNotFound            = "ERROR: cookie not found"
	ErrFileNotFound              = "ERROR: CreateTables.sql not found" // TODO: add handler
	ErrDeleteFileNotFound        = "ERROR: DeleteTables.sql not found" // TODO: add handler
	ErrFillFileNotFound          = "ERROR: Fill.sql not found"         // TODO: add handler
	ErrNotCreateTables           = "ERROR: table not create"           // TODO: add handler
	ErrNotDeleteTables           = "ERROR: table not delete"           // TODO: add handler
	ErrNotFillTables             = "ERROR: table not fill"             // TODO: add handler
)

// Error of profile
const (
	ErrClientScan            = "ERROR: check user on client not scan"
	ErrHostScan              = "ERROR: check user on host not scan"
	ErrCourierScan           = "ERROR: check user on courier not scan"
	ErrGetProfileHostScan    = "ERROR: get profile host not scan"
	ErrGetProfileClientScan  = "ERROR: get profile client not scan"
	ErrGetProfileCourierScan = "ERROR: get profile courier not scan"
	ErrGetBirthdayScan       = "ERROR: birthday not scan"
	ErrUpdateName            = "ERROR: name not update"          // TODO: add handler
	ErrUpdateEmail           = "ERROR: email not update"         // TODO: add handler
	ErrUpdateEmailRepeat     = "ERROR: email already exist"      // TODO: add handler
	ErrUpdatePhone           = "ERROR: phone not update"         // TODO: add handler
	ErrUpdatePhoneRepeat     = "ERROR: phone already not update" // TODO: add handler
	ErrSelectSaltInUpdate    = "ERROR: salt not found in update" // TODO: add handler
	ErrUpdatePassword        = "ERROR: password not update"      // TODO: add handler
	ErrUpdateAvatar          = "ERROR: avatar not update"        // TODO: add handler
)

// Error of restaurant
const (
	ErrRestaurantsNotFound  = "ERROR: restaurants not found"
	ErrRestaurantsScan       = "ERROR: restaurants scan error"
	ErrRestaurantsNotSelect = "ERROR: restaurants not select"
	ErrRestaurantNotFound  = "ERROR: restaurant not found"                    // TODO: add handler
	ErrCategoryRestaurantScan = "ERROR: category restaurants scan error"      // TODO: add handler
	ErrRestaurantsDishesNotSelect = "ERROR: dishes in restaurant not select"  // TODO: add handler
	ErrRestaurantDishesScan = "ERROR: dishes in restaurant not scan"          // TODO: add handler
	ErrRestaurantDishesNotFound = "ERROR: dishes in restaurant not found"     // TODO: add handler
)
