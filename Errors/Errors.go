package Errors

type ResultError struct {
	Status	int         `json:"status"`
	Explain	string		`json:"parsedJSON,omitempty"`
}

const(
	ErrDB                        = "ERROR: database is not responding"
	ErrEncode                    = "ERROR: Encode"
	ErrUnmarshal                 = "ERROR: unmarshal json"
	ErrAuth                      = "ERROR: authorization failed"
	ErrSelectSalt                = "ERROR: salt in login not scan"
	ErrLoginOrPasswordIncorrect  = "Неправильный логин или пароль"
	ErrGeneralInfoScan           = "ERROR: general_user_info not scan"
	ErrInsertHost                = "ERROR: host not insert"
	ErrInsertCourier             = "ERROR: courier not insert"
	ErrInsertClient              = "ERROR: client not insert"
	ErrInsertTransactionCookie   = "ERROR: cookie with transaction not insert"
	ErrDeleteCookie              = "ERROR: cookie not delete"
	ErrInsertCookie              = "ERROR: cookie not insert"
	ErrGeneralInfoUnique         = "Телефон или Email уже зарегистрирован"
	ErrCookieNotScan             = "ERROR: cookie not scan"
	ErrCookieScan                = "ERROR: cookie not scan"
	ErrCheckAccessCookieNotFound = "ERROR: cookie not found in CheckAccess"
	ErrNotConnect                = "ERROR: db not connect"
	ErrUpdateCSRF                = "ERROR: csrf not updated"
	ErrCookieExpired             = "ERROR: cookie expired"
	ErrClientScan                = "ERROR: check user on client not scan"
	ErrHostScan                  = "ERROR: check user on host not scan"
	ErrCourierScan               = "ERROR: check user on courier not scan"
	ErrGetProfileHostScan        = "ERROR: get profile host not scan"
	ErrGetProfileClientScan      = "ERROR: get profile client not scan"
	ErrGetProfileCourierScan     = "ERROR: get profile courier not scan"
	ErrGetBirthdayScan           = "ERROR: birthday not scan"
	ErrRestaurantsNotFound       = "ERROR: restaurants not found"
	ErrRestaurantScan            = "ERROR: restaurant scan error"
	ErrCookieNotFound            = "ERROR: cookie not found"
)
