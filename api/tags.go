package api

type GetTagsParams struct {
	ChallengeID *int    `schema:"challenge_id,omitempty"`
	Value       *string `schema:"value,omitempty"`
	Q           *string `schema:"q,omitempty"`
	Field       *string `schema:"field,omitempty"`
}

func (client *Client) GetTags(params *GetTagsParams, opts ...Option) ([]*Tag, error) {
	tags := []*Tag{}
	if err := client.Get("/tags", params, &tags, opts...); err != nil {
		return nil, err
	}
	return tags, nil
}

type PostTagsParams struct {
	Challenge int    `json:"challenge"`
	Value     string `json:"value"`
}

func (client *Client) PostTags(params *PostTagsParams, opts ...Option) (*Tag, error) {
	tag := &Tag{}
	if err := client.Post("/tags", params, &tag, opts...); err != nil {
		return nil, err
	}
	return tag, nil
}

func (client *Client) GetTag(id string, opts ...Option) (*Tag, error) {
	tag := &Tag{}
	if err := client.Get("/tags/"+id, nil, &tag, opts...); err != nil {
		return nil, err
	}
	return tag, nil
}

func (client *Client) DeleteTag(id string, opts ...Option) error {
	return client.Delete("/tags/"+id, nil, nil, opts...)
}

type PatchTagsParams struct {
	Value string `json:"value"`
}

func (client *Client) PatchTags(id string, params *PatchTagsParams, opts ...Option) (*Tag, error) {
	tag := &Tag{}
	if err := client.Patch("/tags/"+id, params, &tag, opts...); err != nil {
		return nil, err
	}
	return tag, nil
}
