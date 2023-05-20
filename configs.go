package goctfd

type GetConfigsParams struct {
	Key   *string `schema:"key,omitempty"`
	Value *string `schema:"value,omitempty"`
	Q     *string `schema:"q,omitempty"`
	Field *string `schema:"field,omitempty"`
}

func (client *Client) GetConfigs(params *GetConfigsParams, opts ...option) ([]*Config, error) {
	configs := []*Config{}
	if err := get(client, "/configs", params, &configs, opts...); err != nil {
		return nil, err
	}
	return configs, nil
}

func (client *Client) PatchConfigs(params map[string]any, opts ...option) error {
	return patch(client, "/configs", params, nil, opts...)
}

type PostConfigsParams struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (client *Client) PostConfigs(params *PostConfigsParams, opts ...option) (*Config, error) {
	config := &Config{}
	if err := post(client, "/configs", params, &config, opts...); err != nil {
		return nil, err
	}
	return config, nil
}

type GetConfigsFieldsParams struct {
	Type  *string `schema:"type,omitempty"`
	Q     *string `schema:"q,omitempty"`
	Field *string `schema:"field,omitempty"`
}

func (client *Client) GetConfigsFields(params *GetConfigsParams, opts ...option) ([]*ConfigField, error) {
	fields := []*ConfigField{}
	if err := get(client, "/configs/fields", params, &fields, opts...); err != nil {
		return nil, err
	}
	return fields, nil
}

type PostConfigFieldsParams struct {
	ID          float64 `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	FieldType   string  `json:"field_type"`
	Editable    bool    `json:"editable"`
	Public      bool    `json:"public"`
	Required    bool    `json:"required"`
	Type        string  `json:"type"`
}

func (client *Client) PostConfigFields(params *PostConfigFieldsParams, opts ...option) (*ConfigField, error) {
	field := &ConfigField{}
	if err := post(client, "/configs/fields", params, &field, opts...); err != nil {
		return nil, err
	}
	return field, nil
}

func (client *Client) GetConfigsField(id string, opts ...option) (*ConfigField, error) {
	field := &ConfigField{}
	if err := get(client, "/configs/fields/"+id, nil, &field, opts...); err != nil {
		return nil, err
	}
	return field, nil
}

func (client *Client) DeleteConfigsField(id string, opts ...option) error {
	return delete(client, "/configs/fields/"+id, nil, nil, opts...)
}

type PatchConfigsFieldParams struct {
	ID          int    `json:"id"` // XXX duplicated with the ID in URL
	Name        string `json:"name"`
	Description string `json:"description"`
	FieldType   string `json:"field_type"`
	Type        string `json:"type"`
	Editable    bool   `json:"editable"`
	Public      bool   `json:"public"`
	Required    bool   `json:"required"`
}

func (client *Client) PatchConfigsField(id string, params PatchConfigsFieldParams, opts ...option) (*ConfigField, error) {
	field := &ConfigField{}
	if err := patch(client, "/configs/fields/"+id, params, &field, opts...); err != nil {
		return nil, err
	}
	return field, nil
}

// TODO find model
func (client *Client) GetConfigsByKey(key string, opts ...option) (any, error) {
	var config any
	if err := get(client, "/configs/"+key, nil, &config, opts...); err != nil {
		return nil, err
	}
	return config, nil
}

// TODO confirm delete does not take parameters and returns anything
func (client *Client) DeleteConfigsByKey(key string, opts ...option) error {
	return delete(client, "/configs/"+key, nil, nil, opts...)
}

// TODO find input model
func (client *Client) PatchConfigsByKey(key string, params any, opts ...option) (any, error) {
	var config any
	if err := patch(client, "/configs/"+key, params, &config, opts...); err != nil {
		return nil, err
	}
	return config, nil
}
