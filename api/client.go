package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"reflect"

	"github.com/gorilla/schema"
	"github.com/pkg/errors"
)

// NewClient creates a fresh *Client.
// It automatically handles the session and its updates (login, logout...).
func NewClient(url, nonce, session, apiKey string) *Client {
	jar, _ := cookiejar.New(nil)
	return &Client{
		sub: &http.Client{
			Jar: jar,
			// // Don't follow redirections
			// CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			// 	return http.ErrUseLastResponse
			// },
		},
		url:     url,
		nonce:   nonce,
		session: session,
		apiKey:  apiKey,
	}
}

// Client is in charge of interacting with a CTFd instance.
type Client struct {
	sub *http.Client
	url string

	// Used for authentication, apiKey first, session&nonce else
	nonce   string
	session string
	apiKey  string
}

// SetAPIKey enables you to set a mandatory API key, or reset it
// by setting an empty one.
func (client *Client) SetAPIKey(apiKey string) {
	client.apiKey = apiKey
}

func (client *Client) Do(req *http.Request) (*http.Response, error) {
	// Set base URL
	newUrl, err := url.Parse(client.url + req.URL.String())
	if err != nil {
		return nil, err
	}
	req.URL = newUrl

	// Set headers
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json") // XXX this is necessary for GET method but should not, when call uses the APIProvider behaviour
	}
	if client.apiKey != "" {
		req.Header.Set("Authorization", "Token "+client.apiKey)
	} else {
		req.Header.Set("CSRF-Token", client.nonce)
		req.Header.Set("Cookie", "session="+client.session)
	}

	return client.sub.Do(req)
}

// Option represents a functional option.
type Option interface {
	apply(*options)
}

type options struct {
	Ctx context.Context
}

type ctxOption struct {
	Ctx context.Context
}

func (opt ctxOption) apply(opts *options) {
	opts.Ctx = opt.Ctx
}

// WithContext enable providing a context to the HTTP client
// during requests.
func WithContext(ctx context.Context) Option {
	return &ctxOption{
		Ctx: ctx,
	}
}

func applyOpts(req *http.Request, opts ...Option) *http.Request {
	reqopts := &options{
		Ctx: context.Background(),
	}

	for _, opt := range opts {
		opt.apply(reqopts)
	}

	req = req.WithContext(reqopts.Ctx)

	return req
}

// Call is in charge of handling common CTFd API behaviours,
// like dealing with status codes and JSON errors.
//
// It automatically prepends "/api/v1" to each path.
func (client *Client) Call(req *http.Request, dst any, opts ...Option) error {
	req = applyOpts(req, opts...)

	// Set API base URL
	newUrl, err := url.Parse("/api/v1" + req.URL.String())
	if err != nil {
		return err
	}
	req.URL = newUrl

	// Issue HTTP request
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Decode response
	resp := Response{
		Data: dst,
	}
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return errors.Wrapf(err, "CTFd responded with invalid JSON for content")
	}

	// Handle errors if any
	if resp.Errors != nil {
		return fmt.Errorf("CTFd responded with errors: %v", resp.Errors)
	}
	if !resp.Success {
		// This case should not happen, as status code already serves this goal
		// and errors gives the reasons.
		if resp.Message != nil {
			return fmt.Errorf("CTFd responded with no success but no error, got message: %s", *resp.Message)
		}
		return errors.New("CTFd responded with no success but no error, and no message")
	}
	return nil
}

type Response struct {
	Success bool    `json:"success"`
	Data    any     `json:"data,omitempty"`
	Errors  any     `json:"errors,omitempty"` // can't type it to []string due to API model instabilities
	Message *string `json:"message,omitempty"`
}

func (client *Client) Get(edp string, params any, dst any, opts ...Option) error {
	req, _ := http.NewRequest(http.MethodGet, edp, nil)

	// Encode URL parameters
	if params != nil && !reflect.ValueOf(params).IsNil() {
		val := url.Values{}
		if err := schema.NewEncoder().Encode(params, val); err != nil {
			return err
		}
		req.URL.RawQuery = val.Encode()
	}

	return client.Call(req, dst, opts...)
}

func (client *Client) Post(edp string, params any, dst any, opts ...Option) error {
	body, err := json.Marshal(params)
	if err != nil {
		return err
	}
	req, _ := http.NewRequest(http.MethodPost, edp, bytes.NewBuffer(body))

	return client.Call(req, dst, opts...)
}

func (client *Client) Patch(edp string, params any, dst any, opts ...Option) error {
	body, err := json.Marshal(params)
	if err != nil {
		return err
	}
	req, _ := http.NewRequest(http.MethodPatch, edp, bytes.NewBuffer(body))

	return client.Call(req, dst, opts...)
}

func (client *Client) Delete(edp string, params any, dst any, opts ...Option) (err error) {
	var body []byte
	if params != nil {
		body, err = json.Marshal(params)
		if err != nil {
			return
		}
	}
	req, _ := http.NewRequest(http.MethodDelete, edp, bytes.NewBuffer(body))

	return client.Call(req, dst, opts...)
}
