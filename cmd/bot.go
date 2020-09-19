package cmd

import (
	"fmt"

	"github.com/fake-finder/fakefinder/app"
	"github.com/spf13/cobra"
)

var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "Initiate the Bot-API",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("command bot")

		if err := app.Run(); err != nil {
			fmt.Println("error", err)
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(botCmd)
}
