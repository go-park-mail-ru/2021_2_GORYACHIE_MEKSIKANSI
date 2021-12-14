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

type ResultErrorMulti struct {
	Status  int         `json:"status"`
	Explain string      `json:"explain,omitempty"`
	Body    interface{} `json:"body"`
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

// Error of server
const (
	ErrDB              = "database is not responding"
	ErrEncode          = "Encode"
	ErrAtoi            = "func Atoi convert string in int"
	ErrNotStringAndInt = "expected type string or int"
	ErrMarshal         = "marshaling in json"
	ErrCheck           = "err check"
	ErrUnmarshal       = "unmarshal json"
	ErrAuth            = "Вы не авторизированны"
	IntNil             = 0
	ErrCartNull        = "Ваша корзина пустая"
	ErrReviewNull      = "Отзывы не найдены"
	ErrFavoriteNull    = "избранные рестораны не найдены"
	ErrSearchRes       = "Рестораны не найдены"
	ErrGerOrderNull    = "Заказов не  найдено"
	ErrOrderNull       = "Нечего создавать: корзина пустая"
	OrderNotExist      = "Заказ не существует"
	ErrNotSearchAvatar = "Поле avatar не было найдено"
)

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
	ADeleteCookieTransactionNotCreate  = "transaction delete cookie not create" // TODO: add in checkErrors
	ADeleteCookieNotCommit             = "transaction delete cookie not commit" // TODO: add in checkErrors
	AAddCookieTransactionNotCreate     = "transaction add cookie not create"    // TODO: add in checkErrors
	AAddCookieNotCommit                = "transaction add cookie not commit"    // TODO: add in checkErrors
)

// Error of middleware
const (
	MCheckAccessCookieNotScan          = "cookie not scan"
	MGetIdByCookieCookieNotScan        = "cookie not scan"
	MCheckAccessCookieNotFound         = "cookie not found"
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
	MCreateDBNotConnect           = "db not connect"
	MCreateDBCreateFileNotFound   = "createtables.sql not found"
	MCreateDBDeleteFileNotFound   = "deletetables.sql not found"
	MCreateDBFillFileNotFound     = "fill.sql not found"
	MCreateDBNotCreateTables      = "table not create"
	MCreateDBNotDeleteTables      = "table not delete"
	MCreateDBNotFillTables        = "table not fill"
	MCreateDBTransactionNotCreate = "transaction setup not create" // TODO: add in checkErrors
	MCreateDBNotCommit            = "transaction setup not commit" // TODO: add in checkErrors
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
	PUpdateBirthdayNotParse                = "birthday not parse"                         // TODO: add in checkErrors
)

// Error of restaurant
const (
	RGetRestaurantsRestaurantsNotFound            = "restaurants not found"
	RGetRestaurantsRestaurantsNotScan             = "restaurants scan error"
	RGetRestaurantsRestaurantsNotSelect           = "restaurants not select"
	RGetRestaurantRestaurantNotFound              = "restaurant not found"
	RGetTagsRestaurantRestaurantNotScan           = "category restaurants scan error"
	RGetMenuDishesNotFound                        = "dishes not found"
	RGetDishesDishesNotFound                      = "dishes not found"
	RGetStructDishesStructDishesNotSelect         = "struct dishes not select"
	RGetStructDishesStructDishesNotScan           = "dishes not scan"
	RGetRadiosRadiosNotScan                       = "radios not scan"
	RGetTagsRestaurantNotSelect                   = "category not select"
	RGetRadiosRadiosNotSelect                     = "radios not select"
	RGetTagsTagsNotFound                          = "tags not found"
	RGetMenuTransactionNotCreate                  = "transaction menu dishes not create"               // TODO: add in checkErrors
	RGetRadiosTransactionNotCreate                = "transaction get radios not create"                // TODO: add in checkErrors
	RGetRadiosNotCommit                           = "get radios not commit"                            // TODO: add in checkErrors
	RGetRestaurantsTransactionNotCreate           = "transaction get restaurants not create"           // TODO: add in checkErrors
	RGetRestaurantsNotCommit                      = "get restaurants not commit"                       // TODO: add in checkErrors
	RGetRestaurantTransactionNotCreate            = "transaction get restaurant not create"            // TODO: add in checkErrors
	RGetRestaurantNotCommit                       = "get restaurant not commit"                        // TODO: add in checkErrors
	RGetTagsRestaurantTransactionNotCreate        = "transaction get tag restaurant not create"        // TODO: add in checkErrors
	RGetTagsRestaurantNotCommit                   = "get info restaurant not commit"                   // TODO: add in checkErrors
	RGetMenuNotCommit                             = "get menu not commit"                              // TODO: add in checkErrors
	RGetStructDishesTransactionNotCreate          = "transaction get struct dishes not create"         // TODO: add in checkErrors
	RGetStructDishesNotCommit                     = "get get struct dishes not commit"                 // TODO: add in checkErrors
	RGetDishesTransactionNotCreate                = "transaction get dishes not create"                // TODO: add in checkErrors
	RGetDishesNotCommit                           = "get get dishes not commit"                        // TODO: add in checkErrors
	RGetMenuDishesCategoryNotSelect               = "category not select"                              // TODO: add in checkErrors
	RGetReviewTransactionNotCreate                = "transaction get review not create"                // TODO: add in checkErrors
	RGetReviewNotCommit                           = "get get review not commit"                        // TODO: add in checkErrors
	RCreateReviewTransactionNotCreate             = "transaction create review not create"             // TODO: add in checkErrors
	RCreateReviewNotCommit                        = "get create review not commit"                     // TODO: add in checkErrors
	RGetReviewNotSelect                           = "get get review not select"                        // TODO: add in checkErrors
	RGetReviewNotScan                             = "get get review not scan"                          // TODO: add in checkErrors
	RCreateReviewNotInsert                        = "get get review not insert"                        // TODO: add in checkErrors
	RGetReviewEmpty                               = "review is empty"                                  // TODO: add in checkErrors
	RSearchCategoryTransactionNotCreate           = "transaction search category not create"           // TODO: add in checkErrors
	RSearchCategoryNotSelect                      = "search category not select"                       // TODO: add in checkErrors
	RSearchCategoryNotScan                        = "search category not scan"                         // TODO: add in checkErrors
	RSearchCategoryNotCommit                      = "search category not commit"                       // TODO: add in checkErrors
	RSearchRestaurantTransactionNotCreate         = "transaction search restaurant not create"         // TODO: add in checkErrors
	RSearchRestaurantNotSelect                    = "search restaurant not select"                     // TODO: add in checkErrors
	RSearchRestaurantNotScan                      = "search restaurant not scan"                       // TODO: add in checkErrors
	RSearchRestaurantEmpty                        = "search result empty"                              // TODO: add in checkErrors
	RSearchRestaurantNotCommit                    = "search restaurant not commit"                     // TODO: add in checkErrors
	RGetGeneralInfoTransactionNotCreate           = "transaction get general info not create"          // TODO: add in checkErrors
	RGetGeneralInfoNotScan                        = "get general info not scan"                        // TODO: add in checkErrors
	RGetGeneralInfoNotCommit                      = "get general info not commit"                      // TODO: add in checkErrors
	RGetFavoriteRestaurantsTransactionNotCreate   = "transaction get favourite restaurants not create" // TODO: add in checkErrors
	RGetFavoriteRestaurantsRestaurantsNotSelect   = "favourite restaurants not select"                 // TODO: add in checkErrors
	RGetFavoriteRestaurantsRestaurantsNotScan     = "transaction get favourite restaurants not create" // TODO: add in checkErrors
	RGetFavoriteRestaurantsInfoNotCommit          = "transaction get favourite restaurants not create" // TODO: add in checkErrors
	REditRestaurantInFavoriteTransactionNotCreate = "transaction get favourite restaurants not create" // TODO: add in checkErrors
	REditRestaurantInFavoriteRestaurantsNotSelect = "favourite restaurants not select"                 // TODO: add in checkErrors
	REditRestaurantInFavoriteRestaurantsNotScan   = "favourite restaurants not scan"                   // TODO: add in checkErrors
	REditRestaurantInFavoriteInfoNotCommit        = "transaction get favourite restaurants not commit" // TODO: add in checkErrors
	REditRestaurantInFavoriteRestaurantsNotDelete = "favorite restaurant not delete"                   // TODO: add in checkErrors
	RGetFavoriteRestaurantsRestaurantsNotExist    = "restaurant not exist"                             // TODO: add in checkErrors
	RIsFavoriteRestaurantsTransactionNotCreate    = "transaction is favorite not create"               // TODO: add in checkErrors
	RIsFavoriteRestaurantsRestaurantsNotSelect    = "favorite restaurant not select for check"         // TODO: add in checkErrors
	RIsFavoriteRestaurantsInfoNotCommit           = "transaction is favorite not commit"               // TODO: add in checkErrors
)

// Error of cart
const (
	CGetCartDishesNotFound                         = "dishes not found"
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
	CGetCartCartNotFound                           = "cart is void"                              // TODO: add in checkErrors
	CDeleteCartTransactionNotCreate                = "transaction delete cart not create"        // TODO: add in checkErrors
	CDeleteCartNotCommit                           = "transaction delete cart not commit"        // TODO: add in checkErrors
	CGetPriceDeliveryTransactionNotCreate          = "transaction get price delivery not create" // TODO: add in checkErrors
	CGetPriceDeliveryNotCommit                     = "transaction get price delivery not commit" // TODO: add in checkErrors
	CAddPromoCodeTransactionNotCreate              = "transaction add promo code not create"     // TODO: add in checkErrors
	CAddPromoCodeNotCommit                         = "transaction add promo code not commit"     // TODO: add in checkErrors
	CAddPromoCodeNotDelete                         = "promo not delete"                          // TODO: add in checkErrors
	CAddPromoCodeNotInsert                         = "promo not insert"                          // TODO: add in checkErrors
)

// Error of order
const (
	OCreateOrderTransactionNotCreate         = "transaction create order not create"  // TODO: add in checkErrors
	OCreateOrderNotCommit                    = "create order not commit"              // TODO: add in checkErrors
	OGetOrdersTransactionNotCreate           = "transaction get orders not create"    // TODO: add in checkErrors
	OGetOrdersNotCommit                      = "get orders not commit"                // TODO: add in checkErrors
	OCreateOrderOrderUserNotInsert           = "not insert in order_user"             // TODO: add in checkErrors
	OCreateOrderOrderRadiosListUserNotInsert = "not insert in order_radios_list"      // TODO: add in checkErrors
	OCreateOrderOrderStructureListNotInsert  = "not insert in order_structure_list"   // TODO: add in checkErrors
	OCreateOrderOrderListNotInsert           = "not insert in order_list"             // TODO: add in checkErrors
	OCreateOrderCountNotUpdate               = "count dish not update"                // TODO: add in checkErrors
	OCreateOrderCountNotCorrect              = "dishes not enough"                    // TODO: add in checkErrors
	OGetOrdersOrdersIsVoid                   = "orders is void"                       // TODO: add in checkErrors
	OGetOrdersNotSelect                      = "orders not selected"                  // TODO: add in checkErrors
	OGetOrdersNotScan                        = "orders not scan"                      // TODO: add in checkErrors
	OGetOrderTransactionNotCreate            = "transaction get order not create"     // TODO: add in checkErrors
	OGetOrderNotSelect                       = "order not selected"                   // TODO: add in checkErrors
	OGetOrderNotScan                         = "order not scan"                       // TODO: add in checkErrors
	OGetOrderNotCommit                       = "get order not commit"                 // TODO: add in checkErrors
	OUpdateStatusOrderTransactionNotCreate   = "transaction update status not create" // TODO: add in checkErrors
	OUpdateStatusOrderNotUpdate              = "order not update"                     // TODO: add in checkErrors
	OUpdateStatusOrderNotCommit              = "transaction update status not commit" // TODO: add in checkErrors
	OGetCartCartNoActual                     = "cart not valid"                       // TODO: add in checkErrors
	OGetOrderNotExist                        = "order not exist"                      // TODO: add in checkErrors
)

// Error of promo codes
const (
	PGetTypePromoCodeTransactionNotCreate          = "transaction get type promo code not create"        // TODO: add in checkErrors
	PGetTypePromoCodeNotCommit                     = "transaction get type promo code not commit"        // TODO: add in checkErrors
	PGetTypePromoCodeRestaurantsNotFound           = "type not found"                                    // TODO: add in checkErrors
	PGetTypePromoCodeRestaurantsNotSelect          = "type not select"                                   // TODO: add in checkErrors
	PActiveCostForFreeDeliveryTransactionNotCreate = "transaction get cost for free delivery not create" // TODO: add in checkErrors
	PActiveCostForFreeDeliveryNotCommit            = "transaction get cost for free delivery not commit" // TODO: add in checkErrors
	PActiveCostForFreeDeliveryRestaurantsNotFound  = "cost for free delivery not found"                  // TODO: add in checkErrors
	PActiveCostForFreeDeliveryRestaurantsNotSelect = "cost for free delivery not select"                 // TODO: add in checkErrors
	PActiveCostForSaleTransactionNotCreate         = "transaction get cost for sale not create"          // TODO: add in checkErrors
	PActiveCostForSaleNotCommit                    = "transaction get cost for sale not commit"          // TODO: add in checkErrors
	PActiveCostForSaleRestaurantsNotFound          = "cost for sale not found"                           // TODO: add in checkErrors
	PActiveCostForSaleRestaurantsNotSelect         = "cost for sale not select"                          // TODO: add in checkErrors
	PActiveTimeForSaleTransactionNotCreate         = "transaction get Time for sale not create"          // TODO: add in checkErrors
	PActiveTimeForSaleNotCommit                    = "transaction get Time for sale not commit"          // TODO: add in checkErrors
	PActiveTimeForSaleRestaurantsNotFound          = "Time for sale not found"                           // TODO: add in checkErrors
	PActiveTimeForSaleRestaurantsNotSelect         = "Time for sale not select"                          // TODO: add in checkErrors
	PActiveCostForFreeDishTransactionNotCreate     = "transaction for free dish not create"              // TODO: add in checkErrors
	PActiveCostForFreeDishRestaurantsNotFound      = "free dish not found"                               // TODO: add in checkErrors
	PActiveCostForFreeDishRestaurantsNotSelect     = "free dish not select"                              // TODO: add in checkErrors
	PActiveCostForFreeDishNotCommit                = "transaction for free dish not commit"              // TODO: add in checkErrors
)
