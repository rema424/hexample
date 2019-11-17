package controller

// Controllers ...
type Controllers struct {
	Ctrl  *Controller
	Ctrl2 *Controller2
	Ctrl3 *Controller3
}

// NewControllers ...
func NewControllers(ctrl *Controller, ctrl2 *Controller2, ctrl3 *Controller3) *Controllers {
	return &Controllers{ctrl, ctrl2, ctrl3}
}
