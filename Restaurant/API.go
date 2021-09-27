package Restaurant

import (
	auth "2021_2_GORYACHIE_MEKSIKANSI/Authorization"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
	"net/http"
)

type Restaurant struct {
	Id                  int     `json:"id"`
	Img                 string  `json:"imgUrl"`
	Name                string  `json:"name"`
	CostForFreeDelivery int     `json:"cost"`
	MinDelivery         int     `json:"minDeliveryTime"`
	MaxDelivery         int     `json:"maxDeliveryTime"`
	Rating              int 	`json:"rating"`
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
	ctx.SetStatusCode(http.StatusOK)
	json.NewEncoder(ctx).Encode(&auth.Result{
		Status: http.StatusOK,
		Body: restaurant,
	})
	//ctx.Response.SetBody()
	fmt.Printf("Console:  method: %s, url: %s\n", string(ctx.Method()), ctx.URI())
}
