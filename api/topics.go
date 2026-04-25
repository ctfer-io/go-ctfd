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

func (client *Client) GetTopics(params *GetTopicsParams, opts ...Option) ([]*Topic, *MetaResponse, error) {
	topics := []*Topic{}
	meta, err := client.Get("/topics", params, &topics, opts...)
	if err != nil {
		return nil, meta, err
	}
	return topics, meta, nil
}

type PostTopicsParams struct {
	Challenge int    `json:"challenge"`
	Type      string `json:"type"`
	Value     string `json:"value"`
}

func (client *Client) PostTopics(params *PostTopicsParams, opts ...Option) (*Topic, *MetaResponse, error) {
	topic := &Topic{}
	meta, err := client.Post("/topics", params, &topic, opts...)
	if err != nil {
		return nil, meta, err
	}
	return topic, meta, nil
}

func (client *Client) GetTopic(id string, opts ...Option) (*Topic, *MetaResponse, error) {
	topic := &Topic{}
	meta, err := client.Get("/topics/"+id, nil, &topic, opts...)
	if err != nil {
		return nil, meta, err
	}
	return topic, meta, nil
}

type DeleteTopicArgs struct {
	ID   string `schema:"target_id"`
	Type string `schema:"type"`
}

// TODO fix this endpoint API instability, should reconsider using a DELETE method with a JSON body rather that URL-encoded parameters as for all other endpoints
func (client *Client) DeleteTopic(params *DeleteTopicArgs, opts ...Option) (*MetaResponse, error) {
	// Build request
	req, _ := http.NewRequest(http.MethodDelete, "/topics", nil)
	req, client.sub.Transport = applyOpts(req, opts...)

	// Encode parameters
	val := url.Values{}
	if err := schema.NewEncoder().Encode(params, val); err != nil {
		return nil, err
	}
	req.URL.RawQuery = val.Encode()

	// Throw it to CTFd
	return client.Call(req, nil)
}
