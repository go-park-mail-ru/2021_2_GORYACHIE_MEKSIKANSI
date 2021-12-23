package cast

import (
	resPkg "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant"
	resProto "2021_2_GORYACHIE_MEKSIKANSI/internals/microservice/restaurant/proto"
	restaurant "2021_2_GORYACHIE_MEKSIKANSI/internals/restaurant"
)

func CastTagsToTagsProto(tags []restaurant.Tag) []*resProto.Tags {
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

func CastMenuToMenuProto(menu []restaurant.Menu) []*resProto.Menu {
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

func CastRestaurantsProtoToRestaurant(r *resProto.Restaurants) []restaurant.Restaurants {
	var restaurants []restaurant.Restaurants
	for _, rest := range r.Restaurants {
		var res restaurant.Restaurants
		res.Img = rest.Img
		res.Name = rest.Name
		res.CostForFreeDelivery = int(rest.CostForFreeDelivery)
		res.MaxDelivery = int(rest.MaxDelivery)
		res.Rating = rest.Rating
		res.Id = int(rest.Id)
		res.MinDelivery = int(rest.MinDelivery)
		restaurants = append(restaurants, res)
	}
	return restaurants
}

func CastRestaurantsProtoToRestaurants(restaurants []*resProto.Restaurant) []restaurant.Restaurants {
	var r []restaurant.Restaurants
	for _, rest := range restaurants {
		var res restaurant.Restaurants
		res.Id = int(rest.Id)
		res.Img = rest.Img
		res.MaxDelivery = int(rest.MaxDelivery)
		res.MinDelivery = int(rest.MinDelivery)
		res.CostForFreeDelivery = int(rest.CostForFreeDelivery)
		res.Rating = rest.Rating
		res.Name = rest.Name
		r = append(r, res)
	}
	return r
}

func CastRestaurantsTagsProtoToAllRestaurants(restaurants *resProto.RestaurantsTagsPromo) *restaurant.AllRestaurantsPromo {
	var r restaurant.AllRestaurantsPromo
	for _, rest := range restaurants.Restaurants {
		var res restaurant.Restaurants
		res.Id = int(rest.Id)
		res.Img = rest.Img
		res.MaxDelivery = int(rest.MaxDelivery)
		res.MinDelivery = int(rest.MinDelivery)
		res.CostForFreeDelivery = int(rest.CostForFreeDelivery)
		res.Rating = rest.Rating
		res.Name = rest.Name
		r.Restaurant = append(r.Restaurant, res)
	}
	r.AllTags = CastTagsProtoToTags(restaurants.Tags)
	for _, code := range restaurants.Promocode {
		var res restaurant.PromoCode
		res.RestaurantId = int(code.RestId)
		res.Img = code.Img
		res.Name = code.Name
		res.Description = code.Desc
		res.Code = code.Code
		r.AllPromo = append(r.AllPromo, res)
	}
	return &r
}

func CastRecommendedRestaurantsProtoToAllRestaurants(restaurants *resProto.RecommendedRestaurants) *restaurant.AllRestaurants {
	var r restaurant.AllRestaurants
	for _, rest := range restaurants.Restaurants {
		var res restaurant.Restaurants
		res.Id = int(rest.Id)
		res.Img = rest.Img
		res.MaxDelivery = int(rest.MaxDelivery)
		res.MinDelivery = int(rest.MinDelivery)
		res.CostForFreeDelivery = int(rest.CostForFreeDelivery)
		res.Rating = rest.Rating
		res.Name = rest.Name
		r.Restaurant = append(r.Restaurant, res)
	}
	r.AllTags = CastTagsProtoToTags(restaurants.Tags)
	return &r
}

func CastRestaurantInfoToRestaurantIdProto(restInfo *resProto.RestaurantInfo) *restaurant.RestaurantId {
	var p *restaurant.RestaurantId
	p = &restaurant.RestaurantId{}
	p.Id = int(restInfo.Id)
	p.Name = restInfo.Name
	p.Img = restInfo.Img
	p.Rating = restInfo.Rating
	p.CostForFreeDelivery = int(restInfo.CostForFreeDelivery)
	p.MinDelivery = int(restInfo.MinDelivery)
	p.MaxDelivery = int(restInfo.MaxDelivery)
	p.Favourite = restInfo.Favourite

	p.Menu = CastMenuProtoToMenu(restInfo.Menu)
	p.Tags = CastTagsProtoToTags(restInfo.Tags)
	return p
}

func CastTagsProtoToTags(tags []*resProto.Tags) []restaurant.Tag {
	var p []restaurant.Tag
	for _, tag := range tags {
		var rev restaurant.Tag
		rev.Id = int(tag.Id)
		rev.Name = tag.Name
		p = append(p, rev)
	}
	return p
}

func CastMenuProtoToMenu(protoMenu []*resProto.Menu) []restaurant.Menu {
	var p []restaurant.Menu
	for i, m := range protoMenu {
		var menu restaurant.Menu
		menu.Name = m.Name
		for _, me := range protoMenu[i].Dishes {
			var elementMenu restaurant.DishesMenu
			elementMenu.Id = int(me.Id)
			elementMenu.Kilocalorie = int(me.Ccal)
			elementMenu.Cost = int(me.Cost)
			elementMenu.Img = me.Img
			elementMenu.Name = me.Name
			menu.DishesMenu = append(menu.DishesMenu, elementMenu)
		}
		p = append(p, menu)
	}
	return p
}

func CastDishesProtoToDishes(d *resProto.Dishes) *restaurant.Dishes {
	var r restaurant.Dishes
	r.Cost = int(d.Cost)
	r.Title = d.Name
	r.Id = int(d.Id)
	r.Img = d.Img
	r.Ccal = int(d.Ccal)
	r.Description = d.Description
	r.Ingredient = CastIngredientsProtoToIngredients(d.Ingredients)
	r.Radios = CastRadiosProtoToRadios(d.Radios)
	return &r
}

func CastRadiosProtoToRadios(radios []*resProto.Radios) []restaurant.Radios {
	var p []restaurant.Radios
	for _, i := range radios {
		var protoRadios restaurant.Radios
		protoRadios.Id = int(i.Id)
		protoRadios.Title = i.Name
		for _, element := range i.Rows {
			var protoRadiosElement restaurant.ElementRadios
			protoRadiosElement.Id = int(element.Id)
			protoRadiosElement.Name = element.Name
			protoRadios.Rows = append(protoRadios.Rows, protoRadiosElement)
		}
		p = append(p, protoRadios)
	}
	return p
}

func CastIngredientsProtoToIngredients(ingredients []*resProto.Ingredients) []restaurant.Ingredients {
	var p []restaurant.Ingredients
	for _, i := range ingredients {
		var ingredient restaurant.Ingredients
		ingredient.Cost = int(i.Cost)
		ingredient.Id = int(i.Id)
		ingredient.Title = i.Name
		p = append(p, ingredient)
	}
	return p
}

func CastNewReviewToNewReviewProto(review restaurant.NewReview, id int) *resProto.NewReview {
	var p *resProto.NewReview
	p = &resProto.NewReview{}
	p.Text = review.Text
	p.Rate = int64(review.Rate)
	p.Id = int64(id)
	p.Restaurant = CastRestaurantIdToRestaurantInfoProto(&review.Restaurant)
	return p
}

func CastRestaurantIdToRestaurantInfoProto(restInfo *restaurant.RestaurantId) *resProto.RestaurantInfo {
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

func CastResReviewProtoToResReview(review *resProto.ResReview) *restaurant.ResReview {
	var resReview restaurant.ResReview
	resReview.Id = int(review.Id)
	resReview.Name = review.Name
	resReview.Rating = review.Rating

	var protoReview []restaurant.Review
	for _, r := range review.Review {
		var rev restaurant.Review
		rev.Time = r.Time
		rev.Text = r.Text
		rev.Name = r.Name
		rev.Rate = int(r.Rate)
		rev.Date = r.Date
		protoReview = append(protoReview, rev)
	}
	resReview.Reviews = protoReview

	resReview.Tags = CastTagsProtoToTags(review.Tags)

	resReview.MinDelivery = int(review.MinDelivery)
	resReview.MaxDelivery = int(review.MaxDelivery)
	resReview.CostForFreeDelivery = int(review.CostForFreeDelivery)
	resReview.Img = review.Img
	resReview.Status = review.Status
	return &resReview
}

func CastDishHostProtoToDishHost(dishInfo *resPkg.DishHost) *resProto.DishesHost {
	var p *resPkg.DishesHost
	p = &resPkg.DishesHost{}
	p.Id = int64(dishInfo.Dishes.Id)
	p.Title = dishInfo.Dishes.Title
	p.Cost = int64(dishInfo.Dishes.Cost)
	p.Ccal = int64(dishInfo.Dishes.Ccal)
	p.Description = dishInfo.Dishes.Description
	p.Protein = int64(dishInfo.Dishes.Protein)
	p.Falt = int64(dishInfo.Dishes.Falt)
	p.Carbohydrates = int64(dishInfo.Dishes.Carbohydrates)
	p.Weight = int64(dishInfo.Dishes.Weight)
	p.CategoryDishes = dishInfo.Dishes.CategoryDishes
	p.CategoryRestaurant = dishInfo.Dishes.CategoryRestaurant
	p.Count = int64(dishInfo.Dishes.Count)

	p.Radios = CastCreateRadiosToCreateRadiosProto(dishInfo.Dishes.Radios, dishInfo.Dishes.Id).Radios
	p.ingredients = CastCreateIngredientsToCreateIngredientsProto(dishInfo.Dishes.Ingredient, dishInfo.Dishes.Id).Ingredients
	return p
}

func CastCreateRadiosToCreateRadiosProto(radios []resPkg.CreateRadios, dishId int) *resProto.CreateRadiosArray {
	var result *resProto.CreateRadiosArray
	var p []*resProto.CreateRadios
	p.Food = dishId
	for _, i := range radios {
		var protoRadios *resProto.CreateRadios
		protoRadios = &resProto.CreateRadios{}
		protoRadios.Id = int64(i.Id)
		protoRadios.Name = i.Title
		for _, element := range i.Rows {
			var protoRadiosElement *resProto.CreateElementradios
			protoRadiosElement = &resProto.CreateElementradios{}
			protoRadiosElement.Id = int64(element.Id)
			protoRadiosElement.Name = element.Name
			protoRadiosElement.Protein = int64(element.Protein)
			protoRadiosElement.Falt = int64(element.Falt)
			protoRadiosElement.Carbohydrates = int64(element.Carbohydrates)
			protoRadios.Rows = append(protoRadios.Rows, protoRadiosElement)
		}
		p = append(p, protoRadios)
	}
	result.Radios = p
	return result
}

func CastCreateIngredientsToCreateIngredientsProto(ingredients []resPkg.CreateIngredients, dishId int) *resProto.CreateIngredientsArray {
	var result *resProto.CreateIngredientsArray
	var p []*resProto.CreateIngredients
	p.Food = dishId
	for _, i := range ingredients {
		var ingredient *resProto.CreateIngredients
		ingredient = &resProto.CreateIngredients{}
		ingredient.Cost = int64(i.Cost)
		ingredient.Id = int64(i.Id)
		ingredient.Name = i.Title
		ingredient.Protein = int64(i.Protein)
		ingredient.Falt = int64(i.Falt)
		ingredient.Carbohydrates = int64(i.Carbohydrates)
		ingredient.Count = int64(i.Count)
		p = append(p, ingredient)
	}
	result.inggredients = p
	return result
}
