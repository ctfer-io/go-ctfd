package api

type PostSharesParams struct {
	ChallengeID int    `json:"challenge_id"`
	Type        string `json:"type"`
}

func (client *Client) PostShares(params *PostSharesParams, opts ...Option) (*Shares, *MetaResponse, error) {
	sh := &Shares{}
	meta, err := client.Post("/shares", params, sh, opts...)
	if err != nil {
		return nil, meta, err
	}
	return sh, meta, nil
}
