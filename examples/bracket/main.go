package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ctfer-io/go-ctfd/api"
)

func main() {
	url := "http://localhost:8000"
	apiKey := "ctfd_xxx"

	// Connecting to existing CTFd
	fmt.Println("[+] Getting initial nonce and session values")
	nonce, session, err := api.GetNonceAndSession(url)
	if err != nil {
		log.Fatalf("Getting nonce and session: %s", err)
	}
	cli := api.NewClient(url, nonce, session, apiKey)

	// Create a user bracket
	fmt.Println("[+] Creating user bracket")
	bu, err := cli.PostBrackets(&api.PostBracketsParams{
		Name:        "Juniors",
		Description: "Brackets for 14-25 years old players.",
		Type:        "users",
	})
	if err != nil {
		log.Fatalf("Creating bracket: %s", err)
	}

	// Create a user
	fmt.Println("[+] Creating user")
	user, err := cli.PostUsers(&api.PostUsersParams{
		Name:      "player1",
		Email:     "player1@ctfer.io",
		Password:  "password",
		BracketID: ptr(strconv.Itoa(bu.ID)),
		Type:      "user",
		Fields:    []api.Field{},
	})
	if err != nil {
		log.Fatalf("Creating user: %s", err)
	}
	fmt.Printf("Created user %d\n", user.ID)

	// Create a team bracket
	fmt.Println("[+] Creating team bracket")
	bt, err := cli.PostBrackets(&api.PostBracketsParams{
		Name:        "Students",
		Description: "Brackets for students",
		Type:        "teams",
	})
	if err != nil {
		log.Fatalf("Creating bracket: %s", err)
	}

	// Create a team
	team, err := cli.PostTeams(&api.PostTeamsParams{
		Name:      "students",
		Email:     "students@ctfer.io",
		Password:  "password",
		BracketID: ptr(strconv.Itoa(bt.ID)),
		Fields:    []api.Field{},
	})
	if err != nil {
		log.Fatalf("Creating team: %s", err)
	}
	fmt.Printf("Created team %d\n", team.ID)

	if _, err := cli.PostTeamMembers(team.ID, &api.PostTeamsMembersParams{
		UserID: user.ID,
	}); err != nil {
		log.Fatalf("Adding user %d in team %d: %s", user.ID, team.ID, err)
	}
	fmt.Printf("Added user %d in team %d\n", user.ID, team.ID)
}

func ptr[T any](t T) *T {
	return &t
}
