/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks",
	Long: `list all the tasks save for the user`,
	Args: cobra.NoArgs,
	DisableFlagsInUseLine: true,
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

		w := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', 0,)
		for _, v := range tasks{
			fmt.Fprintf(w, "%s\t%s\t%s\n", v[0], v[1], v[2])
		}

		if err := w.Flush(); err != nil{
			log.Fatalln("error while flushing: ", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
