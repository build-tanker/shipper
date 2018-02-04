package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"source.golabs.io/core/shipper/pkg/config"
)

func TestConfigValues(t *testing.T) {
	config.Init()
	assert.Equal(t, "debug", config.LogLevel())
}
