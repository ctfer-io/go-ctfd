package api

import "fmt"

type GetCommentsParams struct {
	ChallengeID *string `schema:"challenge_id,omitempty"`
	UserID      *int    `schema:"user_id,omitempty"`
	TeamID      *int    `schema:"team_id,omitempty"`
	PageID      *int    `schema:"page_id,omitempty"`
	Q           *string `schema:"q,omitempty"`
	Field       *string `schema:"field,omitempty"`
}

func (client *Client) GetComments(params *GetCommentsParams, opts ...Option) ([]*Comment, error) {
	comments := []*Comment{}
	if err := get(client, "/comments", params, &comments, opts...); err != nil {
		return nil, err
	}
	return comments, nil
}

type PostCommentsParams struct {
	PageID  int    `json:"page_id"`
	Content string `json:"content"`
	Type    string `json:"type"`
}

func (client *Client) PostComments(params *PostCommentsParams, opts ...Option) (*Comment, error) {
	comment := &Comment{}
	if err := post(client, "/comments", params, &comment, opts...); err != nil {
		return nil, err
	}
	return comment, nil
}

func (client *Client) DeleteComment(id int, opts ...Option) error {
	return delete(client, fmt.Sprintf("/comments/%d", id), nil, nil, opts...)
}
