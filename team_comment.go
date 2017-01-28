package esa

import "time"

type Comment struct {
	Id        int           `json:"id"`
	BodyMd    string        `json:"body_md"`
	URL       string        `json:"url"`
	CreatedBy CommentMember `json:"created_by"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
