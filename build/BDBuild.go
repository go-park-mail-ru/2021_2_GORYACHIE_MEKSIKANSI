package build

import (
	"2021_2_GORYACHIE_MEKSIKANSI/config"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/MyError"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Util"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"strings"
)

const (
	Restaurants = 53
	Categorys   = 3
)

type Dish struct {
	Name               string
	Cost               int
	Description        string
	Protein            int
	Falt               int
	Kilocalorie        int
	Carbohydrates      int
	CategoryDishes     string
	CategoryRestaurant string
	Count              int
	Weight             int
	Avatar             string
	PlaceCategory      int
	Place              int
	Ingredient         []Ingredient
	Radios             []Radios
}

type Ingredient struct {
	Name          string
	Cost          int
	Protein       int
	Falt          int
	Carbohydrates int
	Kilocalorie   int
	CountElement  int
	Place         int
}

type Radios struct {
	Name          string
	Place         int
	RadiosElement []RadiosElement
}

type RadiosElement struct {
	Name          string
	Protein       int
	Falt          int
	Carbohydrates int
	Kilocalorie   int
	Place         int
}

func CreateDb(configDB config.Database, debug bool) (*pgxpool.Pool, error) {
	var err error
	conn, err := pgxpool.Connect(context.Background(),
		"postgres://"+configDB.UserName+":"+configDB.Password+
			"@"+configDB.Host+":"+configDB.Port+"/"+configDB.SchemaName)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.MCreateDBNotConnect,
		}
	}

	if debug {
		file, err := ioutil.ReadFile("build/PostgreSQL/DeleteTables.sql")
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.MCreateDBDeleteFileNotFound,
			}
		}

		requests := strings.Split(string(file), ";")
		for _, request := range requests {
			_, err = conn.Exec(context.Background(), request)
			if err != nil {
				return nil, &errPkg.Errors{
					Alias: errPkg.MCreateDBNotDeleteTables,
				}
			}
		}
	}

	file, err := ioutil.ReadFile("build/PostgreSQL/CreateTables.sql")
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.MCreateDBCreateFileNotFound,
		}
	}

	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		_, err = conn.Exec(context.Background(), request)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.MCreateDBNotCreateTables,
			}
		}
	}

	if debug {
		file, err := ioutil.ReadFile("build/PostgreSQL/Fill.sql")
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.MCreateDBFillFileNotFound,
			}
		}

		requests := strings.Split(string(file), ";")
		for _, request := range requests {
			_, err = conn.Exec(context.Background(), request)
			if err != nil {
				return nil, &errPkg.Errors{
					Alias: errPkg.MCreateDBNotFillTables,
				}
			}
		}

		fileDishes, err := ioutil.ReadFile("build/PostgreSQL/FillDishes.sql")
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.MCreateDBFillFileNotFound,
			}
		}

		requestsDishes := strings.Split(string(fileDishes), ";")

		fileCategory, err := ioutil.ReadFile("build/PostgreSQL/FillCategory.sql")
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.MCreateDBFillFileNotFound,
			}
		}

		requestsCategory := strings.Split(string(fileCategory), ";")

		var dishesSetups []struct {
			Setup []Dish
		}

		dishes1 := []Dish{
			{
				Name:               "Тако",
				Cost:               60,
				Description:        "То, что нужно настоящему мексиканцу",
				Protein:            1,
				Falt:               1,
				Kilocalorie:        224,
				Carbohydrates:      1,
				CategoryDishes:     "Горячее",
				CategoryRestaurant: "Снеки",
				Count:              1000,
				Weight:             100,
				Avatar:             "https://pbs.twimg.com/media/DtAriH3U8AAD3jV.jpg",
				PlaceCategory:      0,
				Place:              0,
				Ingredient: []Ingredient{
					{
						Name:          "Кетчуп",
						Cost:          5,
						Protein:       1,
						Falt:          1,
						Carbohydrates: 1,
						Kilocalorie:   1,
						CountElement:  1,
						Place:         0,
					},
					{
						Name:          "Горчица",
						Cost:          5,
						Protein:       1,
						Falt:          1,
						Carbohydrates: 1,
						Kilocalorie:   1,
						CountElement:  1,
						Place:         1,
					},
				},
				Radios: nil,
			},
			{
				Name:               "Пряник",
				Cost:               70,
				Description:        "Очень вкусно с чаем",
				Protein:            1,
				Falt:               1,
				Kilocalorie:        126,
				Carbohydrates:      1,
				CategoryDishes:     "К чаю",
				CategoryRestaurant: "К чаю",
				Count:              1000,
				Weight:             70,
				Avatar:             "https://s3.amazonaws.com/images.ecwid.com/images/38011115/2135671623.jpg",
				PlaceCategory:      1,
				Place:              0,
				Ingredient:         nil,
				Radios:             nil,
			},
			{
				Name:               "Универсальное блюдо первого типа",
				Cost:               5,
				Description:        "У этого блюда есть абсолютно всё",
				Protein:            1,
				Falt:               1,
				Kilocalorie:        1,
				Carbohydrates:      1,
				CategoryDishes:     "Универсальное",
				CategoryRestaurant: "Уникальное",
				Count:              1,
				Weight:             5,
				Avatar:             "https://meat-pepper.ru/image/cache/catalog/products/burgers/kotlety-dly-burgerov-black-angus-3-800x667.jpg",
				PlaceCategory:      0,
				Place:              1,
				Ingredient: []Ingredient{
					{
						Name:          "Сахар",
						Cost:          1,
						Protein:       1,
						Falt:          1,
						Carbohydrates: 1,
						Kilocalorie:   1,
						CountElement:  1,
						Place:         0,
					},
				},
				Radios: []Radios{
					{
						Name:  "Напиток",
						Place: 0,
						RadiosElement: []RadiosElement{
							{
								Name:          "Кофе",
								Protein:       1,
								Falt:          1,
								Carbohydrates: 1,
								Kilocalorie:   1,
								Place:         0,
							},
							{
								Name:          "Кола",
								Protein:       2,
								Falt:          2,
								Carbohydrates: 2,
								Kilocalorie:   2,
								Place:         1,
							},
						},
					},
				},
			},
		}

		dishes2 := []Dish{
			{
				Name:               "Универсальное блюдо второго типа",
				Cost:               5,
				Description:        "У этого блюда есть абсолютно всё",
				Protein:            1,
				Falt:               1,
				Kilocalorie:        1,
				Carbohydrates:      1,
				CategoryDishes:     "Универсальное",
				CategoryRestaurant: "Уникальное",
				Count:              1,
				Weight:             5,
				Avatar:             "https://www.avtoall.ru/upload/iblock/49e/multiinstrument_6.jpg",
				PlaceCategory:      0,
				Place:              0,
				Ingredient: []Ingredient{
					{
						Name:          "Сахар",
						Cost:          1,
						Protein:       1,
						Falt:          1,
						Carbohydrates: 1,
						Kilocalorie:   1,
						CountElement:  1,
						Place:         0,
					},
				},
				Radios: []Radios{
					{
						Name:  "Напиток",
						Place: 0,
						RadiosElement: []RadiosElement{
							{
								Name:          "Кофе",
								Protein:       1,
								Falt:          1,
								Carbohydrates: 1,
								Kilocalorie:   1,
								Place:         0,
							},
							{
								Name:          "Кола",
								Protein:       2,
								Falt:          2,
								Carbohydrates: 2,
								Kilocalorie:   2,
								Place:         1,
							},
						},
					},
				},
			},
		}

		dishesSetups = append(dishesSetups, struct{ Setup []Dish }{})
		dishesSetups = append(dishesSetups, struct{ Setup []Dish }{})
		dishesSetups[0].Setup = dishes1
		dishesSetups[1].Setup = dishes2

		categorys := []string{
			"Хороший",
			"Лучший",
			"Единственный",
			"Суши-бар",
			"Кальянная",
			"Пиццерия",
			"Бар",
			"Хенкальная",
			"Общепит",
			"Тестовый",
			"Кафе",
			"Буфеты",
			"Поп-ап",
			"Виртуальный",
		}
		var dishId int
		var radiosId int
		for i := 2; i <= Restaurants; i++ {
			for j := 0; j < Categorys; j++ {
				_, err = conn.Exec(context.Background(), requestsCategory[0], i, categorys[Util.RandomInteger(0, len(categorys))], j)
				if err != nil {
					return nil, &errPkg.Errors{
						Alias: errPkg.MCreateDBNotFillTables,
					}
				}
			}
			_, err = conn.Exec(context.Background(), requestsCategory[1])
			if err != nil {
				return nil, &errPkg.Errors{
					Alias: errPkg.MCreateDBNotFillTables,
				}
			}

			for _, dish := range dishesSetups[Util.RandomInteger(0, len(dishesSetups))].Setup {
				err = conn.QueryRow(context.Background(), requestsDishes[0],
					dish.Name, dish.Cost, i, dish.Description,
					dish.Protein, dish.Falt, dish.Kilocalorie, dish.Carbohydrates,
					dish.CategoryDishes, dish.CategoryRestaurant, dish.Count, dish.Weight,
					dish.Avatar, dish.PlaceCategory, dish.Place).Scan(&dishId)
				if err != nil {
					return nil, &errPkg.Errors{
						Alias: errPkg.MCreateDBNotFillTables,
					}
				}

				if dish.Ingredient != nil {
					for _, ingredient := range dish.Ingredient {
						_, err = conn.Exec(context.Background(), requestsDishes[1],
							ingredient.Name, dishId, ingredient.Cost, ingredient.Protein,
							ingredient.Falt, ingredient.Carbohydrates, ingredient.Kilocalorie,
							ingredient.CountElement, ingredient.Place)
						if err != nil {
							return nil, &errPkg.Errors{
								Alias: errPkg.MCreateDBNotFillTables,
							}
						}
					}
				}

				if dish.Radios != nil {
					for _, radios := range dish.Radios {
						err = conn.QueryRow(context.Background(), requestsDishes[2],
							radios.Name, dishId, radios.Place).Scan(&radiosId)
						if err != nil {
							return nil, &errPkg.Errors{
								Alias: errPkg.MCreateDBNotFillTables,
							}
						}

						for _, elementRadios := range radios.RadiosElement {
							_, err = conn.Exec(context.Background(), requestsDishes[3],
								elementRadios.Name, radiosId, elementRadios.Protein,
								elementRadios.Falt, elementRadios.Carbohydrates, elementRadios.Kilocalorie,
								elementRadios.Place)
							if err != nil {
								return nil, &errPkg.Errors{
									Alias: errPkg.MCreateDBNotFillTables,
								}
							}
						}
					}
				}
			}
		}
	}
	return conn, nil
}
