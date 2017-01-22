package esa_test

import (
	"os"
	"testing"

	"github.com/CotoCoto/go-esa"
	"github.com/k0kubun/pp"
)

func TestGetTeamMembersNoError(t *testing.T) {
	client := esa.New(&esa.Config{AccessToken: os.Getenv("ESA_ACCESS_TOKEN")})
	res, err := client.GetTeamMembers("coto-coto", &esa.GetTeamMembersRequest{})
	if err != nil {
		t.Error(err)
	}
	t.Log(pp.Sprint(res))
}
