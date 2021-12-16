package service

import (
	cartPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/cart/proto"
)

func CastRequestCartDefaultProtoToRequestCartDefault(cart *proto.RequestCartDefault) *cartPkg.RequestCartDefault {
	var protoCart cartPkg.RequestCartDefault
	protoCart.Restaurant.Id = int(cart.Restaurant.Id)

	var dishes []cartPkg.DishesRequest
	for i, dish := range cart.Dishes {
		var newDish cartPkg.DishesRequest
		newDish.Id = int(dish.Id)
		newDish.ItemNumber = int(dish.ItemNumber)
		newDish.Count = int(dish.Count)
		for _, id := range cart.Dishes[i].Ingredients {
			var ingredient cartPkg.IngredientsCartRequest
			ingredient.Id = int(id.Id)
			newDish.Ingredients = append(newDish.Ingredients, ingredient)
		}
		for _, id := range cart.Dishes[i].Radios {
			var radios cartPkg.RadiosCartRequest
			radios.Id = int(id.Id)
			radios.RadiosId = int(id.RadiosId)
			newDish.Radios = append(newDish.Radios, radios)
		}
		dishes = append(dishes, newDish)
	}
	protoCart.Dishes = dishes
	protoCart.PromoCode = cart.PromoCode
	return &protoCart
}

func CastResponseCartErrorsToResponseCartErrorsProto(result *cartPkg.ResponseCartErrors) *proto.ResponseCartErrors {
	var end *proto.ResponseCartErrors
	end = &proto.ResponseCartErrors{}
	end.Restaurant = &proto.RestaurantIdCastResponse{}
	end.Cost = &proto.CostCartResponse{}
	end.PromoCode = &proto.PromoCode{}
	end.Cost.DeliveryCost = int64(result.Cost.DCost)
	end.Cost.SumCost = int64(result.Cost.SumCost)
	end.Restaurant.Id = int64(result.Restaurant.Id)
	end.Restaurant.MaxDelivery = int64(result.Restaurant.MaxDelivery)
	end.Restaurant.MinDelivery = int64(result.Restaurant.MinDelivery)
	end.Restaurant.Name = result.Restaurant.Name
	end.Restaurant.Img = result.Restaurant.Img
	end.Restaurant.Rating = result.Restaurant.Rating
	end.Restaurant.CostForFreeDelivery = int64(result.Restaurant.CostForFreeDelivery)
	for _, dish := range result.Dishes {
		dishes := &proto.DishesCartResponse{}
		var ingredient []*proto.IngredientCartResponse
		var radios []*proto.RadiosCartResponse
		for _, ing := range dish.IngredientCart {
			ingredient = append(ingredient, &proto.IngredientCartResponse{Id: int64(ing.Id), Name: ing.Name, Cost: int64(ing.Cost)})
		}

		for _, rad := range dish.RadiosCart {
			radios = append(radios, &proto.RadiosCartResponse{Id: int64(rad.Id), Name: rad.Name, RadiosId: int64(rad.RadiosId)})
		}
		dishes.Id = int64(dish.Id)
		dishes.Radios = radios
		dishes.Name = dish.Name
		dishes.Cost = int64(dish.Cost)
		dishes.Ingredients = ingredient
		dishes.Weight = int64(dish.Weight)
		dishes.Count = int64(dish.Count)
		dishes.ItemNumber = int64(dish.ItemNumber)
		dishes.Description = dish.Description
		dishes.Ccal = int64(dish.Kilocalorie)
		dishes.Img = dish.Img
		end.Dishes = append(end.Dishes, dishes)
	}
	for _, errDish := range result.DishErr {
		var dish *proto.CastDishesErrs
		dish = &proto.CastDishesErrs{}
		dish.NameDish = errDish.NameDish
		dish.ItemNumber = int64(errDish.ItemNumber)
		dish.CountAvail = int64(errDish.CountAvail)
	}
	end.PromoCode.Code = result.PromoCode.Code
	end.PromoCode.Name = result.PromoCode.Name
	end.PromoCode.Description = result.PromoCode.Description
	return end
}
