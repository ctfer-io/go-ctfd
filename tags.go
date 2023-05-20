package goctfd

type GetTagsParams struct {
	ChallengeID *int    `schema:"challenge_id,omitempty"`
	Value       *string `schema:"value,omitempty"`
	Q           *string `schema:"q,omitempty"`
	Field       *string `schema:"field,omitempty"`
}

func (client *Client) GetTags(params *GetTagsParams, opts ...option) ([]*Tag, error) {
	tags := []*Tag{}
	if err := get(client, "/tags", params, &tags, opts...); err != nil {
		return nil, err
	}
	return tags, nil
}

type PostTagsParams struct {
	Challenge int    `json:"challenge"`
	Value     string `json:"value"`
}

func (client *Client) PostTags(params *PostTagsParams, opts ...option) (*Tag, error) {
	tag := &Tag{}
	if err := post(client, "/tags", params, &tag, opts...); err != nil {
		return nil, err
	}
	return tag, nil
}

func (client *Client) GetTag(id string, opts ...option) (*Tag, error) {
	tag := &Tag{}
	if err := get(client, "/tags/"+id, nil, &tag, opts...); err != nil {
		return nil, err
	}
	return tag, nil
}

func (client *Client) DeleteTag(id string, opts ...option) error {
	return delete(client, "/tags/"+id, nil, nil, opts...)
}

type PatchTagsParams struct {
	Value string `json:"value"`
}

func (client *Client) PatchTags(id string, params *PatchTagsParams, opts ...option) (*Tag, error) {
	tag := &Tag{}
	if err := patch(client, "/tags/"+id, params, &tag, opts...); err != nil {
		return nil, err
	}
	return tag, nil
}
