package cmd

import (
	"fmt"
	"strconv"

	db "../db"
	tasks "../db"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "marks a task as complete.",

	Run: func(cmd *cobra.Command, args []string) {
		var ints []int
		for _, num := range args {
			number, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("couldn't parse %s  \n ", num)
			} else {
				ints = append(ints, number)
			}
		}

		tasks, err := tasks.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		for _, id := range ints {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", id)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
