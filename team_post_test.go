package esa_test

import (
	"os"
	"testing"

	"github.com/k0kubun/pp"
	"github.com/yuichiro-h/go-esa"
)

func TestGetTeamPostsNoError(t *testing.T) {
	client := esa.New(&esa.Config{AccessToken: os.Getenv("ESA_ACCESS_TOKEN")})
	res, err := client.GetTeamPosts("docs", &esa.GetTeamPostsRequest{
		Q:       esa.String("API"),
		Include: esa.String("comments,stargazers"),
		PaginationRequest: esa.PaginationRequest{
			PerPage: esa.Int(1),
			Page:    esa.Int(1),
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(pp.Sprint(res.Posts[0]))
}

func TestGetTeamPostNoErro(t *testing.T) {
	client := esa.New(&esa.Config{AccessToken: os.Getenv("ESA_ACCESS_TOKEN")})
	res, err := client.GetTeamPost("docs", 25)
	if err != nil {
		t.Error(err)
	}
	t.Log(pp.Sprint(res.Post))
}
