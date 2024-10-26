package cmd

import (
	"corrector/app"
	"fmt"

	"github.com/spf13/cobra"
)

func Root(app *app.Server) *cobra.Command {
	root := cobra.Command{
		Use:   "showcase",
		Short: "showcase application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Type a 'showcase help' for usage details")
		},
	}

	//root.AddCommand(
	//	start(app),
	//)
	//root.AddCommand(
	//	importSL(app),
	//)

	return &root
}
