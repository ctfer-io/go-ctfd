package api

type GetHintsParams struct {
	Type        *string `schema:"type,omitempty"`
	ChallengeID *int    `schema:"challenge_id,omitempty"`
	Content     *string `schema:"content,omitempty"`
	Cost        *int    `schema:"cost,omitempty"`
	Q           *string `schema:"q,omitempty"`
	Field       *string `schema:"field,omitempty"`
}

func (client *Client) GetHints(params *GetHintsParams, opts ...Option) ([]*Hint, error) {
	hints := []*Hint{}
	if err := get(client, "/hints", params, &hints, opts...); err != nil {
		return nil, err
	}
	return hints, nil
}

type PostHintsParams struct {
	ChallengeID  int          `json:"challenge_id"`
	Content      string       `json:"content"`
	Cost         string       `json:"cost"` // XXX may be int
	Requirements Requirements `json:"requirements"`
}

func (client *Client) PostHints(params *PostHintsParams, opts ...Option) (*Hint, error) {
	hint := &Hint{}
	if err := post(client, "/hints", params, &hint, opts...); err != nil {
		return nil, err
	}
	return hint, nil
}

func (client *Client) GetHint(id string, opts ...Option) (*Hint, error) {
	hint := &Hint{}
	if err := get(client, "/hints/"+id, nil, &hint, opts...); err != nil {
		return nil, err
	}
	return hint, nil
}

type PatchHintsParams struct {
	ChallengeID  int          `json:"challenge_id"`
	Content      string       `json:"content"`
	Cost         int          `json:"cost"`
	Requirements Requirements `json:"requirements"`
}

func (client *Client) DeleteHint(id string, opts ...Option) error {
	return delete(client, "/hints/"+id, nil, nil, opts...)
}

func (client *Client) PatchHint(id string, params *PatchHintsParams, opts ...Option) (*Hint, error) {
	hint := &Hint{}
	if err := patch(client, "/hints/"+id, params, &hint, opts...); err != nil {
		return nil, err
	}
	return hint, nil
}
