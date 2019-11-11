package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-appengine-clean-architecture/pkg/mysql"
	"go-appengine-clean-architecture/pkg/server"

	"github.com/rema424/sqlxx"
	"google.golang.org/appengine"
	"gopkg.in/go-playground/validator.v9"
)

var e = server.CreateMux()

func init() {
	// DB
	c := mysql.Config{
		Host:                 os.Getenv("DB_HOST"),
		Port:                 os.Getenv("DB_PORT"),
		User:                 os.Getenv("DB_USER"),
		DBName:               os.Getenv("DB_NAME"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		InterpolateParams:    true,
		AllowNatevePasswords: true,
		ParseTime:            true,
		MaxOpenConns:         -1, // use default value
		MaxIdleConns:         -1, // use default value
		ConnMaxLifetime:      -1, // use default value
	}
	db, err := mysql.Connect(c)
	if err != nil {
		log.Fatalln(err)
	}
	acsr, err := sqlxx.Open(db)
	if err != nil {
		log.Fatalln(err)
	}

	// Validator
	e.Validator = server.CreateCustomValidator(map[string]validator.Func{})

	// DI
	ctrl := InitializeControllers(acsr)

	// Routes
	fmt.Println(ctrl)
}

func main() {
	http.Handle("/", e)
	appengine.Main()
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// 	log.Printf("Defaulting to port %s", port)
	// }
	// log.Printf("Listening on port %s", port)
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
