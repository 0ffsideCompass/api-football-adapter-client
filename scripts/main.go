package main

import (
	"fmt"

	client "github.com/0ffsideCompass/api-football-adapter-client"
)

func main() {

	client, err := client.New(
		"http://localhost:4343",
		"api-key",
	)

	if err != nil {
		panic(err)
	}

	league, err := client.AddLeague("39", "2018")
	if err != nil {
		panic(err)
	}

	fmt.Println(league)

	leagueData, err := client.GetLeague("39", "2018")
	if err != nil {
		panic(err)
	}

	fmt.Println(leagueData)
}
