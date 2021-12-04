package application

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/restaurant/application"
	"fmt"
	"github.com/stretchr/testify/require"
)

var ApplicationAllRestaurants = []struct {
	testName string
	out      []rest.Restaurants
	outErr   string
	err      error
}{
	{
		testName: "One",
		out:      []rest.Restaurants{},
		err:      nil,
		outErr:   "",
	},
	{
		testName: "Two",
		out:      nil,
		err:      errors.New("text"),
		outErr:   "text",
	},
}

func TestApplicationAllRestaurants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurant(ctrl)
	for _, tt := range ApplicationAllRestaurants {
		m.
			EXPECT().
			GetRestaurants().
			Return([]rest.Restaurants{}, tt.err)
		test := application.Restaurant{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.AllRestaurants()
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationGetRestaurant = []struct {
	testName                       string
	out                            *rest.RestaurantId
	outErr                         string
	input                          int
	inputGetGeneralInfoRestaurant  int
	resultGetGeneralInfoRestaurant *rest.RestaurantId
	errGetGeneralInfoRestaurant    error
	countGetGeneralInfoRestaurant  int
	inputGetTagsRestaurant         int
	resultGetTagsRestaurant        []rest.Tag
	errGetTagsRestaurant           error
	countGetTagsRestaurant         int
	inputGetMenu                   int
	resultGetMenu                  []rest.Menu
	errGetMenu                     error
	countGetMenu                   int
}{
	{
		testName: "One",
		out: &rest.RestaurantId{Id: 0, Img: "", Name: "", CostForFreeDelivery: 0,
			MinDelivery: 0, MaxDelivery: 0, Rating: 0, Tags: []rest.Tag{}, Menu: []rest.Menu{}},
		outErr:                         "",
		input:                          1,
		inputGetGeneralInfoRestaurant:  1,
		resultGetGeneralInfoRestaurant: &rest.RestaurantId{},
		errGetGeneralInfoRestaurant:    nil,
		countGetGeneralInfoRestaurant:  1,
		inputGetTagsRestaurant:         1,
		resultGetTagsRestaurant:        []rest.Tag{},
		errGetTagsRestaurant:           nil,
		countGetTagsRestaurant:         1,
		inputGetMenu:                   1,
		resultGetMenu:                  []rest.Menu{},
		errGetMenu:                     nil,
		countGetMenu:                   1,
	},
	{
		testName:                       "Two",
		out:                            nil,
		outErr:                         "text",
		input:                          1,
		inputGetGeneralInfoRestaurant:  1,
		resultGetGeneralInfoRestaurant: &rest.RestaurantId{},
		errGetGeneralInfoRestaurant:    errors.New("text"),
		countGetGeneralInfoRestaurant:  1,
		inputGetTagsRestaurant:         1,
		resultGetTagsRestaurant:        []rest.Tag{},
		errGetTagsRestaurant:           nil,
		countGetTagsRestaurant:         0,
		inputGetMenu:                   1,
		resultGetMenu:                  []rest.Menu{},
		errGetMenu:                     nil,
		countGetMenu:                   0,
	},
	{
		testName:                       "Three",
		out:                            nil,
		outErr:                         "text",
		input:                          1,
		inputGetGeneralInfoRestaurant:  1,
		resultGetGeneralInfoRestaurant: &rest.RestaurantId{},
		errGetGeneralInfoRestaurant:    nil,
		countGetGeneralInfoRestaurant:  1,
		inputGetTagsRestaurant:         1,
		resultGetTagsRestaurant:        []rest.Tag{},
		errGetTagsRestaurant:           errors.New("text"),
		countGetTagsRestaurant:         1,
		inputGetMenu:                   1,
		resultGetMenu:                  []rest.Menu{},
		errGetMenu:                     nil,
		countGetMenu:                   0,
	},
	{
		testName:                       "Four",
		out:                            nil,
		outErr:                         "text",
		input:                          1,
		inputGetGeneralInfoRestaurant:  1,
		resultGetGeneralInfoRestaurant: &rest.RestaurantId{},
		errGetGeneralInfoRestaurant:    nil,
		countGetGeneralInfoRestaurant:  1,
		inputGetTagsRestaurant:         1,
		resultGetTagsRestaurant:        []rest.Tag{},
		errGetTagsRestaurant:           nil,
		countGetTagsRestaurant:         1,
		inputGetMenu:                   1,
		resultGetMenu:                  []rest.Menu{},
		errGetMenu:                     errors.New("text"),
		countGetMenu:                   1,
	},
}

func TestApplicationGetRestaurant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurant(ctrl)
	for _, tt := range ApplicationGetRestaurant {
		m.
			EXPECT().
			GetGeneralInfoRestaurant(tt.inputGetGeneralInfoRestaurant).
			Return(tt.resultGetGeneralInfoRestaurant, tt.errGetGeneralInfoRestaurant).
			Times(tt.countGetGeneralInfoRestaurant)
		m.
			EXPECT().
			GetTagsRestaurant(tt.inputGetTagsRestaurant).
			Return(tt.resultGetTagsRestaurant, tt.errGetTagsRestaurant).
			Times(tt.countGetTagsRestaurant)
		m.
			EXPECT().
			GetMenu(tt.inputGetMenu).
			Return(tt.resultGetMenu, tt.errGetMenu).
			Times(tt.countGetMenu)
		test := application.Restaurant{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.GetRestaurant(tt.input)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}

var ApplicationRestaurantDishes = []struct {
	testName              string
	inputRestId           int
	inputDishId           int
	out                   *rest.Dishes
	outErr                string
	inputGetDishesRestId  int
	inputGetDishesDishId  int
	resultGetDishes       *rest.Dishes
	errGetDishes          error
	countGetDishes        int
	inputGetStructDishes  int
	resultGetStructDishes []rest.Ingredients
	errGetStructDishes    error
	countGetStructDishes  int
	inputGetRadios        int
	resultGetRadios       []rest.Radios
	errGetRadios          error
	countGetRadios        int
}{
	{
		testName:              "One",
		out:                   nil,
		outErr:                "text",
		inputGetDishesRestId:  1,
		inputGetDishesDishId:  1,
		resultGetDishes:       &rest.Dishes{},
		errGetDishes:          errors.New("text"),
		countGetDishes:        1,
		inputGetStructDishes:  1,
		resultGetStructDishes: []rest.Ingredients{},
		errGetStructDishes:    nil,
		countGetStructDishes:  0,
		inputGetRadios:        1,
		resultGetRadios:       []rest.Radios{},
		errGetRadios:          nil,
		countGetRadios:        0,
		inputRestId:           1,
		inputDishId:           1,
	},
	{
		testName:              "Two",
		out:                   nil,
		outErr:                "text",
		inputGetDishesRestId:  1,
		inputGetDishesDishId:  1,
		resultGetDishes:       &rest.Dishes{},
		errGetDishes:          nil,
		countGetDishes:        1,
		inputGetStructDishes:  1,
		resultGetStructDishes: []rest.Ingredients{},
		errGetStructDishes:    errors.New("text"),
		countGetStructDishes:  1,
		inputGetRadios:        1,
		resultGetRadios:       []rest.Radios{},
		errGetRadios:          nil,
		countGetRadios:        0,
		inputRestId:           1,
		inputDishId:           1,
	},
	{
		testName:              "Three",
		out:                   nil,
		outErr:                "text",
		inputGetDishesRestId:  1,
		inputGetDishesDishId:  1,
		resultGetDishes:       &rest.Dishes{},
		errGetDishes:          nil,
		countGetDishes:        1,
		inputGetStructDishes:  1,
		resultGetStructDishes: []rest.Ingredients{},
		errGetStructDishes:    nil,
		countGetStructDishes:  1,
		inputGetRadios:        1,
		resultGetRadios:       []rest.Radios{},
		errGetRadios:          errors.New("text"),
		countGetRadios:        1,
		inputRestId:           1,
		inputDishId:           1,
	},
	{
		testName: "Four",
		out: &rest.Dishes{Id: 0, Img: "", Title: "", Cost: 0, Ccal: 0, Description: "",
			Radios: []rest.Radios{}, Ingredient: []rest.Ingredients{}},
		outErr:                "",
		inputGetDishesRestId:  1,
		inputGetDishesDishId:  1,
		resultGetDishes:       &rest.Dishes{},
		errGetDishes:          nil,
		countGetDishes:        1,
		inputGetStructDishes:  1,
		resultGetStructDishes: []rest.Ingredients{},
		errGetStructDishes:    nil,
		countGetStructDishes:  1,
		inputGetRadios:        1,
		resultGetRadios:       []rest.Radios{},
		errGetRadios:          nil,
		countGetRadios:        1,
		inputRestId:           1,
		inputDishId:           1,
	},
}

func TestApplicationRestaurantDishes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockWrapperRestaurant(ctrl)
	for _, tt := range ApplicationRestaurantDishes {
		m.
			EXPECT().
			GetDishes(tt.inputGetDishesRestId, tt.inputGetDishesDishId).
			Return(tt.resultGetDishes, tt.errGetDishes).
			Times(tt.countGetDishes)
		m.
			EXPECT().
			GetStructDishes(tt.inputGetStructDishes).
			Return(tt.resultGetStructDishes, tt.errGetStructDishes).
			Times(tt.countGetStructDishes)
		m.
			EXPECT().
			GetRadios(tt.inputGetRadios).
			Return(tt.resultGetRadios, tt.errGetRadios).
			Times(tt.countGetRadios)
		test := application.Restaurant{DB: m}
		t.Run(tt.testName, func(t *testing.T) {
			result, err := test.RestaurantDishes(tt.inputRestId, tt.inputDishId)
			require.Equal(t, tt.out, result, fmt.Sprintf("Expected: %v\nbut got: %v", tt.out, result))
			if tt.outErr != "" {
				require.EqualError(t, err, tt.outErr, fmt.Sprintf("Expected: %s\nbut got: %s", tt.outErr, err.Error()))
			} else {
				require.Nil(t, err, fmt.Sprintf("Expected: nil\nbut got: %s", err))
			}
		})
	}
}
