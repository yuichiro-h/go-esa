package esa

import (
	"encoding/json"
	"fmt"
)

type Member struct {
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	Icon       string `json:"icon"`
	Email      string `json:"email"`
	PostsCount int    `json:"posts_count"`
}

type GetTeamMembersRequest struct {
	TeamName string
	PaginationRequest
}

type GetTeamMembersResponse struct {
	Members []Member `json:"members"`
	PaginationResponse
}

func (c *Client) GetTeamMembers(teamName string, req *GetTeamMembersRequest) (*GetTeamMembersResponse, error) {
	resp, body, errs := c.setPaginationParams(c.get(fmt.Sprintf("/v1/teams/%s/members", teamName)), &req.PaginationRequest).End()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, c.parseError(body)
	}

	var res GetTeamMembersResponse
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
