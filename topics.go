package goctfd

type GetTopicsParams struct {
	Value *string `schema:"value,omitempty"`
	Q     *string `schema:"q,omitempty"`
	Field *string `schema:"field,omitempty"`
}

func (client *Client) GetTopics(params *GetTopicsParams, opts ...option) ([]*Topic, error) {
	topics := []*Topic{}
	if err := get(client, "/topics", params, &topics, opts...); err != nil {
		return nil, err
	}
	return topics, nil
}

// TODO support DELETE /topics

type PostTopicsParams struct {
	Challenge int    `json:"challenge"`
	Type      string `json:"type"`
	Value     string `json:"value"`
}

func (client *Client) PostTopics(params *PostTopicsParams, opts ...option) (*Topic, error) {
	topic := &Topic{}
	if err := post(client, "/topics", params, &topic, opts...); err != nil {
		return nil, err
	}
	return topic, nil
}

func (client *Client) GetTopic(id string, opts ...option) (*Topic, error) {
	topic := &Topic{}
	if err := get(client, "/topics/"+id, nil, &topic, opts...); err != nil {
		return nil, err
	}
	return topic, nil
}

func (client *Client) DeleteTopic(id string, opts ...option) error {
	return delete(client, "/topics/"+id, nil, nil, opts...)
}
