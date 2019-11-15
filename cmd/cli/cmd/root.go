package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/rema424/hexample/internal/hexample"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		var msg string
		if len(args) != 0 {
			msg = args[0]
		} else {
			msg = "Hello, from cli!"
		}

		arg := hexample.AppCoreLogicIn{
			From:    "cli",
			Message: msg,
		}

		hexample.AppCoreLogic(context.Background(), arg)
	},
}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
