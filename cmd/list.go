package cmd

import (
	"fmt"
	"os"

	tasks "../db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := tasks.AllTasks()

		if err != nil {
			fmt.Println("coudn't get tasks list")
			os.Exit(1)
		}

		if len(tasks) <= 0 {
			fmt.Println("YAY! you have no tasks to complete!")
			return
		}

		fmt.Println("You have the following tasks: ")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
