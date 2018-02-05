package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigValues(t *testing.T) {
	conf := NewConfig([]string{"./testutil"})
	assert.Equal(t, "test1234", conf.AccessKey)
	assert.Equal(t, "http://public.betas.in", conf.Server)
}
