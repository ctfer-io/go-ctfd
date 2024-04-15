package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/schema"
	"github.com/pkg/errors"
)

type ResetParams struct {
	Accounts      *string `schema:"accounts,omitempty"`
	Submissions   *string `schema:"submissions,omitempty"`
	Challenges    *string `schema:"challenges,omitempty"`
	Pages         *string `schema:"pages,omitempty"`
	Notifications *string `schema:"notifications,omitempty"`
	// Nonce is autofilled by the API wrapper.
	// XXX the "nonce" should not be part of the API call but rather be extracted from HTTP headers.
	Nonce string `schema:"nonce"`
}

func (client *Client) Reset(params *ResetParams, opts ...Option) error {
	// Autofill Nonce parameter if not set
	if params != nil {
		params.Nonce = client.nonce
	}

	// Build request
	str := ""
	if params != nil {
		val := url.Values{}
		if err := schema.NewEncoder().Encode(params, val); err != nil {
			return err
		}
		str = val.Encode()
	}
	req, _ := http.NewRequest(http.MethodPost, "/admin/reset", strings.NewReader(str))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "CTFd responded with error")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
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
