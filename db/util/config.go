package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver            string        `mapstructure:"DBDRIVER"`
	DBSource            string        `mapstructure:"DBSOURCE"`
	ServerAddress       string        `mapstructure:"SERVERADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKENSYMMETRICKEY"`
	TokenAccessDuration time.Duration `mapstructure:"TOKENACCESSDURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("App")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
