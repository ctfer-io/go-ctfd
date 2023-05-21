package goctfd

type GetSubmissionsParams struct {
	ChallengeID *int    `schema:"challenge_id,omitempty"`
	UserID      *int    `schema:"user_id,omitempty"`
	TeamID      *int    `schema:"team_id,omitempty"`
	IP          *string `schema:"ip,omitempty"`
	Provided    *string `schema:"provided,omitempty"`
	Type        *string `schema:"type,omitempty"`
	Q           *string `schema:"q,omitempty"`
	Field       *string `schema:"field,omitempty"`
}

// TODO support pagination ? CTFd does not seem to support parameters for this
func (client *Client) GetSubmissions(params *GetSubmissionsParams, opts ...Option) ([]*Submission, error) {
	submissions := []*Submission{}
	if err := get(client, "/submissions", params, &submissions, opts...); err != nil {
		return nil, err
	}
	return submissions, nil
}

// XXX POST /submissions remains usable ?

func (client *Client) GetSubmission(id string, opts ...Option) (*Submission, error) {
	submission := &Submission{}
	if err := get(client, "/submissions/"+id, nil, &submission, opts...); err != nil {
		return nil, err
	}
	return submission, nil
}

func (client *Client) DeleteSubmission(id string, opts ...Option) error {
	return delete(client, "/submissions/"+id, nil, nil, opts...)
}
