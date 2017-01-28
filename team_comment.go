package esa

import (
	"encoding/json"
	"fmt"
)

type GetTeamCommentResponse struct {
	Comment
}

func (c *Client) GetTeamComment(teamName string, commentId int) (*GetTeamCommentResponse, error) {
	resp, body, errs := c.get(fmt.Sprintf("/v1/teams/%s/comments/%d", teamName, commentId)).End()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, c.parseError(body)
	}

	var res GetTeamCommentResponse
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
