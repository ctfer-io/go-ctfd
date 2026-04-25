package api

import "fmt"

type GetTeamsParams struct {
	Affiliation *string `schema:"affiliation,omitempty"`
	Country     *string `schema:"country,omitempty"`
	Bracket     *string `schema:"bracket,omitempty"`
	Q           *string `schema:"q,omitempty"`
	Field       *string `schema:"field,omitempty"`

	Page *int `schema:"page,omitempty"`
	// per_page is not supported but hardcoded to 50
}

func (client *Client) GetTeams(params *GetTeamsParams, opts ...Option) ([]*Team, *MetaResponse, error) {
	teams := []*Team{}
	meta, err := client.Get("/teams", params, &teams, opts...)
	if err != nil {
		return nil, meta, err
	}
	return teams, meta, nil
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
	BracketID   *string `json:"bracket_id,omitempty"`
}

func (client *Client) PostTeams(params *PostTeamsParams, opts ...Option) (*Team, *MetaResponse, error) {
	team := &Team{}
	meta, err := client.Post("/teams", params, &team, opts...)
	if err != nil {
		return nil, meta, err
	}
	return team, meta, nil
}

func (client *Client) GetTeamsMe(opts ...Option) (*Team, *MetaResponse, error) {
	team := &Team{}
	meta, err := client.Get("/teams/me", nil, &team, opts...)
	if err != nil {
		return nil, meta, err
	}
	return team, meta, nil
}

func (client *Client) DeleteTeamsMe(opts ...Option) (*MetaResponse, error) {
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
	BracketID   *string `json:"bracket_id,omitempty"`
}

func (client *Client) PatchTeamsMe(params *PatchTeamsParams, opts ...Option) (*Team, *MetaResponse, error) {
	team := &Team{}
	meta, err := client.Patch("/teams/me", params, &team, opts...)
	if err != nil {
		return nil, meta, err
	}
	return team, meta, nil
}

type PostTeamsMembersParams struct {
	UserID int `json:"user_id"`
}

func (client *Client) PostTeamsMeMembers(params *PostTeamsMembersParams, opts ...Option) (*Team, *MetaResponse, error) {
	team := &Team{}
	meta, err := client.Post("/teams/me/members", params, &team, opts...)
	if err != nil {
		return nil, meta, err
	}
	return team, meta, nil
}

func (client *Client) GetTeamsMeAwards(opts ...Option) ([]*Award, *MetaResponse, error) {
	awards := []*Award{}
	meta, err := client.Get("/teams/me/awards", nil, &awards, opts...)
	if err != nil {
		return nil, meta, err
	}
	return awards, meta, nil
}

func (client *Client) GetTeamsMeFails(opts ...Option) ([]*Submission, *MetaResponse, error) {
	submission := []*Submission{}
	meta, err := client.Get("/teams/me/fails", nil, &submission, opts...)
	if err != nil {
		return nil, meta, err
	}
	return submission, meta, nil
}

func (client *Client) GetTeamsMeSolves(opts ...Option) ([]*Submission, *MetaResponse, error) {
	submission := []*Submission{}
	meta, err := client.Get("/teams/me/solves", nil, &submission, opts...)
	if err != nil {
		return nil, meta, err
	}
	return submission, meta, nil
}

func (client *Client) GetTeam(id int, opts ...Option) (*Team, *MetaResponse, error) {
	team := &Team{}
	meta, err := client.Get(fmt.Sprintf("/teams/%d", id), nil, &team, opts...)
	if err != nil {
		return nil, meta, err
	}
	return team, meta, nil
}

func (client *Client) DeleteTeam(id int, opts ...Option) (*MetaResponse, error) {
	return client.Delete(fmt.Sprintf("/teams/%d", id), nil, nil, opts...)
}

func (client *Client) PatchTeam(id int, params *PatchTeamsParams, opts ...Option) (*Team, *MetaResponse, error) {
	team := &Team{}
	meta, err := client.Patch(fmt.Sprintf("/teams/%d", id), params, &team, opts...)
	if err != nil {
		return nil, meta, err
	}
	return team, meta, nil
}

type DeleteTeamMembersParams struct {
	UserID int `json:"user_id"`
}

// XXX mixture of DELETE and body for control should be cleaned
func (client *Client) DeleteTeamMembers(id int, params *DeleteTeamMembersParams, opts ...Option) ([]int, *MetaResponse, error) {
	v := []int{}
	meta, err := client.Delete(fmt.Sprintf("/teams/%d/members", id), params, &v, opts...)
	if err != nil {
		return nil, meta, err
	}
	return v, meta, nil
}

func (client *Client) PostTeamMembers(id int, params *PostTeamsMembersParams, opts ...Option) (int, *MetaResponse, error) {
	// Use slice as a workaround due to API instabilities
	var team []int
	meta, err := client.Post(fmt.Sprintf("/teams/%d/members", id), params, &team, opts...)
	if err != nil {
		return 0, meta, err
	}
	return team[0], meta, nil
}

func (client *Client) GetTeamAwards(id int, opts ...Option) ([]*Award, *MetaResponse, error) {
	awards := []*Award{}
	meta, err := client.Get(fmt.Sprintf("/teams/%d/awards", id), nil, &awards, opts...)
	if err != nil {
		return nil, meta, err
	}
	return awards, meta, nil
}

func (client *Client) GetTeamFails(id int, opts ...Option) ([]*Submission, *MetaResponse, error) {
	submission := []*Submission{}
	meta, err := client.Get(fmt.Sprintf("/teams/%d/fails", id), nil, &submission, opts...)
	if err != nil {
		return nil, meta, err
	}
	return submission, meta, nil
}

func (client *Client) GetTeamMembers(id int, opts ...Option) ([]int, *MetaResponse, error) {
	members := []int{}
	meta, err := client.Get(fmt.Sprintf("/teams/%d/members", id), nil, &members, opts...)
	if err != nil {
		return nil, meta, err
	}
	return members, meta, nil
}

func (client *Client) GetTeamSolves(id int, opts ...Option) ([]*Submission, *MetaResponse, error) {
	submission := []*Submission{}
	meta, err := client.Get(fmt.Sprintf("/teams/%d/solves", id), nil, &submission, opts...)
	if err != nil {
		return nil, meta, err
	}
	return submission, meta, nil
}
