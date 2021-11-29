package cast

import (
	Cart2 "2021_2_GORYACHIE_MEKSIKANSI/internal/cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservice/cart/proto"
)

func CastRequestCartDefaultToRequestCartDefaultProto(cart *Cart2.RequestCartDefault) *proto.RequestCartDefault {
	var protoCart proto.RequestCartDefault
	protoCart.Restaurant = &proto.RestaurantRequest{}
	protoCart.Restaurant.Id = int64(cart.Restaurant.Id)

	var dishes []*proto.DishesRequest
	for i, dish := range cart.Dishes {
		var newDish *proto.DishesRequest
		newDish = &proto.DishesRequest{}
		newDish.Id = int64(dish.Id)
		newDish.ItemNumber = int64(dish.ItemNumber)
		newDish.Count = int64(dish.Count)
		for _, id := range cart.Dishes[i].Ingredients {
			var ingredient *proto.IngredientsCartRequest
			ingredient = &proto.IngredientsCartRequest{}
			ingredient.Id = int64(id.Id)
			newDish.Ingredients = append(newDish.Ingredients, ingredient)
		}
		for _, id := range cart.Dishes[i].Radios {
			var radios *proto.RadiosCartRequest
			radios = &proto.RadiosCartRequest{}
			radios.Id = int64(id.Id)
			radios.RadiosId = int64(id.RadiosId)
			newDish.Radios = append(newDish.Radios, radios)
		}
		dishes = append(dishes, newDish)
	}
	protoCart.Dishes = dishes
	return &protoCart
}

func CastResponseCartErrorsProtoToResponseCartErrors(result *proto.ResponseCartErrors) *Cart2.ResponseCartErrors {
	var end *Cart2.ResponseCartErrors
	end = &Cart2.ResponseCartErrors{}
	end.Cost.DCost = int(result.Cost.DeliveryCost)
	end.Cost.SumCost = int(result.Cost.SumCost)
	end.Restaurant.Id = int(result.Restaurant.Id)
	end.Restaurant.MaxDelivery = int(result.Restaurant.MaxDelivery)
	end.Restaurant.MinDelivery = int(result.Restaurant.MinDelivery)
	end.Restaurant.Name = result.Restaurant.Name
	end.Restaurant.Img = result.Restaurant.Img
	end.Restaurant.Rating = result.Restaurant.Rating
	end.Restaurant.CostForFreeDelivery = int(result.Restaurant.CostForFreeDelivery)
	dishes := Cart2.DishesCartResponse{}
	for _, dish := range result.Dishes {
		var ingredient []Cart2.IngredientCartResponse
		var radios []Cart2.RadiosCartResponse
		for _, ing := range dish.Ingredients {
			ingredient = append(ingredient, Cart2.IngredientCartResponse{Id: int(ing.Id), Name: ing.Name, Cost: int(ing.Cost)})
		}

		for _, rad := range dish.Radios {
			radios = append(radios, Cart2.RadiosCartResponse{Id: int(rad.Id), Name: rad.Name, RadiosId: int(rad.RadiosId)})
		}
		dishes.Id = int(dish.Id)
		dishes.RadiosCart = radios
		dishes.Name = dish.Name
		dishes.Cost = int(dish.Cost)
		dishes.IngredientCart = ingredient
		dishes.Weight = int(dish.Weight)
		dishes.Count = int(dish.Count)
		dishes.ItemNumber = int(dish.ItemNumber)
		dishes.Description = dish.Description
		dishes.Kilocalorie = int(dish.Ccal)
		dishes.Img = dish.Img
		end.Dishes = append(end.Dishes, dishes)
	}
	for _, errDish := range result.DishesErrors {
		var dish *Cart2.CastDishesErrs
		dish.NameDish = errDish.NameDish
		dish.ItemNumber = int(errDish.ItemNumber)
		dish.CountAvail = int(errDish.CountAvail)
	}
	return end
}
