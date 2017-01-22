package esa

import "time"

type Stargazer struct {
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
	User      struct {
		Name       string `json:"name"`
		ScreenName string `json:"screen_name"`
		Icon       string `json:"icon"`
	} `json:"user"`
}
