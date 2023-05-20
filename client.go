package goctfd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/schema"
	"github.com/pkg/errors"
)

// HTTPClient defines what the sub-client must implement for this
// wrapper to work.
// It enables inter-operability of implementations, like giving a
// compliant circuit-breaker implementation.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

var _ HTTPClient = (*http.Client)(nil)

// AuthProvider defines a source of authentication the client will
// use when calling the API.
type AuthProvider interface {
	Authenticate(*http.Request)
}

// APIProvider is a kind of AuthProvider that uses the API Key, as
// documented by CTFd.
type APIProvider struct {
	ApiKey string
}

var _ AuthProvider = (*APIProvider)(nil)

func (pv *APIProvider) Authenticate(req *http.Request) {
	req.Header.Set("Authorization", "Token "+pv.ApiKey) // XXX the "Token" value should be properly documented in API
}

// FormProvider is a kind of AuthProvider that uses the Session cookie
// and CSRF/Nonce value, behaving such as an authenticated user.
// This is not documented despite being a possibility.
// For instance, you could use this one to setup the CTFd, create an API
// key and then turn to the APIProvider.
type FormProvider struct {
	Session, Nonce string
}

var _ AuthProvider = (*FormProvider)(nil)

func (pv *FormProvider) Authenticate(req *http.Request) {
	req.Header.Set("Cookie", "session="+pv.Session)
	req.Header.Set("CSRF-Token", pv.Nonce)
}

// NewClient creates a fresh *Client.
func NewClient(sub HTTPClient, pv AuthProvider, url string) *Client {
	return &Client{
		sub: sub,
		pv:  pv,
		url: url,
	}
}

// Client is in charge of interacting with a CTFd instance.
type Client struct {
	sub HTTPClient
	pv  AuthProvider
	url string
}

// SetAuthProvider enables you to change authentication provider
// on the fly, for instance when turning from FormProvider to APIProvider.
func (client *Client) SetAuthProvider(pv AuthProvider) {
	client.pv = pv
}

var _ HTTPClient = (*Client)(nil)

func (client *Client) Do(req *http.Request) (*http.Response, error) {
	// Set base URL
	newUrl, err := url.Parse(client.url + "/api/v1" + req.URL.String())
	if err != nil {
		return nil, err
	}
	req.URL = newUrl

	// Set headers
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json") // XXX this is necessary for GET method but should not, when call uses the APIProvider behaviour
	}
	client.pv.Authenticate(req)

	return client.sub.Do(req)
}

type option interface {
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

func WithContext(ctx context.Context) option {
	return &ctxOption{
		Ctx: ctx,
	}
}

// call is in charge of handling common CTFd API behaviours,
// like dealing with status codes and JSON errors.
func call(client *Client, req *http.Request, dst any, opts ...option) error {
	reqopts := &options{
		Ctx: context.Background(),
	}
	for _, opt := range opts {
		opt.apply(reqopts)
	}
	req = req.WithContext(reqopts.Ctx)

	// Issue HTTP request
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Decode response
	type ctfdResponse struct {
		Success bool     `json:"success"`
		Data    any      `json:"data,omitempty"`
		Errors  []string `json:"errors,omitempty"`
	}
	resp := ctfdResponse{
		Data: dst,
	}
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return errors.Wrap(err, "CTFd responded with invalid JSON")
	}

	// Handle errors if any
	if len(resp.Errors) != 0 {
		return fmt.Errorf("CTFd responded with errors: %v", resp.Errors)
	}
	if !resp.Success {
		// This case should not happen, as status code already serves this goal
		// and errors gives the reasons.
		return errors.New("CTFd responded with no success but no error")
	}
	return nil
}

func get(client *Client, edp string, params any, dst any, opts ...option) error {
	req, _ := http.NewRequest(http.MethodGet, edp, nil)

	// Encode URL parameters
	if params != nil {
		val := url.Values{}
		if err := schema.NewEncoder().Encode(params, val); err != nil {
			return err
		}
		req.URL.RawQuery = val.Encode()
	}

	return call(client, req, dst)
}

func post(client *Client, edp string, params any, dst any, opts ...option) error {
	body, err := json.Marshal(params)
	if err != nil {
		return err
	}
	req, _ := http.NewRequest(http.MethodPost, edp, bytes.NewBuffer(body))

	return call(client, req, dst)
}

func patch(client *Client, edp string, params any, dst any, opts ...option) error {
	body, err := json.Marshal(params)
	if err != nil {
		return err
	}
	req, _ := http.NewRequest(http.MethodPatch, edp, bytes.NewBuffer(body))

	return call(client, req, dst)
}

func delete(client *Client, edp string, params any, dst any, opts ...option) (err error) {
	var body []byte
	if params != nil {
		body, err = json.Marshal(params)
		if err != nil {
			return
		}
	}
	req, _ := http.NewRequest(http.MethodDelete, edp, bytes.NewBuffer(body))

	return call(client, req, dst)
}
