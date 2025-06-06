package api_test

import (
	"crypto/rand"
	"encoding/hex"
	"testing"

	"github.com/ctfer-io/go-ctfd/api"
	"github.com/stretchr/testify/assert"
)

func Test_F_CTF(t *testing.T) {
	// Scenario:
	//
	// This test mocks a real CTF, first with a setup, a tutorial page,
	// then a player register and attempt then solve a challenge.
	// For the first blood the admin award the player, then extract the
	// statistics and pause the event.

	assert := assert.New(t)

	// 1a. Get nonce and session to mock a browser first
	nonce, session, err := api.GetNonceAndSession(CTFD_URL)
	if !assert.Nil(err, "got error: %s", err) {
		return
	}
	admin := api.NewClient(CTFD_URL, nonce, session, "")

	t.Cleanup(func() {
		// Due to relicas, forced to unpause the event elseway the test is not reproducible
		_ = admin.PatchConfigs(&api.PatchConfigsParams{
			Paused: ptr(false),
		})

		_ = admin.Reset(&api.ResetParams{
			Accounts:      ptr("y"),
			Submissions:   ptr("y"),
			Challenges:    ptr("y"),
			Pages:         ptr("y"),
			Notifications: ptr("y"),
		})
	})

	// 1b. Configure the CTF
	err = admin.Setup(&api.SetupParams{
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
	if !assert.Nil(err, "got error: %s", err) {
		return
	}

	// 1c. Create an API Key to avoid session/nonce+cookies dance
	token, err := admin.PostTokens(&api.PostTokensParams{
		Expiration:  "2222-01-01",
		Description: "Example API token.",
	})
	if !assert.Nil(err, "got error: %s", err) {
		return
	}
	admin.SetAPIKey(*token.Value)

	// 2. Add a page
	_, err = admin.PostPages(&api.PostPagesParams{
		AuthRequired: false,
		Content:      "# Test",
		Draft:        false,
		Format:       "markdown",
		Hidden:       false,
		Nonce:        nonce,
		Route:        "/tutorials/test",
		Title:        "Test",
	})
	if !assert.Nil(err, "got error: %s", err) {
		return
	}

	// 3. Add a challenge with a flag
	chall, err := admin.PostChallenges(&api.PostChallengesParams{
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
	assert.NotNil(chall)
	if !assert.Nil(err, "got error: %s", err) {
		return
	}
	flag, err := admin.PostFlags(&api.PostFlagsParams{
		Challenge: chall.ID,
		Content:   "24HIUT{IcmpExfiltrationIsEasy}",
		Data:      "case_sensitive",
		Type:      "static",
	})
	assert.NotNil(flag)
	if !assert.Nil(err, "got error: %s", err) {
		return
	}

	// 4. User register
	name := "ctfer-" + randHex()
	nonceUser, sessionUser, _ := api.GetNonceAndSession(CTFD_URL)
	user := api.NewClient(CTFD_URL, nonceUser, sessionUser, "")
	err = user.Register(&api.RegisterParams{
		Name:     name,
		Email:    name + "@example.com",
		Password: "password",
	})
	if !assert.Nil(err, "got error: %s", err) {
		return
	}
	usr, err := user.GetUsersMe()
	if !assert.Nil(err, "got error: %s", err) {
		return
	}

	// 5a. User failed attempt
	att, err := user.PostChallengesAttempt(&api.PostChallengesAttemptParams{
		ChallengeID: chall.ID,
		Submission:  "INVALID-FLAG",
	})
	if !assert.Nil(err, "got error: %s", err) {
		return
	}
	assert.Equal("incorrect", att.Status)

	// 5b. User successfull attempt
	att, err = user.PostChallengesAttempt(&api.PostChallengesAttemptParams{
		ChallengeID: chall.ID,
		Submission:  "24HIUT{IcmpExfiltrationIsEasy}",
	})
	if !assert.Nil(err, "got error: %s", err) {
		return
	}
	assert.Equal("correct", att.Status)

	// 6. Admin gives an award for first blood
	_, err = admin.PostAwards(&api.PostAwardsParams{
		Name:        "First Blood",
		Description: "First Blood for \"Stealing data\"",
		Category:    "first-blood",
		Icon:        "lightning",
		UserID:      usr.ID,
		Value:       50,
	})
	if !assert.Nil(err, "got error: %s", err) {
		return
	}

	// 7. Admin gets some statistics
	_, err = admin.GetStatisticsChallengesSolves()
	if !assert.Nil(err, "got error: %s", err) {
		return
	}
	_, err = admin.GetStatisticsUsers()
	if !assert.Nil(err, "got error: %s", err) {
		return
	}
	// ...

	// 8. Admin pause event
	err = admin.PatchConfigs(&api.PatchConfigsParams{
		Paused: ptr(true),
	})
	if !assert.Nil(err, "got error: %s", err) {
		return
	}

	// Time to open-source your challenges :)
}

func randHex() string {
	buf := make([]byte, 8)
	_, _ = rand.Read(buf)
	return hex.EncodeToString(buf)
}
