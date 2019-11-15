package service1

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
	fmt.Println("--------------------------------------------------")
	fmt.Println("service1:")
	fmt.Println("this is application core logic.")
	fmt.Printf("from: %s, message: %s\n", in.From, in.Message)
	fmt.Println("--------------------------------------------------")
}
