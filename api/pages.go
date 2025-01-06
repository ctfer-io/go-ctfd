package api

type GetPagesParams struct {
	ID           *int    `schema:"id,omitempty"`
	Title        *string `schema:"title,omitempty"`
	Route        *string `schema:"route,omitempty"`
	Draft        *bool   `schema:"draft,omitempty"`
	Hidden       *bool   `schema:"hidden,omitempty"`
	AuthRequired *bool   `schema:"auth_required,omitempty"`
	Q            *string `schema:"q,omitempty"`
	Field        *string `schema:"field,omitempty"`
}

func (client *Client) GetPages(params *GetPagesParams, opts ...Option) ([]*Page, error) {
	pages := []*Page{}
	if err := client.Get("/pages", params, &pages, opts...); err != nil {
		return nil, err
	}
	return pages, nil
}

type PostPagesParams struct {
	AuthRequired bool   `json:"auth_required"`
	Content      string `json:"content"`
	Draft        bool   `json:"draft"`
	Format       string `json:"format"`
	Hidden       bool   `json:"hidden"`
	Nonce        string `json:"nonce"` // XXX API should not use the nonce as you need to parse HTML content to get it, going against the API workflow
	Route        string `json:"route"`
	Title        string `json:"title"`
}

// XXX find if could use constraint error on .route to get a shell using the PIN form on sqlalchemy.exc.IntegrityError
func (client *Client) PostPages(params *PostPagesParams, opts ...Option) (*Page, error) {
	page := &Page{}
	if err := client.Post("/pages", params, &page, opts...); err != nil {
		return nil, err
	}
	return page, nil
}

func (client *Client) GetPage(id string, opts ...Option) (*Page, error) {
	page := &Page{}
	if err := client.Get("/pages/"+id, nil, &page, opts...); err != nil {
		return nil, err
	}
	return page, nil
}

func (client *Client) DeletePage(id string, opts ...Option) error {
	return client.Delete("/pages/"+id, nil, nil, opts...)
}

type PatchPageParams struct {
	Title        string `json:"title"`
	Content      string `json:"content"`
	Format       string `json:"format"`
	Route        string `json:"route"`
	Nonce        string `json:"nonce"` // XXX API should not use the nonce as you need to parse HTML content to get it, going against the API workflow
	AuthRequired bool   `json:"auth_required"`
	Draft        bool   `json:"draft"`
	Hidden       bool   `json:"hidden"`
}

func (client *Client) PatchPage(id string, params *PatchPageParams, opts ...Option) (*Page, error) {
	page := &Page{}
	if err := client.Patch("/pages/"+id, params, &page, opts...); err != nil {
		return nil, err
	}
	return page, nil
}
