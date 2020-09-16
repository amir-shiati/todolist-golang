package cmd

import (
	"fmt"
	"os"
	"strings"

	tasks "../db"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a new task to the todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		id, err := tasks.CreateTask(task)
		if err != nil {
			fmt.Println("couldn't add the task")
			os.Exit(1)
		}
		fmt.Println("tasks added, task id : ", id)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
