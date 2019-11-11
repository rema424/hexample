//+build wireinject

package main

import (
	"go-appengine-clean-architecture/internal/app/default/controller"
	"go-appengine-clean-architecture/internal/pkg/food"

	"github.com/google/wire"
	"github.com/rema424/sqlxx"
)

// InitializeControllers .
func InitializeControllers(acsr *sqlxx.Accessor) *controller.Controller {
	wire.Build(
		food.NewGateway,
		food.NewInteractor,
		controller.NewFoodController,
		controller.NewController,
	)
	return &controller.Controller{}
}
