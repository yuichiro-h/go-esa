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

<<<<<<< HEAD
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
=======
func TestGetTeamPostNoErro(t *testing.T) {
	client := esa.New(&esa.Config{AccessToken: os.Getenv("ESA_ACCESS_TOKEN")})
	res, err := client.GetTeamPost("docs", 25)
	if err != nil {
		t.Error(err)
	}
	t.Log(pp.Sprint(res.Post))
}
>>>>>>> 42ddaf47e5ef749f8606e1167ed2c225f7a71827
