package cmd

import (
	"github.com/fake-finder/fakefinder/botApi"
	"github.com/spf13/cobra"
)

var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "Initiate the Bot-API",
	Run: func(cmd *cobra.Command, args []string) {
		botApi.ReceiveMessage()
	},
}

func init() {
	rootCmd.AddCommand(botCmd)
}
