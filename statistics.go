package goctfd

func (client *Client) GetStatisticsChallengesSolves(opts ...option) ([]*StatChallSubmission, error) {
	stats := []*StatChallSubmission{}
	if err := get(client, "/statistics/challenges/solves", nil, &stats, opts...); err != nil {
		return nil, err
	}
	return stats, nil
}

func (client *Client) GetStatisticsChallengesColumn(column string, opts ...option) (map[string]int, error) {
	stats := map[string]int{}
	if err := get(client, "/statistics/challenges/"+column, nil, &stats, opts...); err != nil {
		return nil, err
	}
	return stats, nil
}

func (client *Client) GetStatisticsScoresDistribution(opts ...option) (*Distribution, error) {
	dist := &Distribution{}
	if err := get(client, "/statistics/scores/distribution", nil, &dist, opts...); err != nil {
		return nil, err
	}
	return dist, nil
}

func (client *Client) GetStatisticsSubmissionsColumn(column string, opts ...option) (map[string]int, error) {
	stats := map[string]int{}
	if err := get(client, "/statistics/submissions/"+column, nil, &stats, opts...); err != nil {
		return nil, err
	}
	return stats, nil
}

func (client *Client) GetStatisticsTeams(opts ...option) (*StatTeams, error) {
	st := &StatTeams{}
	if err := get(client, "/statistics/teams", nil, &st, opts...); err != nil {
		return nil, err
	}
	return st, nil
}

func (client *Client) GetStatisticsUsers(opts ...option) (*StatUsers, error) {
	st := &StatUsers{}
	if err := get(client, "/statistics/users", nil, &st, opts...); err != nil {
		return nil, err
	}
	return st, nil
}

func (client *Client) GetStatisticsUsersColumn(column string, opts ...option) (*StatUsers, error) {
	st := &StatUsers{}
	if err := get(client, "/statistics/users/"+column, nil, &st, opts...); err != nil {
		return nil, err
	}
	return st, nil
}
