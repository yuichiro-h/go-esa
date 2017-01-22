package esa

import "time"

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
