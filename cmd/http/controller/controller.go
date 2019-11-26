package controller

import (
	"net/http"

	"github.com/rema424/hexample/internal/service1"

	"github.com/labstack/echo/v4"
)

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
	return c.String(http.StatusOK, msg)
}
