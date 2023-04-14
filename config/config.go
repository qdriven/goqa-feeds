package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

type DefaultConfigMap struct {
}

var (
	ConfigMapping        = map[string]*viper.Viper{}
	DEFFAULT_CONFIG_FILE = "app.toml"
)

func NewDefaultConfig() *viper.Viper {
	v := viper.New()
	v.SetConfigFile(DEFFAULT_CONFIG_FILE)

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("couldn't load config: %s", err)
		os.Exit(1)
	}
	ConfigMapping["DEFAULT"] = v
	watchConfigChanges(v, &DefaultConfigMap{})
	return v
}

func watchConfigChanges(v *viper.Viper, config any) {
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		fmt.Println(v.AllKeys())
		err := v.Unmarshal(config)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(config)
	})
}

func NewNamedConfig(configFile string, configName string, config any) any {
	v := viper.New()
	v.SetConfigFile(configFile)

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("couldn't load config: %s", err)
		os.Exit(1)
	}
	ConfigMapping[configName] = v
	err := v.Unmarshal(config)
	if err != nil {
		return nil
	}
	fmt.Println(err)
	watchConfigChanges(v, config)
	return config
}

func GetViperByName(name string) *viper.Viper {
	return ConfigMapping[name]
}
