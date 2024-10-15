/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// argsCmd represents the args command
var argsCmd = &cobra.Command{
	Use:   "args",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		num1, err1 := strconv.Atoi(args[0])
		num2, err2 := strconv.Atoi(args[1])
		if err1 != nil || err2 != nil {
			log.Fatalf("Invalid input: %v, %v", err1, err2)
		}
		result := num1 + num2
		fmt.Printf("Result: %d\n", result)
	},
}

func init() {
	rootCmd.AddCommand(argsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// argsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// argsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
