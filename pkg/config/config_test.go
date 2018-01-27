package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sudhanshuraheja/shippers-uploader/pkg/config"
)

func TestConfigValues(t *testing.T) {
	config.Init()
	assert.Equal(t, "debug", config.LogLevel())
}
