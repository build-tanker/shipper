package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigValues(t *testing.T) {
	conf := NewConfig()
	assert.Equal(t, "debug", conf.LogLevel())
}
