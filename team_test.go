package esa_test

import (
	"os"
	"testing"

	"github.com/CotoCoto/go-esa"
	"github.com/k0kubun/pp"
)

var client = esa.New(&esa.Config{AccessToken: os.Getenv("ESA_ACCESS_TOKEN")})

func TestGetTeamsNoError(t *testing.T) {
	res, err := client.GetTeams(&esa.GetTeamsRequest{})
	if err != nil {
		t.Error(err)
	}
	t.Log(pp.Sprint(res))
}

func TestGetTeamNoError(t *testing.T) {
	res, err := client.GetTeam("coto-coto")
	if err != nil {
		t.Error(err)
	}
	t.Log(pp.Sprint(res))
}
