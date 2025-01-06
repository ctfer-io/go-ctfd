package api

import "fmt"

type GetTeamsParams struct {
	Affiliation *string `schema:"affiliation"`
	Country     *string `schema:"country"`
	Bracket     *string `schema:"bracket"`
	Q           *string `schema:"q"`
	Field       *string `schema:"field"`
}

func (client *Client) GetTeams(params *GetTeamsParams, opts ...Option) ([]*Team, error) {
	teams := []*Team{}
	if err := client.Get("/teams", params, &teams, opts...); err != nil {
		return nil, err
	}
	return teams, nil
}

type PostTeamsParams struct {
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	Website     *string `json:"website,omitempty"`
	Affiliation *string `json:"affiliation,omitempty"`
	Country     *string `json:"country,omitempty"`
	Banned      bool    `json:"banned"`
	Hidden      bool    `json:"hidden"`
	Fields      []Field `json:"fields"`
}

func (client *Client) PostTeams(params *PostTeamsParams, opts ...Option) (*Team, error) {
	team := &Team{}
	if err := client.Post("/teams", params, &team, opts...); err != nil {
		return nil, err
	}
	return team, nil
}

func (client *Client) GetTeamsMe(opts ...Option) (*Team, error) {
	team := &Team{}
	if err := client.Get("/teams/me", nil, &team, opts...); err != nil {
		return nil, err
	}
	return team, nil
}

func (client *Client) DeleteTeamsMe(opts ...Option) error {
	return client.Delete("/teams/me", nil, nil, opts...)
}

type PatchTeamsParams struct {
	CaptainID   *int    `json:"captain_id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Email       *string `json:"email,omitempty"`
	Password    *string `json:"password,omitempty"`
	Website     *string `json:"website,omitempty"`
	Affiliation *string `json:"affiliation,omitempty"`
	Country     *string `json:"country,omitempty"`
	Banned      *bool   `json:"banned,omitempty"`
	Hidden      *bool   `json:"hidden,omitempty"`
	Fields      []Field `json:"fields"`
}

func (client *Client) PatchTeamsMe(params *PatchTeamsParams, opts ...Option) (*Team, error) {
	team := &Team{}
	if err := client.Patch("/teams/me", params, &team, opts...); err != nil {
		return nil, err
	}
	return team, nil
}

type PostTeamsMembersParams struct {
	UserID int `json:"user_id"`
}

func (client *Client) PostTeamsMeMembers(params *PostTeamsMembersParams, opts ...Option) (*Team, error) {
	team := &Team{}
	if err := client.Post("/teams/me/members", params, &team, opts...); err != nil {
		return nil, err
	}
	return team, nil
}

func (client *Client) GetTeamsMeAwards(opts ...Option) ([]*Award, error) {
	awards := []*Award{}
	if err := client.Get("/teams/me/awards", nil, &awards, opts...); err != nil {
		return nil, err
	}
	return awards, nil
}

func (client *Client) GetTeamsMeFails(opts ...Option) ([]*Submission, error) {
	submission := []*Submission{}
	if err := client.Get("/teams/me/fails", nil, &submission, opts...); err != nil {
		return nil, err
	}
	return submission, nil
}

func (client *Client) GetTeamsMeSolves(opts ...Option) ([]*Submission, error) {
	submission := []*Submission{}
	if err := client.Get("/teams/me/solves", nil, &submission, opts...); err != nil {
		return nil, err
	}
	return submission, nil
}

func (client *Client) GetTeam(id int, opts ...Option) (*Team, error) {
	team := &Team{}
	if err := client.Get(fmt.Sprintf("/teams/%d", id), nil, &team, opts...); err != nil {
		return nil, err
	}
	return team, nil
}

func (client *Client) DeleteTeam(id int, opts ...Option) error {
	return client.Delete(fmt.Sprintf("/teams/%d", id), nil, nil, opts...)
}

func (client *Client) PatchTeam(id int, params *PatchTeamsParams, opts ...Option) (*Team, error) {
	team := &Team{}
	if err := client.Patch(fmt.Sprintf("/teams/%d", id), params, &team, opts...); err != nil {
		return nil, err
	}
	return team, nil
}

type DeleteTeamMembersParams struct {
	UserID int `json:"user_id"`
}

// XXX mixture of DELETE and body for control should be cleaned
func (client *Client) DeleteTeamMembers(id int, params *DeleteTeamMembersParams, opts ...Option) ([]int, error) {
	v := []int{}
	if err := client.Delete(fmt.Sprintf("/teams/%d/members", id), params, &v, opts...); err != nil {
		return nil, err
	}
	return v, nil
}

func (client *Client) PostTeamMembers(id int, params *PostTeamsMembersParams, opts ...Option) (int, error) {
	// Use slice as a workaround due to API instabilities
	var team []int
	if err := client.Post(fmt.Sprintf("/teams/%d/members", id), params, &team, opts...); err != nil {
		return 0, err
	}
	return team[0], nil
}

func (client *Client) GetTeamAwards(id int, opts ...Option) ([]*Award, error) {
	awards := []*Award{}
	if err := client.Get(fmt.Sprintf("/teams/%d/awards", id), nil, &awards, opts...); err != nil {
		return nil, err
	}
	return awards, nil
}

func (client *Client) GetTeamFails(id int, opts ...Option) ([]*Submission, error) {
	submission := []*Submission{}
	if err := client.Get(fmt.Sprintf("/teams/%d/fails", id), nil, &submission, opts...); err != nil {
		return nil, err
	}
	return submission, nil
}

func (client *Client) GetTeamMembers(id int, opts ...Option) ([]int, error) {
	members := []int{}
	if err := client.Get(fmt.Sprintf("/teams/%d/members", id), nil, &members, opts...); err != nil {
		return nil, err
	}
	return members, nil
}

func (client *Client) GetTeamSolves(id int, opts ...Option) ([]*Submission, error) {
	submission := []*Submission{}
	if err := client.Get(fmt.Sprintf("/teams/%d/solves", id), nil, &submission, opts...); err != nil {
		return nil, err
	}
	return submission, nil
}
