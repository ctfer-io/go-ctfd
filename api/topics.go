package api

import (
	"net/http"
	"net/url"

	"github.com/gorilla/schema"
)

type GetTopicsParams struct {
	Value *string `schema:"value,omitempty"`
	Q     *string `schema:"q,omitempty"`
	Field *string `schema:"field,omitempty"`
}

func (client *Client) GetTopics(params *GetTopicsParams, opts ...Option) ([]*Topic, error) {
	topics := []*Topic{}
	if err := get(client, "/topics", params, &topics, opts...); err != nil {
		return nil, err
	}
	return topics, nil
}

// TODO support DELETE /topics

type PostTopicsParams struct {
	Challenge string `json:"challenge"`
	Type      string `json:"type"`
	Value     string `json:"value"`
}

func (client *Client) PostTopics(params *PostTopicsParams, opts ...Option) (*Topic, error) {
	topic := &Topic{}
	if err := post(client, "/topics", params, &topic, opts...); err != nil {
		return nil, err
	}
	return topic, nil
}

func (client *Client) GetTopic(id string, opts ...Option) (*Topic, error) {
	topic := &Topic{}
	if err := get(client, "/topics/"+id, nil, &topic, opts...); err != nil {
		return nil, err
	}
	return topic, nil
}

// XXX this API endpoint should be aligned with the others.
func (client *Client) DeleteTopic(id string, opts ...Option) error {
	req, _ := http.NewRequest(http.MethodGet, "/topics", nil)

	type deleteParams struct {
		Type     string `schema:"type"`
		TargetID string `schema:"target_id"`
	}
	params := deleteParams{
		Type:     "challenge",
		TargetID: id,
	}

	val := url.Values{}
	if err := schema.NewEncoder().Encode(params, val); err != nil {
		return err
	}
	req.URL.RawQuery = val.Encode()

	return call(client, req, nil, opts...)
}
