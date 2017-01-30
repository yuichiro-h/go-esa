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
	Number          int         `json:"number"`
	Name            string      `json:"name"`
	FullName        string      `json:"full_name"`
	WIP             bool        `json:"wip"`
	BodyMD          string      `json:"body_md"`
	BodyHTML        string      `json:"body_html"`
	CreatedAt       time.Time   `json:"created_at"`
	Message         string      `json:"message"`
	URL             string      `json:"url"`
	UpdatedAt       time.Time   `json:"updated_at"`
	Tags            []string    `json:"tags"`
	Category        *string     `json:"category"`
	RevisionNumber  int         `json:"revision_number"`
	CreatedBy       PostMember  `json:"created_by"`
	UpdatedBy       PostMember  `json:"updated_by"`
	Kind            string      `json:"kind"`
	CommentsCount   int         `json:"comments_countr"`
	TaskCount       int         `json:"task_count"`
	DoneTasksCount  int         `json:"done_tasks_count"`
	StargazersCount int         `json:"stargazers_count"`
	WatchersCount   int         `json:"watchers_count"`
	Star            bool        `json:"star"`
	Watch           bool        `json:"watch"`
	Comments        []Comment   `json:"comments"`
	Stargazers      []Stargazer `json:"stargazers"`
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

type GetTeamPostResponse struct {
	Post
}

func (c *Client) GetTeamPosts(teamName string, req *GetTeamPostsRequest) (*GetTeamPostsResponse, error) {
	r := c.get(fmt.Sprintf("/v1/teams/%s/posts", teamName))

	if req.Q != nil {
		r = r.Param("q", *req.Q)
	}
	if req.Include != nil {
		r = r.Param("include", *req.Include)
	}
	if req.Sort != nil {
		r = r.Param("sort", *req.Sort)
	}
	if req.Order != nil {
		r = r.Param("order", *req.Order)
	}

	resp, body, errs := c.setPaginationParams(r, &req.PaginationRequest).End()
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

func (c *Client) GetTeamPost(teamName string, number int) (*GetTeamPostResponse, error) {
	resp, body, errs := c.get(fmt.Sprintf("/v1/teams/%s/posts/%d", teamName, number)).End()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, c.parseError(body)
	}

	var res GetTeamPostResponse
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type CreateTeamPostRequest struct {
	Name     string    `json:"name"`
	BodyMD   *string   `json:"body_md"`
	Tags     *[]string `json:"tags"`
	Category *string   `json:"category"`
	WIP      *bool     `json:"wip"`
	Message  *string   `json:"message"`
	User     *string   `json:"user"`
}

type CreateTeamPostResponse struct {
	Post
}

func (c *Client) CreateTeamPost(teamName string, req *CreateTeamPostRequest) (*CreateTeamPostResponse, error) {
	resp, body, errs := c.post(fmt.Sprintf("/v1/teams/%s/posts", teamName)).
		Send(&struct {
			Post *CreateTeamPostRequest `json:"post"`
		}{
			Post: req,
		}).End()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, c.parseError(body)
	}

	var res CreateTeamPostResponse
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type UpdateTeamPostRequest struct {
	Name             string            `json:"name"`
	BodyMD           *string           `json:"body_md"`
	Tags             *[]string         `json:"tags"`
	Category         *string           `json:"category"`
	WIP              bool              `json:"wip"`
	Message          *string           `json:"message"`
	CreatedBy        string            `json:"created_by"`
	UpdatedBy        string            `json:"updated_by"`
	OriginalRevision *OriginalRevision `json:"original_revision"`
}

type OriginalRevision struct {
	BodyMD string `json:"body_md"`
	Number int    `json:"number"`
	User   string `json:"user"`
}

type UpdateTeamPostResponse struct {
	Post
	Overlapped *bool `json:"overlapped"`
}

func (c *Client) UpdateTeamPost(teamName string, postNumber int, req *UpdateTeamPostRequest) (*UpdateTeamPostResponse, error) {
	resp, body, errs := c.patch(fmt.Sprintf("/v1/teams/%s/posts/%d", teamName, postNumber)).
		Send(&struct {
			Post *UpdateTeamPostRequest `json:"post"`
		}{
			Post: req,
		}).End()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, c.parseError(body)
	}

	var res UpdateTeamPostResponse
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
