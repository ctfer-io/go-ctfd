package api

type GetUnlocksParams struct {
	UserID *int    `schema:"user_id,omitempty"`
	TeamID *int    `schema:"team_id,omitempty"`
	Target *int    `schema:"target,omitempty"`
	Type   *string `schema:"type,omitempty"`
	Q      *string `schema:"q,omitempty"`
	Field  *string `schema:"field,omitempty"`
}

func (client *Client) GetUnlocks(params *GetUnlocksParams, opts ...Option) ([]*Unlock, *MetaResponse, error) {
	unlocks := []*Unlock{}
	meta, err := client.Get("/unlocks", params, &unlocks, opts...)
	if err != nil {
		return nil, meta, err
	}
	return unlocks, meta, nil
}

type PostUnlocksParams struct {
	Target int    `json:"target"`
	Type   string `json:"type"`
}

func (client *Client) PostUnlocks(params *PostUnlocksParams, opts ...Option) (*Unlock, *MetaResponse, error) {
	unlock := &Unlock{}
	meta, err := client.Post("/unlocks", params, &unlock, opts...)
	if err != nil {
		return nil, meta, err
	}
	return unlock, meta, nil
}
