package main

import (
	"fmt"
	"log"

	ctfd "github.com/pandatix/go-ctfd/api"
)

func main() {
	url := "http://127.0.0.1:8080"
	apiKey := "8480991a05f6d0ef7bf34500c257676905f8c7ce516c5a86129d83214532b20c"

	client := ctfd.NewClient(url, "", "", apiKey)

	// Query all challenges
	challs, err := client.GetChallenges(&ctfd.GetChallengesParams{})
	if err != nil {
		log.Fatal(err)
	}
	for _, chall := range challs {
		fmt.Printf("[%d] %s\n", chall.ID, chall.Name)
	}
}
