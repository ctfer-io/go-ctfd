package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type ResetParams struct {
	Accounts      *string `json:"accounts,omitempty"`
	Submissions   *string `json:"submissions,omitempty"`
	Challenges    *string `json:"challenges,omitempty"`
	Pages         *string `json:"pages,omitempty"`
	Notifications *string `json:"notifications,omitempty"`
	// Nonce is autofilled by the API wrapper.
	// XXX the "nonce" should not be part of the API call but rather be extracted from HTTP headers.
	Nonce string `json:"nonce"`
}

func (client *Client) Reset(params *ResetParams, opts ...Option) error {
	// Autofill Nonce parameter if not set
	if params != nil {
		params.Nonce = client.nonce
	}

	// Build request
	body, err := json.Marshal(params)
	if err != nil {
		return errors.Wrap(err, "during JSON marshalling")
	}
	req, _ := http.NewRequest(http.MethodPost, "/admin/reset", bytes.NewBuffer(body))

	res, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "CTFd responded with error")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusFound {
		return fmt.Errorf("CTFd responded with status code %d", res.StatusCode)
	}

	// Update nonce and session (~ de-authenticate)
	nonce, session, err := GetNonceAndSession(client.url, opts...)
	if err != nil {
		return err
	}
	client.nonce = nonce
	client.session = session
	return nil
}
