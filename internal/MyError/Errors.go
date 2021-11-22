package MyError

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
	ALoginByEmailTransactionNotCreate  = "transaction login by email not create"
	ALoginByEmailNotCommit             = "login by email not commit"
	ALoginByPhoneTransactionNotCreate  = "transaction login by phone not create"
	ALoginByPhoneNotCommit             = "login by phone not commit"
	ADeleteCookieTransactionNotCreate  = "transaction delete cookie not create" // TODO: add in checkErrors
	ADeleteCookieNotCommit             = "transaction delete cookie not commit" // TODO: add in checkErrors
	AAddCookieTransactionNotCreate     = "transaction add cookie not create"    // TODO: add in checkErrors
	AAddCookieNotCommit                = "transaction add cookie not commit"    // TODO: add in checkErrors
)

// Error of Middleware
const (
	MCheckAccessCookieNotScan          = "cookie not scan"
	MGetIdByCookieCookieNotScan        = "cookie not scan"
	MCheckAccessCookieNotFound         = "cookie not found"
	MNewCSRFCSRFNotUpdate              = "csrf not updated"
	MGetIdByCookieCookieExpired        = "cookie expired"
	MGetIdByCookieCookieNotFound       = "cookie not found"
	MCheckAccessTransactionNotCreate   = "transaction check access not create" // TODO: add in checkErrors
	MCheckAccessNotCommit              = "transaction check access not commit" // TODO: add in checkErrors
	MNewCSRFCSRFTransactionNotCreate   = "transaction new csrf not create"     // TODO: add in checkErrors
	MNewCSRFCSRFNotCommit              = "transaction new csrf not commit"     // TODO: add in checkErrors
	MGetIdByCookieTransactionNotCreate = "transaction get id not create"       // TODO: add in checkErrors
	MGetIdByCookieNotCommit            = "transaction get id not commit"       // TODO: add in checkErrors
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
	PGetRoleByIdClientNotScan              = "check user on client not scan"
	PGetRoleByIdHostNotScan                = "check user on host not scan"
	PGetRoleByIdCourierNotScan             = "check user on courier not scan"
	PGetProfileHostHostNotScan             = "get profile host not scan"
	PGetProfileClientClientNotScan         = "get profile client not scan"
	PGetProfileCourierCourierNotScan       = "get profile courier not scan"
	PGetProfileClientBirthdayNotScan       = "birthday not scan"
	PUpdateNameNameNotUpdate               = "name not update"
	PUpdateEmailEmailNotUpdate             = "email not update"
	PUpdateEmailEmailRepeat                = "email already exist"
	PUpdatePhonePhoneNotUpdate             = "phone not update"
	PUpdatePhonePhoneRepeat                = "phone already exist"
	PUpdatePasswordSaltNotSelect           = "salt not found"
	PUpdatePasswordPasswordNotUpdate       = "password not update"
	PUpdateAvatarAvatarNotUpdate           = "avatar not update"
	PUpdateBirthdayBirthdayNotUpdate       = "birthday not update"
	PUpdateAddressAddressNotUpdate         = "address not update"
	PGetProfileUnknownRole                 = "unknown role of user"
	PUpdatePhoneIncorrectPhoneFormat       = "incorrect format phone"
	PUpdateAvatarAvatarNotOpen             = "file not open"
	PUpdateAvatarAvatarNotUpload           = "avatar not send"
	PUpdateAvatarFileNameEmpty             = "file name is empty"
	PUpdateAvatarFileWithoutExtension      = "file without extension"
	PGetRoleByIdTransactionNotCreate       = "transaction role by id not create"          // TODO: add in checkErrors
	PGetRoleByIdNotCommit                  = "role by id not commit"                      // TODO: add in checkErrors
	PGetProfileClientTransactionNotCreate  = "transaction get profile client not create"  // TODO: add in checkErrors
	PGetProfileClientNotCommit             = "get profile client not commit"              // TODO: add in checkErrors
	PUpdatePasswordTransactionNotCreate    = "transaction update password not create"     // TODO: add in checkErrors
	PUpdatePasswordNotCommit               = "update password not commit"                 // TODO: add in checkErrors
	PAddAddressAddressNotAdd               = "address not insert"                         // TODO: add in checkErrors
	PAddDeleteAddressNotDelete             = "address not delete"                         // TODO: add in checkErrors
	PGetProfileHostTransactionNotCreate    = "transaction get profile host not create"    // TODO: add in checkErrors
	PGetProfileHostNotCommit               = "transaction get profile host not commit"    // TODO: add in checkErrors
	PGetProfileCourierTransactionNotCreate = "transaction get profile courier not create" // TODO: add in checkErrors
	PGetProfileCourierNotCommit            = "transaction get profile courier not commit" // TODO: add in checkErrors
	PUpdateNameTransactionNotCreate        = "transaction update name not create"         // TODO: add in checkErrors
	PUpdateNameNotCommit                   = "transaction update name not commit"         // TODO: add in checkErrors
	PUpdateEmailTransactionNotCreate       = "transaction update email not create"        // TODO: add in checkErrors
	PUpdateEmailNotCommit                  = "transaction update email not commit"        // TODO: add in checkErrors
	PUpdatePhoneTransactionNotCreate       = "transaction update phone not create"        // TODO: add in checkErrors
	PUpdatePhoneNotCommit                  = "transaction update phone not commit"        // TODO: add in checkErrors
	PUpdateAvatarTransactionNotCreate      = "transaction update avatar not create"       // TODO: add in checkErrors
	PUpdateAvatarNotCommit                 = "transaction update avatar not commit"       // TODO: add in checkErrors
	PUpdateBirthdayTransactionNotCreate    = "transaction update birthday not create"     // TODO: add in checkErrors
	PUpdateBirthdayNotCommit               = "transaction update birthday not commit"     // TODO: add in checkErrors
	PUpdateAddressTransactionNotCreate     = "transaction update address not create"      // TODO: add in checkErrors
	PUpdateAddressNotCommit                = "transaction update address not commit"      // TODO: add in checkErrors
	PAddAddressNotCreate                   = "transaction add address not create"         // TODO: add in checkErrors
	PAddAddressNotCommit                   = "transaction add address not commit"         // TODO: add in checkErrors
	PAddDeleteAddressTransactionNotCreate  = "transaction delete address not create"      // TODO: add in checkErrors
	PAddDeleteAddressNotCommit             = "transaction delete address not commit"      // TODO: add in checkErrors
)

// Error of restaurant
const (
	RGetRestaurantsRestaurantsNotFound     = "restaurants not found"
	RGetRestaurantsRestaurantsNotScan      = "restaurants scan error"
	RGetRestaurantsRestaurantsNotSelect    = "restaurants not select"
	RGetGeneralInfoRestaurantNotFound      = "restaurant not found"
	RGetTagsRestaurantRestaurantNotScan    = "category restaurants scan error"
	RGetMenuDishesNotSelect                = "dishes not select"
	RGetDishesRestaurantDishesNotScan      = "dishes not scan"
	RGetMenuDishesNotFound                 = "dishes not found"
	RGetDishesDishesNotFound               = "dishes not found"
	RGetDishesDishesNotScan                = "dishes not scan"
	RGetStructDishesStructDishesNotSelect  = "struct dishes not select"
	RGetStructDishesStructDishesNotScan    = "dishes not scan"
	RGetRadiosRadiosNotScan                = "radios not scan"
	RGetTagsRestaurantNotSelect            = "category not select"
	RGetRadiosRadiosNotSelect              = "radios not select"
	RGetTagsTagsNotFound                   = "tags not found"
	RGetMenuTransactionNotCreate           = "transaction menu dishes not create"         // TODO: add in checkErrors
	RGetRadiosNotCreate                    = "transaction get radios not create"          // TODO: add in checkErrors
	RGetRadiosNotCommit                    = "get radios not commit"                      // TODO: add in checkErrors
	RGetRestaurantsTransactionNotCreate    = "transaction get restaurants not create"     // TODO: add in checkErrors
	RGetRestaurantsNotCommit               = "get restaurants not commit"                 // TODO: add in checkErrors
	RGetGeneralInfoTransactionNotCreate    = "transaction get info restaurant not create" // TODO: add in checkErrors
	RGetGeneralInfoNotCommit               = "get info restaurant not commit"             // TODO: add in checkErrors
	RGetTagsRestaurantTransactionNotCreate = "transaction get tag restaurant not create"  // TODO: add in checkErrors
	RGetTagsRestaurantNotCommit            = "get info restaurant not commit"             // TODO: add in checkErrors
	RGetMenuNotCommit                      = "get menu not commit"                        // TODO: add in checkErrors
	RGetStructDishesTransactionNotCreate   = "transaction get struct dishes not create"   // TODO: add in checkErrors
	RGetStructDishesNotCommit              = "get get struct dishes not commit"           // TODO: add in checkErrors
	RGetDishesTransactionNotCreate         = "transaction get dishes not create"          // TODO: add in checkErrors
	RGetDishesNotCommit                    = "get get dishes not commit"                  // TODO: add in checkErrors
	RGetMenuDishesCategoryNotSelect        = "category not select"                        // TODO: add in checkErrors
)

// Error of Cart
const (
	CGetCartDishesNotFound                         = "dishes not found"
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
	CGetCartTransactionNotCreate                   = "transaction get not create"
	CGetCartNotSelect                              = "cart not select"
	CGetCartNotCommit                              = "transaction get not commit"
	CGetCartNotScan                                = "cart not scan"
	CDeleteCartTransactionNotCreate                = "transaction delete cart not create"        // TODO: add in checkErrors
	CDeleteCartNotCommit                           = "transaction delete cart not commit"        // TODO: add in checkErrors
	CGetPriceDeliveryTransactionNotCreate          = "transaction get price delivery not create" // TODO: add in checkErrors
	CGetPriceDeliveryNotCommit                     = "transaction get price delivery not commit" // TODO: add in checkErrors
)

// Error of Order
const (
	OCreateOrderTransactionNotCreate         = "transaction create order not create" // TODO: add in checkErrors
	OCreateOrderNotCommit                    = "create order not commit"             // TODO: add in checkErrors
	OGetOrdersTransactionNotCreate           = "transaction get orders not create"   // TODO: add in checkErrors
	OGetOrdersDishesNotCommit                = "get orders not commit"               // TODO: add in checkErrors
	OCreateOrderOrderUserNotInsert           = "not insert in order_user"            // TODO: add in checkErrors
	OCreateOrderOrderRadiosListUserNotInsert = "not insert in order_radios_list"     // TODO: add in checkErrors
	OCreateOrderOrderStructureListNotInsert  = "not insert in order_structure_list"  // TODO: add in checkErrors
	OCreateOrderOrderListNotInsert           = "not insert in order_list"            // TODO: add in checkErrors
	OCreateOrderCountNotUpdate               = "count dish not update"               // TODO: add in checkErrors
	OCreateOrderCountNotCorrect              = "dishes not enough"                   // TODO: add in checkErrors
	OCreateOrderCartIsVoid                   = "cart is void"                        // TODO: add in checkErrors
	OGetOrdersNotSelect                      = "all not selected"                    // TODO: add in checkErrors
)
