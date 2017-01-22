package esa

import (
	"encoding/json"
	"fmt"
)

type Pagination struct {
	PrevPage   *int `json:"prev_page"`
	NextPage   *int `json:"next_page"`
	TotalCount int  `json:"total_count"`
	Page       int  `json:"page"`
	PerPage    int  `json:"per_page"`
	MaxPerPage int  `json:"max_per_page"`
}

type Team struct {
	Name        string `json:"name"`
	Privacy     string `json:"privacy"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	URL         string `json:"url"`
}

type GetTeamsRequest struct {
	PaginationRequest
}

type GetTeamsResponse struct {
	Teams []Team `json:"teams"`
	PaginationResponse
}

func (c *Client) GetTeams(req *GetTeamsRequest) (*GetTeamsResponse, error) {
	resp, body, errs := c.setPaginationParams(c.get("/v1/teams"), &req.PaginationRequest).End()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, c.parseError(body)
	}

	var res GetTeamsResponse
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type GetTeamResponse struct {
	Team
}

func (c *Client) GetTeam(teamName string) (*GetTeamResponse, error) {
	resp, body, errs := c.get(fmt.Sprintf("/v1/teams/%s", teamName)).End()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, c.parseError(body)
	}

	var res GetTeamResponse
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	return &res, nil

}
