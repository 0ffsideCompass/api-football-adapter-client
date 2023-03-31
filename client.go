package client

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	url    string
	client *http.Client
}

var (
	fixtureEndpoint = "fixture/get"
)

func NewClient(url string) (*Client, error) {
	if url == "" {
		return nil, errors.New("url is empty")
	}
	return &Client{
		url:    url,
		client: &http.Client{},
	}, nil
}

func (c *Client) get(endpoint string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.url, endpoint)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return body, err
}

func (c *Client) Fixtures() ([]byte, error) {
	return c.get(fmt.Sprintf(fixtureEndpoint))
}
