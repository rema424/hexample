//+build wireinject

package main

import (
	"github.com/rema424/hexample/cmd/httpwire/controller"
	"github.com/rema424/hexample/internal/service2"
	"github.com/rema424/hexample/internal/service3"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/rema424/sqlxx"
)

// InitializeControllers .
func InitializeControllers(db *sqlx.DB, acsr *sqlxx.Accessor) *controller.Controllers {
	wire.Build(
		service2.NewGateway,
		service3.NewGateway,
		service2.NewProvider,
		service3.NewProvider,
		controller.NewController,
		controller.NewController2,
		controller.NewController3,
		controller.NewControllers,
	)
	return &controller.Controllers{}
}
