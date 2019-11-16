package hexample

import (
	"context"

	"github.com/rema424/hexample/internal/service1"
)

// Run ...
func Run(ctx context.Context, msg string) {
	if msg == "" {
		msg = "Hello, from external pkg!"
	}
	arg := service1.AppCoreLogicIn{
		From:    "external pkg",
		Message: msg,
	}
	service1.AppCoreLogic(ctx, arg)
}
