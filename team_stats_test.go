package esa_test

import (
	"os"
	"testing"

	"github.com/CotoCoto/go-esa"
	"github.com/k0kubun/pp"
)

var client = esa.New(&esa.Config{AccessToken: os.Getenv("ESA_ACCESS_TOKEN")})

func TestGetTeamStatsNoError(t *testing.T) {
	res, err := client.GetTeamStats("coto-coto")
	if err != nil {
		t.Error(err)
	}
	t.Log(pp.Sprint(res))
}
