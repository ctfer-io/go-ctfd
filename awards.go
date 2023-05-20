package goctfd

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

func (client *Client) GetAwards(params *GetAwardsParams, opts ...option) ([]*Award, error) {
	awards := []*Award{}
	if err := get(client, "/awards", params, &awards, opts...); err != nil {
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
	Value       string `json:"value"`
}

func (client *Client) PostAwards(params *PostAwardsParams, opts ...option) (*Award, error) {
	award := &Award{}
	if err := post(client, "/awards", params, &award, opts...); err != nil {
		return nil, err
	}
	return award, nil
}

func (client *Client) GetAward(id string, opts ...option) (*Award, error) {
	award := &Award{}
	if err := get(client, "/awards/"+id, nil, &award, opts...); err != nil {
		return nil, err
	}
	return award, nil
}

func (client *Client) DeleteAward(id string, opts ...option) error {
	return delete(client, "/awards/"+id, nil, nil, opts...)
}
