package controller

import (
	"net/http"

	"github.com/rema424/hexample/internal/service3"

	"github.com/labstack/echo/v4"
)

// Controller3 ...
type Controller3 struct {
	p *service3.Provider
}

// NewController3 ...
func NewController3(p *service3.Provider) *Controller3 {
	return &Controller3{p}
}

// HandleAccountOpen ...
// curl -X POST -H 'Content-type: application/json' -d '{"ammount": 1000}' localhost:8080/accounts
func (ctrl *Controller3) HandleAccountOpen(c echo.Context) error {
	in := struct {
		Ammount int `json:"ammount"`
	}{}

	if err := c.Bind(&in); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// TODO: implement
	// if err := c.Validate(&in); err != nil {
	// 	return c.JSON(http.StatusUnprocessableEntity, err.Error())
	// }

	ctx := c.Request().Context()
	psn, err := ctrl.p.OpenAccount(ctx, in.Ammount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, psn)
}

// HandleMoneyTransfer ...
// curl -X POST -H 'Content-type: application/json' -d '{"fromId": , "toId": , "ammount": 1000}' localhost:8080/accounts/transfer
func (ctrl *Controller3) HandleMoneyTransfer(c echo.Context) error {
	in := struct {
		FromAccountID int64 `json:"fromId"`
		ToAccountID   int64 `json:"toId"`
		Ammount       int   `json:"ammount"`
	}{}

	if err := c.Bind(&in); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// TODO: implement
	// if err := c.Validate(&in); err != nil {
	// 	return c.JSON(http.StatusUnprocessableEntity, err.Error())
	// }

	ctx := c.Request().Context()
	from, to, err := ctrl.p.Transfer(ctx, in.Ammount, in.FromAccountID, in.ToAccountID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"from": from, "to": to})
}
