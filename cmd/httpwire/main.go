package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rema424/hexample/pkg/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rema424/sqlxx"
)

var e = createMux()

func main() {
	http.Handle("/", e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func init() {
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
	acsr, err := sqlxx.Open(db)
	if err != nil {
		log.Fatalln(err)
	}

	// --------------------------------------------------
	// 手動でDI
	// --------------------------------------------------

	// Service2
	// gateway2 := service2.NewGateway(db)
	// provider2 := service2.NewProvider(gateway2)
	// mockGateway2 := service2.NewMockGateway(service2.NewMockDB())
	// provider2 := service2.NewProvider(mockGateway2)

	// Service3
	// gateway3 := service3.NewGateway(acsr)
	// provider3 := service3.NewProvider(gateway3)
	// mockGateway3 := service3.NewMockGateway(service3.NewMockDB())
	// provider3 := service3.NewProvider(mockGateway3)

	// ctrl := &controller.Controller{}
	// ctrl2 := controller.NewController2(provider2)
	// ctrl3 := controller.NewController3(provider3)

	// --------------------------------------------------
	// wireでDI
	// --------------------------------------------------

	ctrls := InitializeControllers(db, acsr)

	e.GET("/:message", ctrls.Ctrl.HandleMessage)
	e.GET("/people/:personID", ctrls.Ctrl2.HandlePersonGet)
	e.POST("/people", ctrls.Ctrl2.HandlePersonRegister)
	e.POST("/accounts", ctrls.Ctrl3.HandleAccountOpen)
	e.POST("/accounts/transfer", ctrls.Ctrl3.HandleMoneyTransfer)
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}
