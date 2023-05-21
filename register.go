package goctfd

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
	Nonce    string // XXX this should not be part of the API
}

func (client *Client) Register(params *RegisterParams, opts ...Option) error {
	val := url.Values{}
	val.Set("name", params.Name)
	val.Set("email", params.Email)
	val.Set("password", params.Password)
	val.Set("nonce", params.Nonce)
	val.Set("_submit", "Submit")

	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(val.Encode()))
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
