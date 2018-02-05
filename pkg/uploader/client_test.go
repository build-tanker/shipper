package uploader

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"source.golabs.io/core/shipper/pkg/requester"
)

var testState string

type MockRequester struct {
}

func (m MockRequester) Get(url string) ([]byte, error) {
	var bytes []byte
	var err error
	switch testState {
	default:
		bytes = []byte{}
		err = nil
	}
	return bytes, err
}

func (m MockRequester) Post(url string) ([]byte, error) {
	var bytes []byte
	var err error
	switch testState {
	case "GetAccessKey":
		bytes = []byte(`{ "data":{ "id":2, "access_key":"testAccessKey"},"success":"true" }`)
		err = nil
	default:
		bytes = []byte{}
		err = nil
	}
	return bytes, err
}

func (m MockRequester) Put(url string) ([]byte, error) {
	var bytes []byte
	var err error
	switch testState {
	default:
		bytes = []byte{}
		err = nil
	}
	return bytes, err
}

func (m MockRequester) Delete(url string) ([]byte, error) {
	var bytes []byte
	var err error
	switch testState {
	case "DeleteAccessKeySuccess":
		bytes = []byte(`{ "success": "true" }`)
		err = nil
	case "DeleteAccessKeyFailure":
		bytes = []byte(`{ "success": "false" }`)
		err = nil
	default:
		bytes = []byte{}
		err = nil
	}
	return bytes, err
}

func NewMockRequester() requester.Requester {
	return MockRequester{}
}

func NewTestClient() Client {
	ctx := NewTestContext()
	r := NewMockRequester()
	return &client{
		ctx: ctx,
		r:   r,
	}
}

func TestGetAccessKey(t *testing.T) {
	c := NewTestClient()
	testState = "GetAccessKey"
	accessKey, err := c.GetAccessKey("mockServer")
	assert.Nil(t, err)
	assert.Equal(t, "testAccessKey", accessKey)
}

func TestDeleteAccessKey(t *testing.T) {
	c := NewTestClient()
	testState = "DeleteAccessKeySuccess"
	err := c.DeleteAccessKey("mockServer", "mockAccessKey")
	assert.Nil(t, err)

	testState = "DeleteAccessKeyFailure"
	err = c.DeleteAccessKey("mockServer", "mockAccessKey")
	assert.Equal(t, "Could not delete AccessKey from the server", err.Error())
}
