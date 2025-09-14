package s2s

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	BaseURL      string
	ServiceToken string
}

func NewClient(baseURL, token string) *Client {
	return &Client{
		BaseURL:      baseURL,
		ServiceToken: token,
	}
}

func (c *Client) PostJSON(path string, data interface{}) (*http.Response, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s%s", c.BaseURL, path)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.ServiceToken)
	req.Header.Set("Content-Type", "application/json")

	return (&http.Client{}).Do(req)
}

func (c *Client) Get(path string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.ServiceToken)

	return (&http.Client{}).Do(req)
}
