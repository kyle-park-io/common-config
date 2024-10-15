/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"viper-cobra-example/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// flagsCmd represents the flags command
var flagsCmd = &cobra.Command{
	Use:   "flags",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		executeApp2()
	},
}

func executeApp2() {
	port := viper.GetInt("server.port")
	debug := viper.GetBool("server.debug")

	log.Printf("Server running on port %d, debug mode: %v\n", port, debug)
}

func init() {
	rootCmd.AddCommand(flagsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// flagsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// flagsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	cobra.OnInitialize(config.InitConfig)

	// flag
	rootCmd.PersistentFlags().Int("port", 8080, "Server port")
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug mode")
	viper.BindPFlag("server.port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("server.debug", rootCmd.PersistentFlags().Lookup("debug"))
}
