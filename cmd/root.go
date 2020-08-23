package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fakefinder-bot-api",
	Short: "Check for errors in the Bot-API",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("An error ocurred, exiting execution.")
		os.Exit(1)
	}
}
