package logger

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gojekfarm/shipper/pkg/config"
)

func TestLogger(t *testing.T) {
	conf := config.NewConfig([]string{"$HOME"})
	var b bytes.Buffer

	log := NewLogger(conf, &b)
	log.Debugln("Testing", "Debug")
	assert.Contains(t, b.String(), "Testing Debug")

	log.Errorln("Testing", "Error")
	assert.Contains(t, b.String(), "Testing Error")

	log.Infoln("Testing", "Info")
	assert.Contains(t, b.String(), "Testing Info")
}
