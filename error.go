package esa

import (
	"encoding/json"
	"fmt"
)

type EsaError struct {
	Code    string `json:"error"`
	Message string `json:"message"`
}

func (e EsaError) Error() string {
	return fmt.Sprintf("Esa API Error. error=%s, message=%s", e.Code, e.Message)
}

func (c *Client) parseError(response string) error {
	var esaError EsaError
	if err := json.Unmarshal([]byte(response), &esaError); err != nil {
		return err
	}

	return &esaError
}
