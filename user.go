package esa

import (
	"encoding/json"
	"time"
)

type User struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	ScreenName string    `json:"screen_name"`
	Icon       string    `json:"icon"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GetUserResponse struct {
	User
}

func (c *Client) GetUser() (*GetUserResponse, error) {
	resp, body, errs := c.get("/v1/user").End()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, c.parseError(body)
	}

	var res GetUserResponse
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
