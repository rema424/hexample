// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/rema424/hexample/cmd/httpwire/controller"
	"github.com/rema424/hexample/internal/service2"
	"github.com/rema424/hexample/internal/service3"
	"github.com/rema424/sqlxx"
)

// Injectors from wire.go:

func InitializeControllers(db *sqlx.DB, acsr *sqlxx.Accessor) *controller.Controllers {
	controllerController := controller.NewController()
	repository := service2.NewGateway(db)
	provider := service2.NewProvider(repository)
	controller2 := controller.NewController2(provider)
	service3Repository := service3.NewGateway(acsr)
	service3Provider := service3.NewProvider(service3Repository)
	controller3 := controller.NewController3(service3Provider)
	controllers := controller.NewControllers(controllerController, controller2, controller3)
	return controllers
}
