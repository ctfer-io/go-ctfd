package main

import (
	"fmt"
	"log"

	ctfd "github.com/pandatix/go-ctfd/api"
)

func main() {
	url := "http://127.0.0.1:8080"
	apiKey := "bcdef68b05cb834cdb5cd81e19ded1c3f5507634e43757689e35f3d43e5c38dc"

	client := ctfd.NewClient(url, "", "", apiKey)

	// Add challenge
	fmt.Println("[+] Creating challenge")
	ch, err := client.PostChallenges(&ctfd.PostChallengesParams{
		Name:        "Stealing data",
		Category:    "netwokr",
		Description: "...",
		Initial:     ptr(500),
		Decay:       ptr(20),
		Minimum:     ptr(50),
		State:       "visible",
		Type:        "dynamic",
	})
	if err != nil {
		log.Fatalf("Creating challenge: %s", err)
	}
	fmt.Printf("Created challenge %d\n", ch.ID)

	// Update challenge
	fmt.Println("[~] Updating challenge")
	ch, err = client.PatchChallenge(ch.ID, &ctfd.PatchChallengeParams{
		Category: ptr("network"),
		Decay:    ptr(17),
	})
	if err != nil {
		log.Fatalf("Updating challenge: %s", err)
	}
	fmt.Printf("Updated challenge %d\n", ch.ID)
}

func ptr[T any](t T) *T {
	return &t
}
