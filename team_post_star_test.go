package esa_test

import (
	"os"
	"testing"

	"github.com/k0kubun/pp"
	"github.com/yuichiro-h/go-esa"
)

func TestGetTeamPostStartgazars(t *testing.T) {
	client := esa.New(&esa.Config{AccessToken: os.Getenv("ESA_ACCESS_TOKEN")})
	res, err := client.GetTeamPostStarGazers("coto-coto", 2387, &esa.GetTeamPostStargazersRequest{})
	if err != nil {
		t.Error(err)
	}
	t.Log(pp.Sprint(res))
}

func TestCreateTeamPostStar(t *testing.T) {
	client := esa.New(&esa.Config{AccessToken: os.Getenv("ESA_ACCESS_TOKEN")})
	err := client.CreateTeamPostStar("coto-coto", 2387, &esa.CreateTeamPostStartRequest{
		Body: esa.String("go-esa"),
	})
	if err != nil {
		t.Error(err)
	}
}
