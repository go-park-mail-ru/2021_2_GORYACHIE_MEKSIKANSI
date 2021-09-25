package main

import (
	"context"
	"fmt"
/*	"github.com/gorilla/mux"*/
	"os"
)

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

const password string = ""
type Restaurant struct {
	img string
	name string
	costForFreeDelivery int
	minDeliveryTime int
	maxDeliveryTime int
	rating float32
}

type Wrapper struct {
	conn *pgxpool.Pool
}

type Handler struct {
	conn *pgxpool.Pool
}

type SignUp struct {
	TypeIn       string        `json:"type"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Phone        string        `json:"phone"`
	Password     string        `json:"password"`
}


func (db *Wrapper) getRestaurants() ([]Restaurant, error) {
	row, _ := db.conn.Query(context.Background(), "SELECT id FROM general_user_info")
	p := Restaurant{}
	var result []Restaurant
	for row.Next() {
		//err := row.Scan(&p.img, &p.name, &p.costForFreeDelivery, &p.minDeliveryTime, &p.maxDeliveryTime, &p.rating)
		err := row.Scan(&p.costForFreeDelivery)
		if err != nil {
			panic(err)
		}
		result = append(result, p)
	}

	return result, nil
}


func allRestaurants(db Wrapper) []Restaurant {
	result, _ := db.getRestaurants()
	return result
}


func productsHandler(ctx *fasthttp.RequestCtx) {

	// TODO: make response
	fmt.Fprintf(ctx, "Hi there! RequestURI is ")

}

func signUpHandler(ctx *fasthttp.RequestCtx, conn *pgxpool.Pool) {
	var db = Wrapper{conn}
	restaurants := allRestaurants(db)
	// TODO: make response
	fmt.Fprintf(ctx, "Hi there! RequestURI is %d", restaurants[0].costForFreeDelivery)
}

func loginHandler(ctx *fasthttp.RequestCtx, conn *pgxpool.Pool) {
	var db = Wrapper{conn}
	restaurants := allRestaurants(db)
	// TODO: make response
	fmt.Fprintf(ctx, "Hi there! RequestURI is %d", restaurants[0].costForFreeDelivery)
}

func logoutHandler(ctx *fasthttp.RequestCtx, conn *pgxpool.Pool) {
	var db = Wrapper{conn}
	restaurants := allRestaurants(db)
	// TODO: make response
	fmt.Fprintf(ctx, "Hi there! RequestURI is %d", restaurants[0].costForFreeDelivery)
}

func (h *Handler) Routing(ctx *fasthttp.RequestCtx) {
/*	router := mux.NewRouter()
	api := "/api"
	router.HandleFunc(api + "/restaurants", productsHandler).Methods("GET")
	router.HandleFunc(api + "/login", loginHandler).Methods("POST")
	router.HandleFunc(api + "/logout", logoutHandler).Methods("POST")
	router.HandleFunc(api + "/signup", signUpHandler).Methods("POST")*/

	switch string(ctx.Path()) {
	case "/api/restaurants":
		productsHandler(ctx)
	case "/api/":
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}

}

func runServer(port string) {
	conn, err := pgxpool.Connect(context.Background(), "postgres://Captain-matroskin:" + password + "@localhost:5432/hot_mexican")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()


	handler := &Handler{conn: conn}
	err = fasthttp.ListenAndServe(port, handler.Routing)
	if err != nil {
		return
	}
}

func main() {
	runServer(":8080")
}


