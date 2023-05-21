package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type ResetParams struct {
	Accounts      string `json:"accounts"`
	Submissions   string `json:"submissions"`
	Challenges    string `json:"challenges"`
	Pages         string `json:"pages"`
	Notifications string `json:"notifications"`
	Nonce         string `json:"nonce"` // XXX this should not be part of the API
}

func (client *Client) Reset(params *ResetParams, opts ...Option) error {
	body, err := json.Marshal(params)
	if err != nil {
		return err
	}
	req, _ := http.NewRequest(http.MethodPost, "/reset", bytes.NewBuffer(body))

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
			break
		}
	}
	return errors.New("session cookie not found")
}
