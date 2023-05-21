package goctfd

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
	if err := get(client, "/teams", params, &teams, opts...); err != nil {
		return nil, err
	}
	return teams, nil
}

type PostTeamsParams struct {
	Name     string   `json:"name"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	Banned   bool     `json:"banned"`
	Hidden   bool     `json:"hidden"`
	Fields   []string `json:"fields"`
}

func (client *Client) PostTeams(params *PostTeamsParams, opts ...Option) (*Team, error) {
	team := &Team{}
	if err := post(client, "/teams", params, &team, opts...); err != nil {
		return nil, err
	}
	return team, nil
}

func (client *Client) GetTeamsMe(opts ...Option) (*Team, error) {
	team := &Team{}
	if err := get(client, "/teams/me", nil, &team, opts...); err != nil {
		return nil, err
	}
	return team, nil
}

func (client *Client) DeleteTeamsMe(opts ...Option) error {
	return delete(client, "/teams/me", nil, nil, opts...)
}

type PatchTeamsParams struct {
	CaptainID *int    `json:"captain_id,omitempty"`
	Banned    *bool   `json:"banned,omitempty"`
	Fields    []Field `json:"fields,omitempty"`
	Hidden    *bool   `json:"hidden,omitempty"`
	Name      *string `json:"name,omitempty"`
}

func (client *Client) PatchTeamsMe(params *PatchTeamsParams, opts ...Option) (*Team, error) {
	team := &Team{}
	if err := patch(client, "/teams/me", params, &team, opts...); err != nil {
		return nil, err
	}
	return team, nil
}

type PostTeamsMembers struct {
	UserID int `json:"user_id"`
}

func (client *Client) PostTeamsMeMembers(params *PostTeamsMembers, opts ...Option) (*Team, error) {
	team := &Team{}
	if err := post(client, "/teams/me/members", params, &team, opts...); err != nil {
		return nil, err
	}
	return team, nil
}

func (client *Client) GetTeamsMeAwards(opts ...Option) ([]*Award, error) {
	awards := []*Award{}
	if err := get(client, "/teams/me/awards", nil, &awards, opts...); err != nil {
		return nil, err
	}
	return awards, nil
}

func (client *Client) GetTeamsMeFails(opts ...Option) ([]*Submission, error) {
	submission := []*Submission{}
	if err := get(client, "/teams/me/fails", nil, &submission, opts...); err != nil {
		return nil, err
	}
	return submission, nil
}

func (client *Client) GetTeamsMeSolves(opts ...Option) ([]*Submission, error) {
	submission := []*Submission{}
	if err := get(client, "/teams/me/solves", nil, &submission, opts...); err != nil {
		return nil, err
	}
	return submission, nil
}

func (client *Client) GetTeam(id int, opts ...Option) (*Team, error) {
	team := &Team{}
	if err := get(client, fmt.Sprintf("/teams/%d", id), nil, &team, opts...); err != nil {
		return nil, err
	}
	return team, nil
}

func (client *Client) DeleteTeam(id int, opts ...Option) error {
	return delete(client, fmt.Sprintf("/teams/%d", id), nil, nil, opts...)
}

func (client *Client) PatchTeam(id int, params *PatchTeamsParams, opts ...Option) (*Team, error) {
	team := &Team{}
	if err := patch(client, fmt.Sprintf("/teams/%d", id), params, &team, opts...); err != nil {
		return nil, err
	}
	return team, nil
}

// XXX this could cause errors easily, play with JSON variants

type DeleteTeamMembersParams struct {
	UserID int `json:"user_id"`
}

// XXX mixture of DELETE and body for control should be cleaned
func (client *Client) DeleteTeamMembers(id int, params *DeleteTeamMembersParams, opts ...Option) ([]int, error) {
	v := []int{}
	if err := delete(client, fmt.Sprintf("/teams/%d/members", id), params, &v, opts...); err != nil {
		return nil, err
	}
	return v, nil
}

func (client *Client) PostTeamMembers(id int, params *PostTeamsMembers, opts ...Option) (*Team, error) {
	team := &Team{}
	if err := post(client, fmt.Sprintf("/teams/%d/members", id), params, &team, opts...); err != nil {
		return nil, err
	}
	return team, nil
}

func (client *Client) GetTeamAwards(id int, opts ...Option) ([]*Award, error) {
	awards := []*Award{}
	if err := get(client, fmt.Sprintf("/teams/%d/awards", id), nil, &awards, opts...); err != nil {
		return nil, err
	}
	return awards, nil
}

func (client *Client) GetTeamFails(id int, opts ...Option) ([]*Submission, error) {
	submission := []*Submission{}
	if err := get(client, fmt.Sprintf("/teams/%d/fails", id), nil, &submission, opts...); err != nil {
		return nil, err
	}
	return submission, nil
}

func (client *Client) GetTeamMembers(id int, opts ...Option) ([]int, error) {
	members := []int{}
	if err := get(client, fmt.Sprintf("/teams/%d/members", id), nil, &members, opts...); err != nil {
		return nil, err
	}
	return members, nil
}

func (client *Client) GetTeamSolves(id int, opts ...Option) ([]*Submission, error) {
	submission := []*Submission{}
	if err := get(client, fmt.Sprintf("/teams/%d/solves", id), nil, &submission, opts...); err != nil {
		return nil, err
	}
	return submission, nil
}
