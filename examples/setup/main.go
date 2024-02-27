package main

import (
	"fmt"
	"log"

	ctfd "github.com/ctfer-io/go-ctfd/api"
)

func main() {
	url := "http://192.168.49.2:32052"

	// Note: add /setup so won't have to follow redirect ot work
	fmt.Println("[+] Getting initial nonce and session values")
	nonce, session, err := ctfd.GetNonceAndSession(url)
	if err != nil {
		log.Fatalf("Getting nonce and session: %s", err)
	}

	// Setup CTFd
	fmt.Println("[+] Setting up CTFd")
	client := ctfd.NewClient(url, nonce, session, "")
	if err := client.Setup(&ctfd.SetupParams{
		CTFName:                "24h IUT",
		CTFDescription:         "24h IUT annual Cybersecurity CTF.",
		UserMode:               "users",
		Name:                   "PandatiX",
		Email:                  "lucastesson@protonmail.com",
		Password:               "password",
		ChallengeVisibility:    "public",
		AccountVisibility:      "public",
		ScoreVisibility:        "public",
		RegistrationVisibility: "public",
		VerifyEmails:           false,
		TeamSize:               nil,
		CTFLogo:                nil,
		CTFBanner:              nil,
		CTFSmallIcon:           nil,
		CTFTheme:               "core",
		ThemeColor:             "",
		Start:                  "",
		End:                    "",
		Nonce:                  nonce,
	}); err != nil {
		log.Fatalf("Setting up CTFd: %s", err)
	}

	// Create API Key
	fmt.Println("[+] Creating API Token")
	token, err := client.PostTokens(&ctfd.PostTokensParams{
		Expiration:  "2024-05-14",
		Description: "Example API token.",
	})
	if err != nil {
		log.Fatalf("Creating API token: %s", err)
	}
	client.SetAPIKey(*token.Value)

	// Add challenge
	fmt.Println("[+] Creating challenge")
	ch, err := client.PostChallenges(&ctfd.PostChallengesParams{
		Name:           "Break The License 1/2",
		Category:       "crypto",
		Description:    "...",
		Function:       "logarithmic",
		ConnectionInfo: ptr("ssh -l user@crypto1.ctfer.io"),
		MaxAttempts:    ptr(3),
		Initial:        ptr(500),
		Decay:          ptr(17),
		Minimum:        ptr(50),
		State:          "visible",
		Type:           "dynamic",
	})
	if err != nil {
		log.Fatalf("Creating challenge: %s", err)
	}
	fmt.Printf("    Created challenge %d\n", ch.ID)

	// Add files to it
	files, err := client.PostFiles(&ctfd.PostFilesParams{
		Files: []*ctfd.InputFile{
			{
				Name:    "file1",
				Content: []byte(`toto 1`),
			}, {
				Name:    "file2",
				Content: []byte(`toto 2`),
			},
		},
		Challenge: ch.ID,
	})
	if err != nil {
		log.Fatalf("Creating files: %s", err)
	}
	fmt.Printf("    Created %d files\n", len(files))

	// Add a flag to solve it
	fmt.Println("[~] Updating challenge")
	f, err := client.PostFlags(&ctfd.PostFlagsParams{
		Challenge: ch.ID,
		Content:   "24HIUT{content}",
		Type:      "static",
	})
	if err != nil {
		log.Fatalf("Creating flag: %s", err)
	}

	// Solve it
	fmt.Println("[+] Creating attempt")
	att, err := client.PostChallengesAttempt(&ctfd.PostChallengesAttemptParams{
		ChallengeID: ch.ID,
		Submission:  f.Content,
	})
	if err != nil {
		log.Fatalf("Creating attempt: %s", err)
	}
	fmt.Printf("    Result: %s\n", att.Status)

	// Make it loop on itself, deadlock :imp:
	fmt.Println("[~] Making the challenge require itself...")
	ch, err = client.PatchChallenge(ch.ID, &ctfd.PatchChallengeParams{
		Name:           ch.Name,
		Category:       ch.Category,
		Description:    ch.Description,
		Function:       ch.Function,
		ConnectionInfo: ch.ConnectionInfo,
		Initial:        ch.Initial,
		Decay:          ch.Decay,
		Minimum:        ch.Minimum,
		MaxAttempts:    ch.MaxAttempts,
		State:          ch.State,
		Requirements: &ctfd.Requirements{
			Anonymize:     ptr(false),
			Prerequisites: []int{ch.ID},
		},
	})
	if err != nil {
		log.Fatalf("   Failed: %s", err)
	}
	ch.Requirements, err = client.GetChallengeRequirements(ch.ID)
	if err != nil {
		log.Fatalf("    Failed: %s", err)
	}

	fmt.Printf("ch: %+v\n", ch)
	fmt.Printf("ch.Requirements: %+v\n", ch.Requirements)
}

func ptr[T any](t T) *T {
	return &t
}
