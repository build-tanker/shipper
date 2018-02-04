package config

import (
	"github.com/spf13/viper"
)

// Config - structure to hold the configuration for shipper
type Config struct {
	Server    string
	AccessKey string
}

// NewConfig - create a new configuration
func NewConfig() *Config {
	config := &Config{}

	viper.AddConfigPath("$HOME")
	// viper.AddConfigPath(".")
	// viper.AddConfigPath("..")
	// viper.AddConfigPath("../..")

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

// IsMissing - check if none of the values in the config could be read
func (c *Config) IsMissing() bool {
	return c.Server == "" && c.AccessKey == ""
}
