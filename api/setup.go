package api

import (
	"fmt"
	"net/http"
	"strconv"
)

type SetupParams struct {
	CTFName                string
	CTFDescription         string
	UserMode               string
	ChallengeVisibility    string
	AccountVisibility      string
	ScoreVisibility        string
	RegistrationVisibility string
	VerifyEmails           bool
	TeamSize               *int
	Name                   string
	Email                  string
	Password               string
	CTFLogo                *InputFile
	CTFBanner              *InputFile
	CTFSmallIcon           *InputFile
	CTFTheme               string
	ThemeColor             string
	Start                  string
	End                    string
	Nonce                  string // XXX this should not be part of the API, but is required
}

// WARNING: this endpoint is not officially supported.
func (client *Client) Setup(params *SetupParams, opts ...Option) error {
	mp := map[string]any{
		"ctf_name":                params.CTFName,
		"ctf_description":         params.CTFDescription,
		"user_mode":               params.UserMode,
		"challenge_visibility":    params.ChallengeVisibility,
		"account_visibility":      params.AccountVisibility,
		"score_visibility":        params.ScoreVisibility,
		"registration_visibility": params.RegistrationVisibility,
		"verify_emails":           fmt.Sprintf("%t", params.VerifyEmails),
		"name":                    params.Name,
		"email":                   params.Email,
		"password":                params.Password,
		"ctf_logo":                params.CTFLogo,
		"ctf_banner":              params.CTFBanner,
		"ctf_smallicon":           params.CTFSmallIcon,
		"ctf_theme":               params.CTFTheme,
		"theme_color":             params.ThemeColor,
		"start":                   params.Start,
		"end":                     params.End,
		"_submit":                 "Finish",
		"nonce":                   params.Nonce,
	}
	if params.TeamSize != nil {
		mp["team_size"] = strconv.Itoa(*params.TeamSize)
	}
	b, ct, err := encodeMultipart(mp)
	if err != nil {
		return err
	}

	req, _ := http.NewRequest(http.MethodPost, "/setup", b)
	req.Header.Set("Content-Type", ct)

	// Enable redirection as it is necessary for this request
	oldCR := client.sub.CheckRedirect
	client.sub.CheckRedirect = nil
	defer func() {
		client.sub.CheckRedirect = oldCR
	}()

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("CTFd responded with status code %d", res.StatusCode)
	}

	// Update nonce, session is handled by the cookie jar
	nonce, err := getNonce(res.Body)
	if err != nil {
		return err
	}
	client.nonce = nonce
	return nil
}
