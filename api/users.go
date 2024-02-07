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
	if err := get(client, "/users", params, &users, opts...); err != nil {
		return nil, err
	}
	return users, nil
}

type PostUsersParams struct {
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Language    *string `json:"language,omitempty"`
	Password    string  `json:"password"`
	Website     *string `json:"website,omitempty"`
	Affiliation *string `json:"affiliation,omitempty"`
	Country     *string `json:"country,omitempty"`
	Type        string  `json:"type"` // "user" or "admin"
	Verified    bool    `json:"verified"`
	Hidden      bool    `json:"hidden"`
	Banned      bool    `json:"banned"`
	Fields      []Field `json:"fields"`
}

func (client *Client) PostUsers(params *PostUsersParams, opts ...Option) (*User, error) {
	user := &User{}
	if err := post(client, "/users", params, &user, opts...); err != nil {
		return nil, err
	}
	return user, nil
}

func (client *Client) GetUsersMe(opts ...Option) (*User, error) {
	user := &User{}
	if err := get(client, "/users/me", nil, &user, opts...); err != nil {
		return nil, err
	}
	return user, nil
}

type PatchUsersParams struct {
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Language    *string `json:"language,omitempty"`
	Password    *string `json:"password,omitempty"`
	Website     *string `json:"website,omitempty"`
	Affiliation *string `json:"affiliation,omitempty"`
	Country     *string `json:"country,omitempty"`
	Type        string  `json:"type"`
	Verified    bool    `json:"verified"`
	Hidden      bool    `json:"hidden"`
	Banned      bool    `json:"banned"`
	Fields      []Field `json:"fields"`
}

func (client *Client) PatchUsersMe(params *PatchUsersParams, opts ...Option) (*User, error) {
	user := &User{}
	if err := patch(client, "/users/me", params, &user, opts...); err != nil {
		return nil, err
	}
	return user, nil
}

func (client *Client) GetUsersMeAwards(opts ...Option) ([]*Award, error) {
	awards := []*Award{}
	if err := get(client, "/users/me/awards", nil, &awards, opts...); err != nil {
		return nil, err
	}
	return awards, nil
}

func (client *Client) GetUsersMeFails(opts ...Option) ([]*Submission, error) {
	submissions := []*Submission{}
	if err := get(client, "/users/me/fails", nil, &submissions, opts...); err != nil {
		return nil, err
	}
	return submissions, nil
}

func (client *Client) GetUsersMeSolves(opts ...Option) ([]*Submission, error) {
	submissions := []*Submission{}
	if err := get(client, "/users/me/solves", nil, &submissions, opts...); err != nil {
		return nil, err
	}
	return submissions, nil
}

func (client *Client) GetUser(id int, opts ...Option) (*User, error) {
	user := &User{}
	if err := get(client, fmt.Sprintf("/users/%d", id), nil, &user, opts...); err != nil {
		return nil, err
	}
	return user, nil
}

func (client *Client) DeleteUser(id int, opts ...Option) error {
	return delete(client, fmt.Sprintf("/users/%d", id), nil, nil, opts...)
}

func (client *Client) PatchUser(id int, params *PatchUsersParams, opts ...Option) (*User, error) {
	user := &User{}
	if err := patch(client, fmt.Sprintf("/users/%d", id), params, &user, opts...); err != nil {
		return nil, err
	}
	return user, nil
}

func (client *Client) GetUserAwards(id int, opts ...Option) ([]*Award, error) {
	awards := []*Award{}
	if err := get(client, fmt.Sprintf("/users/%d/awards", id), nil, &awards, opts...); err != nil {
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
	if err := post(client, fmt.Sprintf("/users/%d/email", id), params, &res, opts...); err != nil {
		return nil, err
	}
	return res, nil
}

func (client *Client) GetUserFails(id int, opts ...Option) ([]*Submission, error) {
	submisions := []*Submission{}
	if err := get(client, fmt.Sprintf("/users/%d/fails", id), nil, &submisions, opts...); err != nil {
		return nil, err
	}
	return submisions, nil
}

func (client *Client) GetUserSolves(id int, opts ...Option) ([]*Submission, error) {
	submisions := []*Submission{}
	if err := get(client, fmt.Sprintf("/users/%d/solves", id), nil, &submisions, opts...); err != nil {
		return nil, err
	}
	return submisions, nil
}
