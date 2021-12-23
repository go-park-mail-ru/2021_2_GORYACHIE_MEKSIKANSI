package service

import (
	resPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant"
	resProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant/proto"
)

func CastRestaurantsToRestaurantProto(r *resPkg.Restaurants) *resProto.Restaurant {
	var p resProto.Restaurant
	p.Img = r.Img
	p.Name = r.Name
	p.CostForFreeDelivery = int64(r.CostForFreeDelivery)
	p.MaxDelivery = int64(r.MaxDelivery)
	p.Rating = r.Rating
	p.Id = int64(r.Id)
	p.MinDelivery = int64(r.MinDelivery)
	return &p
}

func CastTagsToTagsProto(tags []resPkg.Tag) []*resProto.Tags {
	var p []*resProto.Tags
	for _, tag := range tags {
		var rev *resProto.Tags
		rev = &resProto.Tags{}
		rev.Id = int64(tag.Id)
		rev.Name = tag.Name
		p = append(p, rev)
	}
	return p
}

func CastResReviewToResReviewProto(review *resPkg.ResReview) *resProto.ResReview {
	var p resProto.ResReview
	p.Id = int64(review.Id)
	p.Name = review.Name
	p.Rating = review.Rating

	var protoReview []*resProto.Review
	for _, r := range review.Reviews {
		var rev *resProto.Review
		rev = &resProto.Review{}
		rev.Time = r.Time
		rev.Text = r.Text
		rev.Name = r.Name
		rev.Rate = int64(r.Rate)
		rev.Date = r.Date
		protoReview = append(protoReview, rev)
	}
	p.Review = protoReview

	p.Tags = CastTagsToTagsProto(review.Tags)

	p.MinDelivery = int64(review.MinDelivery)
	p.MaxDelivery = int64(review.MaxDelivery)
	p.CostForFreeDelivery = int64(review.CostForFreeDelivery)
	p.Img = review.Img
	p.Status = review.Status
	return &p
}

func CastRadiosToRadiosProto(radios []resPkg.Radios) []*resProto.Radios {
	var p []*resProto.Radios
	for _, i := range radios {
		var protoRadios *resProto.Radios
		protoRadios = &resProto.Radios{}
		protoRadios.Id = int64(i.Id)
		protoRadios.Name = i.Title
		for _, element := range i.Rows {
			var protoRadiosElement *resProto.CheckboxesRows
			protoRadiosElement = &resProto.CheckboxesRows{}
			protoRadiosElement.Id = int64(element.Id)
			protoRadiosElement.Name = element.Name
			protoRadios.Rows = append(protoRadios.Rows, protoRadiosElement)
		}
		p = append(p, protoRadios)
	}
	return p
}

func CastIngredientsToIngredientsProto(ingredients []resPkg.Ingredients) []*resProto.Ingredients {
	var p []*resProto.Ingredients
	for _, i := range ingredients {
		var ingredient *resProto.Ingredients
		ingredient = &resProto.Ingredients{}
		ingredient.Cost = int64(i.Cost)
		ingredient.Id = int64(i.Id)
		ingredient.Name = i.Title
		p = append(p, ingredient)
	}
	return p
}

func CastDishesToDishesProto(d *resPkg.Dishes) *resProto.Dishes {
	var p resProto.Dishes
	p.Cost = int64(d.Cost)
	p.Name = d.Title
	p.Id = int64(d.Id)
	p.Img = d.Img
	p.Ccal = int64(d.Ccal)
	p.Description = d.Description
	return &p
}

func CastMenuToMenuProto(menu []resPkg.Menu) []*resProto.Menu {
	var p []*resProto.Menu
	for i, m := range menu {
		var protoMenu *resProto.Menu
		protoMenu = &resProto.Menu{}
		protoMenu.Name = m.Name
		for _, me := range menu[i].DishesMenu {
			var elementMenu *resProto.DishesMenu
			elementMenu = &resProto.DishesMenu{}
			elementMenu.Id = int64(me.Id)
			elementMenu.Ccal = int64(me.Kilocalorie)
			elementMenu.Cost = int64(me.Cost)
			elementMenu.Img = me.Img
			elementMenu.Name = me.Name
			protoMenu.Dishes = append(protoMenu.Dishes, elementMenu)
		}
		p = append(p, protoMenu)
	}
	return p
}

func CastAllRestaurantsToRecommendedRestaurantsProto(restaurants *resPkg.AllRestaurants) *resProto.RecommendedRestaurants {
	var p *resProto.RecommendedRestaurants
	p = &resProto.RecommendedRestaurants{}
	var protoRestaurants []*resProto.Restaurant
	for _, restaurant := range restaurants.Restaurant {
		var res *resProto.Restaurant
		res = &resProto.Restaurant{}
		res.Id = int64(restaurant.Id)
		res.Img = restaurant.Img
		res.MaxDelivery = int64(restaurant.MaxDelivery)
		res.MinDelivery = int64(restaurant.MinDelivery)
		res.CostForFreeDelivery = int64(restaurant.CostForFreeDelivery)
		res.Rating = restaurant.Rating
		res.Name = restaurant.Name
		protoRestaurants = append(protoRestaurants, res)
	}

	p.Restaurants = protoRestaurants
	p.Tags = CastTagsToTagsProto(restaurants.AllTags)
	return p
}

func CastAllRestaurantsPromoToRestaurantsTagsPromoProto(restaurants *resPkg.AllRestaurantsPromo) *resProto.RestaurantsTagsPromo {
	var p *resProto.RestaurantsTagsPromo
	p = &resProto.RestaurantsTagsPromo{}
	var protoRestaurants []*resProto.Restaurant
	for _, restaurant := range restaurants.Restaurant {
		var res *resProto.Restaurant
		res = &resProto.Restaurant{}
		res.Id = int64(restaurant.Id)
		res.Img = restaurant.Img
		res.MaxDelivery = int64(restaurant.MaxDelivery)
		res.MinDelivery = int64(restaurant.MinDelivery)
		res.CostForFreeDelivery = int64(restaurant.CostForFreeDelivery)
		res.Rating = restaurant.Rating
		res.Name = restaurant.Name
		protoRestaurants = append(protoRestaurants, res)
	}

	p.Restaurants = protoRestaurants
	p.Tags = CastTagsToTagsProto(restaurants.AllTags)
	var protoPromoCodes []*resProto.Promocode
	for _, code := range restaurants.AllPromo {
		var res *resProto.Promocode
		res = &resProto.Promocode{}
		res.RestId = int64(code.RestaurantId)
		res.Img = code.Img
		res.Desc = code.Description
		res.Name = code.Name
		res.Code = code.Code
		protoPromoCodes = append(protoPromoCodes, res)
	}
	p.Promocode = protoPromoCodes
	return p
}

func CastRestaurantsToRestaurantsProto(restaurants []resPkg.Restaurants) *resProto.Restaurants {
	var p *resProto.Restaurants
	p = &resProto.Restaurants{}
	var protoRestaurants []*resProto.Restaurant
	for _, restaurant := range restaurants {
		var res *resProto.Restaurant
		res = &resProto.Restaurant{}
		res.Id = int64(restaurant.Id)
		res.Img = restaurant.Img
		res.MaxDelivery = int64(restaurant.MaxDelivery)
		res.MinDelivery = int64(restaurant.MinDelivery)
		res.CostForFreeDelivery = int64(restaurant.CostForFreeDelivery)
		res.Rating = restaurant.Rating
		res.Name = restaurant.Name
		protoRestaurants = append(protoRestaurants, res)
	}
	p.Restaurants = protoRestaurants
	return p
}

func CastRestaurantIdToRestaurantInfoProto(restInfo *resPkg.RestaurantId) *resProto.RestaurantInfo {
	var p *resProto.RestaurantInfo
	p = &resProto.RestaurantInfo{}
	p.Id = int64(restInfo.Id)
	p.Name = restInfo.Name
	p.Img = restInfo.Img
	p.Rating = restInfo.Rating
	p.CostForFreeDelivery = int64(restInfo.CostForFreeDelivery)
	p.MinDelivery = int64(restInfo.MinDelivery)
	p.MaxDelivery = int64(restInfo.MaxDelivery)
	p.Favourite = restInfo.Favourite

	p.Menu = CastMenuToMenuProto(restInfo.Menu)
	p.Tags = CastTagsToTagsProto(restInfo.Tags)
	return p
}

func CastDishHostProtoToDishHost(dishInfo *resProto.DishesHost) *resPkg.DishHost {
	var p *resPkg.DishHost
	p = &resPkg.DishHost{}
	p.Dishes.Id = int(dishInfo.Id)
	p.Dishes.Title = dishInfo.Name
	p.Dishes.Cost = int(dishInfo.Cost)
	p.Dishes.Ccal = int(dishInfo.Ccal)
	p.Dishes.Description = dishInfo.Description
	p.Dishes.Protein = int(dishInfo.Protein)
	p.Dishes.Falt = int(dishInfo.Falt)
	p.Dishes.Carbohydrates = int(dishInfo.Carbohydrates)
	p.Dishes.Weight = int(dishInfo.Weight)
	p.Dishes.CategoryDishes = dishInfo.CategoryDishes
	p.Dishes.CategoryRestaurant = dishInfo.CategoryRestaurant
	p.Dishes.Count = int(dishInfo.Count)

	for _, radio := range p.Dishes.Radios {
		t := &resProto.CreateRadiosArray{Radios: radio}
		_, p.Dishes.Radios = CastCreateRadiosProtoToCreateRadios(t)

	}
	for _, ingredient := range p.Dishes.Ingredient {
		t := &resProto.CreateIngredientsArray{Ingredients: ingredient}
		_, p.Dishes.Ingredient = CastCreateIngredientsProtoToCreateIngredients(t)
	}
	return p
}

func CastCreateRadiosProtoToCreateRadios(radios *resProto.CreateRadiosArray) (int, []resPkg.CreateRadios) {
	var p []resPkg.CreateRadios
	for _, i := range radios.Radios.Radios {
		var protoRadios resPkg.CreateRadios
		protoRadios = resPkg.CreateRadios{}
		protoRadios.Id = int(i.Id)
		protoRadios.Title = i.Name
		for _, element := range i.Rows {
			var protoRadiosElement resPkg.CreateElementRadios
			protoRadiosElement = resPkg.CreateElementRadios{}
			protoRadiosElement.Id = int(element.Id)
			protoRadiosElement.Name = element.Name
			protoRadiosElement.Protein = int(element.Protein)
			protoRadiosElement.Falt = int(element.Falt)
			protoRadiosElement.Carbohydrates = int(element.Carbohydrates)
			protoRadios.Rows = append(protoRadios.Rows, protoRadiosElement)
		}
		p = append(p, protoRadios)
	}
	return radios.Food, p
}

func CastCreateIngredientsProtoToCreateIngredients(ingredients *resProto.CreateIngredientsArray) (int, []resPkg.CreateIngredients) {
	var p []resPkg.CreateIngredients
	for _, i := range ingredients.Ingredients.Ingredients {
		var ingredient resPkg.CreateIngredients
		ingredient = resPkg.CreateIngredients{}
		ingredient.Cost = int(i.Cost)
		ingredient.Id = int(i.Id)
		ingredient.Title = i.Name
		ingredient.Protein = int(i.Protein)
		ingredient.Falt = int(i.Falt)
		ingredient.Carbohydrates = int(i.Carbohydrates)
		ingredient.Count = int(i.Count)
		p = append(p, ingredient)
	}
	return ingredients.Food, p
}
