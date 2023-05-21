package api

import (
	"io"
	"net/http"
	"regexp"

	"github.com/pkg/errors"
)

var nonceRegex = regexp.MustCompile(`([0-9a-f]{64})`)

// GetNonceAndSession is provided for simplicity.
// It uses the default HTTP client under the hood.
func GetNonceAndSession(url string, opts ...Option) (string, string, error) {
	req, _ := http.NewRequest(http.MethodGet, url+"/setup", nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()

	nonce, err := getNonce(res.Body)
	if err != nil {
		return "", "", err
	}

	for _, cookie := range res.Cookies() {
		if cookie.Name == "session" {
			return nonce, cookie.Value, nil
		}
	}
	return "", "", errors.New("session cookie not found")
}

func getNonce(r io.Reader) (string, error) {
	body, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	match := nonceRegex.Find(body)
	if match == nil {
		return "", errors.New("nonce not found")
	}
	return string(match), nil
}
