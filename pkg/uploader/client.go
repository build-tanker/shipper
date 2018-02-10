package uploader

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/sudhanshuraheja/shipper/pkg/appcontext"
	"github.com/sudhanshuraheja/shipper/pkg/requester"
)

// Client - interface to talk to tanker service
type Client interface {
	GetAccessKey(server string) (string, error)
	DeleteAccessKey(server, accessKey string) error
	GetUploadURL(server, accessKey, bundle string) (string, error)
	UploadFile(string, string) error
}

type client struct {
	ctx *appcontext.AppContext
	r   requester.Requester
}

// NewClient - create a new client to talk to tanker service
func NewClient(ctx *appcontext.AppContext) Client {
	r := requester.NewRequester(2 * time.Second)
	return &client{
		ctx: ctx,
		r:   r,
	}
}

func (c *client) GetAccessKey(server string) (string, error) {
	url := fmt.Sprintf("%s/v1/shippers?name=%s&machineName=%s", server, uuid.NewV4(), "")

	bytes, err := c.r.Post(url)
	if err != nil {
		return "", errors.Wrap(err, "Could not handle post request")
	}

	var o ShipperAdd
	err = json.Unmarshal(bytes, &o)
	if err != nil {
		return "", errors.Wrap(err, "Could not unmarshall json")
	}

	return o.Data.AccessKey, nil
}

func (c *client) DeleteAccessKey(server, accessKey string) error {
	url := fmt.Sprintf("%s/v1/shippers/%s", server, accessKey)

	bytes, err := c.r.Delete(url)
	if err != nil {
		return errors.Wrap(err, "Could not handle post request")
	}

	var o ShipperDelete
	err = json.Unmarshal(bytes, &o)
	if err != nil {
		return errors.Wrap(err, "Could not unmarshall json")
	}

	if o.Success == "false" {
		return errors.New("Could not delete AccessKey from the server")
	}

	return nil
}

// /v1/builds?accessKey=a1b2c3&bundle=com.me.app
func (c *client) GetUploadURL(server, accessKey, bundle string) (string, error) {
	url := fmt.Sprintf("%s/v1/builds?accessKey=%s&bundle=%s", server, accessKey, bundle)

	bytes, err := c.r.Post(url)
	if err != nil {
		return "", errors.Wrap(err, "Could not handle post request")
	}

	var o BuildAdd
	err = json.Unmarshal(bytes, &o)
	if err != nil {
		return "", errors.Wrap(err, "Could not unmarshall json")
	}

	if o.Success == "false" {
		return "", errors.New("Could not get uploadURL from the server")
	}

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
