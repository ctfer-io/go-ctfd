package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type LoginParams struct {
	Name     string
	Password string
	Nonce    string // XXX this should not be part of the API
}

// Returns the session value or an error.
//
// WARNING 1: this endpoint is not officially supported.
// WARNING 2: provided client must have a no-follow-redirect behaviour,
// or a cookie jar. Else, it won't detect the login worked properly thus
// won't extract and save the new session id.
func (client *Client) Login(params *LoginParams, opts ...Option) error {
	val := url.Values{}
	val.Set("name", params.Name)
	val.Set("password", params.Password)
	val.Set("nonce", params.Nonce)
	val.Set("_submit", "Submit")

	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(val.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusFound {
		return fmt.Errorf("CTFd responded with status code %d", res.StatusCode)
	}

	// Update nonce and session
	nonce, err := getNonce(res.Body)
	if err != nil {
		return err
	}
	client.nonce = nonce

	for _, cookie := range res.Cookies() {
		if cookie.Name == "session" {
			client.session = cookie.Value
			return nil
		}
	}
	return errors.New("session cookie not found")
}
