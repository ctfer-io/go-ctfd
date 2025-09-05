package api

type PostSharesParams struct {
	ChallengeID int    `json:"challenge_id"`
	Type        string `json:"type"`
}

func (client *Client) PostShares(params *PostSharesParams, opts ...Option) (*Shares, error) {
	sh := &Shares{}
	if err := client.Post("/shares", params, sh, opts...); err != nil {
		return nil, err
	}
	return sh, nil
}
