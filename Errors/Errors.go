package Errors

import "time"

type ResultError struct {
	Status	int         `json:"status"`
	Explain	string		`json:"explain,omitempty"`
}

type Errors struct {
	Text string
	Time time.Time
}

func (e *Errors) Error() string {
	return e.Text
}

// Error of server
const(
	ErrDB                        = "ERROR: database is not responding"
	ErrEncode                    = "ERROR: Encode"
	ErrUnmarshal                 = "ERROR: unmarshal json"
	ErrAuth                      = "ERROR: authorization failed"
)

// Error of Authorization
const(
	ErrSelectSaltInLogin         = "ERROR: salt in login not scan"
	ErrLoginOrPasswordIncorrect  = "Неправильный логин или пароль"
	ErrGeneralInfoScan           = "ERROR: general_user_info not scan"
	ErrInsertHost                = "ERROR: host not insert"
	ErrInsertCourier             = "ERROR: courier not insert"
	ErrInsertClient              = "ERROR: client not insert"
	ErrInsertTransactionCookie   = "ERROR: cookie with transaction not insert"
	ErrDeleteCookie              = "ERROR: cookie not delete"
	ErrInsertCookie              = "ERROR: cookie not insert"
	ErrGeneralInfoUnique         = "Телефон или Email уже зарегистрирован"
)

// Error of Middleware
const(
	ErrNotConnect                = "ERROR: db not connect"
	ErrCookieNotScan             = "ERROR: cookie not scan"
	ErrCookieScan                = "ERROR: cookie not scan"
	ErrCheckAccessCookieNotFound = "ERROR: cookie not found in CheckAccess"
	ErrUpdateCSRF                = "ERROR: csrf not updated"
	ErrCookieExpired             = "ERROR: cookie expired"
	ErrCookieNotFound            = "ERROR: cookie not found"
)

// Error of profile
const(
	ErrClientScan                = "ERROR: check user on client not scan"
	ErrHostScan                  = "ERROR: check user on host not scan"
	ErrCourierScan               = "ERROR: check user on courier not scan"
	ErrGetProfileHostScan        = "ERROR: get profile host not scan"
	ErrGetProfileClientScan      = "ERROR: get profile client not scan"
	ErrGetProfileCourierScan     = "ERROR: get profile courier not scan"
	ErrGetBirthdayScan           = "ERROR: birthday not scan"
)

// Error of restaurant
const(
	ErrRestaurantsNotFound       = "ERROR: restaurants not found"
	ErrRestaurantScan            = "ERROR: restaurant scan error"
	ErrRestaurantsNotSelect      = "ERROR: restaurant not select"
)