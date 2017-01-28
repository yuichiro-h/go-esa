package esa

import (
	"encoding/json"
	"fmt"
)

type GetTeamCommentStargazersRequest struct {
	PaginationRequest
}
type GetTeamCommentStargazersResponse struct {
	Stargazers []Stargazer `json:"stargazers"`
	PaginationResponse
}

func (c *Client) GetTeamCommentStargazers(teamName string, commentID int, req *GetTeamCommentStargazersRequest) (*GetTeamCommentStargazersResponse, error) {
	buildReq := c.get(fmt.Sprintf("/v1/teams/%s/comments/%d/stargazers", teamName, commentID))

	resp, body, errs := c.setPaginationParams(buildReq, &req.PaginationRequest).End()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, c.parseError(body)
	}

	var res GetTeamCommentStargazersResponse
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
