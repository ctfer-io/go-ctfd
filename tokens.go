package goctfd

func (client *Client) GetTokens(opts ...option) ([]*Token, error) {
	tokens := []*Token{}
	if err := get(client, "/tokens", nil, &tokens, opts...); err != nil {
		return nil, err
	}
	return tokens, nil
}

type PostTokensParams struct {
	Expiration string `json:"expiration"`
}

func (client *Client) PostTokens(params *PostTokensParams, opts ...option) (*Token, error) {
	token := &Token{}
	if err := post(client, "/tokens", params, &token, opts...); err != nil {
		return nil, err
	}
	return token, nil
}

// XXX Using this endpoint, you could get back the token value which is not a desired behaviour !
func (client *Client) GetToken(id string, opts ...option) (*Token, error) {
	token := &Token{}
	if err := get(client, "/tokens/"+id, nil, &token, opts...); err != nil {
		return nil, err
	}
	return token, nil
}

func (client *Client) DeleteToken(id string, opts ...option) error {
	return delete(client, "/tokens/"+id, nil, nil, opts...)
}
