package api

func (client *Client) GetStatisticsChallengesSolves(opts ...Option) ([]*StatChallSubmission, *MetaResponse, error) {
	stats := []*StatChallSubmission{}
	meta, err := client.Get("/statistics/challenges/solves", nil, &stats, opts...)
	if err != nil {
		return nil, meta, err
	}
	return stats, meta, nil
}

func (client *Client) GetStatisticsChallengesColumn(column string, opts ...Option) (map[string]int, *MetaResponse, error) {
	stats := map[string]int{}
	meta, err := client.Get("/statistics/challenges/"+column, nil, &stats, opts...)
	if err != nil {
		return nil, meta, err
	}
	return stats, meta, nil
}

func (client *Client) GetStatisticsScoresDistribution(opts ...Option) (*Distribution, *MetaResponse, error) {
	dist := &Distribution{}
	meta, err := client.Get("/statistics/scores/distribution", nil, &dist, opts...)
	if err != nil {
		return nil, meta, err
	}
	return dist, meta, nil
}

func (client *Client) GetStatisticsSubmissionsColumn(column string, opts ...Option) (map[string]int, *MetaResponse, error) {
	stats := map[string]int{}
	meta, err := client.Get("/statistics/submissions/"+column, nil, &stats, opts...)
	if err != nil {
		return nil, meta, err
	}
	return stats, meta, nil
}

func (client *Client) GetStatisticsTeams(opts ...Option) (*StatTeams, *MetaResponse, error) {
	st := &StatTeams{}
	meta, err := client.Get("/statistics/teams", nil, &st, opts...)
	if err != nil {
		return nil, meta, err
	}
	return st, meta, nil
}

func (client *Client) GetStatisticsUsers(opts ...Option) (*StatUsers, *MetaResponse, error) {
	st := &StatUsers{}
	meta, err := client.Get("/statistics/users", nil, &st, opts...)
	if err != nil {
		return nil, meta, err
	}
	return st, meta, nil
}

func (client *Client) GetStatisticsUsersColumn(column string, opts ...Option) (*StatUsers, *MetaResponse, error) {
	st := &StatUsers{}
	meta, err := client.Get("/statistics/users/"+column, nil, &st, opts...)
	if err != nil {
		return nil, meta, err
	}
	return st, meta, nil
}

func (client *Client) GetProgressionMatrix(opts ...Option) (*ProgressionMatrix, *MetaResponse, error) {
	pm := &ProgressionMatrix{}
	meta, err := client.Get("/statistics/progression/matrix", nil, &pm, opts...)
	if err != nil {
		return nil, meta, err
	}
	return pm, meta, nil
}
