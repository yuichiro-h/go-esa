package esa_test

import (
	"os"
	"testing"

	"github.com/k0kubun/pp"
	"github.com/yuichiro-h/go-esa"
)

func TestGetPostComments(t *testing.T) {
	client := esa.New(&esa.Config{AccessToken: os.Getenv("ESA_ACCESS_TOKEN")})
	res, err := client.GetTeamPostComments("coto-coto", &esa.GetTeamPostCommentRequest{
		PostNumber: 7,
		PaginationRequest: esa.PaginationRequest{
			PerPage: esa.Int(1),
			Page:    esa.Int(1),
		},
	})

	if err != nil {
		t.Error(err)
	}
	t.Log(pp.Sprint(res))
}

// func TestCreateTeamPostCommentNoError(t *testing.T) {
// 	client := esa.New(&esa.Config{AccessToken: os.Getenv("ESA_ACCESS_TOKEN")})
// 	res, err := client.CreateTeamPostComment("coto-coto", 2387, &esa.CreateTeamPostCommentRequest{
// 		BodyMD: "LGTM!!",
// 		User:   esa.String("esa_bot"),
// 	})
//
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Log(pp.Sprint(res))
// }
