package esa

import (
	"strconv"

	"github.com/parnurzeal/gorequest"
)

type PaginationResponse struct {
	PrevPage   *int `json:"prev_page"`
	NextPage   *int `json:"next_page"`
	TotalCount int  `json:"total_count"`
	Page       int  `json:"page"`
	PerPage    int  `json:"per_page"`
	MaxPerPage int  `json:"max_per_page"`
}

type PaginationRequest struct {
	Page    *int `json:"page"`
	PerPage *int `json:"per_page"`
}

func (c *Client) setPaginationParams(request *gorequest.SuperAgent, pagination *PaginationRequest) *gorequest.SuperAgent {
	req := request
	if pagination == nil {
		return req
	}

	if pagination.Page != nil {
		req = req.Param("page", strconv.Itoa(*pagination.Page))
	}
	if pagination.PerPage != nil {
		req = req.Param("per_page", strconv.Itoa(*pagination.PerPage))
	}

	return req
}
