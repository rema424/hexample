package controller

import "github.com/labstack/echo/v4"

import "github.com/rema424/hexample/internal/service1"

// Controller ...
type Controller struct{}

// HandleMessage ...
func (ctrl *Controller) HandleMessage(c echo.Context) error {
	msg := c.Param("message")
	if msg == "" {
		msg = "Hello, from http!"
	}

	arg := service1.AppCoreLogicIn{
		From:    "http",
		Message: msg,
	}

	service1.AppCoreLogic(c.Request().Context(), arg)
	return nil
}
