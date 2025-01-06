package api

func (client *Client) GetStatisticsChallengesSolves(opts ...Option) ([]*StatChallSubmission, error) {
	stats := []*StatChallSubmission{}
	if err := client.Get("/statistics/challenges/solves", nil, &stats, opts...); err != nil {
		return nil, err
	}
	return stats, nil
}

func (client *Client) GetStatisticsChallengesColumn(column string, opts ...Option) (map[string]int, error) {
	stats := map[string]int{}
	if err := client.Get("/statistics/challenges/"+column, nil, &stats, opts...); err != nil {
		return nil, err
	}
	return stats, nil
}

func (client *Client) GetStatisticsScoresDistribution(opts ...Option) (*Distribution, error) {
	dist := &Distribution{}
	if err := client.Get("/statistics/scores/distribution", nil, &dist, opts...); err != nil {
		return nil, err
	}
	return dist, nil
}

func (client *Client) GetStatisticsSubmissionsColumn(column string, opts ...Option) (map[string]int, error) {
	stats := map[string]int{}
	if err := client.Get("/statistics/submissions/"+column, nil, &stats, opts...); err != nil {
		return nil, err
	}
	return stats, nil
}

func (client *Client) GetStatisticsTeams(opts ...Option) (*StatTeams, error) {
	st := &StatTeams{}
	if err := client.Get("/statistics/teams", nil, &st, opts...); err != nil {
		return nil, err
	}
	return st, nil
}

func (client *Client) GetStatisticsUsers(opts ...Option) (*StatUsers, error) {
	st := &StatUsers{}
	if err := client.Get("/statistics/users", nil, &st, opts...); err != nil {
		return nil, err
	}
	return st, nil
}

func (client *Client) GetStatisticsUsersColumn(column string, opts ...Option) (*StatUsers, error) {
	st := &StatUsers{}
	if err := client.Get("/statistics/users/"+column, nil, &st, opts...); err != nil {
		return nil, err
	}
	return st, nil
}
