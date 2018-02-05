package requester

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type Requester interface {
	Get(url string) ([]byte, error)
	Post(url string) ([]byte, error)
	Put(url string) ([]byte, error)
	Delete(url string) ([]byte, error)
}

type requester struct {
	c *http.Client
}

func NewRequester(timeout time.Duration) Requester {
	c := &http.Client{
		Timeout: timeout,
	}
	return &requester{
		c: c,
	}
}

func (r *requester) Get(url string) ([]byte, error) {
	return r.call(http.MethodGet, url)
}

func (r *requester) Post(url string) ([]byte, error) {
	return r.call(http.MethodPost, url)
}

func (r *requester) Put(url string) ([]byte, error) {
	return r.call(http.MethodPut, url)
}

func (r *requester) Delete(url string) ([]byte, error) {
	return r.call(http.MethodDelete, url)
}

func (r *requester) call(method string, url string) ([]byte, error) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Could not create request")
	}

	response, err := r.c.Do(request)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Could not complete request")
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, errors.Wrap(err, "Could not read response")
	}

	return bytes, nil
}
