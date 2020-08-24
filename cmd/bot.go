package cmd

import (
	"fmt"
	"github.com/fake-finder/fakefinder/bot"
	"github.com/spf13/cobra"
)

var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "Initiate the Bot-API",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("command bot")

		bot, err := bot.NewBot()
		if err != nil {
			fmt.Println("erro new instance")
			return err
		}
		if err := bot.ReceiveMessage(); err != nil {
			fmt.Println("error", err)
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(botCmd)
}
