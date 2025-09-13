package s2s

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

func (c *Client) PostJSON(path string, data interface{}, tenantID, userID string, roles []string) (*http.Response, error) {
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
	req.Header.Set("X-Internal-Tenant-ID", tenantID)
	req.Header.Set("X-Internal-User-ID", userID)
	req.Header.Set("X-Internal-Roles", strings.Join(roles, ","))

	req.Header.Set("Content-Type", "application/json")

	return (&http.Client{}).Do(req)
}
