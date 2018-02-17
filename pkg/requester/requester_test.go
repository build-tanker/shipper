package requester

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRequester(t *testing.T) {
	mockServer := MockServer{}
	mockServer.Start("9000")

	r := NewRequester(5 * time.Minute)

	// Fail with http / https
	_, err := r.Get("localhost:9000/get")
	assert.Equal(t, `requester:call Could not complete request: Get localhost:9000/get: unsupported protocol scheme "localhost"`, err.Error())

	bytes, err := r.Get("http://localhost:9000/get")
	assert.Nil(t, err)
	assert.Equal(t, `{ "data": { "method": "get" }, "success": "true" }`, string(bytes))

	bytes, err = r.Post("http://localhost:9000/post")
	assert.Nil(t, err)
	assert.Equal(t, `{ "data": { "method": "post" }, "success": "true" }`, string(bytes))

	bytes, err = r.Put("http://localhost:9000/put")
	assert.Nil(t, err)
	assert.Equal(t, `{ "data": { "method": "put" }, "success": "true" }`, string(bytes))

	bytes, err = r.Delete("http://localhost:9000/delete")
	assert.Nil(t, err)
	assert.Equal(t, `{ "data": { "method": "delete" }, "success": "true" }`, string(bytes))

	bytes, err = r.Upload("http://localhost:9000/upload", "../../external/test.fakeFile")
	assert.Equal(t, "requester:call Could not open file: open ../../external/test.fakeFile: no such file or directory", err.Error())

	bytes, err = r.Upload("http://localhost:9000/upload", "../../external/test.txt")
	assert.Nil(t, err)
	assert.Equal(t, `{ "data": { "method": "upload" }, "success": "true" }`, string(bytes))

	mockServer.Stop()
}
