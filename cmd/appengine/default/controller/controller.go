package controller

import (
	"go-appengine-clean-architecture/internal/pkg/food"
)

// Controller .
type Controller struct {
	Food *FoodController
}

// NewController .
func NewController(f *FoodController) *Controller {
	return &Controller{f}
}

// FoodController .
type FoodController struct {
	fi *food.Interactor
}

// NewFoodController ...
func NewFoodController(fi *food.Interactor) *FoodController {
	return &FoodController{fi}
}
