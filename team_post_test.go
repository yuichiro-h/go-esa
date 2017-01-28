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

// func TestCreateTeamPostNoError(t *testing.T) {
// 	client := esa.New(&esa.Config{AccessToken: os.Getenv("ESA_ACCESS_TOKEN")})
// 	res, err := client.CreateTeamPost("coto-coto", &esa.CreateTeamPostRequest{
// 		Name:     "go-esa test post name",
// 		BodyMD:   esa.String("# go-esa test post body"),
// 		Tags:     esa.SliceString([]string{"go-esa"}),
// 		Category: esa.String("go-esa"),
// 		WIP:      esa.Bool(true),
// 		Message:  esa.String("go-esa test message"),
// 		User:     esa.String("esa_bot"),
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Log(pp.Sprint(res))
// }
