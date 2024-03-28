package conf

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitViper(ctx context.Context, confFile string) *Config {
	if ok := pathExists(confFile); !ok {
		panic("config file path not exists")
	}

	v := viper.New()
	v.SetConfigFile(confFile)

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %s", err))
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %s", err))
	}

	return &config
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}
