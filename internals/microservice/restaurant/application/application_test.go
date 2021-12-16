package application

import (
	resPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant/application/mocks"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var AllRestaurants = []struct {
	testName             string
	out                  *resPkg.AllRestaurantsPromo
	outErr               string
	outQueryRestaurant   *resPkg.AllRestaurantsPromo
	errQueryRestaurant   error
	countQueryRestaurant int
	outQueryPromoCodes   []resPkg.Promocode
	errQueryPromoCodes   error
	countQueryPromoCodes int
}{
	{
		testName: "First",
		out: &resPkg.AllRestaurantsPromo{
			Restaurant: []resPkg.Restaurants(nil),
			AllTags:    []resPkg.Tag(nil),
			AllPromo:   []resPkg.Promocode{},
		},
		outErr:               "",
		outQueryRestaurant:   &resPkg.AllRestaurantsPromo{},
		outQueryPromoCodes:   []resPkg.Promocode{},
		errQueryRestaurant:   nil,
		errQueryPromoCodes:   nil,
		countQueryRestaurant: 1,
		countQueryPromoCodes: 1,
	},
	{
		testName:             "Second",
		out:                  nil,
		outErr:               "text",
		outQueryRestaurant:   &resPkg.AllRestaurantsPromo{},
		outQueryPromoCodes:   []resPkg.Promocode{},
		errQueryRestaurant:   errors.New("text"),
		errQueryPromoCodes:   nil,
		countQueryRestaurant: 1,
		countQueryPromoCodes: 0,
	},
}

func TestAllRestaurants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurantInterface(ctrl)
	for _, tt := range AllRestaurants {
		m.
			EXPECT().
			GetRestaurants().
			Return(tt.outQueryRestaurant, tt.errQueryRestaurant).
			Times(tt.countQueryRestaurant)
		m.
			EXPECT().
			GetPromoCodes().
			Return(tt.outQueryPromoCodes, tt.errQueryPromoCodes).
			Times(tt.countQueryPromoCodes)
		test := Restaurant{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.AllRestaurantsPromo()
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				if err == nil {
					require.NotNil(t, err, fmt.Sprintf("Expected: %s\nbut got: nil", tt.outErr))
				}
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var GetRestaurant = []struct {
	testName                              string
	out                                   *resPkg.RestaurantId
	outErr                                string
	inputRestaurantId                     int
	inputClientId                         int
	inputGetRestaurantRestaurantId        int
	inputGetRestaurantClientId            int
	outGetRestaurant                      *resPkg.RestaurantId
	errGetRestaurant                      error
	countGetRestaurant                    int
	inputIsFavoriteRestaurantClientId     int
	inputIsFavoriteRestaurantRestaurantId int
	outIsFavoriteRestaurant               bool
	errIsFavoriteRestaurant               error
	countIsFavoriteRestaurant             int
	inputGetTagsRestaurant                int
	outGetTagsRestaurant                  []resPkg.Tag
	errGetTagsRestaurant                  error
	countGetTagsRestaurant                int
	inputGetMenu                          int
	outGetMenu                            []resPkg.Menu
	errGetMenu                            error
	countGetMenu                          int
}{
	{
		testName: "First",
		out: &resPkg.RestaurantId{
			Id:                  0,
			Img:                 "",
			Name:                "",
			CostForFreeDelivery: 0,
			MinDelivery:         0,
			MaxDelivery:         0,
			Rating:              0,
			Favourite:           true,
			Tags:                []resPkg.Tag{},
			Menu:                []resPkg.Menu{},
		},
		outErr:                                "",
		inputRestaurantId:                     1,
		inputClientId:                         1,
		inputGetRestaurantRestaurantId:        1,
		inputGetRestaurantClientId:            1,
		outGetRestaurant:                      &resPkg.RestaurantId{},
		errGetRestaurant:                      nil,
		countGetRestaurant:                    1,
		inputIsFavoriteRestaurantClientId:     1,
		inputIsFavoriteRestaurantRestaurantId: 1,
		outIsFavoriteRestaurant:               true,
		errIsFavoriteRestaurant:               nil,
		countIsFavoriteRestaurant:             1,
		inputGetTagsRestaurant:                1,
		outGetTagsRestaurant:                  []resPkg.Tag{},
		errGetTagsRestaurant:                  nil,
		countGetTagsRestaurant:                1,
		inputGetMenu:                          1,
		outGetMenu:                            []resPkg.Menu{},
		errGetMenu:                            nil,
		countGetMenu:                          1,
	},
	{
		testName:                              "Second",
		out:                                   nil,
		outErr:                                "text",
		inputRestaurantId:                     1,
		inputClientId:                         1,
		inputGetRestaurantRestaurantId:        1,
		inputGetRestaurantClientId:            1,
		outGetRestaurant:                      &resPkg.RestaurantId{},
		errGetRestaurant:                      errors.New("text"),
		countGetRestaurant:                    1,
		inputIsFavoriteRestaurantClientId:     1,
		inputIsFavoriteRestaurantRestaurantId: 1,
		outIsFavoriteRestaurant:               true,
		errIsFavoriteRestaurant:               nil,
		countIsFavoriteRestaurant:             0,
		inputGetTagsRestaurant:                1,
		outGetTagsRestaurant:                  []resPkg.Tag{},
		errGetTagsRestaurant:                  nil,
		countGetTagsRestaurant:                0,
		inputGetMenu:                          1,
		outGetMenu:                            []resPkg.Menu{},
		errGetMenu:                            nil,
		countGetMenu:                          0,
	},
	{
		testName:                              "Third",
		out:                                   nil,
		outErr:                                "text",
		inputRestaurantId:                     1,
		inputClientId:                         1,
		inputGetRestaurantRestaurantId:        1,
		inputGetRestaurantClientId:            1,
		outGetRestaurant:                      &resPkg.RestaurantId{},
		errGetRestaurant:                      nil,
		countGetRestaurant:                    1,
		inputIsFavoriteRestaurantClientId:     1,
		inputIsFavoriteRestaurantRestaurantId: 1,
		outIsFavoriteRestaurant:               true,
		errIsFavoriteRestaurant:               nil,
		countIsFavoriteRestaurant:             1,
		inputGetTagsRestaurant:                1,
		outGetTagsRestaurant:                  []resPkg.Tag{},
		errGetTagsRestaurant:                  errors.New("text"),
		countGetTagsRestaurant:                1,
		inputGetMenu:                          1,
		outGetMenu:                            []resPkg.Menu{},
		errGetMenu:                            nil,
		countGetMenu:                          0,
	},
	{
		testName:                              "Fourth",
		out:                                   nil,
		outErr:                                "text",
		inputRestaurantId:                     1,
		inputClientId:                         1,
		inputGetRestaurantRestaurantId:        1,
		inputGetRestaurantClientId:            1,
		outGetRestaurant:                      &resPkg.RestaurantId{},
		errGetRestaurant:                      nil,
		countGetRestaurant:                    1,
		inputIsFavoriteRestaurantClientId:     1,
		inputIsFavoriteRestaurantRestaurantId: 1,
		outIsFavoriteRestaurant:               true,
		errIsFavoriteRestaurant:               nil,
		countIsFavoriteRestaurant:             1,
		inputGetTagsRestaurant:                1,
		outGetTagsRestaurant:                  []resPkg.Tag{},
		errGetTagsRestaurant:                  nil,
		countGetTagsRestaurant:                1,
		inputGetMenu:                          1,
		outGetMenu:                            []resPkg.Menu{},
		errGetMenu:                            errors.New("text"),
		countGetMenu:                          1,
	},
	{
		testName:                              "Sixth",
		out:                                   nil,
		outErr:                                "text",
		inputRestaurantId:                     1,
		inputClientId:                         1,
		inputGetRestaurantRestaurantId:        1,
		inputGetRestaurantClientId:            1,
		outGetRestaurant:                      &resPkg.RestaurantId{},
		errGetRestaurant:                      nil,
		countGetRestaurant:                    1,
		inputIsFavoriteRestaurantClientId:     1,
		inputIsFavoriteRestaurantRestaurantId: 1,
		outIsFavoriteRestaurant:               false,
		errIsFavoriteRestaurant:               errors.New("text"),
		countIsFavoriteRestaurant:             1,
		inputGetTagsRestaurant:                1,
		outGetTagsRestaurant:                  []resPkg.Tag{},
		errGetTagsRestaurant:                  nil,
		countGetTagsRestaurant:                0,
		inputGetMenu:                          1,
		outGetMenu:                            []resPkg.Menu{},
		errGetMenu:                            nil,
		countGetMenu:                          0,
	},
}

func TestGetRestaurant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurantInterface(ctrl)
	for _, tt := range GetRestaurant {
		m.
			EXPECT().
			GetRestaurant(tt.inputGetRestaurantRestaurantId).
			Return(tt.outGetRestaurant, tt.errGetRestaurant).
			Times(tt.countGetRestaurant)
		m.
			EXPECT().
			IsFavoriteRestaurant(tt.inputIsFavoriteRestaurantClientId, tt.inputIsFavoriteRestaurantRestaurantId).
			Return(tt.outIsFavoriteRestaurant, tt.errIsFavoriteRestaurant).
			Times(tt.countIsFavoriteRestaurant)
		m.
			EXPECT().
			GetTagsRestaurant(tt.inputGetTagsRestaurant).
			Return(tt.outGetTagsRestaurant, tt.errGetTagsRestaurant).
			Times(tt.countGetTagsRestaurant)
		m.
			EXPECT().
			GetMenu(tt.inputGetMenu).
			Return(tt.outGetMenu, tt.errGetMenu).
			Times(tt.countGetMenu)
		test := Restaurant{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.GetRestaurant(tt.inputRestaurantId, tt.inputClientId)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				if err == nil {
					require.NotNil(t, err, fmt.Sprintf("Expected: %s\nbut got: nil", tt.outErr))
				}
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var RestaurantDishes = []struct {
	testName             string
	inputRestId          int
	inputDishId          int
	out                  *resPkg.Dishes
	outErr               string
	inputGetDishesRestId int
	inputGetDishesDishId int
	outGetDishes         *resPkg.Dishes
	errGetDishes         error
	countGetDishes       int
}{
	{
		testName:    "First",
		inputRestId: 1,
		inputDishId: 1,
		out: &resPkg.Dishes{
			Id:          0,
			Img:         "",
			Title:       "",
			Cost:        0,
			Ccal:        0,
			Description: "",
			Radios:      []resPkg.Radios(nil),
			Ingredient:  []resPkg.Ingredients(nil),
		},
		outErr:               "text",
		inputGetDishesRestId: 1,
		inputGetDishesDishId: 1,
		outGetDishes: &resPkg.Dishes{
			Id:          0,
			Img:         "",
			Title:       "",
			Cost:        0,
			Ccal:        0,
			Description: "",
			Radios:      []resPkg.Radios(nil),
			Ingredient:  []resPkg.Ingredients(nil),
		},
		errGetDishes:   errors.New("text"),
		countGetDishes: 1,
	},
	{
		testName:    "Second",
		inputRestId: 1,
		inputDishId: 1,
		out: &resPkg.Dishes{
			Id:          0,
			Img:         "",
			Title:       "",
			Cost:        0,
			Ccal:        0,
			Description: "",
			Radios:      []resPkg.Radios(nil),
			Ingredient:  []resPkg.Ingredients(nil),
		},
		outErr:               "",
		inputGetDishesRestId: 1,
		inputGetDishesDishId: 1,
		outGetDishes:         &resPkg.Dishes{},
		errGetDishes:         nil,
		countGetDishes:       1,
	},
}

func TestRestaurantDishes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurantInterface(ctrl)
	for _, tt := range RestaurantDishes {
		m.
			EXPECT().
			GetDishes(tt.inputGetDishesRestId, tt.inputGetDishesDishId).
			Return(tt.outGetDishes, tt.errGetDishes).
			Times(tt.countGetDishes)
		test := Restaurant{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.RestaurantDishes(tt.inputRestId, tt.inputDishId)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				if err == nil {
					require.NotNil(t, err, fmt.Sprintf("Expected: %s\nbut got: nil", tt.outErr))
				}
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var CreateReview = []struct {
	testName                   string
	inputId                    int
	inputNewReview             resPkg.NewReview
	out                        *resPkg.NewReview
	outErr                     string
	inputCreateReviewId        int
	inputCreateReviewNewReview resPkg.NewReview
	errCreateReview            error
	countCreateReview          int
}{
	{
		testName:                   "First",
		inputId:                    1,
		inputNewReview:             resPkg.NewReview{},
		out:                        nil,
		outErr:                     "text",
		inputCreateReviewId:        1,
		inputCreateReviewNewReview: resPkg.NewReview{},
		errCreateReview:            errors.New("text"),
		countCreateReview:          1,
	},
	{
		testName:                   "Second",
		inputId:                    1,
		inputNewReview:             resPkg.NewReview{},
		out:                        &resPkg.NewReview{},
		outErr:                     "",
		inputCreateReviewId:        1,
		inputCreateReviewNewReview: resPkg.NewReview{},
		errCreateReview:            nil,
		countCreateReview:          1,
	},
}

func TestCreateReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurantInterface(ctrl)
	for _, tt := range CreateReview {
		m.
			EXPECT().
			CreateReview(tt.inputCreateReviewId, tt.inputCreateReviewNewReview).
			Return(tt.errCreateReview).
			Times(tt.countCreateReview)
		test := Restaurant{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			err := test.CreateReview(tt.inputId, tt.inputNewReview)
			if tt.outErr != "" {
				if err == nil {
					require.NotNil(t, err, fmt.Sprintf("Expected: %s\nbut got: nil", tt.outErr))
				}
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var GetReview = []struct {
	testName                              string
	inputIdRestaurant                     int
	inputIdClient                         int
	out                                   *resPkg.ResReview
	outErr                                string
	inputGetReview                        int
	outGetReview                          []resPkg.Review
	errGetReview                          error
	countGetReview                        int
	inputIsFavoriteRestaurantIdClient     int
	inputIsFavoriteRestaurantIdRestaurant int
	outIsFavoriteRestaurant               bool
	errIsFavoriteRestaurant               error
	countIsFavoriteRestaurant             int
	inputGetRestaurantIdRestaurant        int
	inputGetRestaurantIdClient            int
	outGetRestaurant                      *resPkg.RestaurantId
	errGetRestaurant                      error
	countGetRestaurant                    int
	inputGetTagsRestaurant                int
	outGetTagsRestaurant                  []resPkg.Tag
	errGetTagsRestaurant                  error
	countGetTagsRestaurant                int
}{
	{
		testName:          "First",
		inputIdRestaurant: 1,
		inputIdClient:     1,
		out: &resPkg.ResReview{
			Id:                  0,
			Img:                 "",
			Name:                "",
			CostForFreeDelivery: 0,
			MinDelivery:         0,
			MaxDelivery:         0,
			Rating:              0,
			Tags:                []resPkg.Tag{},
			Reviews:             []resPkg.Review{},
			Status:              true,
		},
		outErr:                                "",
		inputGetReview:                        1,
		outGetReview:                          []resPkg.Review{},
		errGetReview:                          nil,
		countGetReview:                        1,
		inputIsFavoriteRestaurantIdClient:     1,
		inputIsFavoriteRestaurantIdRestaurant: 1,
		outIsFavoriteRestaurant:               true,
		errIsFavoriteRestaurant:               nil,
		countIsFavoriteRestaurant:             1,
		inputGetRestaurantIdRestaurant:        1,
		inputGetRestaurantIdClient:            1,
		outGetRestaurant:                      &resPkg.RestaurantId{},
		errGetRestaurant:                      nil,
		countGetRestaurant:                    1,
		inputGetTagsRestaurant:                1,
		outGetTagsRestaurant:                  []resPkg.Tag{},
		errGetTagsRestaurant:                  nil,
		countGetTagsRestaurant:                1,
	},
}

func TestGetReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurantInterface(ctrl)
	for _, tt := range GetReview {
		m.
			EXPECT().
			GetReview(tt.inputGetReview).
			Return(tt.outGetReview, tt.errGetReview).
			Times(tt.countGetReview)
		m.
			EXPECT().
			IsFavoriteRestaurant(tt.inputIsFavoriteRestaurantIdClient, tt.inputIsFavoriteRestaurantIdRestaurant).
			Return(tt.outIsFavoriteRestaurant, tt.errIsFavoriteRestaurant).
			Times(tt.countIsFavoriteRestaurant)
		m.
			EXPECT().
			GetRestaurant(tt.inputGetRestaurantIdRestaurant).
			Return(tt.outGetRestaurant, tt.errGetRestaurant).
			Times(tt.countGetRestaurant)
		m.
			EXPECT().
			GetTagsRestaurant(tt.inputGetTagsRestaurant).
			Return(tt.outGetTagsRestaurant, tt.errGetTagsRestaurant).
			Times(tt.countGetTagsRestaurant)
		test := Restaurant{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.GetReview(tt.inputIdRestaurant, tt.inputIdClient)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				if err == nil {
					require.NotNil(t, err, fmt.Sprintf("Expected: %s\nbut got: nil", tt.outErr))
				}
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var SearchRestaurant = []struct {
	testName                      string
	input                         string
	out                           []resPkg.Restaurants
	outErr                        string
	inputSearchCategory           string
	outSearchCategory             []int
	errSearchCategory             error
	countSearchCategory           int
	inputSearchRestaurant         string
	outSearchRestaurant           []int
	errSearchRestaurant           error
	countSearchRestaurant         int
	inputGetGeneralInfoRestaurant int
	outGetGeneralInfoRestaurant   *resPkg.Restaurants
	errGetGeneralInfoRestaurant   error
	countGetGeneralInfoRestaurant int
}{
	{
		testName:                      "First",
		input:                         "Cafe",
		out:                           []resPkg.Restaurants{{}},
		outErr:                        "",
		inputSearchCategory:           "Cafe",
		outSearchCategory:             []int{1},
		errSearchCategory:             nil,
		countSearchCategory:           1,
		inputSearchRestaurant:         "",
		outSearchRestaurant:           []int{},
		errSearchRestaurant:           nil,
		countSearchRestaurant:         0,
		inputGetGeneralInfoRestaurant: 1,
		outGetGeneralInfoRestaurant:   &resPkg.Restaurants{},
		errGetGeneralInfoRestaurant:   nil,
		countGetGeneralInfoRestaurant: 1,
	},
}

func TestSearchRestaurant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurantInterface(ctrl)
	for _, tt := range SearchRestaurant {
		m.
			EXPECT().
			SearchCategory(tt.inputSearchCategory).
			Return(tt.outSearchCategory, tt.errSearchCategory).
			Times(tt.countSearchCategory)
		m.
			EXPECT().
			SearchRestaurant(tt.inputSearchRestaurant).
			Return(tt.outSearchRestaurant, tt.errSearchRestaurant).
			Times(tt.countSearchRestaurant)
		m.
			EXPECT().
			GetGeneralInfoRestaurant(tt.inputGetGeneralInfoRestaurant).
			Return(tt.outGetGeneralInfoRestaurant, tt.errGetGeneralInfoRestaurant).
			Times(tt.countGetGeneralInfoRestaurant)
		test := Restaurant{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.SearchRestaurant(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				if err == nil {
					require.NotNil(t, err, fmt.Sprintf("Expected: %s\nbut got: nil", tt.outErr))
				}
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var GetFavoriteRestaurants = []struct {
	testName                    string
	input                       int
	out                         []resPkg.Restaurants
	outErr                      string
	inputGetFavoriteRestaurants int
	outGetFavoriteRestaurants   []resPkg.Restaurants
	errGetFavoriteRestaurants   error
	countGetFavoriteRestaurants int
}{
	{
		testName:                    "First",
		input:                       1,
		out:                         nil,
		outErr:                      "text",
		inputGetFavoriteRestaurants: 1,
		outGetFavoriteRestaurants:   nil,
		errGetFavoriteRestaurants:   errors.New("text"),
		countGetFavoriteRestaurants: 1,
	},
	{
		testName:                    "Second",
		input:                       1,
		out:                         []resPkg.Restaurants{},
		outErr:                      "",
		inputGetFavoriteRestaurants: 1,
		outGetFavoriteRestaurants:   []resPkg.Restaurants{},
		errGetFavoriteRestaurants:   nil,
		countGetFavoriteRestaurants: 1,
	},
}

func TestGetFavoriteRestaurants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurantInterface(ctrl)
	for _, tt := range GetFavoriteRestaurants {
		m.
			EXPECT().
			GetFavoriteRestaurants(tt.inputGetFavoriteRestaurants).
			Return(tt.outGetFavoriteRestaurants, tt.errGetFavoriteRestaurants).
			Times(tt.countGetFavoriteRestaurants)
		test := Restaurant{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.GetFavoriteRestaurants(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				if err == nil {
					require.NotNil(t, err, fmt.Sprintf("Expected: %s\nbut got: nil", tt.outErr))
				}
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var EditRestaurantInFavorite = []struct {
	testName                                  string
	inputIdRestaurant                         int
	inputIdClient                             int
	out                                       bool
	outErr                                    string
	inputEditRestaurantInFavoriteRestaurantId int
	inputEditRestaurantInFavoriteClientId     int
	outEditRestaurantInFavorite               bool
	errEditRestaurantInFavorite               error
	countEditRestaurantInFavorite             int
}{
	{
		testName:          "First",
		inputIdRestaurant: 1,
		inputIdClient:     1,
		out:               false,
		outErr:            "text",
		inputEditRestaurantInFavoriteRestaurantId: 1,
		inputEditRestaurantInFavoriteClientId:     1,
		outEditRestaurantInFavorite:               false,
		errEditRestaurantInFavorite:               errors.New("text"),
		countEditRestaurantInFavorite:             1,
	},
	{
		testName:          "Second",
		inputIdRestaurant: 1,
		inputIdClient:     1,
		out:               true,
		outErr:            "",
		inputEditRestaurantInFavoriteRestaurantId: 1,
		inputEditRestaurantInFavoriteClientId:     1,
		outEditRestaurantInFavorite:               true,
		errEditRestaurantInFavorite:               nil,
		countEditRestaurantInFavorite:             1,
	},
}

func TestEditRestaurantInFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurantInterface(ctrl)
	for _, tt := range EditRestaurantInFavorite {
		m.
			EXPECT().
			EditRestaurantInFavorite(tt.inputEditRestaurantInFavoriteRestaurantId, tt.inputEditRestaurantInFavoriteClientId).
			Return(tt.outEditRestaurantInFavorite, tt.errEditRestaurantInFavorite).
			Times(tt.countEditRestaurantInFavorite)
		test := Restaurant{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.EditRestaurantInFavorite(tt.inputIdRestaurant, tt.inputIdClient)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				if err == nil {
					require.NotNil(t, err, fmt.Sprintf("Expected: %s\nbut got: nil", tt.outErr))
				}
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %v\nbut got: %v", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
