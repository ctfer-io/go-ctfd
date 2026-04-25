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
	// If view is set to admin, returns all challenges (not only the visible ones).
	View *string `schema:"view,omitempty"`
}

func (client *Client) GetChallenges(params *GetChallengesParams, opts ...Option) ([]*Challenge, *MetaResponse, error) {
	challs := []*Challenge{}
	meta, err := client.Get("/challenges", params, &challs, opts...)
	if err != nil {
		return nil, meta, err
	}
	return challs, meta, nil
}

type PostChallengesParams struct {
	Name           string        `json:"name"`
	Category       string        `json:"category"`
	Description    string        `json:"description"`
	Attribution    *string       `json:"attribution,omitempty"`
	Function       *string       `json:"function,omitempty"`
	ConnectionInfo *string       `json:"connection_info,omitempty"`
	Value          int           `json:"value"`
	Initial        *int          `json:"initial,omitempty"`
	Decay          *int          `json:"decay,omitempty"`
	Minimum        *int          `json:"minimum,omitempty"`
	Logic          string        `json:"logic"`
	MaxAttempts    *int          `json:"max_attempts,omitempty"`
	Position       *int          `json:"position,omitempty"`
	NextID         *int          `json:"next_id,omitempty"`
	Requirements   *Requirements `json:"requirements,omitempty"`
	State          string        `json:"state"`
	Type           string        `json:"type"`
}

func (client *Client) PostChallenges(params *PostChallengesParams, opts ...Option) (*Challenge, *MetaResponse, error) {
	chall := &Challenge{}
	meta, err := client.Post("/challenges", params, &chall, opts...)
	if err != nil {
		return nil, meta, err
	}
	return chall, meta, nil
}

type PostChallengesAttemptParams struct {
	// TODO support parameter (e.g. "preview=true")
	ChallengeID int    `json:"challenge_id"`
	Submission  string `json:"submission"`
}

func (client *Client) PostChallengesAttempt(params *PostChallengesAttemptParams, opts ...Option) (*Attempt, *MetaResponse, error) {
	attempt := &Attempt{}
	meta, err := client.Post("/challenges/attempt", params, &attempt, opts...)
	if err != nil {
		return nil, meta, err
	}
	return attempt, meta, nil
}

func (client *Client) GetChallengesTypes(opts ...Option) (map[string]*Type, *MetaResponse, error) {
	types := map[string]*Type{}
	meta, err := client.Get("/challenges/types", nil, &types, opts...)
	if err != nil {
		return nil, meta, err
	}
	return types, meta, nil
}

func (client *Client) GetChallenge(id int, opts ...Option) (*Challenge, *MetaResponse, error) {
	chall := &Challenge{}
	meta, err := client.Get(fmt.Sprintf("/challenges/%d", id), nil, &chall, opts...)
	if err != nil {
		return nil, meta, err
	}
	return chall, meta, nil
}

type PatchChallengeParams struct {
	Name           string  `json:"name"`
	Category       string  `json:"category"`
	Description    string  `json:"description"`
	Attribution    *string `json:"attribution,omitempty"`
	Function       *string `json:"function,omitempty"`
	ConnectionInfo *string `json:"connection_info,omitempty"`
	Value          *int    `json:"value,omitempty"`
	Initial        *int    `json:"initial,omitempty"`
	Decay          *int    `json:"decay,omitempty"`
	Minimum        *int    `json:"minimum,omitempty"`
	Logic          *string `json:"logic,omitempty"`
	MaxAttempts    *int    `json:"max_attempts,omitempty"`
	Position       *int    `json:"position,omitempty"`
	NextID         *int    `json:"next_id,omitempty"`
	// Requirements can update the challenge's behavior and prerequisites i.e.
	// the other challenges the team/user must have solved before.
	// WARNING: it won't return those in the response body, so updating this
	// field requires you to do it manually through *Client.GetChallengeRequirements
	Requirements *Requirements `json:"requirements,omitempty"`
	State        string        `json:"state"`
}

func (client *Client) DeleteChallenge(id int, opts ...Option) (*MetaResponse, error) {
	return client.Delete(fmt.Sprintf("/challenges/%d", id), nil, nil, opts...)
}

func (client *Client) PatchChallenge(id int, params *PatchChallengeParams, opts ...Option) (*Challenge, *MetaResponse, error) {
	ch := &Challenge{}
	meta, err := client.Patch(fmt.Sprintf("/challenges/%d", id), params, &ch, opts...)
	if err != nil {
		return nil, meta, err
	}
	return ch, meta, nil
}

func (client *Client) GetChallengeFiles(id int, opts ...Option) ([]*File, *MetaResponse, error) {
	cf := []*File{}
	meta, err := client.Get(fmt.Sprintf("/challenges/%d/files", id), nil, &cf, opts...)
	if err != nil {
		return nil, meta, err
	}
	return cf, meta, nil
}

func (client *Client) GetChallengeFlags(id int, opts ...Option) ([]*Flag, *MetaResponse, error) {
	cf := []*Flag{}
	meta, err := client.Get(fmt.Sprintf("/challenges/%d/flags", id), nil, &cf, opts...)
	if err != nil {
		return nil, meta, err
	}
	return cf, meta, nil
}

func (client *Client) GetChallengeHints(id int, opts ...Option) ([]*Hint, *MetaResponse, error) {
	ch := []*Hint{}
	meta, err := client.Get(fmt.Sprintf("/challenges/%d/hints", id), nil, &ch, opts...)
	if err != nil {
		return nil, meta, err
	}
	return ch, meta, nil
}

func (client *Client) GetChallengeRequirements(id int, opts ...Option) (*Requirements, *MetaResponse, error) {
	req := &Requirements{}
	meta, err := client.Get(fmt.Sprintf("/challenges/%d/requirements", id), nil, &req, opts...)
	if err != nil {
		return nil, meta, err
	}
	return req, meta, nil
}

// TODO find content to determine model
func (client *Client) GetChallengeSolves(id int, opts ...Option) (*Challenge, *MetaResponse, error) {
	chall := &Challenge{}
	meta, err := client.Get(fmt.Sprintf("/challenges/%d/solves", id), nil, &chall, opts...)
	if err != nil {
		return nil, meta, err
	}
	return chall, meta, nil
}

func (client *Client) GetChallengeTags(id int, opts ...Option) ([]*Tag, *MetaResponse, error) {
	ct := []*Tag{}
	meta, err := client.Get(fmt.Sprintf("/challenges/%d/tags", id), nil, &ct, opts...)
	if err != nil {
		return nil, meta, err
	}
	return ct, meta, nil
}

func (client *Client) GetChallengeTopics(id int, opts ...Option) ([]*Topic, *MetaResponse, error) {
	ct := []*Topic{}
	meta, err := client.Get(fmt.Sprintf("/challenges/%d/topics", id), nil, &ct, opts...)
	if err != nil {
		return nil, meta, err
	}
	return ct, meta, nil
}

func (client *Client) GetChallengeRatings(id int, opts ...Option) ([]*Rating, *MetaResponse, error) {
	ratings := []*Rating{}
	meta, err := client.Get(fmt.Sprintf("/challenges/%d/ratings", id), nil, &ratings, opts...)
	if err != nil {
		return nil, meta, err
	}
	return ratings, meta, nil
}

type PutChallengeRatingsParams struct {
	Value  int    `json:"value"` // either 1 for thumbsup or -1 for thumbsdown
	Review string `json:"review"`
}

func (client *Client) PutChallengeRatings(id int, params *PutChallengeRatingsParams, opts ...Option) (*Rating, *MetaResponse, error) {
	rat := &Rating{}
	meta, err := client.Put(fmt.Sprintf("/challenges/%d/ratings", id), params, rat, opts...)
	if err != nil {
		return nil, meta, err
	}
	return rat, meta, nil
}
