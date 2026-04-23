package api

import "fmt"

type GetCommentsParams struct {
	ChallengeID *int    `schema:"challenge_id,omitempty"`
	UserID      *int    `schema:"user_id,omitempty"`
	TeamID      *int    `schema:"team_id,omitempty"`
	PageID      *int    `schema:"page_id,omitempty"`
	Q           *string `schema:"q,omitempty"`
	Field       *string `schema:"field,omitempty"`
}

func (client *Client) GetComments(params *GetCommentsParams, opts ...Option) ([]*Comment, *MetaResponse, error) {
	comments := []*Comment{}
	meta, err := client.Get("/comments", params, &comments, opts...)
	if err != nil {
		return nil, meta, err
	}
	return comments, meta, nil
}

type PostCommentsParams struct {
	PageID  int    `json:"page_id"`
	Content string `json:"content"`
	Type    string `json:"type"`
}

func (client *Client) PostComments(params *PostCommentsParams, opts ...Option) (*Comment, *MetaResponse, error) {
	comment := &Comment{}
	meta, err := client.Post("/comments", params, &comment, opts...)
	if err != nil {
		return nil, meta, err
	}
	return comment, meta, nil
}

func (client *Client) DeleteComment(id int, opts ...Option) (*MetaResponse, error) {
	return client.Delete(fmt.Sprintf("/comments/%d", id), nil, nil, opts...)
}
