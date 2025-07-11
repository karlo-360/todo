/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a task",
	Long: `delete any of you tasks`,
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

		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("failed to convert string to integer %s: %s\n", args[0], err)
		}

		if args[0] != tasks[id][0] {
			fmt.Println("these id not exist")
			return
		}

		newList := tasks[:id]

		newList = append(newList, tasks[id+1:len(tasks)]...)

		fw, err := os.OpenFile("/home/karlo/Documents/tasks.csv", os.O_WRONLY|os.O_TRUNC, 0777)
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
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
