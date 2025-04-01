package api

import (
	"fmt"
	"net/http"
	"net/url"
)

func (client *Client) Logout(opts ...Option) error {
	req, _ := http.NewRequest(http.MethodGet, "/logout", nil)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("CTFd responded with status code %d", res.StatusCode)
	}

	// Update nonce and session
	nonce, err := getNonce(res.Body)
	if err != nil {
		return err
	}
	client.nonce = nonce

	u, _ := url.Parse(client.url)
	hds := client.sub.Jar.Cookies(u)
	for _, hd := range hds {
		if hd.Name == "session" {
			client.session = hd.Value
			break
		}
	}

	return nil
}
