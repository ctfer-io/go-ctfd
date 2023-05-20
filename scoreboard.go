package goctfd

import "fmt"

func (client *Client) GetScoreboard(opts ...option) ([]*Scoreboard, error) {
	sb := []*Scoreboard{}
	if err := get(client, "/scoreboard", nil, &sb, opts...); err != nil {
		return nil, err
	}
	return sb, nil
}

func (client *Client) GetScoreboardTop(count int, opts ...option) ([]*Scoreboard, error) {
	sb := []*Scoreboard{}
	if err := get(client, fmt.Sprintf("/scoreboard/top/%d", count), nil, &sb, opts...); err != nil {
		return nil, err
	}
	return sb, nil
}
