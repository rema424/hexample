package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rema424/hexample/internal/service2"
)

// Controller2 ...
type Controller2 struct {
	p *service2.Provider
}

// NewController2 ...
func NewController2(p *service2.Provider) *Controller2 {
	return &Controller2{p}
}

// HandlePersonRegister ...
// curl -X POST -H 'Content-type: application/json' -d '{"name": "Alice", "email": "alice@example.com"}' localhost:8080/people
func (ctrl *Controller2) HandlePersonRegister(c echo.Context) error {
	in := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{}

	if err := c.Bind(&in); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// TODO: implement
	// if err := c.Validate(&in); err != nil {
	// 	return c.JSON(http.StatusUnprocessableEntity, err.Error())
	// }

	ctx := c.Request().Context()
	psn, err := ctrl.p.RegisterPerson(ctx, in.Name, in.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, psn)
}

// HandlePersonGet ...
// curl localhost:8080/people/999
func (ctrl *Controller2) HandlePersonGet(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("personID"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	psn, err := ctrl.p.GetPersonByID(ctx, int64(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, psn)
}
