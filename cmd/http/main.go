package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rema424/hexample/cmd/http/controller"
	"github.com/rema424/hexample/internal/service2"
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
	// c := mysql.Config{
	// 	Host:                 os.Getenv("DB_HOST"),
	// 	Port:                 os.Getenv("DB_PORT"),
	// 	User:                 os.Getenv("DB_USER"),
	// 	DBName:               os.Getenv("DB_NAME"),
	// 	Passwd:               os.Getenv("DB_PASSWORD"),
	// 	AllowNativePasswords: true,
	// }
	// db, err := mysql.Connect(c)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// svc2Repo := service2.NewRepositoryImpl(db)
	// svc2Pvdr := service2.NewProvider(svc2Repo)

	// MockDB
	mockDB := service2.NewMockDB()
	svc2RepoMock := service2.NewRepositoryImplMock(mockDB)
	svc2Pvdr := service2.NewProvider(svc2RepoMock)

	ctrl := &controller.Controller{}
	ctrl2 := controller.NewController2(svc2Pvdr)

	e.GET("/:message", ctrl.HandleMessage)
	e.GET("/people/:personID", ctrl2.HandlePersonGet)
	e.POST("/people", ctrl2.HandlePersonRegister)
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}
