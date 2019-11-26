package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/rema424/hexample/cmd/http/controller"
	"github.com/rema424/hexample/internal/service2"
	"github.com/rema424/hexample/internal/service3"
	"github.com/rema424/hexample/pkg/mysql"
	"github.com/rema424/sqlxx"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	// Setup
	e := newEcho()
	dbx := newDBx()
	dbxx := newDBxx(dbx)
	handler := newHandler(e, dbx, dbxx)

	// Server
	log.Printf("Listening on port %s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
	if err != nil {
		log.Fatal(err)
	}
}

func newHandler(e *echo.Echo, dbx *sqlx.DB, dbxx *sqlxx.Accessor) http.Handler {
	// --------------------
	// DI
	// --------------------
	// Service2
	gateway2 := service2.NewGateway(dbx)
	provider2 := service2.NewProvider(gateway2)
	// mockGateway2 := service2.NewMockGateway(service2.NewMockDB())
	// provider2 := service2.NewProvider(mockGateway2)

	// Service3
	// gateway3 := service3.NewGateway(dbxx)
	// provider3 := service3.NewProvider(gateway3)
	mockGateway3 := service3.NewMockGateway(service3.NewMockDB())
	provider3 := service3.NewProvider(mockGateway3)

	ctrl := &controller.Controller{}
	ctrl2 := controller.NewController2(provider2)
	ctrl3 := controller.NewController3(provider3)

	// --------------------
	// HandlerFunc
	// --------------------
	e.GET("/:message", ctrl.HandleMessage)
	e.GET("/people/:personID", ctrl2.HandlePersonGet)
	e.POST("/people", ctrl2.HandlePersonRegister)
	e.POST("/accounts", ctrl3.HandleAccountOpen)
	e.POST("/accounts/transfer", ctrl3.HandleMoneyTransfer)
	return e
}

func newEcho() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

func newDBx() *sqlx.DB {
	// Mysql
	c := mysql.Config{
		Host:                 os.Getenv("DB_HOST"),
		Port:                 os.Getenv("DB_PORT"),
		User:                 os.Getenv("DB_USER"),
		DBName:               os.Getenv("DB_NAME"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		AllowNativePasswords: true,
	}
	db, err := mysql.Connect(c)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func newDBxx(dbx *sqlx.DB) *sqlxx.Accessor {
	dbxx, err := sqlxx.Open(dbx)
	if err != nil {
		log.Fatalln(err)
	}
	return dbxx
}
