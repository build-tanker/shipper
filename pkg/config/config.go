package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	name     string
	version  string
	logLevel string
	server   string
}

func NewConfig() *Config {
	config := &Config{}

	viper.AutomaticEnv()

	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../..")

	viper.SetConfigName(".shipper")
	viper.SetConfigType("toml")

	viper.SetDefault("application.name", "shipper")
	viper.SetDefault("application.version", "0.0.0")
	viper.SetDefault("application.logLevel", "debug")
	viper.SetDefault("application.server", "http://localhost:8080")

	viper.ReadInConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file %s was edited, reloading config\n", e.Name)
		config.readLatestConfig()
	})

	config.readLatestConfig()

	return config
}

func (c *Config) Name() string {
	return c.name
}

func (c *Config) Version() string {
	return c.version
}

func (c *Config) LogLevel() string {
	return c.logLevel
}

func (c *Config) Server() string {
	return c.server
}

func (c *Config) readLatestConfig() {
	c.name = viper.GetString("application.name")
	c.version = viper.GetString("application.version")
	c.logLevel = viper.GetString("application.logLevel")
	c.server = viper.GetString("application.server")
}
