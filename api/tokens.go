package api

func (client *Client) GetTokens(opts ...Option) ([]*Token, *MetaResponse, error) {
	tokens := []*Token{}
	meta, err := client.Get("/tokens", nil, &tokens, opts...)
	if err != nil {
		return nil, meta, err
	}
	return tokens, meta, nil
}

type PostTokensParams struct {
	Description string `json:"description"`
	Expiration  string `json:"expiration"`
}

func (client *Client) PostTokens(params *PostTokensParams, opts ...Option) (*Token, *MetaResponse, error) {
	token := &Token{}
	meta, err := client.Post("/tokens", params, &token, opts...)
	if err != nil {
		return nil, meta, err
	}
	return token, meta, nil
}

// XXX Using this endpoint, you could get back the token value which is not a desired behaviour ! Issue #2309
func (client *Client) GetToken(id string, opts ...Option) (*Token, *MetaResponse, error) {
	token := &Token{}
	meta, err := client.Get("/tokens/"+id, nil, &token, opts...)
	if err != nil {
		return nil, meta, err
	}
	return token, meta, nil
}

func (client *Client) DeleteToken(id string, opts ...Option) (*MetaResponse, error) {
	return client.Delete("/tokens/"+id, nil, nil, opts...)
}
