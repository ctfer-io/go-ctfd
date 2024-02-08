package api

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type RegisterParams struct {
	Name     string
	Email    string
	Password string
}

func (client *Client) Register(params *RegisterParams, opts ...Option) error {
	val := url.Values{}
	val.Set("name", params.Name)
	val.Set("email", params.Email)
	val.Set("password", params.Password)
	val.Set("nonce", client.nonce) // XXX this should not be part of the API
	val.Set("_submit", "Submit")

	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(val.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusFound {
		return fmt.Errorf("CTFd responded with status code %d, which could be due to email reuse", res.StatusCode)
	}

	// Update session to track user then fetch nonce for later API calls
	cookieFound := false
	for _, cookie := range res.Cookies() {
		if cookie.Name == "session" {
			client.session = cookie.Value
			cookieFound = true
			break
		}
	}
	if !cookieFound {
		return errors.New("session cookie not found, may be due to server misconfiguration (not setup yet) or API instability")
	}
	req, _ = http.NewRequest(http.MethodGet, "/", nil)
	res, err = client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	nonce, err := getNonce(res.Body)
	if err != nil {
		return err
	}
	client.nonce = nonce

	return nil
}
