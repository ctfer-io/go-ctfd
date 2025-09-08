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
	if err := client.Get("/hints", params, &hints, opts...); err != nil {
		return nil, err
	}
	return hints, nil
}

type PostHintsParams struct {
	ChallengeID  int          `json:"challenge_id"`
	Title        *string      `json:"title,omitempty"`
	Content      string       `json:"content"`
	Cost         int          `json:"cost"`
	Requirements Requirements `json:"requirements"`
}

func (client *Client) PostHints(params *PostHintsParams, opts ...Option) (*Hint, error) {
	hint := &Hint{}
	if err := client.Post("/hints", params, &hint, opts...); err != nil {
		return nil, err
	}
	return hint, nil
}

type GetHintParams struct {
	// As per CTFd commit ed5dbb762a013800edb1c322cbe0779b25c7daec, you can only get the hint data
	// if you are admin and specify a "preview" argument in the request
	Preview *bool `schema:"preview,omitempty"`
}

func (client *Client) GetHint(id string, params *GetHintParams, opts ...Option) (*Hint, error) {
	hint := &Hint{}
	if err := client.Get("/hints/"+id, params, &hint, opts...); err != nil {
		return nil, err
	}
	return hint, nil
}

type PatchHintsParams struct {
	ChallengeID  int          `json:"challenge_id"`
	Title        *string      `json:"title,omitempty"`
	Content      string       `json:"content"`
	Cost         int          `json:"cost"`
	Requirements Requirements `json:"requirements"`
}

func (client *Client) PatchHint(id string, params *PatchHintsParams, opts ...Option) (*Hint, error) {
	hint := &Hint{}
	if err := client.Patch("/hints/"+id, params, &hint, opts...); err != nil {
		return nil, err
	}
	return hint, nil
}

func (client *Client) DeleteHint(id string, opts ...Option) error {
	return client.Delete("/hints/"+id, nil, nil, opts...)
}
