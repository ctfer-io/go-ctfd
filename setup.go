package goctfd

import (
	"fmt"
	"net/http"
)

type SetupParams struct {
	CTFName        string
	CTFDescription string
	UserMode       string
	Name           string
	Email          string
	Password       string
	CTFLogo        *InputFile
	CTFBanner      *InputFile
	CTFSmallIcon   *InputFile
	CTFTheme       string
	ThemeColor     string
	Start          string
	End            string
	Nonce          string // XXX this should not be part of the API
}

// WARNING 1: this endpoint is not officially supported.
// WARNING 2: provided client must have a no-follow-redirect behaviour,
// or a cookie jar. Else, it won't detect the setup worked properly.
func (client *Client) Setup(params *SetupParams, opts ...Option) error {
	b, ct, err := encodeMultipart(map[string]any{
		"ctf_name":        params.CTFName,
		"ctf_description": params.CTFDescription,
		"user_mode":       params.UserMode,
		"name":            params.Name,
		"email":           params.Email,
		"password":        params.Password,
		"ctf_logo":        params.CTFLogo,
		"ctf_banner":      params.CTFBanner,
		"ctf_smallicon":   params.CTFSmallIcon,
		"ctf_theme":       params.CTFTheme,
		"theme_color":     params.ThemeColor,
		"start":           params.Start,
		"end":             params.End,
		"_submit":         "Submit",
		"nonce":           params.Nonce,
	})
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
