package Requests

import (
	"3x-ui-aggregator/Api"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

type Resp struct {
	Obj struct {
		Settings string `json:"settings"`
	} `json:"obj"`
}

type Settings struct {
	Clients []struct {
		SubID string `json:"subId"`
	} `json:"clients"`
}

type GetInboundRequest struct {
	ID string
}

func (r *GetInboundRequest) EndPoint() string {
	return fmt.Sprintf("/panel/api/inbounds/get/%s", r.ID)
}

func (r *GetInboundRequest) SendRequest(client *Api.Client) ([]byte, error) {
	u, err := url.Parse(client.PanelUrl)
	if err != nil {
		return nil, err
	}
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

func (r *GetInboundRequest) GetSubID(jsonString []byte) (string, error) {
	resp := &Resp{}
	if err := json.Unmarshal([]byte(jsonString), &resp); err != nil {
		return "", err
	}
	settings := &Settings{}
	if err := json.Unmarshal([]byte(resp.Obj.Settings), &settings); err != nil {
		return "", err
	}

	if len(settings.Clients) == 0 {
		return "", errors.New("no clients found")
	}

	return settings.Clients[0].SubID, nil
}
