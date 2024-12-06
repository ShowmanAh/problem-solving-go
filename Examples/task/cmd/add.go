package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add new task...",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add new task....")
		fmt.Printf("Adding a new task: %s ... \n", strings.Join(args, " "))
	},
}
