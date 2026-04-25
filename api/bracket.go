package api

import "fmt"

type GetBracketsParams struct {
	Name        *string `schema:"name,omitempty"`
	Description *string `schema:"description,omitempty"`
	Type        *string `schema:"type,omitempty"`
	Q           *string `schema:"q,omitempty"`
}

func (client *Client) GetBrackets(params *GetBracketsParams, opts ...Option) ([]*Bracket, *MetaResponse, error) {
	bks := []*Bracket{}
	meta, err := client.Get("/brackets", params, &bks, opts...)
	if err != nil {
		return nil, meta, err
	}
	return bks, meta, nil
}

type PostBracketsParams struct {
	ID          float64 `json:"id"` // XXX Why is that a float64 ?? Why is it even sent by the client as CTFd will return a new one ?
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
}

func (client *Client) PostBrackets(params *PostBracketsParams, opts ...Option) (*Bracket, *MetaResponse, error) {
	bk := &Bracket{}
	meta, err := client.Post("/brackets", params, bk, opts...)
	if err != nil {
		return nil, meta, err
	}
	return bk, meta, nil
}

type PatchBracketsParams struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Type        *string `json:"type,omitempty"`
}

func (client *Client) PatchBrackets(id int, params *PatchBracketsParams, opts ...Option) (*Bracket, *MetaResponse, error) {
	bk := &Bracket{}
	meta, err := client.Patch(fmt.Sprintf("/brackets/%d", id), params, bk, opts...)
	if err != nil {
		return nil, meta, err
	}
	return bk, meta, nil
}

func (client *Client) DeleteBrackets(id int, opts ...Option) (*MetaResponse, error) {
	return client.Delete(fmt.Sprintf("/brackets/%d", id), nil, nil, opts...)
}
