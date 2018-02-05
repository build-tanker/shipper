package logger

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"source.golabs.io/core/shipper/pkg/config"
)

func TestLogger(t *testing.T) {
	conf := config.NewConfig()
	var b bytes.Buffer
	log := NewLogger(conf, &b)
	log.Errorln("Testing", "hello")
	assert.Contains(t, b.String(), "Testing hello")
}
