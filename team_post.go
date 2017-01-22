package esa

import (
	"encoding/json"
	"fmt"
	"time"
)

type PostMember struct {
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	Icon       string `json:"icon"`
}

type Post struct {
	Number          int        `json:"number"`
	Name            string     `json:"name"`
	FullName        string     `json:"full_name"`
	WIP             bool       `json:"wip"`
	BodyMD          string     `json:"body_md"`
	BodyHTML        string     `json:"body_html"`
	CreatedAt       time.Time  `json:"created_at"`
	Message         string     `json:"message"`
	URL             string     `json:"url"`
	UpdatedAt       time.Time  `json:"updated_at"`
	Tags            []string   `json:"tags"`
	Category        *string    `json:"category"`
	RevisionNumber  int        `json:"revision_number"`
	CreatedBy       PostMember `json:"created_by"`
	UpdatedBy       PostMember `json:"updated_by"`
	Kind            string     `json:"kind"`
	CommentsCount   int        `json:"comments_countr"`
	TaskCount       int        `json:"task_count"`
	DoneTasksCount  int        `json:"done_tasks_count"`
	StargazersCount int        `json:"stargazers_count"`
	WatchersCount   int        `json:"watchers_count"`
	Star            bool       `json:"star"`
	Watch           bool       `json:"watch"`
}

type GetTeamPostsRequest struct {
	Q       *string
	Include *string
	Sort    *string
	Order   *string
	PaginationRequest
}

type GetTeamPostsResponse struct {
	Posts []Post `json:"posts"`
	PaginationResponse
}

func (c *Client) GetTeamPosts(teamName string, req *GetTeamPostsRequest) (*GetTeamPostsResponse, error) {
	buildReq := c.get(fmt.Sprintf("/v1/teams/%s/posts", teamName))

	if req.Q != nil {
		buildReq = buildReq.Param("q", *req.Q)
	}
	if req.Include != nil {
		buildReq = buildReq.Param("include", *req.Include)
	}
	if req.Sort != nil {
		buildReq = buildReq.Param("sort", *req.Sort)
	}
	if req.Order != nil {
		buildReq = buildReq.Param("order", *req.Order)
	}

	resp, body, errs := c.setPaginationParams(buildReq, &req.PaginationRequest).End()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, c.parseError(body)
	}

	var res GetTeamPostsResponse
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
