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
	Img                 string  `json:"img"`
	Name                string  `json:"name"`
	CostForFreeDelivery int     `json:"сostForFreeDelivery"`
	MinDelivery         int     `json:"minDeliveryTime"`
	MaxDelivery         int     `json:"maxDeliveryTime"`
	Rating              float32 `json:"rating"`
}

type RestaurantInfo struct {
	ConnectionDB *pgxpool.Pool
}

func (r *RestaurantInfo) ProductsHandler(ctx *fasthttp.RequestCtx) {
	WrapperDB := Wrapper{Conn: r.ConnectionDB}
	restaurant, err := AllRestaurants(WrapperDB)
	err = CheckErrorRestaurant(err, ctx, restaurant)
	if err != nil {
		return
	}

	ctx.SetStatusCode(http.StatusOK)
	err = json.NewEncoder(ctx).Encode(&auth.Result{
		Status: http.StatusOK,
		Body: restaurant,
	})
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusOK)
		fmt.Printf("Console: %s\n", auth.ERRENCODE)
		return
	}
	fmt.Printf("Console:  method: %s, url: %s\n", string(ctx.Method()), ctx.URI())
}
