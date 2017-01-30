package esa_test

import (
	"os"
	"testing"

	"github.com/k0kubun/pp"
	"github.com/yuichiro-h/go-esa"
)

func TestGetTeamPostWatchersNotError(t *testing.T) {
	client := esa.New(&esa.Config{AccessToken: os.Getenv("ESA_ACCESS_TOKEN")})
	res, err := client.GetTeamPostWatchers("coto-coto", 2387, &esa.GetTeamPostWatchersRequest{})
	if err != nil {
		t.Error(err)
	}
	t.Log(pp.Sprint(res))
}

func TestCreateTeamPostWatch(t *testing.T) {
	client := esa.New(&esa.Config{AccessToken: os.Getenv("ESA_ACCESS_TOKEN")})
	err := client.CreateTeamPostWatch("coto-coto", 2387)
	if err != nil {
		t.Error(err)
	}
}
