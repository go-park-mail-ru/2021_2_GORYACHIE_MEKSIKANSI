package cast

import (
	Cart2 "2021_2_GORYACHIE_MEKSIKANSI/internal/cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/microservices/cart/proto"
)

func CastRequestCartDefaultProtoToRequestCartDefault(cart *proto.RequestCartDefault) *Cart2.RequestCartDefault {
	var protoCart Cart2.RequestCartDefault
	protoCart.Restaurant.Id = int(cart.Restaurant.Id)

	var dishes []Cart2.DishesRequest
	for i, dish := range cart.Dishes {
		var newDish Cart2.DishesRequest
		newDish.Id = int(dish.Id)
		newDish.ItemNumber = int(dish.ItemNumber)
		newDish.Count = int(dish.Count)
		for _, id := range cart.Dishes[i].Ingredients {
			var ingredient Cart2.IngredientsCartRequest
			ingredient.Id = int(id.Id)
			newDish.Ingredients = append(newDish.Ingredients, ingredient)
		}
		for _, id := range cart.Dishes[i].Radios {
			var radios Cart2.RadiosCartRequest
			radios.Id = int(id.Id)
			radios.RadiosId = int(id.RadiosId)
			newDish.Radios = append(newDish.Radios, radios)
		}
		dishes = append(dishes, newDish)
	}
	protoCart.Dishes = dishes
	return &protoCart
}

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

func CastResponseCartErrorsToResponseCartErrorsProto(result *Cart2.ResponseCartErrors) *proto.ResponseCartErrors {
	var end *proto.ResponseCartErrors
	end = &proto.ResponseCartErrors{}
	end.Restaurant = &proto.RestaurantIdCastResponse{}
	end.Cost = &proto.CostCartResponse{}
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
	return end
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
