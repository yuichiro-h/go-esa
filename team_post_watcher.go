package esa

import (
	"encoding/json"
	"fmt"
	"time"
)

type Watcher struct {
	CreatedAt time.Time `json:"created_at"`
	User      struct {
		Name       string `json:"name"`
		ScreenName string `json:"screen_name"`
		Icon       string `json:"icon"`
	} `json:"user"`
}

type GetTeamPostWatchersRequest struct {
	PaginationRequest
}

type GetTeamPostWatchersResponse struct {
	Watchers []Watcher `json:"watchers"`
	PaginationResponse
}

func (c *Client) GetTeamPostWatchers(teamName string, postNumber int, req *GetTeamPostWatchersRequest) (*GetTeamPostWatchersResponse, error) {
	r := c.get(fmt.Sprintf("/v1/teams/%s/posts/%d/watchers", teamName, postNumber))
	resp, body, errs := c.setPaginationParams(r, &req.PaginationRequest).End()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, c.parseError(body)
	}

	var res GetTeamPostWatchersResponse
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
