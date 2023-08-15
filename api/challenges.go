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
	ConnectionInfo *string       `json:"connection_info,omitempty"`
	MaxAttempts    *int          `json:"max_attempts,omitempty"`
	Value          *int          `json:"value,omitempty"`
	Initial        *int          `json:"initial,omitempty"`
	Decay          *int          `json:"decay,omitempty"`
	Minimum        *int          `json:"minimum,omitempty"`
	State          string        `json:"state"`
	Type           string        `json:"type"`
	Requirements   *Requirements `json:"requirements,omitempty"`
	NextID         *string       `json:"next_id,omitempty"`
}

func (client *Client) PostChallenges(params *PostChallengesParams, opts ...Option) (*Challenge, error) {
	chall := &Challenge{}
	if err := post(client, "/challenges", params, &chall, opts...); err != nil {
		return nil, err
	}
	return chall, nil
}

type PostChallengesAttemptParams struct {
	ChallengeID string `json:"challenge_id"`
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

func (client *Client) GetChallenge(id string, opts ...Option) (*Challenge, error) {
	chall := &Challenge{}
	if err := get(client, fmt.Sprintf("/challenges/%s", id), nil, &chall, opts...); err != nil {
		return nil, err
	}
	return chall, nil
}

type PatchChallengeParams struct {
	Name           *string       `json:"name,omitempty"`
	Category       *string       `json:"category,omitempty"`
	Description    *string       `json:"description,omitempty"`
	ConnectionInfo *string       `json:"connection_info,omitempty"`
	MaxAttempts    *int          `json:"max_attempts,omitempty"`
	Value          *int          `json:"value,omitempty"`
	Initial        *int          `json:"initial,omitempty"`
	Decay          *int          `json:"decay,omitempty"`
	Minimum        *int          `json:"minimum,omitempty"`
	State          *string       `json:"state,omitempty"`
	Requirements   *Requirements `json:"requirements,omitempty"`
	NextID         *string       `json:"next_id,omitempty"`
}

func (client *Client) DeleteChallenge(id string, opts ...Option) error {
	return delete(client, fmt.Sprintf("/challenges/%s", id), nil, nil, opts...)
}

func (client *Client) PatchChallenge(id string, params *PatchChallengeParams, opts ...Option) (*Challenge, error) {
	ch := &Challenge{}
	if err := patch(client, fmt.Sprintf("/challenges/%s", id), params, &ch, opts...); err != nil {
		return nil, err
	}
	return ch, nil
}

func (client *Client) GetChallengeFiles(id string, opts ...Option) ([]*File, error) {
	cf := []*File{}
	if err := get(client, fmt.Sprintf("/challenges/%s/files", id), nil, &cf, opts...); err != nil {
		return nil, err
	}
	return cf, nil
}

func (client *Client) GetChallengeFlags(id string, opts ...Option) ([]*Flag, error) {
	cf := []*Flag{}
	if err := get(client, fmt.Sprintf("/challenges/%s/flags", id), nil, &cf, opts...); err != nil {
		return nil, err
	}
	return cf, nil
}

func (client *Client) GetChallengeHints(id string, opts ...Option) ([]*Hint, error) {
	ch := []*Hint{}
	if err := get(client, fmt.Sprintf("/challenges/%s/hints", id), nil, &ch, opts...); err != nil {
		return nil, err
	}
	return ch, nil
}

func (client *Client) GetChallengeRequirements(id string, opts ...Option) (*Requirements, error) {
	req := &Requirements{}
	if err := get(client, fmt.Sprintf("/challenges/%s/requirements", id), nil, &req, opts...); err != nil {
		return nil, err
	}
	return req, nil
}

// TODO find content to determine model
func (client *Client) GetChallengeSolves(id string, opts ...Option) (*Challenge, error) {
	chall := &Challenge{}
	if err := get(client, fmt.Sprintf("/challenges/%s/solves", id), nil, &chall, opts...); err != nil {
		return nil, err
	}
	return chall, nil
}

func (client *Client) GetChallengeTags(id string, opts ...Option) ([]*Tag, error) {
	ct := []*Tag{}
	if err := get(client, fmt.Sprintf("/challenges/%s/tags", id), nil, &ct, opts...); err != nil {
		return nil, err
	}
	return ct, nil
}

func (client *Client) GetChallengeTopics(id string, opts ...Option) ([]*Topic, error) {
	ct := []*Topic{}
	if err := get(client, fmt.Sprintf("/challenges/%s/topics", id), nil, &ct, opts...); err != nil {
		return nil, err
	}
	return ct, nil
}
