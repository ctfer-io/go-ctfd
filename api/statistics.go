package api

func (client *Client) GetStatisticsChallengesSolves(opts ...Option) ([]*StatChallSubmission, error) {
	stats := []*StatChallSubmission{}
	if err := get(client, "/statistics/challenges/solves", nil, &stats, opts...); err != nil {
		return nil, err
	}
	return stats, nil
}

func (client *Client) GetStatisticsChallengesColumn(column string, opts ...Option) (map[string]int, error) {
	stats := map[string]int{}
	if err := get(client, "/statistics/challenges/"+column, nil, &stats, opts...); err != nil {
		return nil, err
	}
	return stats, nil
}

func (client *Client) GetStatisticsScoresDistribution(opts ...Option) (*Distribution, error) {
	dist := &Distribution{}
	if err := get(client, "/statistics/scores/distribution", nil, &dist, opts...); err != nil {
		return nil, err
	}
	return dist, nil
}

func (client *Client) GetStatisticsSubmissionsColumn(column string, opts ...Option) (map[string]int, error) {
	stats := map[string]int{}
	if err := get(client, "/statistics/submissions/"+column, nil, &stats, opts...); err != nil {
		return nil, err
	}
	return stats, nil
}

func (client *Client) GetStatisticsTeams(opts ...Option) (*StatTeams, error) {
	st := &StatTeams{}
	if err := get(client, "/statistics/teams", nil, &st, opts...); err != nil {
		return nil, err
	}
	return st, nil
}

func (client *Client) GetStatisticsUsers(opts ...Option) (*StatUsers, error) {
	st := &StatUsers{}
	if err := get(client, "/statistics/users", nil, &st, opts...); err != nil {
		return nil, err
	}
	return st, nil
}

func (client *Client) GetStatisticsUsersColumn(column string, opts ...Option) (*StatUsers, error) {
	st := &StatUsers{}
	if err := get(client, "/statistics/users/"+column, nil, &st, opts...); err != nil {
		return nil, err
	}
	return st, nil
}
