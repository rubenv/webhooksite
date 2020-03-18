// Simple library to test webhook calls, ideal for testing callbacks in a CI
// server that doesn't accept calls from the outside world.
//
// Features:
//
// - Create tokens, inspect results
//
// - Can be configured to use a self-hosted webhook.site instance
//
// - Zero dependencies!
package webhooksite

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// The default endpoint, but obviously you can use any self-hosted instance as
// well.
const Endpoint = "https://webhook.site"

type Client struct {
	endpoint string

	Client *http.Client
}

func New() *Client {
	return NewWithEndpoint(Endpoint)
}

func NewWithEndpoint(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
		Client:   &http.Client{},
	}
}

func (c *Client) url(p string) string {
	return fmt.Sprintf("%s/%s", strings.TrimSuffix(c.endpoint, "/"), p)
}

func (c *Client) CreateToken() (*Token, error) {
	return c.CreateTokenWithOptions(TokenOptions{})
}

func (c *Client) CreateTokenWithOptions(opts TokenOptions) (*Token, error) {
	in, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.url("token"), bytes.NewReader(in))
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("CreateTokenWithOptions: Unexpected status code: %d", resp.StatusCode)
	}

	r := &Token{}
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	r.URL = c.url(r.UUID)
	return r, nil
}

func (c *Client) GetRequests(id string) (*Requests, error) {
	req, err := http.NewRequest("GET", c.url(fmt.Sprintf("/token/%s/requests", id)), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GetRequests: Unexpected status code: %d", resp.StatusCode)
	}

	r := &Requests{}
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
