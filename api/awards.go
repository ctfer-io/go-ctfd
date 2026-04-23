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

func (client *Client) GetAwards(params *GetAwardsParams, opts ...Option) ([]*Award, *MetaResponse, error) {
	awards := []*Award{}
	meta, err := client.Get("/awards", params, &awards, opts...)
	if err != nil {
		return nil, meta, err
	}
	return awards, meta, nil
}

type PostAwardsParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Icon        string `json:"icon"`
	UserID      int    `json:"user_id"`
	Value       int    `json:"value"`
}

func (client *Client) PostAwards(params *PostAwardsParams, opts ...Option) (*Award, *MetaResponse, error) {
	award := &Award{}
	meta, err := client.Post("/awards", params, &award, opts...)
	if err != nil {
		return nil, meta, err
	}
	return award, meta, nil
}

func (client *Client) GetAward(id string, opts ...Option) (*Award, *MetaResponse, error) {
	award := &Award{}
	meta, err := client.Get("/awards/"+id, nil, &award, opts...)
	if err != nil {
		return nil, meta, err
	}
	return award, meta, nil
}

func (client *Client) DeleteAward(id string, opts ...Option) (*MetaResponse, error) {
	return client.Delete("/awards/"+id, nil, nil, opts...)
}
