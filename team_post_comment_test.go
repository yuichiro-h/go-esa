package esa_test

import (
	"os"
	"testing"

	"github.com/k0kubun/pp"
	"github.com/yuichiro-h/go-esa"
)

func TestGetPostComments(t *testing.T) {
	client := esa.New(&esa.Config{AccessToken: os.Getenv("ESA_ACCESS_TOKEN")})
	res, err := client.GetTeamPostComments("docs", &esa.GetTeamPostCommentRequest{
		PostNumber: 2,
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
