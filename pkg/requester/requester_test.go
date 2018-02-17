package requester

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRequester(t *testing.T) {
	r := NewRequester(5 * time.Minute)

	// Fail with http / https
	bytes, err := r.Get("httpbin.org/ip")
	assert.Equal(t, `Could not complete request: Get httpbin.org/ip: unsupported protocol scheme ""`, err.Error())

	bytes, err = r.Get("http://httpbin.org/get")
	assert.Nil(t, err)
	assert.Equal(t, true, strings.Contains(string(bytes), "http://httpbin.org/get"))

	bytes, err = r.Upload("https://requestb.in/10w4sue1", "../../external/test.txt")
	assert.Nil(t, err)
	assert.Equal(t, "ok", string(bytes))

	bytes, err = r.Get("https://requestb.in/10w4sue1?inspect")
	assert.Equal(t, true, strings.Contains(string(bytes), "hello-this-is-a-file"))

}
