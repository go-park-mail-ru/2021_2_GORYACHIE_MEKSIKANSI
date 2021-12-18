//go:generate easyjson -no_std_marshalers errors.go
package myerror

type MultiLogger interface {
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Sync() error
}

//easyjson:json
type ResultError struct {
	Status  int    `json:"status"`
	Explain string `json:"explain,omitempty"`
}

//easyjson:json
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
	MCreateDBNotConnect           = "db not connect"
	MCreateDBCreateFileNotFound   = "createtables.sql not found"
	MCreateDBDeleteFileNotFound   = "deletetables.sql not found"
	MCreateDBFillFileNotFound     = "fill.sql not found"
	MCreateDBNotCreateTables      = "table not create"
	MCreateDBNotDeleteTables      = "table not delete"
	MCreateDBNotFillTables        = "table not fill"
	MCreateDBTransactionNotCreate = "transaction setup not create"
	MCreateDBNotCommit            = "transaction setup not commit"
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
	PGetRoleByIdTransactionNotCreate       = "transaction role by id not create"
	PGetRoleByIdNotCommit                  = "role by id not commit"
	PGetProfileClientTransactionNotCreate  = "transaction get profile client not create"
	PGetProfileClientNotCommit             = "get profile client not commit"
	PUpdatePasswordTransactionNotCreate    = "transaction update password not create"
	PUpdatePasswordNotCommit               = "update password not commit"
	PAddAddressAddressNotAdd               = "address not insert"
	PAddDeleteAddressNotDelete             = "address not delete"
	PGetProfileHostTransactionNotCreate    = "transaction get profile host not create"
	PGetProfileHostNotCommit               = "transaction get profile host not commit"
	PGetProfileCourierTransactionNotCreate = "transaction get profile courier not create"
	PGetProfileCourierNotCommit            = "transaction get profile courier not commit"
	PUpdateNameTransactionNotCreate        = "transaction update name not create"
	PUpdateNameNotCommit                   = "transaction update name not commit"
	PUpdateEmailTransactionNotCreate       = "transaction update email not create"
	PUpdateEmailNotCommit                  = "transaction update email not commit"
	PUpdatePhoneTransactionNotCreate       = "transaction update phone not create"
	PUpdatePhoneNotCommit                  = "transaction update phone not commit"
	PUpdateAvatarTransactionNotCreate      = "transaction update avatar not create"
	PUpdateAvatarNotCommit                 = "transaction update avatar not commit"
	PUpdateBirthdayTransactionNotCreate    = "transaction update birthday not create"
	PUpdateBirthdayNotCommit               = "transaction update birthday not commit"
	PUpdateAddressTransactionNotCreate     = "transaction update address not create"
	PUpdateAddressNotCommit                = "transaction update address not commit"
	PAddAddressNotCreate                   = "transaction add address not create"
	PAddAddressNotCommit                   = "transaction add address not commit"
	PAddDeleteAddressTransactionNotCreate  = "transaction delete address not create"
	PAddDeleteAddressNotCommit             = "transaction delete address not commit"
	PUpdateBirthdayNotParse                = "birthday not parse"
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
	RGetMenuTransactionNotCreate                  = "transaction menu dishes not create"
	RGetRadiosTransactionNotCreate                = "transaction get radios not create"
	RGetRadiosNotCommit                           = "get radios not commit"
	RGetRestaurantsTransactionNotCreate           = "transaction get restaurants not create"
	RGetRestaurantsNotCommit                      = "get restaurants not commit"
	RGetRestaurantTransactionNotCreate            = "transaction get restaurant not create"
	RGetRestaurantNotCommit                       = "get restaurant not commit"
	RGetTagsRestaurantTransactionNotCreate        = "transaction get tag restaurant not create"
	RGetTagsRestaurantNotCommit                   = "get info restaurant not commit"
	RGetMenuNotCommit                             = "get menu not commit"
	RGetStructDishesTransactionNotCreate          = "transaction get struct dishes not create"
	RGetStructDishesNotCommit                     = "get get struct dishes not commit"
	RGetDishesTransactionNotCreate                = "transaction get dishes not create"
	RGetDishesNotCommit                           = "get get dishes not commit"
	RGetMenuDishesCategoryNotSelect               = "category not select"
	RGetReviewTransactionNotCreate                = "transaction get review not create"
	RGetReviewNotCommit                           = "get get review not commit"
	RCreateReviewTransactionNotCreate             = "transaction create review not create"
	RCreateReviewNotCommit                        = "get create review not commit"
	RGetReviewNotSelect                           = "get get review not select"
	RGetReviewNotScan                             = "get get review not scan"
	RCreateReviewNotInsert                        = "get get review not insert"
	RGetReviewEmpty                               = "review is empty"
	RSearchCategoryTransactionNotCreate           = "transaction search category not create"
	RSearchCategoryNotSelect                      = "search category not select"
	RSearchCategoryNotScan                        = "search category not scan"
	RSearchCategoryNotCommit                      = "search category not commit"
	RSearchRestaurantTransactionNotCreate         = "transaction search restaurant not create"
	RSearchRestaurantNotSelect                    = "search restaurant not select"
	RSearchRestaurantNotScan                      = "search restaurant not scan"
	RSearchRestaurantEmpty                        = "search result empty"
	RSearchRestaurantNotCommit                    = "search restaurant not commit"
	RGetGeneralInfoTransactionNotCreate           = "transaction get general info not create"
	RGetGeneralInfoNotScan                        = "get general info not scan"
	RGetGeneralInfoNotCommit                      = "get general info not commit"
	RGetFavoriteRestaurantsTransactionNotCreate   = "transaction get favourite restaurants not create"
	RGetFavoriteRestaurantsRestaurantsNotSelect   = "favourite restaurants not select"
	RGetFavoriteRestaurantsRestaurantsNotScan     = "transaction get favourite restaurants not create"
	RGetFavoriteRestaurantsInfoNotCommit          = "transaction get favourite restaurants not create"
	REditRestaurantInFavoriteTransactionNotCreate = "transaction get favourite restaurants not create"
	REditRestaurantInFavoriteRestaurantsNotSelect = "favourite restaurants not select"
	REditRestaurantInFavoriteRestaurantsNotScan   = "favourite restaurants not scan"
	REditRestaurantInFavoriteInfoNotCommit        = "transaction get favourite restaurants not commit"
	REditRestaurantInFavoriteRestaurantsNotDelete = "favorite restaurant not delete"
	RGetFavoriteRestaurantsRestaurantsNotExist    = "restaurant not exist"
	RIsFavoriteRestaurantsTransactionNotCreate    = "transaction is favorite not create"
	RIsFavoriteRestaurantsRestaurantsNotSelect    = "favorite restaurant not select for check"
	RIsFavoriteRestaurantsInfoNotCommit           = "transaction is favorite not commit"
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
	CGetPromoCodeTransactionNotCreate              = "transaction get promo code not create"
	CGetPromoCodeNotSelect                         = "promo code not select"
	CGetPromoCodeNotCommit                         = "transaction get promo code not commit"
)

// Error of order
const (
	OCreateOrderTransactionNotCreate         = "transaction create order not create"
	OCreateOrderNotCommit                    = "create order not commit"
	OGetOrdersTransactionNotCreate           = "transaction get orders not create"
	OGetOrdersNotCommit                      = "get orders not commit"
	OCreateOrderOrderUserNotInsert           = "not insert in order_user"
	OCreateOrderOrderRadiosListUserNotInsert = "not insert in order_radios_list"
	OCreateOrderOrderStructureListNotInsert  = "not insert in order_structure_list"
	OCreateOrderOrderListNotInsert           = "not insert in order_list"
	OCreateOrderCountNotUpdate               = "count dish not update"
	OCreateOrderCountNotCorrect              = "dishes not enough"
	OGetOrdersOrdersIsVoid                   = "orders is void"
	OGetOrdersNotSelect                      = "orders not selected"
	OGetOrdersNotScan                        = "orders not scan"
	OGetOrderTransactionNotCreate            = "transaction get order not create"
	OGetOrderNotSelect                       = "order not selected"
	OGetOrderNotScan                         = "order not scan"
	OGetOrderNotCommit                       = "get order not commit"
	OUpdateStatusOrderTransactionNotCreate   = "transaction update status not create"
	OUpdateStatusOrderNotUpdate              = "order not update"
	OUpdateStatusOrderNotCommit              = "transaction update status not commit"
	OGetCartCartNoActual                     = "cart not valid"
	OGetOrderNotExist                        = "order not exist"
)

// Error of promo codes
const (
	PGetTypePromoCodeTransactionNotCreate          = "transaction get type promo code not create"
	PGetTypePromoCodeNotCommit                     = "transaction get type promo code not commit"
	PGetTypePromoCodeRestaurantsNotFound           = "type not found"
	PGetTypePromoCodeRestaurantsNotSelect          = "type not select"
	PActiveCostForFreeDeliveryTransactionNotCreate = "transaction get cost for free delivery not create"
	PActiveCostForFreeDeliveryNotCommit            = "transaction get cost for free delivery not commit"
	PActiveCostForFreeDeliveryRestaurantsNotFound  = "cost for free delivery not found"
	PActiveCostForFreeDeliveryRestaurantsNotSelect = "cost for free delivery not select"
	PActiveCostForSaleTransactionNotCreate         = "transaction get cost for sale not create"
	PActiveCostForSaleNotCommit                    = "transaction get cost for sale not commit"
	PActiveCostForSaleRestaurantsNotFound          = "cost for sale not found"
	PActiveCostForSaleRestaurantsNotSelect         = "cost for sale not select"
	PActiveTimeForSaleTransactionNotCreate         = "transaction get Time for sale not create"
	PActiveTimeForSaleNotCommit                    = "transaction get Time for sale not commit"
	PActiveTimeForSaleRestaurantsNotFound          = "Time for sale not found"
	PActiveTimeForSaleRestaurantsNotSelect         = "Time for sale not select"
	PActiveCostForFreeDishTransactionNotCreate     = "transaction for free dish not create"
	PActiveCostForFreeDishRestaurantsNotFound      = "free dish not found"
	PActiveCostForFreeDishRestaurantsNotSelect     = "free dish not select"
	PActiveCostForFreeDishNotCommit                = "transaction for free dish not commit"
)
