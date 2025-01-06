package api

type GetFlagsParams struct {
	ChallengeID *int    `schema:"challenge_id,omitempty"`
	Type        *string `schema:"type,omitempty"`
	Content     *string `schema:"content,omitempty"`
	Data        *string `schema:"data,omitempty"`
	Q           *string `schema:"q,omitempty"`
	Field       *string `schema:"field,omitempty"`
}

func (client *Client) GetFlags(params *GetFlagsParams, opts ...Option) ([]*Flag, error) {
	flags := []*Flag{}
	if err := client.Get("/flags", params, &flags, opts...); err != nil {
		return nil, err
	}
	return flags, nil
}

type PostFlagsParams struct {
	Challenge int    `json:"challenge"`
	Content   string `json:"content"`
	Data      string `json:"data"`
	Type      string `json:"type"`
}

func (client *Client) PostFlags(params *PostFlagsParams, opts ...Option) (*Flag, error) {
	flag := &Flag{}
	if err := client.Post("/flags", params, &flag, opts...); err != nil {
		return nil, err
	}
	return flag, nil
}

func (client *Client) GetFlagsTypes(opts ...Option) (map[string]*Type, error) {
	types := map[string]*Type{}
	if err := client.Get("/flags/types", nil, &types, opts...); err != nil {
		return nil, err
	}
	return types, nil
}

func (client *Client) GetFlagsType(typename string, opts ...Option) (*Type, error) {
	tp := &Type{}
	if err := client.Get("/flags/types/"+typename, nil, &tp, opts...); err != nil {
		return nil, err
	}
	return tp, nil
}

func (client *Client) GetFlag(id string, opts ...Option) (*Flag, error) {
	flag := &Flag{}
	if err := client.Get("/flags/"+id, nil, &flag, opts...); err != nil {
		return nil, err
	}
	return flag, nil
}

func (client *Client) DeleteFlag(id string, opts ...Option) error {
	return client.Delete("/flags/"+id, nil, nil, opts...)
}

type PatchFlagParams struct {
	Content string `json:"content"`
	Data    string `json:"data"`
	ID      string `json:"id"` // XXX should be int + duplicated with id in URL
	Type    string `json:"type"`
}

func (client *Client) PatchFlag(id string, params *PatchFlagParams, opts ...Option) (*Flag, error) {
	flag := &Flag{}
	if err := client.Patch("/flags/"+id, params, &flag, opts...); err != nil {
		return nil, err
	}
	return flag, nil
}
