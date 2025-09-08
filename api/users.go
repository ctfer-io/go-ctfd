package api

import "fmt"

type GetUsersParams struct {
	Affiliation *string `schema:"affiliation"`
	Country     *string `schema:"country"`
	Bracket     *string `schema:"bracket"`
	Q           *string `schema:"q"`
	Field       *string `schema:"field"`
}

// TODO handle pagination, but don't seem supported by CTFd API
func (client *Client) GetUsers(params *GetUsersParams, opts ...Option) ([]*User, error) {
	users := []*User{}
	if err := client.Get("/users", params, &users, opts...); err != nil {
		return nil, err
	}
	return users, nil
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

func (client *Client) PostUsers(params *PostUsersParams, opts ...Option) (*User, error) {
	user := &User{}
	if err := client.Post("/users", params, &user, opts...); err != nil {
		return nil, err
	}
	return user, nil
}

func (client *Client) GetUsersMe(opts ...Option) (*User, error) {
	user := &User{}
	if err := client.Get("/users/me", nil, &user, opts...); err != nil {
		return nil, err
	}
	return user, nil
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

func (client *Client) PatchUsersMe(params *PatchUsersParams, opts ...Option) (*User, error) {
	user := &User{}
	if err := client.Patch("/users/me", params, &user, opts...); err != nil {
		return nil, err
	}
	return user, nil
}

func (client *Client) GetUsersMeAwards(opts ...Option) ([]*Award, error) {
	awards := []*Award{}
	if err := client.Get("/users/me/awards", nil, &awards, opts...); err != nil {
		return nil, err
	}
	return awards, nil
}

func (client *Client) GetUsersMeFails(opts ...Option) ([]*Submission, error) {
	submissions := []*Submission{}
	if err := client.Get("/users/me/fails", nil, &submissions, opts...); err != nil {
		return nil, err
	}
	return submissions, nil
}

func (client *Client) GetUsersMeSolves(opts ...Option) ([]*Submission, error) {
	submissions := []*Submission{}
	if err := client.Get("/users/me/solves", nil, &submissions, opts...); err != nil {
		return nil, err
	}
	return submissions, nil
}

func (client *Client) GetUser(id int, opts ...Option) (*User, error) {
	user := &User{}
	if err := client.Get(fmt.Sprintf("/users/%d", id), nil, &user, opts...); err != nil {
		return nil, err
	}
	return user, nil
}

func (client *Client) DeleteUser(id int, opts ...Option) error {
	return client.Delete(fmt.Sprintf("/users/%d", id), nil, nil, opts...)
}

func (client *Client) PatchUser(id int, params *PatchUsersParams, opts ...Option) (*User, error) {
	user := &User{}
	if err := client.Patch(fmt.Sprintf("/users/%d", id), params, &user, opts...); err != nil {
		return nil, err
	}
	return user, nil
}

func (client *Client) GetUserAwards(id int, opts ...Option) ([]*Award, error) {
	awards := []*Award{}
	if err := client.Get(fmt.Sprintf("/users/%d/awards", id), nil, &awards, opts...); err != nil {
		return nil, err
	}
	return awards, nil
}

type PostUserMailParams struct {
	Text string `json:"text"`
}

// TODO find model when email turned on
func (client *Client) PostUserMail(params *PostUserMailParams, id int, opts ...Option) (any, error) {
	var res any
	if err := client.Post(fmt.Sprintf("/users/%d/email", id), params, &res, opts...); err != nil {
		return nil, err
	}
	return res, nil
}

func (client *Client) GetUserFails(id int, opts ...Option) ([]*Submission, error) {
	submisions := []*Submission{}
	if err := client.Get(fmt.Sprintf("/users/%d/fails", id), nil, &submisions, opts...); err != nil {
		return nil, err
	}
	return submisions, nil
}

func (client *Client) GetUserSolves(id int, opts ...Option) ([]*Submission, error) {
	submisions := []*Submission{}
	if err := client.Get(fmt.Sprintf("/users/%d/solves", id), nil, &submisions, opts...); err != nil {
		return nil, err
	}
	return submisions, nil
}
