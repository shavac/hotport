/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	cfgPath   string
	cfgFormat string
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config", // [-h] [-json <path>] [-toml <path>] [-net <ip:port>]",
	Short: "Read or display port config from file or network",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			cfgPath = args[0]
		}
		fmt.Println(cfgPath)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	toml := *configCmd.Flags().BoolP("toml", "t", true, "toml config file")
	json := *configCmd.Flags().BoolP("json", "j", false, "json config file")
	net := *configCmd.Flags().BoolP("net", "n", false, "config from net")
	fmt.Println(toml, json, net)
}
