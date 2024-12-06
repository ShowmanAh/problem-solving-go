package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "do command marking task as complete...",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Missing task ID")
			os.Exit(1)
		}
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID")
			os.Exit(1)
		}
		fmt.Printf("Marking task %d as a complete....\n", taskID)
	},
}
