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

type CheckError struct {
	RequestId int
	Logger    MultiLogger
}

// Error of authorization
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
	ALoginByEmailTransactionNotCreate  = "transaction login by email not create"
	ALoginByEmailNotCommit             = "login by email not commit"
	ALoginByPhoneTransactionNotCreate  = "transaction login by phone not create"
	ALoginByPhoneNotCommit             = "login by phone not commit"
	ADeleteCookieTransactionNotCreate  = "transaction delete cookie not create"
	ADeleteCookieNotCommit             = "transaction delete cookie not commit"
	AAddCookieTransactionNotCreate     = "transaction add cookie not create"
	AAddCookieNotCommit                = "transaction add cookie not commit"
)

// Error of middleware
const (
	MCheckAccessCookieNotScan          = "cookie not scan"
	MGetIdByCookieCookieNotScan        = "cookie not scan"
	MCheckAccessCookieNotFound         = "cookie not found"
	MNewCSRFCSRFNotUpdate              = "csrf not updated"
	MGetIdByCookieCookieExpired        = "cookie expired"
	MGetIdByCookieCookieNotFound       = "cookie not found"
	MCheckAccessTransactionNotCreate   = "transaction check access not create"
	MCheckAccessNotCommit              = "transaction check access not commit"
	MNewCSRFCSRFTransactionNotCreate   = "transaction new csrf not create"
	MNewCSRFCSRFNotCommit              = "transaction new csrf not commit"
	MGetIdByCookieTransactionNotCreate = "transaction get id not create"
	MGetIdByCookieNotCommit            = "transaction get id not commit"
)

// Error of main
const (
	MCreateDBNotConnect = "db not connect"
)
