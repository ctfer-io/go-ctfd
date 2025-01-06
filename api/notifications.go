package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/schema"
)

type GetNotificationsParams struct {
	Title   *string `schema:"title,omitempty"`
	Content *string `schema:"content,omitempty"`
	UserID  *int    `schema:"user_id,omitempty"`
	TeamID  *int    `schema:"team_id,omitempty"`
	Q       *string `schema:"q,omitempty"`
	Field   *string `schema:"field,omitempty"`
	SinceID *int    `schema:"since_id,omitempty"`
}

func (client *Client) GetNotifications(params *GetNotificationsParams, opts ...Option) ([]*Notification, error) {
	notifs := []*Notification{}
	if err := client.Get("/notifications", params, &notifs, opts...); err != nil {
		return nil, err
	}
	return notifs, nil
}

type PostNotificationsParams struct {
	Content string `json:"content"`
	Sound   bool   `json:"sound"`
	Title   string `json:"title"`
	Type    string `json:"type"`
}

func (client *Client) PostNotifications(params *PostNotificationsParams, opts ...Option) (*Notification, error) {
	notif := &Notification{}
	if err := client.Post("/notifications", params, &notif, opts...); err != nil {
		return nil, err
	}
	return notif, nil
}

type HeadNotificationsParams struct {
	Title   *string `schema:"title,omitempty"`
	Content *string `schema:"content,omitempty"`
	UserID  *int    `schema:"user_id,omitempty"`
	TeamID  *int    `schema:"team_id,omitempty"`
	Q       *string `schema:"q,omitempty"`
	Field   *string `schema:"field,omitempty"`
	SinceID *int    `schema:"since_id,omitempty"`
}

// XXX does not need to be authenticated. Issue #2310
func (client *Client) HeadNotifications(params *HeadNotificationsParams, opts ...Option) (int, error) {
	req, _ := http.NewRequest(http.MethodHead, "/notifications", nil)

	// Encode URL parameters
	if params != nil {
		val := url.Values{}
		if err := schema.NewEncoder().Encode(params, val); err != nil {
			return 0, err
		}
		req.URL.RawQuery = val.Encode()
	}

	// Compile options
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
		return 0, err
	}
	defer res.Body.Close()

	// Check status code
	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("CTFd responded with status code %d", res.StatusCode)
	}

	return strconv.Atoi(res.Header.Get("Result-Count"))
}

func (client *Client) GetNotification(id string, opts ...Option) (*Notification, error) {
	notif := &Notification{}
	if err := client.Get("/notifications/"+id, nil, &notif, opts...); err != nil {
		return nil, err
	}
	return notif, nil
}

func (client *Client) DeleteNotification(id string, opts ...Option) error {
	return client.Delete("/notifications/"+id, nil, nil, opts...)
}
