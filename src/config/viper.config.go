package config

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

func GetEnv(key string) string {
	return fmt.Sprint(viper.Get(key))
}

func PrepareViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(filepath.Join(".", "src", "config"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
