package hexample

import (
	"context"
	"fmt"
)

// AppCoreLogicIn .
type AppCoreLogicIn struct {
	From    string
	Message string
}

// AppCoreLogic .
func AppCoreLogic(ctx context.Context, in AppCoreLogicIn) {
	fmt.Println("this is application core logic.")
	fmt.Println("from:", in.From, "message:", in.Message)
}
