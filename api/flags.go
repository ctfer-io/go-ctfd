package api

type GetFlagsParams struct {
	ChallengeID *int    `schema:"challenge_id,omitempty"`
	Type        *string `schema:"type,omitempty"`
	Content     *string `schema:"content,omitempty"`
	Data        *string `schema:"data,omitempty"`
	Q           *string `schema:"q,omitempty"`
	Field       *string `schema:"field,omitempty"`
}

func (client *Client) GetFlags(params *GetFlagsParams, opts ...Option) ([]*Flag, *MetaResponse, error) {
	flags := []*Flag{}
	meta, err := client.Get("/flags", params, &flags, opts...)
	if err != nil {
		return nil, meta, err
	}
	return flags, meta, nil
}

type PostFlagsParams struct {
	Challenge int    `json:"challenge"`
	Content   string `json:"content"`
	Data      string `json:"data"`
	Type      string `json:"type"`
}

func (client *Client) PostFlags(params *PostFlagsParams, opts ...Option) (*Flag, *MetaResponse, error) {
	flag := &Flag{}
	meta, err := client.Post("/flags", params, &flag, opts...)
	if err != nil {
		return nil, meta, err
	}
	return flag, meta, nil
}

func (client *Client) GetFlagsTypes(opts ...Option) (map[string]*Type, *MetaResponse, error) {
	types := map[string]*Type{}
	meta, err := client.Get("/flags/types", nil, &types, opts...)
	if err != nil {
		return nil, meta, err
	}
	return types, meta, nil
}

func (client *Client) GetFlagsType(typename string, opts ...Option) (*Type, *MetaResponse, error) {
	tp := &Type{}
	meta, err := client.Get("/flags/types/"+typename, nil, &tp, opts...)
	if err != nil {
		return nil, meta, err
	}
	return tp, meta, nil
}

func (client *Client) GetFlag(id string, opts ...Option) (*Flag, *MetaResponse, error) {
	flag := &Flag{}
	meta, err := client.Get("/flags/"+id, nil, &flag, opts...)
	if err != nil {
		return nil, meta, err
	}
	return flag, meta, nil
}

func (client *Client) DeleteFlag(id string, opts ...Option) (*MetaResponse, error) {
	return client.Delete("/flags/"+id, nil, nil, opts...)
}

type PatchFlagParams struct {
	Content string `json:"content"`
	Data    string `json:"data"`
	ID      string `json:"id"` // XXX should be int + duplicated with id in URL
	Type    string `json:"type"`
}

func (client *Client) PatchFlag(id string, params *PatchFlagParams, opts ...Option) (*Flag, *MetaResponse, error) {
	flag := &Flag{}
	meta, err := client.Patch("/flags/"+id, params, &flag, opts...)
	if err != nil {
		return nil, meta, err
	}
	return flag, meta, nil
}
