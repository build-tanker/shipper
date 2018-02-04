package uploader

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"source.golabs.io/core/shipper/pkg/appcontext"
	"source.golabs.io/core/shipper/pkg/config"
	"source.golabs.io/core/shipper/pkg/logger"
)

var testContext *appcontext.AppContext

func NewTestContext() *appcontext.AppContext {
	if testContext == nil {
		conf := config.NewConfig()
		log := logger.NewLogger(conf)
		testContext = appcontext.NewAppContext(conf, log)
	}
	return testContext
}

type MockClient struct{}

func NewTestService() *service {
	ctx := NewTestContext()
	client := &MockClient{}
	return &service{
		ctx:    ctx,
		client: client,
	}
}

func TestServiceInstall(t *testing.T) {
	s := NewTestService()

	s.ctx.GetConfig().AccessKey = "testAccessKey"
	s.ctx.GetConfig().Server = "testServer"
	err := s.Install("http://localhost:8000")
	assert.Equal(t, "Non empty config already present", err.Error())

	s.ctx.GetConfig().AccessKey = ""
	s.ctx.GetConfig().Server = ""
	err = s.Install("")
	assert.Equal(t, "Server flag missing", err.Error())

	err = s.Install("http://localhost:8000")
	assert.Nil(t, err)

}

func TestServiceUninstall(t *testing.T) {
	s := NewTestService()

	s.ctx.GetConfig().AccessKey = ""
	s.ctx.GetConfig().Server = ""
	err := s.Uninstall()
	assert.Equal(t, "No config file found", err.Error())

	s.ctx.GetConfig().AccessKey = "testAccessKey"
	s.ctx.GetConfig().Server = "testServer"
	err = s.Uninstall()
	assert.Nil(t, err)
}
