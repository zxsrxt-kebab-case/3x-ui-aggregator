package Api

type IRequest interface {
	EndPoint() string
	SendRequest(client *Client) ([]byte, error)
}
