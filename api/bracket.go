package api

type GetBracketsParams struct{}

func (client *Client) GetBrackets(params *GetBracketsParams, opts ...Option) ([]*Bracket, error) {
	bks := []*Bracket{}
	if err := client.Get("/brackets", params, &bks, opts...); err != nil {
		return nil, err
	}
	return bks, nil
}

type PostBracketsParams struct {
	ID          float64 `json:"id"` // XXX Why is that a float64 ?? Why is it even sent by the client as CTFd will return a new one ?
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
}

func (client *Client) PostBrackets(params *PostBracketsParams, opts ...Option) (*Bracket, error) {
	bk := &Bracket{}
	if err := client.Post("/brackets", params, bk, opts...); err != nil {
		return nil, err
	}
	return bk, nil
}
