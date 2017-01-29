package esa

import "fmt"

type CreateTeamPostCommentStarRequest struct {
	Body *string `json:"body"`
}

func (c *Client) CreateTeamPostCommentStar(teamName string, commentID int, req *CreateTeamPostCommentStarRequest) error {
	resp, body, errs := c.post(fmt.Sprintf("/v1/teams/%s/comments/%d/star", teamName, commentID)).Send(req).End()
	if len(errs) > 0 {
		return errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return c.parseError(body)
	}

	return nil
}
