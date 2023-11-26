package integration_test

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"testing"

	"github.com/ctfer-io/go-ctfd/api"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/assert"
)

func Test_F_Setup(t *testing.T) {
	// Scenario:
	//
	// As a ChallOps, your job is to setup the CTF and add the challenges.
	// The fine-grained CTF's configuration is in the backlog of someone else,
	// so no need to worry: a minimal configuration is enough.
	// Given this task, you have to setup a challenge, add a file, set hints,
	// flags and topics.
	// For test purposes, you will need to solve the challenge, and wanting to
	// be the first in the scoreboard (for once), you will look at the scoreboard.
	// Once done, you wipe the challenge for later tests.

	cwd, _ := os.Getwd()
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Quick:       true,
		SkipRefresh: true,
		Dir:         path.Join(cwd, ".."),
		Config:      map[string]string{},
		ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
			assert := assert.New(t)

			// XXX Pulumi store/convert int values as float64
			port := int(stack.Outputs["port"].(float64))
			url := fmt.Sprintf("http://%s:%d", Base, port)

			// 1a. Get nonce and session to mock a browser first
			nonce, session, err := api.GetNonceAndSession(url)
			if !assert.Nil(err, "got error: %s", err) {
				return
			}
			client := api.NewClient(url, nonce, session, "")

			// 1b. Configure the CTF
			err = client.Setup(&api.SetupParams{
				CTFName:                "CTFer",
				CTFDescription:         "Ephemeral CTFd running for API tests purposes.",
				UserMode:               "users",
				Name:                   "ctfer",
				Email:                  "ctfer-io@protonmail.com",
				Password:               "password", // This is not real, don't bother trying x)
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
			})
			if !assert.Nil(err, "got error: %s", err) {
				return
			}

			// 1c. Create an API Key to avoid session/nonce+cookies dance
			token, err := client.PostTokens(&api.PostTokensParams{
				Expiration:  "2222-01-01",
				Description: "Example API token.",
			})
			if !assert.Nil(err, "got error: %s", err) {
				return
			}
			client.SetAPIKey(*token.Value)

			// 2. Create a challenge
			chall, err := client.PostChallenges(&api.PostChallengesParams{
				Name:        "Stealing data",
				Category:    "network",
				Description: "The network administrator just sent you the info that some strange packets where going out of a server.\nAt first glance, it is an internal one.\nCan you tell us what it is ?",
				Function:    "logarithmic",
				Initial:     500,
				Decay:       ptr(17),
				Minimum:     ptr(50),
				State:       "visible",
				Type:        "dynamic",
			})
			assert.NotNil(chall)
			if !assert.Nil(err, "got error: %s", err) {
				return
			}

			// 3. Push a file
			files, err := client.PostFiles(&api.PostFilesParams{
				File: &api.InputFile{
					Name:    "icmp.pcap",
					Content: []byte("bla bla bla CTFER{flag} bip boop"),
				},
				Challenge: chall.ID,
			})
			assert.NotEmpty(files)
			if !assert.Nil(err, "got error: %s", err) {
				return
			}

			// 4. Update the challenge, give it hints, flags and topics
			// XXX the strconv should not occur
			chall, err = client.PatchChallenge(strconv.Itoa(chall.ID), &api.PatchChallengeParams{
				Name:        chall.Name,
				Category:    chall.Category,
				Description: chall.Description,
				MaxAttempts: "3",
				Initial:     strconv.Itoa(*chall.Initial),
				Decay:       strconv.Itoa(*chall.Decay),
				Minimum:     strconv.Itoa(*chall.Minimum),
				State:       chall.State,
			})
			if !assert.Nil(err, "got error: %s", err) {
				return
			}
			assert.NotNil(chall)
			hint, err := client.PostHints(&api.PostHintsParams{
				ChallengeID: chall.ID,
				Content:     "C'mon dude...",
				Cost:        "50",
			})
			if !assert.Nil(err, "got error: %s", err) {
				return
			}
			_, err = client.PostHints(&api.PostHintsParams{
				ChallengeID: chall.ID,
				Content:     "Nop.",
				Cost:        "100",
				Requirements: api.Requirements{
					Prerequisites: []int{hint.ID},
				},
			})
			if !assert.Nil(err, "got error: %s", err) {
				return
			}
			_, err = client.PostFlags(&api.PostFlagsParams{
				Challenge: chall.ID,
				Content:   "CTFER{flag}",
				Type:      "static",
			})
			if !assert.Nil(err, "got error: %s", err) {
				return
			}
			_, err = client.PostTopics(&api.PostTopicsParams{
				Challenge: chall.ID,
				Type:      "challenge", // required as the resource can't be determined by CTFd
				Value:     "Inspection",
			})
			if !assert.Nil(err, "got error: %s", err) {
				return
			}

			// 5. Solve the challenge (but first fail)
			att1, err := client.PostChallengesAttempt(&api.PostChallengesAttemptParams{
				ChallengeID: chall.ID,
				Submission:  "CTFER{fla}",
			})
			if !assert.Nil(err, "got error: %s", err) {
				return
			}
			assert.Equal("incorrect", att1.Status)
			att2, err := client.PostChallengesAttempt(&api.PostChallengesAttemptParams{
				ChallengeID: chall.ID,
				Submission:  "CTFER{flag}",
			})
			if !assert.Nil(err, "got error: %s", err) {
				return
			}
			assert.Equal("correct", att2.Status)

			// 6. Get statistics
			stats, err := client.GetStatisticsChallengesSolves()
			if !assert.Nil(err, "got error: %s", err) {
				return
			}
			assert.NotEmpty(stats)

			// 7. Delete the challenge
			// XXX the strconv should not occur
			err = client.DeleteChallenge(strconv.Itoa(chall.ID))
			if !assert.Nil(err, "got error: %s", err) {
				return
			}

			// 8. Check no challenge remain
			challs, err := client.GetChallenges(nil)
			assert.Empty(challs)
			if !assert.Nil(err, "got error: %s", err) {
				return
			}
		},
	})
}

func Test_F_AdvancedSetup(t *testing.T) {
	// Scenario:
	//
	// As the Ops in chief, your job is to setup the whole CTF.
	// Given this task, you have to do a fine-grained configuration for a future
	// CTF, add a page, and send a notification to announce the end of your job.
	// This is part of a procedure that you are testing, so once your job is
	// completed you reset the instance.

	cwd, _ := os.Getwd()
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Quick:       true,
		SkipRefresh: true,
		Dir:         path.Join(cwd, ".."),
		Config:      map[string]string{},
		ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
			assert := assert.New(t)

			// XXX Pulumi store/convert int values as float64
			port := int(stack.Outputs["port"].(float64))
			url := fmt.Sprintf("http://%s:%d", Base, port)

			// 1a. Get nonce and session to mock a browser first
			nonce, session, err := api.GetNonceAndSession(url)
			if !assert.Nil(err, "got error: %s", err) {
				return
			}
			client := api.NewClient(url, nonce, session, "")

			// 1b. Configure the CTF
			err = client.Setup(&api.SetupParams{
				CTFName:                "CTFer",
				CTFDescription:         "Ephemeral CTFd running for API tests purposes.",
				UserMode:               "teams",
				Name:                   "ctfer",
				Email:                  "ctfer-io@protonmail.com",
				Password:               "password", // This is not real, don't bother trying x)
				ChallengeVisibility:    "admins",
				AccountVisibility:      "private",
				ScoreVisibility:        "hidden",
				RegistrationVisibility: "mlc",
				VerifyEmails:           false,
				TeamSize:               ptr(4),
				CTFLogo:                nil,
				CTFBanner:              nil,
				CTFSmallIcon:           nil,
				CTFTheme:               "core",
				ThemeColor:             "",
				Start:                  "",
				End:                    "",
				Nonce:                  nonce,
			})
			if !assert.Nil(err, "got error: %s", err) {
				return
			}

			// 1c. Create an API Key to avoid session/nonce+cookies dance
			token, err := client.PostTokens(&api.PostTokensParams{
				Expiration:  "2222-01-01",
				Description: "Example API token.",
			})
			if !assert.Nil(err, "got error: %s", err) {
				return
			}
			client.SetAPIKey(*token.Value)

			// 2. Fine-configuration
			err = client.PatchConfigs(&api.PatchConfigsParams{})
			if !assert.Nil(err, "got error: %s", err) {
				return
			}

			// 3. Add a page
			_, err = client.PostPages(&api.PostPagesParams{
				Title:   "Production",
				Route:   "/prod",
				Format:  "markdown",
				Content: "## Production\n\nThis CTFd is now configured, all the ChallMakers and ChallOps can work on it !\n",
			})
			if !assert.Nil(err, "got error: %s", err) {
				return
			}

			// 4. Send a notification
			_, err = client.PostNotifications(&api.PostNotificationsParams{
				Title:   "CTFd is ready to go !",
				Content: "After a lot of effort, and thanks to our procedure, the CTF is now up, running and ready-to-go :D\nEnjoy !",
				Sound:   true,
				Type:    "toast",
			})
			if !assert.Nil(err, "got error: %s", err) {
				return
			}

			// 5. Reset the instance
			// nonce, _, _ = api.GetNonceAndSession(url)
			err = client.Reset(&api.ResetParams{
				Accounts:      ptr("y"),
				Submissions:   ptr("y"),
				Challenges:    ptr("y"),
				Pages:         ptr("y"),
				Notifications: ptr("y"),
			})
			if !assert.Nil(err, "got error: %s", err) {
				return
			}
		},
	})
}
