package application

import (
	Cart2 "2021_2_GORYACHIE_MEKSIKANSI/internal/Cart"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Cart/Interface"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Cart/proto"
	resInterface "2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Restaurant/Interface"
	Utils2 "2021_2_GORYACHIE_MEKSIKANSI/internal/Restaurant"
	"context"
)

type CartManager struct {
	DB    Interface.WrapperCart
	DBRes resInterface.WrapperRestaurant
}

func (cm *CartManager) CalculatePriceDelivery(ctx context.Context, id *proto.CalculatePriceDeliveryId) (*proto.CalculatePriceDeliveryResponse, error) {
	// TODO: add convert func
	delivery, err := cm.DB.GetPriceDelivery(int(id.Id))
	if err != nil {
		return nil, err
	}
	return &proto.CalculatePriceDeliveryResponse{Id: int64(delivery)}, nil
}

func (cm *CartManager) CalculateCost(result *Cart2.ResponseCartErrors, rest *Utils2.RestaurantId) (*Cart2.CostCartResponse, error) {
	// TODO: add convert func
	var cost Cart2.CostCartResponse
	sumCost := 0
	for i, dish := range result.Dishes {
		ingredientCost := 0
		for _, ingredient := range dish.IngredientCart {
			ingredientCost += ingredient.Cost
		}
		dishCost := (dish.Cost + ingredientCost) * dish.Count
		sumCost += dishCost
		result.Dishes[i].Cost = dishCost
	}
	cost.SumCost = sumCost
	if sumCost >= int(rest.CostForFreeDelivery) {
		cost.DCost = 0
	} else {
		var err error
		var price *proto.CalculatePriceDeliveryResponse
		var cast *proto.CalculatePriceDeliveryId
		cast.Id = int64(rest.Id)
		price, err = cm.CalculatePriceDelivery(context.Background(), cast)
		cost.DCost = int(price.Id)
		if err != nil {
			return nil, err
		}
	}
	cost.SumCost += cost.DCost
	return &cost, nil
}

func (cm *CartManager) GetCart(ctx context.Context, id *proto.CartId) (*proto.ResponseCartErrors, error) {
	// TODO: add convert func
	result, errorDishes, err := cm.DB.GetCart(int(id.Id))
	if err != nil {
		return nil, err
	}

	rest, err := cm.DBRes.GetRestaurant(result.Restaurant.Id)
	if err != nil {
		return nil, err
	}

	result.CastFromRestaurantId(*rest)

	cost, err := cm.CalculateCost(result, rest)
	if err != nil {
		return nil, err
	}
	result.Cost = *cost
	result.DishErr = errorDishes

	var end *proto.ResponseCartErrors
	dishes := &proto.DishesCartResponse{}
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
	return end, nil
}

func (cm *CartManager) UpdateCart(ctx context.Context, id *proto.RequestCartDefault) (*proto.ResponseCartErrors, error) {
	// TODO: add convert func
	result, errorDishes, err := cm.DB.GetCart(int(id.ClientId))
	if err != nil {
		return nil, err
	}

	rest, err := cm.DBRes.GetRestaurant(result.Restaurant.Id)
	if err != nil {
		return nil, err
	}

	result.CastFromRestaurantId(*rest)

	cost, err := cm.CalculateCost(result, rest)
	if err != nil {
		return nil, err
	}
	result.Cost = *cost
	result.DishErr = errorDishes

	var end *proto.ResponseCartErrors
	dishes := &proto.DishesCartResponse{}
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
	return end, nil
}

func (cm *CartManager) DeleteCart(ctx context.Context, id *proto.DeleteCartId) (*proto.DeleteCartResponse, error) {
	// TODO: add convert func
	err := cm.DB.DeleteCart(int(id.Id))
	if err != nil {
		return nil, err
	}
	return &proto.DeleteCartResponse{Error: ""}, nil
}
