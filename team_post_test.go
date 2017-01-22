package esa_test

import (
	"os"
	"testing"

	"github.com/CotoCoto/go-esa"
	"github.com/k0kubun/pp"
)

func TestGetTeamPostsNoError(t *testing.T) {
	client := esa.New(&esa.Config{AccessToken: os.Getenv("ESA_ACCESS_TOKEN")})
	res, err := client.GetTeamPosts("coto-coto", &esa.GetTeamPostsRequest{
		Q:       esa.String("日報"),
		Include: esa.String("comments,stargazers"),
		PaginationRequest: esa.PaginationRequest{
			PerPage: esa.Int(10),
			Page:    esa.Int(5),
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(pp.Sprint(res.Posts[0]))
}
