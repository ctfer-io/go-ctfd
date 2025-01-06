package api

type GetUnlocksParams struct {
	UserID *int    `schema:"user_id,omitempty"`
	TeamID *int    `schema:"team_id,omitempty"`
	Target *int    `schema:"target,omitempty"`
	Type   *string `schema:"type,omitempty"`
	Q      *string `schema:"q,omitempty"`
	Field  *string `schema:"field,omitempty"`
}

func (client *Client) GetUnlocks(params *GetUnlocksParams, opts ...Option) ([]*Unlock, error) {
	unlocks := []*Unlock{}
	if err := client.Get("/unlocks", params, &unlocks, opts...); err != nil {
		return nil, err
	}
	return unlocks, nil
}

type PostUnlocksParams struct {
	Target int    `json:"target"`
	Type   string `json:"type"`
}

func (client *Client) PostUnlocks(params *PostUnlocksParams, opts ...Option) (*Unlock, error) {
	unlock := &Unlock{}
	if err := client.Post("/unlocks", params, &unlock, opts...); err != nil {
		return nil, err
	}
	return unlock, nil
}
