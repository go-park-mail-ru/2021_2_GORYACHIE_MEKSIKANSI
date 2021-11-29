package build

import (
	"2021_2_GORYACHIE_MEKSIKANSI/config"
	errPkg "2021_2_GORYACHIE_MEKSIKANSI/internal/myerror"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/util"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"strings"
)

const (
	Restaurants = 53
	Categories  = 3
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
	addressPostgres := "postgres://" + configDB.UserName + ":" + configDB.Password +
		"@" + configDB.Host + ":" + configDB.Port + "/" + configDB.SchemaName

	conn, err := pgxpool.Connect(context.Background(), addressPostgres)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.MCreateDBNotConnect,
		}
	}

	if debug {
		file, err := ioutil.ReadFile("./build/postgresql/deletetables.sql")
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

	file, err := ioutil.ReadFile("./build/postgresql/createtables.sql")
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
		file, err := ioutil.ReadFile("./build/postgresql/fill.sql")
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

		fileDishes, err := ioutil.ReadFile("./build/postgresql/filldishes.sql")
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.MCreateDBFillFileNotFound,
			}
		}

		requestsDishes := strings.Split(string(fileDishes), ";")

		fileCategory, err := ioutil.ReadFile("./build/postgresql/fillcategory.sql")
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
				CategoryRestaurant: "Снеки",
				Count:              1000,
				Weight:             70,
				Avatar:             "https://s3.amazonaws.com/images.ecwid.com/images/38011115/2135671623.jpg",
				PlaceCategory:      0,
				Place:              1,
				Ingredient:         nil,
				Radios:             nil,
			},
			{
				Name:               "Чёрный бургер",
				Cost:               139,
				Description:        "Получен в угольных шахтах",
				Protein:            1,
				Falt:               1,
				Kilocalorie:        361,
				Carbohydrates:      1,
				CategoryDishes:     "Универсальное",
				CategoryRestaurant: "Снеки",
				Count:              1000,
				Weight:             220,
				Avatar:             "https://meat-pepper.ru/image/cache/catalog/products/burgers/kotlety-dly-burgerov-black-angus-3-800x667.jpg",
				PlaceCategory:      0,
				Place:              2,
				Ingredient:         nil,
				Radios:             nil,
			},
			{
				Name:               "Пицца Ассорти",
				Cost:               429,
				Description:        "Просто",
				Protein:            1,
				Falt:               1,
				Kilocalorie:        1024,
				Carbohydrates:      1,
				CategoryDishes:     "Универсальное",
				CategoryRestaurant: "Пиццы",
				Count:              1000,
				Weight:             600,
				Avatar:             "https://www.koolinar.ru/all_image/recipes/156/156543/recipe_7b4bb7f7-1d42-428a-bb0a-3db8df03093a.jpg",
				PlaceCategory:      1,
				Place:              0,
				Ingredient: []Ingredient{
					{
						Name:          "Сырные бортики",
						Cost:          1,
						Protein:       1,
						Falt:          1,
						Carbohydrates: 1,
						Kilocalorie:   1,
						CountElement:  1,
						Place:         0,
					},
					{
						Name:          "Колбаса",
						Cost:          1,
						Protein:       1,
						Falt:          1,
						Carbohydrates: 1,
						Kilocalorie:   1,
						CountElement:  1,
						Place:         1,
					},
					{
						Name:          "Сыр Пармезан",
						Cost:          1,
						Protein:       1,
						Falt:          1,
						Carbohydrates: 1,
						Kilocalorie:   1,
						CountElement:  1,
						Place:         2,
					},
					{
						Name:          "Сыр Моцарелла",
						Cost:          1,
						Protein:       1,
						Falt:          1,
						Carbohydrates: 1,
						Kilocalorie:   1,
						CountElement:  1,
						Place:         3,
					},
				},
				Radios: nil,
			},
			{
				Name:               "Кофе",
				Cost:               149,
				Description:        "Горячий, ароматный кофе",
				Protein:            1,
				Falt:               1,
				Kilocalorie:        90,
				Carbohydrates:      1,
				CategoryDishes:     "Универсальное",
				CategoryRestaurant: "Напитки",
				Count:              1000,
				Weight:             100,
				Avatar:             "https://traveltimes.ru/wp-content/uploads/2021/08/kofe-caska-penka-scaled.jpg",
				PlaceCategory:      2,
				Place:              0,
				Ingredient: []Ingredient{
					{
						Name:          "Сахар",
						Cost:          5,
						Protein:       1,
						Falt:          1,
						Carbohydrates: 1,
						Kilocalorie:   1,
						CountElement:  1,
						Place:         0,
					},
				},
				Radios: nil,
			},
			{
				Name:               "Coca-cola",
				Cost:               65,
				Description:        "Горячий, ароматный кофе",
				Protein:            1,
				Falt:               1,
				Kilocalorie:        230,
				Carbohydrates:      1,
				CategoryDishes:     "Универсальное",
				CategoryRestaurant: "Напитки",
				Count:              1000,
				Weight:             500,
				Avatar:             "https://traveltimes.ru/wp-content/uploads/2021/08/kofe-caska-penka-scaled.jpg",
				PlaceCategory:      2,
				Place:              1,
				Ingredient:         nil,
				Radios:             nil,
			},
			{
				Name:               "Fanta",
				Cost:               60,
				Description:        "Горячий, ароматный кофе",
				Protein:            1,
				Falt:               1,
				Kilocalorie:        225,
				Carbohydrates:      1,
				CategoryDishes:     "Универсальное",
				CategoryRestaurant: "Напитки",
				Count:              1000,
				Weight:             500,
				Avatar:             "https://traveltimes.ru/wp-content/uploads/2021/08/kofe-caska-penka-scaled.jpg",
				PlaceCategory:      2,
				Place:              2,
				Ingredient:         nil,
				Radios:             nil,
			},
			{
				Name:               "Sprite",
				Cost:               65,
				Description:        "Горячий, ароматный кофе",
				Protein:            1,
				Falt:               1,
				Kilocalorie:        215,
				Carbohydrates:      1,
				CategoryDishes:     "Универсальное",
				CategoryRestaurant: "Напитки",
				Count:              1000,
				Weight:             500,
				Avatar:             "https://traveltimes.ru/wp-content/uploads/2021/08/kofe-caska-penka-scaled.jpg",
				PlaceCategory:      2,
				Place:              3,
				Ingredient:         nil,
				Radios:             nil,
			},
			{
				Name:               "Картошка Фри",
				Cost:               60,
				Description:        "Классический картофель фри",
				Protein:            1,
				Falt:               1,
				Kilocalorie:        232,
				Carbohydrates:      1,
				CategoryDishes:     "Универсальное",
				CategoryRestaurant: "Снеки",
				Count:              1000,
				Weight:             120,
				Avatar:             "https://вести35.рф/images/2020/07/06/5f2775ffddc94d76a57605479b3f02e0.jpg",
				PlaceCategory:      0,
				Place:              3,
				Ingredient:         nil,
				Radios:             nil,
			},
			{
				Name:               "Картошка по деревенски",
				Cost:               60,
				Description:        "Горячий, ароматный кофе",
				Protein:            1,
				Falt:               1,
				Kilocalorie:        172,
				Carbohydrates:      1,
				CategoryDishes:     "Универсальное",
				CategoryRestaurant: "Снеки",
				Count:              1000,
				Weight:             130,
				Avatar:             "https://cherkessk.crazybrothers.ru/wp-content/uploads/Kartofel-po-derevenski.jpg",
				PlaceCategory:      0,
				Place:              4,
				Ingredient:         nil,
				Radios:             nil,
			},
			{
				Name:               "Блюдо со стейком",
				Cost:               256,
				Description:        "У этого блюда есть абсолютно всё",
				Protein:            1,
				Falt:               1,
				Kilocalorie:        756,
				Carbohydrates:      1,
				CategoryDishes:     "Универсальное",
				CategoryRestaurant: "Комбо",
				Count:              1000,
				Weight:             400,
				Avatar:             "https://www.islandresortandcasino.com/sites/default/wp/blog/wp-content/uploads/2015/02/DSC_2162.jpg",
				PlaceCategory:      3,
				Place:              0,
				Ingredient:         nil,
				Radios: []Radios{
					{
						Name:  "Картофель",
						Place: 0,
						RadiosElement: []RadiosElement{
							{
								Name:          "Картофель Фри",
								Protein:       1,
								Falt:          1,
								Carbohydrates: 1,
								Kilocalorie:   1,
								Place:         0,
							},
							{
								Name:          "Картофель по Деревенски",
								Protein:       1,
								Falt:          1,
								Carbohydrates: 1,
								Kilocalorie:   1,
								Place:         1,
							},
						},
					},
					{
						Name:  "Основное блюдо",
						Place: 1,
						RadiosElement: []RadiosElement{
							{
								Name:          "Стейк",
								Protein:       10,
								Falt:          1,
								Carbohydrates: 1,
								Kilocalorie:   1,
								Place:         0,
							},
						},
					},
				},
			},
			{
				Name:               "Сёмга",
				Cost:               800,
				Description:        "У этого блюда есть абсолютно всё",
				Protein:            1,
				Falt:               1,
				Kilocalorie:        700,
				Carbohydrates:      1,
				CategoryDishes:     "Универсальное",
				CategoryRestaurant: "Комбо",
				Count:              1,
				Weight:             5,
				Avatar:             "http://www.t-h.ru/photomenu/public/img/c/c6dc9d517433930dddcbb26e4824382c.jpg",
				PlaceCategory:      3,
				Place:              1,
				Ingredient:         nil,
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
						},
					},
				},
			},
			{
				Name:               "Цезарь",
				Cost:               5,
				Description:        "У этого блюда есть абсолютно всё",
				Protein:            1,
				Falt:               1,
				Kilocalorie:        1,
				Carbohydrates:      1,
				CategoryDishes:     "Универсальное",
				CategoryRestaurant: "Комбо",
				Count:              1,
				Weight:             5,
				Avatar:             "https://pizza-house18.ru/wp-content/uploads/2021/01/IMG_9061.jpg",
				PlaceCategory:      3,
				Place:              2,
				Ingredient:         nil,
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
						},
					},
				},
			},
			{
				Name:               "Бефстроганов",
				Cost:               1000,
				Description:        "С грибами и картофельным пюре",
				Protein:            1,
				Falt:               1,
				Kilocalorie:        560,
				Carbohydrates:      1,
				CategoryDishes:     "Универсальное",
				CategoryRestaurant: "Комбо",
				Count:              1000,
				Weight:             1000,
				Avatar:             "http://eday-cafe.ru/wp-content/uploads/2020/03/Бефстроганов-с-грибами-и-толченым-картофелем.jpg",
				PlaceCategory:      3,
				Place:              3,
				Ingredient:         nil,
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
			for j := 0; j < Categories; j++ {
				_, err = conn.Exec(context.Background(), requestsCategory[0], i, categorys[util.RandomInteger(0, len(categorys))], j)
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

			for _, dish := range dishesSetups[util.RandomInteger(0, len(dishesSetups))].Setup {
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
