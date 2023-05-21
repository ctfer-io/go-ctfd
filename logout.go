package goctfd

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func (client *Client) Logout(opts ...Option) error {
	req, _ := http.NewRequest(http.MethodGet, "/logout", nil)
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
