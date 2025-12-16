package Api

import (
	"bufio"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
)

type ClientConfig struct {
	Domain   string
	PanelUrl string
	Username string
	Password string
}
type Client struct {
	Domain   string
	PanelUrl string
	Username string
	Password string
	Http     *http.Client
}

var ClientStorage []*Client

func LoadClientsFromFile(path string) ([]*Client, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var clients []*Client

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, "|")
		if len(parts) != 4 {
			continue
		}

		jar, _ := cookiejar.New(nil)
		client := &Client{
			Domain:   parts[0],
			PanelUrl: parts[0] + parts[1],
			Username: parts[2],
			Password: parts[3],
			Http: &http.Client{
				Jar: jar,
			},
		}

		clients = append(clients, client)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return clients, nil
}
