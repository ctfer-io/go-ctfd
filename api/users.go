package api

import "fmt"

type GetUsersParams struct {
	Affiliation *string `schema:"affiliation,omitempty"`
	Country     *string `schema:"country,omitempty"`
	Bracket     *string `schema:"bracket,omitempty"`
	Q           *string `schema:"q,omitempty"`
	Field       *string `schema:"field,omitempty"`

	Page *int `schema:"page,omitempty"`
	// per_page is not supported but hardcoded to 50
}

func (client *Client) GetUsers(params *GetUsersParams, opts ...Option) ([]*User, *MetaResponse, error) {
	users := []*User{}
	meta, err := client.Get("/users", params, &users, opts...)
	if err != nil {
		return nil, meta, err
	}
	return users, meta, nil
}

type PostUsersParams struct {
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	Language       *string `json:"language,omitempty"`
	Password       string  `json:"password"`
	Website        *string `json:"website,omitempty"`
	Affiliation    *string `json:"affiliation,omitempty"`
	Country        *string `json:"country,omitempty"`
	Type           string  `json:"type"` // "user" or "admin"
	Verified       bool    `json:"verified"`
	Hidden         bool    `json:"hidden"`
	Banned         bool    `json:"banned"`
	ChangePassword bool    `json:"change_password"`
	Fields         []Field `json:"fields"`
	BracketID      *string `json:"bracket_id,omitempty"`
}

func (client *Client) PostUsers(params *PostUsersParams, opts ...Option) (*User, *MetaResponse, error) {
	user := &User{}
	meta, err := client.Post("/users", params, &user, opts...)
	if err != nil {
		return nil, meta, err
	}
	return user, meta, nil
}

func (client *Client) GetUsersMe(opts ...Option) (*User, *MetaResponse, error) {
	user := &User{}
	meta, err := client.Get("/users/me", nil, &user, opts...)
	if err != nil {
		return nil, meta, err
	}
	return user, meta, nil
}

type PatchUsersParams struct {
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	Language       *string `json:"language,omitempty"`
	Password       *string `json:"password,omitempty"`
	Website        *string `json:"website,omitempty"`
	Affiliation    *string `json:"affiliation,omitempty"`
	Country        *string `json:"country,omitempty"`
	Type           *string `json:"type,omitempty"`
	Verified       *bool   `json:"verified,omitempty"`
	Hidden         *bool   `json:"hidden,omitempty"`
	Banned         *bool   `json:"banned,omitempty"`
	ChangePassword *bool   `json:"change_password,omitempty"`
	Fields         []Field `json:"fields"`
	BracketID      *string `json:"bracket_id,omitempty"`
}

func (client *Client) PatchUsersMe(params *PatchUsersParams, opts ...Option) (*User, *MetaResponse, error) {
	user := &User{}
	meta, err := client.Patch("/users/me", params, &user, opts...)
	if err != nil {
		return nil, meta, err
	}
	return user, meta, nil
}

func (client *Client) GetUsersMeAwards(opts ...Option) ([]*Award, *MetaResponse, error) {
	awards := []*Award{}
	meta, err := client.Get("/users/me/awards", nil, &awards, opts...)
	if err != nil {
		return nil, meta, err
	}
	return awards, meta, nil
}

func (client *Client) GetUsersMeFails(opts ...Option) ([]*Submission, *MetaResponse, error) {
	submissions := []*Submission{}
	meta, err := client.Get("/users/me/fails", nil, &submissions, opts...)
	if err != nil {
		return nil, meta, err
	}
	return submissions, meta, nil
}

func (client *Client) GetUsersMeSolves(opts ...Option) ([]*Submission, *MetaResponse, error) {
	submissions := []*Submission{}
	meta, err := client.Get("/users/me/solves", nil, &submissions, opts...)
	if err != nil {
		return nil, meta, err
	}
	return submissions, meta, nil
}

func (client *Client) GetUser(id int, opts ...Option) (*User, *MetaResponse, error) {
	user := &User{}
	meta, err := client.Get(fmt.Sprintf("/users/%d", id), nil, &user, opts...)
	if err != nil {
		return nil, meta, err
	}
	return user, meta, nil
}

func (client *Client) DeleteUser(id int, opts ...Option) (*MetaResponse, error) {
	return client.Delete(fmt.Sprintf("/users/%d", id), nil, nil, opts...)
}

func (client *Client) PatchUser(id int, params *PatchUsersParams, opts ...Option) (*User, *MetaResponse, error) {
	user := &User{}
	meta, err := client.Patch(fmt.Sprintf("/users/%d", id), params, &user, opts...)
	if err != nil {
		return nil, meta, err
	}
	return user, meta, nil
}

func (client *Client) GetUserAwards(id int, opts ...Option) ([]*Award, *MetaResponse, error) {
	awards := []*Award{}
	meta, err := client.Get(fmt.Sprintf("/users/%d/awards", id), nil, &awards, opts...)
	if err != nil {
		return nil, meta, err
	}
	return awards, meta, nil
}

type PostUserMailParams struct {
	Text string `json:"text"`
}

// TODO find model when email turned on
func (client *Client) PostUserMail(params *PostUserMailParams, id int, opts ...Option) (any, *MetaResponse, error) {
	var res any
	meta, err := client.Post(fmt.Sprintf("/users/%d/email", id), params, &res, opts...)
	if err != nil {
		return nil, meta, err
	}
	return res, meta, nil
}

func (client *Client) GetUserFails(id int, opts ...Option) ([]*Submission, *MetaResponse, error) {
	submisions := []*Submission{}
	meta, err := client.Get(fmt.Sprintf("/users/%d/fails", id), nil, &submisions, opts...)
	if err != nil {
		return nil, meta, err
	}
	return submisions, meta, nil
}

func (client *Client) GetUserSolves(id int, opts ...Option) ([]*Submission, *MetaResponse, error) {
	submisions := []*Submission{}
	meta, err := client.Get(fmt.Sprintf("/users/%d/solves", id), nil, &submisions, opts...)
	if err != nil {
		return nil, meta, err
	}
	return submisions, meta, nil
}
