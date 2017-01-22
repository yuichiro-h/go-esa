package esa

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

const baseURL = "https://api.esa.io"

type Client struct {
	config *Config
}

type Config struct {
	AccessToken string
}

func New(config *Config) *Client {
	return &Client{
		config,
	}
}

func (c *Client) get(URI string) *gorequest.SuperAgent {
	return c.setAuthorizationHeader(gorequest.New().Get(baseURL + URI))
}

func (c *Client) post(URI string) *gorequest.SuperAgent {
	return c.setAuthorizationHeader(gorequest.New().Post(baseURL + URI))
}

func (c *Client) put(URI string) *gorequest.SuperAgent {
	return c.setAuthorizationHeader(gorequest.New().Put(baseURL + URI))
}

func (c *Client) delete(URI string) *gorequest.SuperAgent {
	return c.setAuthorizationHeader(gorequest.New().Delete(baseURL + URI))
}

func (c *Client) patch(URI string) *gorequest.SuperAgent {
	return c.setAuthorizationHeader(gorequest.New().Patch(baseURL + URI))
}

func (c *Client) setAuthorizationHeader(request *gorequest.SuperAgent) *gorequest.SuperAgent {
	return request.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.AccessToken))
}
