package api

import "fmt"

type GetSolutionsParams struct{}

func (client *Client) GetSolutions(challID int, params *GetSolutionsParams, opts ...Option) (*Solution, error) {
	sol := &Solution{}
	if err := client.Get(fmt.Sprintf("/solutions/%d", challID), params, sol, opts...); err != nil {
		return nil, err
	}
	return sol, nil
}

type PostSolutionsParams struct {
	ChallengeID int    `json:"challenge_id"`
	Content     string `json:"content"`
	State       string `json:"state"`
}

func (client *Client) PostSolutions(params *PostSolutionsParams, opts ...Option) (*Solution, error) {
	sol := &Solution{}
	if err := client.Post("/solutions", params, sol, opts...); err != nil {
		return nil, err
	}
	return sol, nil
}

type PatchSolutionsParams struct {
	Content string `json:"content"`
	State   string `json:"state"`
}

func (client *Client) PatchSolutions(id int, params *PatchSolutionsParams, opts ...Option) (*Solution, error) {
	sol := &Solution{}
	if err := client.Patch(fmt.Sprintf("/solutions/%d", id), params, sol, opts...); err != nil {
		return nil, err
	}
	return sol, nil
}

func (client *Client) DeleteSolutions(id int, opts ...Option) error {
	return client.Delete(fmt.Sprintf("/solutions/%d", id), nil, nil, opts...)
}
