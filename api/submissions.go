package api

type GetSubmissionsParams struct {
	ChallengeID *int    `schema:"challenge_id,omitempty"`
	UserID      *int    `schema:"user_id,omitempty"`
	TeamID      *int    `schema:"team_id,omitempty"`
	IP          *string `schema:"ip,omitempty"`
	Provided    *string `schema:"provided,omitempty"`
	Type        *string `schema:"type,omitempty"`
	Q           *string `schema:"q,omitempty"`
	Field       *string `schema:"field,omitempty"`

	Page    *int `schema:"page,omitempty"`
	PerPage *int `schema:"per_page,omitempty"`
}

func (client *Client) GetSubmissions(params *GetSubmissionsParams, opts ...Option) ([]*Submission, *MetaResponse, error) {
	submissions := []*Submission{}
	meta, err := client.Get("/submissions", params, &submissions, opts...)
	if err != nil {
		return nil, meta, err
	}
	return submissions, meta, nil
}

// XXX POST /submissions remains usable ?

func (client *Client) GetSubmission(id string, opts ...Option) (*Submission, *MetaResponse, error) {
	submission := &Submission{}
	meta, err := client.Get("/submissions/"+id, nil, &submission, opts...)
	if err != nil {
		return nil, meta, err
	}
	return submission, meta, nil
}

func (client *Client) DeleteSubmission(id string, opts ...Option) (*MetaResponse, error) {
	return client.Delete("/submissions/"+id, nil, nil, opts...)
}
