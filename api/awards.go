package api

type GetAwardsParams struct {
	UserID   *int    `schema:"user_id,omitempty"`
	TeamID   *int    `schema:"team_id,omitempty"`
	Type     *string `schema:"type,omitempty"`
	Value    *int    `schema:"value,omitempty"`
	Category *int    `schema:"category,omitempty"`
	Icon     *int    `schema:"icon,omitempty"`
	Q        *string `schema:"q,omitempty"`
	Field    *string `schema:"field,omitempty"`
}

func (client *Client) GetAwards(params *GetAwardsParams, opts ...Option) ([]*Award, error) {
	awards := []*Award{}
	if err := client.Get("/awards", params, &awards, opts...); err != nil {
		return nil, err
	}
	return awards, nil
}

type PostAwardsParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Icon        string `json:"icon"`
	UserID      int    `json:"user_id"`
	Value       int    `json:"value"`
}

func (client *Client) PostAwards(params *PostAwardsParams, opts ...Option) (*Award, error) {
	award := &Award{}
	if err := client.Post("/awards", params, &award, opts...); err != nil {
		return nil, err
	}
	return award, nil
}

func (client *Client) GetAward(id string, opts ...Option) (*Award, error) {
	award := &Award{}
	if err := client.Get("/awards/"+id, nil, &award, opts...); err != nil {
		return nil, err
	}
	return award, nil
}

func (client *Client) DeleteAward(id string, opts ...Option) error {
	return client.Delete("/awards/"+id, nil, nil, opts...)
}
