package cmd

import (
	"fmt"
	"github.com/qdriven/go-for-qa/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "qago",
	Short: "qago is a qa daily cli application",
	Long:  `A QA tools for daily testing usage`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		println(args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type CliBaseInfo struct {
	CfgFile     string `default: goqa.toml`
	ProjectBase string `default: ./`
	Author      string `default: Patrick`
	License     string `default: NDV`
	UseViper    bool   `default: true`
}

var appInfo = &CliBaseInfo{}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&appInfo.CfgFile, "config", "", "config file (default is ./goqa.yaml)")
	rootCmd.PersistentFlags().StringVarP(&appInfo.ProjectBase, "workspace", "b", "", "base project directory eg. github.com/spf13/")
	rootCmd.PersistentFlags().StringP("author", "a", appInfo.Author, "Author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&appInfo.License, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("workspace", rootCmd.PersistentFlags().Lookup("workspace"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", appInfo.Author)
	viper.SetDefault("license", appInfo.License)
}

func initConfig() {
	config.InitConfig(appInfo.CfgFile)
}
