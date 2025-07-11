/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"

	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a task",
	Long: `add a task to the todo list`,
	Args: cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {

		fr, err := os.Open("/home/karlo/Documents/tasks.csv")
		if err != nil {
			log.Fatalln("error opening the file for reading: ", err)
		}

		r := csv.NewReader(fr)

		tasks, err := r.ReadAll()
		if err != nil {
			log.Fatalln("error reading all: ", err)
		}

		if err := fr.Close(); err != nil{
			log.Fatalln("error closing file reader: ", err)
		}

		var id int
		if len(tasks) == 0 {
			tasks = append(tasks, []string{"ID", "NAME", "DATE", "COMPLETED"})
			id = 1
		} else if len(tasks) == 1 {
			id = 1
		} else {
			for i, v := range tasks{

				if i == 0 {
					continue
				}

				taskId, err := strconv.Atoi(v[0])
				if err != nil {
					log.Fatalf("failed to convert string to integer at index %d: %s\n", i, err)
				}

				if i != taskId {
					id = i
					break
				} else {
					id = taskId + 1
				}
			}
		}


		newId := strconv.Itoa(id)
		name := strings.Join(args, " ")
		date := time.Now().Format(time.DateOnly)

		originTasks := append([][]string(nil), tasks[id:]...)
		newList := append(
			tasks[:id],
			[]string{
				newId,
				name,
				date,
				"NO",
			},
		)

		if id != 0 {
			newList = append(newList, originTasks...)
		}

		fw, err := os.OpenFile("/home/karlo/Documents/tasks.csv", os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Fatalln("error opnening the file for writing: ", err)
		}
		defer fw.Close()

		w := csv.NewWriter(fw)

		for _, v := range newList{
			if err := w.Write(v); err != nil {
				log.Fatalln("error wrinting newTask to csv: ", err)
			}
		}

		w.Flush()
		if err := w.Error(); err != nil {
			log.Fatalln("error while flushing: ", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
