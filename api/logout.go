package api

import (
	"fmt"
	"net/http"
)

func (client *Client) Logout(opts ...Option) error {
	req, _ := http.NewRequest(http.MethodGet, "/logout", nil)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("CTFd responded with status code %d", res.StatusCode)
	}

	// Update nonce and session
	nonce, err := getNonce(res.Body)
	if err != nil {
		return err
	}
	client.nonce = nonce
	return nil
}
