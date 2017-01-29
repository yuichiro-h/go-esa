package esa_test

import (
	"os"
	"testing"

	"github.com/yuichiro-h/go-esa"
)

func TestCreateTeamPostCommentStar(t *testing.T) {
	client := esa.New(&esa.Config{AccessToken: os.Getenv("ESA_ACCESS_TOKEN")})
	err := client.CreateTeamPostCommentStar("coto-coto", 235531, &esa.CreateTeamPostCommentStarRequest{
		Body: esa.String("LGTM"),
	})
	if err != nil {
		t.Error(err)
	}
}
