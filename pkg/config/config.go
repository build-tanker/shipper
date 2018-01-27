package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config object
type Config struct {
	name         string
	version      string
	logLevel     string
	uploadServer string
}

var config *Config

// Init config from file
func Init() {
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../..")
	viper.SetConfigName("application")
	viper.SetConfigType("toml")

	viper.SetDefault("application.name", "shipper")
	viper.SetDefault("application.version", "NotDefined")
	viper.SetDefault("application.logLevel", "debug")
	viper.SetDefault("application.uploadServer", "http://localhost:8000")

	viper.ReadInConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file %s was edited, reloading config\n", e.Name)
		readLatestConfig()
	})

	readLatestConfig()
}

func readLatestConfig() {
	config = &Config{
		name:         viper.GetString("application.name"),
		version:      viper.GetString("application.version"),
		logLevel:     viper.GetString("application.logLevel"),
		uploadServer: viper.GetString("application.uploadServer"),
	}

}

// Name : Exporting Name
func Name() string {
	return config.name
}

// Version : Export application version
func Version() string {
	return config.version
}

// LogLevel : Export the log level
func LogLevel() string {
	return config.logLevel
}

// UploadServer : Get the upload server
func UploadServer() string {
	return config.uploadServer
}
