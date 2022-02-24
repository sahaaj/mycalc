/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"strconv"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		f, _ := cmd.Flags().GetBool("float")
		if f {
			addFloat(args)
		} else {
			addInt(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("float", "f", false, "To add floating point numbers")
}

func addInt(args []string) {
	var sum int
	for _, a := range args {
		itemp, err := strconv.Atoi(a)
		if err != nil {
			log.Fatal(err)
		}
		sum += itemp
	}
	fmt.Printf("The addition of the numbers %s is %d\n", args, sum)
}

func addFloat(args []string) {
	var sum float64
	for _, a := range args {
		itemp, err := strconv.ParseFloat(a, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += itemp
	}
	fmt.Printf("The sum of the numbers %s is %f\n", args, sum)
}
