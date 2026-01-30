package api

import (
	"bytes"
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
	req, client.sub.Transport = applyOpts(req, opts...)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("CTFd responded with status code %d, which could be due to email reuse", res.StatusCode)
	}

	// Update session to track user then fetch nonce for later API calls
	req, _ = http.NewRequest(http.MethodGet, "/", nil)
	req, client.sub.Transport = applyOpts(req, opts...)
	res, err = client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	nonce, err := getNonce(res.Body)
	if err != nil {
		return err
	}
	client.nonce = nonce

	return nil
}
