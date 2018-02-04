package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server    string
	AccessKey string
}

func NewConfig() *Config {
	config := &Config{}

	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../..")

	viper.SetConfigName(".shipper")
	viper.SetConfigType("toml")

	viper.ReadInConfig()
	config.readLatestConfig()

	return config
}

func (c *Config) readLatestConfig() {
	c.Server = viper.GetString("application.server")
	c.AccessKey = viper.GetString("application.accessKey")
}

func (c *Config) IsMissing() bool {
	return c.Server == "" && c.AccessKey == ""
}

func (c *Config) Write(server string, accessKey string) error {
	return nil
}
