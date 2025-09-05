package api_test

import (
	"strconv"
	"testing"

	"github.com/ctfer-io/go-ctfd/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	// 1a. Get nonce and session to mock a browser first
	nonce, session, err := api.GetNonceAndSession(CTFD_URL)
	require.NoError(t, err)
	client := api.NewClient(CTFD_URL, nonce, session, "")

	t.Cleanup(func() {
		_ = client.Reset(&api.ResetParams{
			Accounts:      ptr("y"),
			Submissions:   ptr("y"),
			Challenges:    ptr("y"),
			Pages:         ptr("y"),
			Notifications: ptr("y"),
		})
	})

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
	})
	require.NoError(t, err)

	// 1c. Create an API Key to avoid session/nonce+cookies dance
	token, err := client.PostTokens(&api.PostTokensParams{
		Expiration:  "2222-01-01",
		Description: "Example API token.",
	})
	require.NoError(t, err)
	client.SetAPIKey(*token.Value)

	// 1d. Logout because we don't know what could happen with a mouse on the UI
	err = client.Logout()
	require.NoError(t, err)

	// 1e. Relog, and drop the use of the API token because... why not ¯\_(ツ)_/¯
	err = client.Login(&api.LoginParams{
		Name:     "ctfer",
		Password: "password",
	})
	require.NoError(t, err)
	client.SetAPIKey("")

	// 2. Create a challenge
	chall, err := client.PostChallenges(&api.PostChallengesParams{
		Name:           "Stealing data",
		Category:       "network",
		Description:    "The network administrator just sent you the info that some strange packets where going out of a server.\nAt first glance, it is an internal one.\nCan you tell us what it is ?",
		Function:       ptr("logarithmic"),
		ConnectionInfo: ptr("ssh -l pandatix@master.pandatix.dev"),
		MaxAttempts:    ptr(3),
		Initial:        ptr(500),
		Decay:          ptr(17),
		Minimum:        ptr(50),
		State:          "visible",
		Type:           "dynamic",
	})
	assert.NotNil(t, chall)
	require.NoError(t, err)

	// 3. Push a file
	files, err := client.PostFiles(&api.PostFilesParams{
		Files: []*api.InputFile{
			{
				Name:    "icmp.pcap",
				Content: []byte("bla bla bla CTFER{flag} bip boop"),
			},
		},
		Challenge: &chall.ID,
	})
	assert.NotEmpty(t, files)
	require.NoError(t, err)

	// Check it has been properly pushed
	c, err := client.GetFileContent(files[0])
	assert.NotEmpty(t, c)
	require.NoError(t, err)

	// 4. Update the challenge, give it hints, flags and topics
	chall, err = client.PatchChallenge(chall.ID, &api.PatchChallengeParams{
		Name:        chall.Name,
		Category:    chall.Category,
		Description: chall.Description,
		Function:    chall.Function,
		MaxAttempts: ptr(3),
		Initial:     chall.Initial,
		Decay:       chall.Decay,
		Minimum:     chall.Minimum,
		State:       chall.State,
	})
	require.NotNil(t, chall)
	require.NoError(t, err)

	hint, err := client.PostHints(&api.PostHintsParams{
		ChallengeID: chall.ID,
		Content:     "C'mon dude...",
		Cost:        50,
	})
	require.NoError(t, err)

	_, err = client.PostHints(&api.PostHintsParams{
		ChallengeID: chall.ID,
		Content:     "Nop.",
		Cost:        100,
		Requirements: api.Requirements{
			Prerequisites: []int{hint.ID},
		},
	})
	require.NoError(t, err)

	_, err = client.PostFlags(&api.PostFlagsParams{
		Challenge: chall.ID,
		Content:   "CTFER{flag}",
		Type:      "static",
	})
	require.NoError(t, err)

	topic, err := client.PostTopics(&api.PostTopicsParams{
		Challenge: chall.ID,
		Type:      "challenge", // required as the resource can't be determined by CTFd
		Value:     "Inspection",
	})
	require.NoError(t, err)

	// 5. Solve the challenge (but first fail)
	att1, err := client.PostChallengesAttempt(&api.PostChallengesAttemptParams{
		ChallengeID: chall.ID,
		Submission:  "CTFER{fla}",
	})
	assert.Equal(t, "incorrect", att1.Status)
	require.NoError(t, err)

	att2, err := client.PostChallengesAttempt(&api.PostChallengesAttemptParams{
		ChallengeID: chall.ID,
		Submission:  "CTFER{flag}",
	})
	assert.Equal(t, "correct", att2.Status)
	require.NoError(t, err)

	// 6. Get statistics
	stats, err := client.GetStatisticsChallengesSolves()
	assert.NotEmpty(t, stats)
	require.NoError(t, err)

	// 7. Delete the challenge
	// XXX the strconv should not occur
	err = client.DeleteTopic(&api.DeleteTopicArgs{
		ID:   strconv.Itoa(topic.ID),
		Type: "challenge",
	})
	require.NoError(t, err)

	err = client.DeleteChallenge(chall.ID)
	require.NoError(t, err)

	// 8. Check no challenge remain
	challs, err := client.GetChallenges(nil)
	assert.Empty(t, challs)
	require.NoError(t, err)
}

func Test_F_AdvancedSetup(t *testing.T) {
	// Scenario:
	//
	// As the Ops in chief, your job is to setup the whole CTF.
	// Given this task, you have to do a fine-grained configuration for a future
	// CTF, add a page, and send a notification to announce the end of your job.
	// This is part of a procedure that you are testing, so once your job is
	// completed you reset the instance.

	// 1a. Get nonce and session to mock a browser first
	nonce, session, err := api.GetNonceAndSession(CTFD_URL)
	require.NoError(t, err)
	client := api.NewClient(CTFD_URL, nonce, session, "")

	t.Cleanup(func() {
		_ = client.Reset(&api.ResetParams{
			Accounts:      ptr("y"),
			Submissions:   ptr("y"),
			Challenges:    ptr("y"),
			Pages:         ptr("y"),
			Notifications: ptr("y"),
		})
	})

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
	})
	require.NoError(t, err)

	// 1c. Create an API Key to avoid session/nonce+cookies dance
	token, err := client.PostTokens(&api.PostTokensParams{
		Expiration:  "2222-01-01",
		Description: "Example API token.",
	})
	require.NoError(t, err)
	client.SetAPIKey(*token.Value)

	// 2. Fine-configuration
	err = client.PatchConfigs(&api.PatchConfigsParams{})
	require.NoError(t, err)

	// 3. Add a page
	_, err = client.PostPages(&api.PostPagesParams{
		Title:   "Production",
		Route:   "/prod",
		Format:  "markdown",
		Content: "## Production\n\nThis CTFd is now configured, all the ChallMakers and ChallOps can work on it !\n",
	})
	require.NoError(t, err)

	// 4. Send a notification
	_, err = client.PostNotifications(&api.PostNotificationsParams{
		Title:   "CTFd is ready to go !",
		Content: "After a lot of effort, and thanks to our procedure, the CTF is now up, running and ready-to-go :D\nEnjoy !",
		Sound:   true,
		Type:    "toast",
	})
	require.NoError(t, err)

	// 5. Reset the instance
	err = client.Reset(&api.ResetParams{
		Accounts:      ptr("y"),
		Submissions:   ptr("y"),
		Challenges:    ptr("y"),
		Pages:         ptr("y"),
		Notifications: ptr("y"),
	})
	require.NoError(t, err)
}

func Test_F_UsersAndTeams(t *testing.T) {
	// Scenario:
	//
	// As an Ops, your job is to import all the registered users and teams
	// before the event such that at the very beginning you are sure no one
	// is lost.

	// 1a. Get nonce and session to mock a browser first
	nonce, session, err := api.GetNonceAndSession(CTFD_URL)
	require.NoError(t, err)
	client := api.NewClient(CTFD_URL, nonce, session, "")

	t.Cleanup(func() {
		_ = client.Reset(&api.ResetParams{
			Accounts:      ptr("y"),
			Submissions:   ptr("y"),
			Challenges:    ptr("y"),
			Pages:         ptr("y"),
			Notifications: ptr("y"),
		})
	})

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
	})
	require.NoError(t, err)

	// 1c. Create an API Key to avoid session/nonce+cookies dance
	token, err := client.PostTokens(&api.PostTokensParams{
		Expiration:  "2222-01-01",
		Description: "Example API token.",
	})
	require.NoError(t, err)
	client.SetAPIKey(*token.Value)

	// Define all users and teams
	type User struct {
		name, email, password string
	}
	type Team struct {
		name, email, password string
		users                 []User
	}
	var teams = []Team{
		{
			name:     "MILF CTF Team",
			email:    "milfctf@example.com",
			password: "password",
			users: []User{
				{
					name:     "hashp4",
					email:    "hashp4@example.com",
					password: "password",
				},
				// ...
			},
		},
	}

	// 2. Create all the users and their teams
	for _, team := range teams {
		// 2a. Create team
		tm, err := client.PostTeams(&api.PostTeamsParams{
			Name:     team.name,
			Email:    team.email,
			Password: team.password,
			Banned:   false,
			Hidden:   false,
			Fields:   []api.Field{},
		})
		require.NoError(t, err)

		for _, user := range team.users {
			// 2b. Create user
			usr, err := client.PostUsers(&api.PostUsersParams{
				Name:     user.name,
				Email:    user.email,
				Password: user.password,
				Type:     "user",
				Verified: false,
				Hidden:   false,
				Banned:   false,
				Fields:   []api.Field{},
			})
			require.NoError(t, err)

			// 2c. Join user to team
			_, err = client.PostTeamMembers(tm.ID, &api.PostTeamsMembersParams{
				UserID: usr.ID,
			})
			require.NoError(t, err)
		}
	}
}
