package config

import "github.com/spf13/viper"

func MustLoad() {
	viper.SetConfigFile("cfg.yml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
