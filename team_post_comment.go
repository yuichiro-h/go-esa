package esa

import (
	"encoding/json"
	"fmt"
	"time"
)

type CommentMember struct {
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	Icon       string `json:"icon"`
}

type Comment struct {
	ID              int           `json:"id"`
	BodyMD          string        `json:"body_md"`
	BodyHTML        string        `json:"body_html"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	URL             string        `json:"url"`
	CreatedBy       CommentMember `json:"created_by"`
	StargazersCount int           `json:"stargazers_count"`
	Star            bool          `json:"star"`
}

type GetTeamPostCommentsResponse struct {
	Comments []Comment `json:"comments"`
	PaginationResponse
}

type GetTeamPostCommentResponse struct {
	Comment
}

type GetTeamPostCommentRequest struct {
	PostNumber int
	PaginationRequest
}

func (c *Client) GetTeamPostComments(teamName string, req *GetTeamPostCommentRequest) (*GetTeamPostCommentsResponse, error) {
	buildReq := c.get(fmt.Sprintf("/v1/teams/%s/posts/%d/comments", teamName, req.PostNumber))

	resp, body, errs := c.setPaginationParams(buildReq, &req.PaginationRequest).End()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, c.parseError(body)
	}
	var res GetTeamPostCommentsResponse
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
