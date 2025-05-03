package api

import "fmt"

type GetBracketsParams struct {
	Name        *string `schema:"name,omitempty"`
	Description *string `schema:"description,omitempty"`
	Type        *string `schema:"type,omitempty"`
	Q           *string `schema:"q,omitempty"`
}

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

type PatchBracketsParams struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Type        *string `json:"type,omitempty"`
}

func (client *Client) PatchBrackets(id int, params *PatchBracketsParams, opts ...Option) (*Bracket, error) {
	bk := &Bracket{}
	if err := client.Patch(fmt.Sprintf("/brackets/%d", id), params, bk, opts...); err != nil {
		return nil, err
	}
	return bk, nil
}

func (client *Client) DeleteBrackets(id int, opts ...Option) error {
	return client.Delete(fmt.Sprintf("/brackets/%d", id), nil, nil, opts...)
}
