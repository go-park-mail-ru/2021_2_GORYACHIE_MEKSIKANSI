package application

import (
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Restaurant/Interface"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Microservices/Restaurant/proto"
	restaurant "2021_2_GORYACHIE_MEKSIKANSI/internal/Restaurant"
	"context"
)

type RestaurantManager struct {
	DB Interface.WrapperRestaurant
}

func (rm *RestaurantManager) AllRestaurants(ctx context.Context, _ *resProto.Empty) (*resProto.Restaurants, error) {
	var a resProto.Restaurants
	result, err := rm.DB.GetRestaurants()
	if err != nil {
		return nil, err
	}
	for i, r := range result {
		a.Restaurants[i].Id = int64(r.Id)
		a.Restaurants[i].Img = r.Img
		a.Restaurants[i].MaxDelivery = int64(r.MaxDelivery)
		a.Restaurants[i].MinDelivery = int64(r.MinDelivery)
		a.Restaurants[i].CostForFreeDelivery = int64(r.CostForFreeDelivery)
		a.Restaurants[i].Rating = r.Rating
		a.Restaurants[i].Name = r.Name
	}
	return &a, nil
}

func (rm *RestaurantManager) GetRestaurant(ctx context.Context, id *resProto.RestaurantId) (*resProto.RestaurantInfo, error) {
	restInfo, err := rm.DB.GetRestaurant(int(id.Id))
	if err != nil {
		return nil, err
	}

	tags, err := rm.DB.GetTagsRestaurant(int(id.Id))
	if err != nil {
		return nil, err
	}

	dishes, err := rm.DB.GetMenu(int(id.Id))
	if err != nil {
		return nil, err
	}

	restInfo.Menu = dishes
	restInfo.Tags = tags
	var a *resProto.RestaurantInfo
	a.Name = restInfo.Name
	a.Img = restInfo.Img
	a.Rating = restInfo.Rating
	a.CostForFreeDelivery = int64(restInfo.CostForFreeDelivery)
	a.MinDelivery = int64(restInfo.MinDelivery)
	a.MaxDelivery = int64(restInfo.MaxDelivery)
	for i, m := range restInfo.Menu {
		a.Menu[i].Name = m.Name
		for j, me := range restInfo.Menu[i].DishesMenu {
			a.Menu[i].Dishes[j].Id = int64(me.Id)
			a.Menu[i].Dishes[j].Ccal = int64(me.Kilocalorie)
			a.Menu[i].Dishes[j].Cost = int64(me.Cost)
			a.Menu[i].Dishes[j].Img = me.Img
			a.Menu[i].Dishes[j].Name = me.Name
		}
	}
	for _, m := range restInfo.Tags {
		var c *resProto.Tag
		c.Name = m.Name
		c.Id = int64(m.Id)
		a.Tags = append(a.Tags, c)
	}
	return a, nil
}

func (rm *RestaurantManager) RestaurantDishes(ctx context.Context, id *resProto.DishInfo) (*resProto.Dishes, error) {
	var c resProto.Dishes
	dishes, err := rm.DB.GetDishes(int(id.RestaurantId), int(id.DishId))
	if err != nil {
		return nil, err
	}
	c.Cost = int64(dishes.Cost)
	c.Id = int64(dishes.Id)
	c.Img = dishes.Img
	c.Ccal = int64(dishes.Ccal)

	dishes.Ingredient, err = rm.DB.GetStructDishes(int(id.DishId))
	if err != nil {
		return nil, err
	}

	for j, i := range dishes.Ingredient {
		c.Ingredients[j].Cost = int64(i.Cost)
		c.Ingredients[j].Id = int64(i.Id)
		c.Ingredients[j].Name = i.Title
	}

	dishes.Radios, err = rm.DB.GetRadios(int(id.DishId))
	if err != nil {
		return nil, err
	}

	for j, i := range dishes.Radios {
		c.Radios[j].Id = int64(i.Id)
		c.Radios[j].Name = i.Title
		for k, e := range i.Rows {
			c.Radios[j].Rows[k].Id = int64(e.Id)
			c.Radios[j].Rows[k].Name = e.Name
		}
	}

	return &c, nil
}

func (rm *RestaurantManager) CreateReview(ctx context.Context, rev *resProto.NewReview) (*resProto.Error, error) {
	var r restaurant.RestaurantId
	r.Id = int(rev.Restaurant.Id)
	err := rm.DB.CreateReview(int(rev.Id), restaurant.NewReview{Rate: int(rev.Rate),
		Restaurant: r, Text: rev.Text})
	if err != nil {
		return nil, err
	}
	return &resProto.Error{}, nil
}

func (rm *RestaurantManager) GetReview(ctx context.Context, id *resProto.RestaurantId) (*resProto.ResReview, error) {
	var review restaurant.ResReview
	reviewInfo, err := rm.DB.GetReview(int(id.Id))
	if err != nil {
		return nil, err
	}

	restInfo, err := rm.DB.GetRestaurant(int(id.Id))
	if err != nil {
		return nil, err
	}

	tags, err := rm.DB.GetTagsRestaurant(int(id.Id))
	if err != nil {
		return nil, err
	}
	restInfo.Tags = tags

	review.CastFromRestaurantId(*restInfo)
	review.Tags = tags
	review.Reviews = reviewInfo

	end := &resProto.ResReview{}
	end.Id = int64(review.Id)
	end.Name = review.Name
	end.Rating = review.Rating

	var a []*resProto.Review
	for _, r := range review.Reviews {
		var rev *resProto.Review
		rev.Time = r.Time
		rev.Text = r.Text
		rev.Name = r.Name
		rev.Rate = int64(r.Rate)
		rev.Date = r.Date
	}
	end.Review = a

	var b []*resProto.Tag
	for _, r := range review.Tags {
		var rev *resProto.Tag
		rev.Id = int64(r.Id)
		rev.Name = r.Name
	}
	end.Tags = b

	end.MinDelivery = int64(review.MinDelivery)
	end.MaxDelivery = int64(review.MaxDelivery)
	end.CostForFreeDelivery = int64(review.CostForFreeDelivery)
	end.Img = review.Img
	return end, nil
}

func (rm *RestaurantManager) SearchRestaurant(ctx context.Context, search *resProto.SearchRestaurantText) (*resProto.Restaurants, error) {
	result, err := rm.DB.SearchCategory(search.Text)
	if err != nil {
		return nil, err
	}

	if result == nil {
		result, err = rm.DB.SearchRestaurant(search.Text)
		if err != nil {
			return nil, err
		}
	}

	var searchResult *resProto.Restaurants
	for _, id := range result {
		var rest resProto.Restaurant
		restaurantInfo, err := rm.DB.GetGeneralInfoRestaurant(id)
		if err != nil {
			return nil, err
		}
		rest.Img = restaurantInfo.Img
		rest.Name = restaurantInfo.Name
		rest.CostForFreeDelivery = int64(restaurantInfo.CostForFreeDelivery)
		rest.MaxDelivery = int64(restaurantInfo.MaxDelivery)
		rest.Rating = restaurantInfo.Rating
		rest.Id = int64(restaurantInfo.Id)
		rest.MinDelivery = int64(restaurantInfo.MinDelivery)
		searchResult.Restaurants = append(searchResult.Restaurants, &rest)
	}
	return searchResult, nil
}
