package Requests

import (
	"3x-ui-aggregator/Api"
	"io"
	"net/http"
	"net/url"
	"path"
)

type GetSubscriptionRequest struct {
	SubscriptionId string
}

func (r GetSubscriptionRequest) EndPoint() string {
	return path.Join("detour", r.SubscriptionId)
}

func (r GetSubscriptionRequest) SendRequest(client *Api.Client) ([]byte, error) {
	u, err := url.Parse(client.Domain)
	if err != nil {
		return nil, err
	}
	u.Host = u.Hostname() + ":2096"
	u.Path = path.Join(u.Path, r.EndPoint())

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
