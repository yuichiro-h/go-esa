package esa

import (
	"encoding/json"
	"fmt"
)

type TeamStats struct {
	Members            int `json:"members"`
	Posts              int `json:"posts"`
	PostsWIP           int `json:"posts_wip"`
	PostsShipped       int `json:"posts_shipped"`
	Comments           int `json:"comments"`
	Stars              int `json:"stars"`
	DailyActiveUsers   int `json:"daily_active_users"`
	WeeklyActiveusers  int `json:"weekly_active_users"`
	MonthlyActiveUsers int `json:"monthly_active_users"`
}

type GetTeamStatsResponse struct {
	TeamStats
}

func (c *Client) GetTeamStats(teamName string) (*GetTeamStatsResponse, error) {
	resp, body, errs := c.get(fmt.Sprintf("/v1/teams/%s/stats", teamName)).End()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, c.parseError(body)
	}

	var res GetTeamStatsResponse
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
