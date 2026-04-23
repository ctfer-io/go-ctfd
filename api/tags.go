package api

type GetTagsParams struct {
	ChallengeID *int    `schema:"challenge_id,omitempty"`
	Value       *string `schema:"value,omitempty"`
	Q           *string `schema:"q,omitempty"`
	Field       *string `schema:"field,omitempty"`
}

func (client *Client) GetTags(params *GetTagsParams, opts ...Option) ([]*Tag, *MetaResponse, error) {
	tags := []*Tag{}
	meta, err := client.Get("/tags", params, &tags, opts...)
	if err != nil {
		return nil, meta, err
	}
	return tags, meta, nil
}

type PostTagsParams struct {
	Challenge int    `json:"challenge"`
	Value     string `json:"value"`
}

func (client *Client) PostTags(params *PostTagsParams, opts ...Option) (*Tag, *MetaResponse, error) {
	tag := &Tag{}
	meta, err := client.Post("/tags", params, &tag, opts...)
	if err != nil {
		return nil, meta, err
	}
	return tag, meta, nil
}

func (client *Client) GetTag(id string, opts ...Option) (*Tag, *MetaResponse, error) {
	tag := &Tag{}
	meta, err := client.Get("/tags/"+id, nil, &tag, opts...)
	if err != nil {
		return nil, meta, err
	}
	return tag, meta, nil
}

func (client *Client) DeleteTag(id string, opts ...Option) (*MetaResponse, error) {
	return client.Delete("/tags/"+id, nil, nil, opts...)
}

type PatchTagsParams struct {
	Value string `json:"value"`
}

func (client *Client) PatchTags(id string, params *PatchTagsParams, opts ...Option) (*Tag, *MetaResponse, error) {
	tag := &Tag{}
	meta, err := client.Patch("/tags/"+id, params, &tag, opts...)
	if err != nil {
		return nil, meta, err
	}
	return tag, meta, nil
}
