package uploader

type Client interface{}
type client struct{}

func NewClient() Client {
	return &client{}
}
