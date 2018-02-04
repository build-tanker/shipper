package uploader

// Client - interfaces with the tanker web service
type Client interface {
	GetAccessKey() (string, error)
	DeleteAccessKey() error
	GetUploadURL() (string, error)
	UploadFile(string, string) error
}
type client struct{}

// NewClient - create a new client
func NewClient() Client {
	return &client{}
}

func (c *client) GetAccessKey() (string, error) {
	return "", nil
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
