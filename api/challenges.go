package api

import "fmt"

type GetChallengesParams struct {
	Name        *string `schema:"name,omitempty"`
	MaxAttempts *int    `schema:"max_attempts,omitempty"`
	Value       *int    `schema:"value,omitempty"`
	Category    *int    `schema:"category,omitempty"`
	Type        *string `schema:"type,omitempty"`
	State       *string `schema:"state,omitempty"`
	Q           *string `schema:"q,omitempty"`
	// TODO cover "view=admin" parameter that shows hidden challenges
}

func (client *Client) GetChallenges(params *GetChallengesParams, opts ...Option) ([]*Challenge, error) {
	challs := []*Challenge{}
	if err := get(client, "/challenges", params, &challs, opts...); err != nil {
		return nil, err
	}
	return challs, nil
}

type PostChallengesParams struct {
	Name           string        `json:"name"`
	Category       string        `json:"category"`
	Description    string        `json:"description"`
	Function       string        `json:"function"`
	ConnectionInfo *string       `json:"connection_info,omitempty"`
	Value          int           `json:"value"`
	Initial        *int          `json:"initial,omitempty"`
	Decay          *int          `json:"decay,omitempty"`
	Minimum        *int          `json:"minimum,omitempty"`
	MaxAttempts    *int          `json:"max_attempts,omitempty"`
	Requirements   *Requirements `json:"requirements,omitempty"`
	State          string        `json:"state"`
	Type           string        `json:"type"`
}

func (client *Client) PostChallenges(params *PostChallengesParams, opts ...Option) (*Challenge, error) {
	chall := &Challenge{}
	if err := post(client, "/challenges", params, &chall, opts...); err != nil {
		return nil, err
	}
	return chall, nil
}

type PostChallengesAttemptParams struct {
	// TODO support parameter (e.g. "preview=true")
	ChallengeID int    `json:"challenge_id"`
	Submission  string `json:"submission"`
}

func (client *Client) PostChallengesAttempt(params *PostChallengesAttemptParams, opts ...Option) (*Attempt, error) {
	attempt := &Attempt{}
	if err := post(client, "/challenges/attempt", params, &attempt, opts...); err != nil {
		return nil, err
	}
	return attempt, nil
}

func (client *Client) GetChallengesTypes(opts ...Option) (map[string]*Type, error) {
	types := map[string]*Type{}
	if err := get(client, "/challenges/types", nil, &types, opts...); err != nil {
		return nil, err
	}
	return types, nil
}

func (client *Client) GetChallenge(id int, opts ...Option) (*Challenge, error) {
	chall := &Challenge{}
	if err := get(client, fmt.Sprintf("/challenges/%d", id), nil, &chall, opts...); err != nil {
		return nil, err
	}
	return chall, nil
}

type PatchChallengeParams struct {
	Name           string  `json:"name"`
	Category       string  `json:"category"`
	Description    string  `json:"description"`
	Function       string  `json:"function"`
	ConnectionInfo *string `json:"connection_info,omitempty"`
	Initial        *int    `json:"initial,omitempty"`
	Decay          *int    `json:"decay,omitempty"`
	Minimum        *int    `json:"minimum,omitempty"`
	MaxAttempts    *int    `json:"max_attempts,omitempty"`
	// Requirements can update the challenge's behavior and prerequisites i.e.
	// the other challenges the team/user must have solved before.
	// WARNING: it won't return those in the response body, so updating this
	// field requires you to do it manually through *Client.GetChallengeRequirements
	Requirements *Requirements `json:"requirements,omitempty"`
	State        string        `json:"state"`
}

func (client *Client) DeleteChallenge(id int, opts ...Option) error {
	return delete(client, fmt.Sprintf("/challenges/%d", id), nil, nil, opts...)
}

func (client *Client) PatchChallenge(id int, params *PatchChallengeParams, opts ...Option) (*Challenge, error) {
	ch := &Challenge{}
	if err := patch(client, fmt.Sprintf("/challenges/%d", id), params, &ch, opts...); err != nil {
		return nil, err
	}
	return ch, nil
}

func (client *Client) GetChallengeFiles(id int, opts ...Option) ([]*File, error) {
	cf := []*File{}
	if err := get(client, fmt.Sprintf("/challenges/%d/files", id), nil, &cf, opts...); err != nil {
		return nil, err
	}
	return cf, nil
}

func (client *Client) GetChallengeFlags(id int, opts ...Option) ([]*Flag, error) {
	cf := []*Flag{}
	if err := get(client, fmt.Sprintf("/challenges/%d/flags", id), nil, &cf, opts...); err != nil {
		return nil, err
	}
	return cf, nil
}

func (client *Client) GetChallengeHints(id int, opts ...Option) ([]*Hint, error) {
	ch := []*Hint{}
	if err := get(client, fmt.Sprintf("/challenges/%d/hints", id), nil, &ch, opts...); err != nil {
		return nil, err
	}
	return ch, nil
}

func (client *Client) GetChallengeRequirements(id int, opts ...Option) (*Requirements, error) {
	req := &Requirements{}
	if err := get(client, fmt.Sprintf("/challenges/%d/requirements", id), nil, &req, opts...); err != nil {
		return nil, err
	}
	return req, nil
}

// TODO find content to determine model
func (client *Client) GetChallengeSolves(id int, opts ...Option) (*Challenge, error) {
	chall := &Challenge{}
	if err := get(client, fmt.Sprintf("/challenges/%d/solves", id), nil, &chall, opts...); err != nil {
		return nil, err
	}
	return chall, nil
}

func (client *Client) GetChallengeTags(id int, opts ...Option) ([]*Tag, error) {
	ct := []*Tag{}
	if err := get(client, fmt.Sprintf("/challenges/%d/tags", id), nil, &ct, opts...); err != nil {
		return nil, err
	}
	return ct, nil
}

func (client *Client) GetChallengeTopics(id int, opts ...Option) ([]*Topic, error) {
	ct := []*Topic{}
	if err := get(client, fmt.Sprintf("/challenges/%d/topics", id), nil, &ct, opts...); err != nil {
		return nil, err
	}
	return ct, nil
}
