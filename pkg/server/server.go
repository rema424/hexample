package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/go-playground/validator.v9"
)

// CreateMux .
func CreateMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

// CustomValidator .
type CustomValidator struct {
	validator *validator.Validate
}

// Validate .
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// CreateCustomValidator .
func CreateCustomValidator(m map[string]validator.Func) *CustomValidator {
	v := validator.New()
	for key, val := range m {
		v.RegisterValidation(key, val)
	}
	return &CustomValidator{validator: v}
}

func sampleValidationRule(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	if len(val) > 30 {
		return true
	}
	return false
}
