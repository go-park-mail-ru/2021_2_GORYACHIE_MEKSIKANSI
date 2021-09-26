package Restaurant

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"net/http"
	middleware "project/Middleware"
)

type Restaurant struct {
	Id                  int     `json:"id"`
	Img                 string  `json:"imgUrl"`
	Name                string  `json:"restaurantName"`
	CostForFreeDelivery int     `json:"costForFreeDelivery"`
	MinDelivery         int     `json:"minDelivery"`
	MaxDelivery         int     `json:"maxDelivery"`
	Rating              float32 `json:"rating"`
}

type RestaurantInfo struct {
	ConnectionDB *pgxpool.Pool
}

func (r *RestaurantInfo) ProductsHandler(ctx *fasthttp.RequestCtx) {
	WrapperDB := Wrapper{Conn: r.ConnectionDB}
	restaurant, _ /*err*/ := AllRestaurants(WrapperDB) // TODO: проверки на ошибки
	if restaurant != nil {
		ctx.Response.SetStatusCode(http.StatusBadRequest) // TODO: только 200 вернуть
	}
	middleware.SetHeaders(ctx)

	fmt.Printf("Console:  method: %s, url: %s\n", string(ctx.Method()), ctx.URI())
}
