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
	Alias string
	Text  string
}

func (e *Errors) Error() string {
	return e.Alias
}

// Error of main
const (
	MCreateDBNotConnect = "db not connect"
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
	RGetDishesDishesNotScan                       = "dishes not scan"
	RGetStructDishesStructDishesNotSelect         = "struct dishes not select"
	RGetStructDishesStructDishesNotScan           = "dishes not scan"
	RGetRadiosRadiosNotScan                       = "radios not scan"
	RGetTagsRestaurantNotSelect                   = "category not select"
	RGetRadiosRadiosNotSelect                     = "radios not select"
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
	RGetFavoriteRestaurantsRestaurantsNotSelect   = "transaction get favourite restaurants not create"
	RGetFavoriteRestaurantsRestaurantsNotScan     = "transaction get favourite restaurants not create"
	RGetFavoriteRestaurantsInfoNotCommit          = "transaction get favourite restaurants not create"
	REditRestaurantInFavoriteTransactionNotCreate = "transaction get favourite restaurants not create"
	REditRestaurantInFavoriteRestaurantsNotSelect = "favourite restaurants not select"
	REditRestaurantInFavoriteRestaurantsNotScan   = "favourite restaurants not scan"
	REditRestaurantInFavoriteInfoNotCommit        = "transaction get favourite restaurants not commmit"
	REditRestaurantInFavoriteRestaurantsNotDelete = "favorite restaurant not delete"
	RGetFavoriteRestaurantsRestaurantsNotExist    = "restaurant not exist"
	RGetStatusRestaurantTransactionNotCreate      = "transaction get restaurant not create"
	RGetStatusRestaurantNotSelect                 = "status not select"
	RGetStatusRestaurantNotCommit                 = "transaction get restaurant not commit"
)
