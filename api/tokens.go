package api

func (client *Client) GetTokens(opts ...Option) ([]*Token, error) {
	tokens := []*Token{}
	if err := client.Get("/tokens", nil, &tokens, opts...); err != nil {
		return nil, err
	}
	return tokens, nil
}

type PostTokensParams struct {
	Description string `json:"description"`
	Expiration  string `json:"expiration"`
}

func (client *Client) PostTokens(params *PostTokensParams, opts ...Option) (*Token, error) {
	token := &Token{}
	if err := client.Post("/tokens", params, &token, opts...); err != nil {
		return nil, err
	}
	return token, nil
}

// XXX Using this endpoint, you could get back the token value which is not a desired behaviour ! Issue #2309
func (client *Client) GetToken(id string, opts ...Option) (*Token, error) {
	token := &Token{}
	if err := client.Get("/tokens/"+id, nil, &token, opts...); err != nil {
		return nil, err
	}
	return token, nil
}

func (client *Client) DeleteToken(id string, opts ...Option) error {
	return client.Delete("/tokens/"+id, nil, nil, opts...)
}
