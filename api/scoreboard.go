package api

import "fmt"

func (client *Client) GetScoreboard(opts ...Option) ([]*Scoreboard, error) {
	sb := []*Scoreboard{}
	if err := get(client, "/scoreboard", nil, &sb, opts...); err != nil {
		return nil, err
	}
	return sb, nil
}

func (client *Client) GetScoreboardTop(count int, opts ...Option) ([]*Scoreboard, error) {
	sb := []*Scoreboard{}
	if err := get(client, fmt.Sprintf("/scoreboard/top/%d", count), nil, &sb, opts...); err != nil {
		return nil, err
	}
	return sb, nil
}
