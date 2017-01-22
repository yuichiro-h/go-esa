package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CotoCoto/go-esa"
)

func main() {
	client := esa.New(&esa.Config{AccessToken: os.Getenv("ESA_ACCESS_TOKEN")})
	res, err := client.GetTeams(&esa.GetTeamsRequest{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", res)
}
