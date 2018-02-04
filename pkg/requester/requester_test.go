package requester

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequester(t *testing.T) {
	r := NewRequester()

	bytes, err := r.Get("httpbin.org/ip")
	assert.Equal(t, `Could not complete request: Get httpbin.org/ip: unsupported protocol scheme ""`, err.Error())

	bytes, err = r.Get("http://httpbin.org/ip")
	assert.Nil(t, err)
	assert.Equal(t, "{\n  \"origin\": \"140.0.251.125\"\n}\n", string(bytes))
}
