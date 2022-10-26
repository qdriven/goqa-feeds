package config

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
)

const (
	DefaultConfigName = ".goqa"
	DefaultConfigType = "toml"
)

func InitConfig(cfgFile string) {
	viper.SetConfigType(DefaultConfigType)
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		fmt.Println("home dir is {}", home)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(DefaultConfigName)
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
