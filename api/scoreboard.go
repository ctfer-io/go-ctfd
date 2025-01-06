package api

import "fmt"

func (client *Client) GetScoreboard(opts ...Option) ([]*Scoreboard, error) {
	sb := []*Scoreboard{}
	if err := client.Get("/scoreboard", nil, &sb, opts...); err != nil {
		return nil, err
	}
	return sb, nil
}

// GetScoreboardTopParams holds the parameters for the scoreboard top count endpoint.
type GetScoreboardTopParams struct {
	// Count is the top number of players to get the info.
	Count int `schema:"-"`

	// BracketID is an optional parameter to filter on a specific bracket.
	BracketID *int `schema:"bracket_id,omitempty"`
}

// GetScoreboardTop returns the scoreboard top for the given count as a map
// of the rank by the entry.
func (client *Client) GetScoreboardTop(params *GetScoreboardTopParams, opts ...Option) (map[string]*Scoreboard, error) {
	sb := map[string]*Scoreboard{}
	if err := client.Get(fmt.Sprintf("/scoreboard/top/%d", params.Count), params, &sb, opts...); err != nil {
		return nil, err
	}
	return sb, nil
}
