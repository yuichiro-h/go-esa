package esa_test

import (
	"os"
	"testing"

	"github.com/k0kubun/pp"
	"github.com/yuichiro-h/go-esa"
)

func TestGetTeamPostsNoError(t *testing.T) {
	client := esa.New(&esa.Config{AccessToken: os.Getenv("ESA_ACCESS_TOKEN")})
	res, err := client.GetTeamPosts("coto-coto", &esa.GetTeamPostsRequest{
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

func TestGetTeamPostNoError(t *testing.T) {
	client := esa.New(&esa.Config{AccessToken: os.Getenv("ESA_ACCESS_TOKEN")})
	res, err := client.GetTeamPost("coto-coto", 2387)
	if err != nil {
		t.Error(err)
	}
	t.Log(pp.Sprint(res.Post))
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

func TestUpdateTeamPostNoError(t *testing.T) {
	client := esa.New(&esa.Config{AccessToken: os.Getenv("ESA_ACCESS_TOKEN")})
	res, err := client.UpdateTeamPost("coto-coto", 2387, &esa.UpdateTeamPostRequest{
		Name:      "updated go-esa test post name",
		WIP:       true,
		CreatedBy: "esa_bot",
		UpdatedBy: "esa_bot",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(pp.Sprint(res))
}
