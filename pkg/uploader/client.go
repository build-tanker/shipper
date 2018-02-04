package uploader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"source.golabs.io/core/shipper/pkg/appcontext"
)

// Client - interface to talk to tanker service
type Client interface {
	GetAccessKey(server string) (string, error)
	DeleteAccessKey() error
	GetUploadURL() (string, error)
	UploadFile(string, string) error
}

type client struct {
	ctx *appcontext.AppContext
	h   *http.Client
}

// NewClient - create a new client to talk to tanker service
func NewClient(ctx *appcontext.AppContext) Client {
	h := &http.Client{
		Timeout: time.Millisecond * 100,
	}
	return &client{
		ctx: ctx,
		h:   h,
	}
}

func (c *client) GetAccessKey(server string) (string, error) {
	url := fmt.Sprintf("%s/v1/shippers?name=%s&machineName=%s", server, uuid.NewV4(), "")

	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return "", errors.Wrap(err, "Could not create request")
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := c.h.Do(request)
	if err != nil {
		return "", errors.Wrap(err, "Could not complete request")
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.Wrap(err, "Could not read response")
	}

	fmt.Println("bytes", string(bytes))

	var o ShipperAdd
	err = json.Unmarshal(bytes, &o)
	if err != nil {
		return "", errors.Wrap(err, "Could not unmarshall json")
	}

	return o.Data.AccessKey, nil
}

func (c *client) DeleteAccessKey() error {
	return nil
}

func (c *client) GetUploadURL() (string, error) {
	return "", nil
}

func (c *client) UploadFile(url string, file string) error {
	return nil
}

// 	toUpload, err := os.Open(file)
// 	if err != nil {
// 		return err
// 	}
// 	defer toUpload.Close()

// 	serverURL := fmt.Sprintf("%s?key=%s&bundle=%s&file=%s", config.UploadServer(), key, bundle, file)
// 	logger.Infof(serverURL)

// 	response, err := http.Post(serverURL, "binary/octet-stream", toUpload)
// 	if err != nil {
// 		return err
// 	}
// 	defer response.Body.Close()

// 	message, _ := ioutil.ReadAll(response.Body)
// 	logger.Infoln(string(message))

// 	return nil
// }
