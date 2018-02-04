package uploader

type Client interface {
	GetAccessKey() (string, error)
	DeleteAccessKey() error
}
type client struct{}

func NewClient() Client {
	return &client{}
}

func (c *client) GetAccessKey() (string, error) {
	return "", nil
}

func (c *client) DeleteAccessKey() error {
	return nil
}
