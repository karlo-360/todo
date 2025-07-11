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

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "complete a tasks",
	Long: `mark as completed any of you tasks`,
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
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

		if tasks[id][3] != "NO" {
			fmt.Println("this task is alredy complete")
			return
		}

		tasks[id][3] = "YES"
		fmt.Println("task completed")

		fw, err := os.OpenFile("/home/karlo/Documents/tasks.csv", os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatalln("error opnening the file for writing: ", err)
		}
		defer fw.Close()

		w := csv.NewWriter(fw)

		for _, v := range tasks{
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
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
