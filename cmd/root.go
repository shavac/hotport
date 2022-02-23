/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/shavac/mp1p/cfg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	CfgPath string
	CfgFmt  string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mp1p",
	Short: "The network protocol proxy, support \"Multiple Protocols on 1 Port\" in particular. ",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if err := cfg.ReadFromFile(CfgPath, CfgFmt); err != nil {
			log.Errorln(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mp1p.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&CfgPath, "config", "c", "/etc/mp1p/mp1p.toml", "config file path")
	rootCmd.PersistentFlags().StringVarP(&CfgFmt, "type", "t", "toml", "config file format <toml|json|net>")
}
