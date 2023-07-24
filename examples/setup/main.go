package main

import (
	"fmt"
	"log"

	ctfd "github.com/pandatix/go-ctfd/api"
)

func main() {
	url := "http://127.0.0.1:8080"

	// Note: add /setup so won't have to follow redirect ot work
	fmt.Println("[+] Getting initial nonce and session values")
	nonce, session, err := ctfd.GetNonceAndSession(url)
	if err != nil {
		log.Fatalf("Getting nonce and session: %s", err)
	}

	// Setup CTFd
	fmt.Println("[+] Setting up CTFd")
	client := ctfd.NewClient(url, session, nonce, "")
	if err := client.Setup(&ctfd.SetupParams{
		CTFName:        "24h IUT",
		CTFDescription: "24h IUT annual Cybersecurity CTF.",
		UserMode:       "teams",
		Name:           "PandatiX",
		Email:          "lucastesson@protonmail.com",
		Password:       "password",
		CTFLogo:        nil,
		CTFBanner:      nil,
		CTFSmallIcon:   nil,
		CTFTheme:       "core",
		ThemeColor:     "",
		Start:          "",
		End:            "",
		Nonce:          nonce,
	}); err != nil {
		log.Fatalf("Setting up CTFd: %s", err)
	}

	// Create API Key
	fmt.Println("[+] Creating API Token")
	token, err := client.PostTokens(&ctfd.PostTokensParams{
		Expiration: "2024-05-14",
	})
	if err != nil {
		log.Fatalf("Creating API token: %s", err)
	}
	client.SetAPIKey(*token.Value)

	// Add challenge
	fmt.Println("[+] Creating challenge")
	ch, err := client.PostChallenges(&ctfd.PostChallengesParams{
		Name:        "Break The License 1/2",
		Category:    "crypto",
		Description: "...",
		Initial:     ptr(500),
		Decay:       ptr(17),
		Minimum:     ptr(50),
		State:       "visible",
		Type:        "dynamic",
	})
	if err != nil {
		log.Fatalf("Creating challenge: %s", err)
	}
	fmt.Printf("Created challenge %d\n", ch.ID)
}

func ptr[T any](t T) *T {
	return &t
}
