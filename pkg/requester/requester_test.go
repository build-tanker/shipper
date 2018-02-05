package requester

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRequester(t *testing.T) {
	r := NewRequester(5 * time.Second)

	bytes, err := r.Get("httpbin.org/ip")
	assert.Equal(t, `Could not complete request: Get httpbin.org/ip: unsupported protocol scheme ""`, err.Error())

	bytes, err = r.Get("http://httpbin.org/get")

	assert.Nil(t, err)
	assert.Equal(t, true, strings.Contains(string(bytes), "http://httpbin.org/get"))

}
