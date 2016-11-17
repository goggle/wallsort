package imsort

import (
	"fmt"

	"github.com/spf13/viper"
)

func Initialize() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.config/imsort")
	viper.AddConfigPath(".")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	panic(fmt.Errorf("Fatal error config file: %s\n", err))
	//
	// }
	err := viper.ReadInConfig()
	return err
}

func ReadConfiguration() error {
	// Config.Directory = viper.GetString("directory")
	// d := viper.GetStringMap("dir")
	// _ = d
	viper.Unmarshal(&Config)
	fmt.Println(Config)
	fmt.Println(viper.IsSet("height"))
	return nil
}
