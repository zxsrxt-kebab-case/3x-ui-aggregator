package Requests

import (
	"3x-ui-aggregator/Api"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
)

type LoginPostBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Obj     interface{} `json:"obj"`
}

type LoginRequest struct {
	LoginBody LoginPostBody
}

func (r *LoginRequest) EndPoint() string {
	return "/login"
}

func (r *LoginRequest) SendRequest(client *Api.Client) ([]byte, error) {

	u, err := url.Parse(client.PanelUrl)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, r.EndPoint())

	jsonBody, err := json.Marshal(r.LoginBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		u.String(),
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Http.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
