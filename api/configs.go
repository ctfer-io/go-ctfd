package api

type GetConfigsParams struct {
	Key   *string `schema:"key,omitempty"`
	Value *string `schema:"value,omitempty"`
	Q     *string `schema:"q,omitempty"`
	Field *string `schema:"field,omitempty"`
}

func (client *Client) GetConfigs(params *GetConfigsParams, opts ...Option) ([]*Config, error) {
	configs := []*Config{}
	if err := client.Get("/configs", params, &configs, opts...); err != nil {
		return nil, err
	}
	return configs, nil
}

type PatchConfigsParams struct {
	// Appearance

	CTFDescription *string `json:"ctf_description,omitempty"`
	CTFName        *string `json:"ctf_name,omitempty"`

	// Theme

	CTFTheme      *string `json:"ctf_theme,omitempty"`
	ThemeFooter   *string `json:"theme_footer,omitempty"`
	ThemeHeader   *string `json:"theme_header,omitempty"`
	ThemeSettings *string `json:"theme_settings,omitempty"`

	// Localization

	DefaultLocale *string `json:"default_locale,omitempty"`

	// Accounts

	DomainWhitelist            *string `json:"domain_whitelist,omitempty"`
	IncorrectSubmissionsPerMin *int    `json:"incorrect_submissions_per_min,omitempty"`
	NameChanges                *bool   `json:"name_changes,omitempty"`
	NumTeams                   *int    `json:"num_teams,omitempty"`
	NumUsers                   *int    `json:"num_users,omitempty"`
	TeamCreation               *bool   `json:"team_creation,omitempty"`
	TeamDisbanding             *string `json:"team_disbanding,omitempty"`
	TeamSize                   *int    `json:"team_size,omitempty"`
	VerifyEmails               *bool   `json:"verify_emails,omitempty"`

	// Pages

	RobotsTxt *string `json:"robots_txt,omitempty"`

	// MajorLeagueCyber

	OauthClientID     *string `json:"oauth_client_id,omitempty"`
	OauthClientSecret *string `json:"oauth_client_secret,omitempty"`

	// Settings

	AccountVisibility      *string `json:"account_visibility,omitempty"`
	ChallengeVisibility    *string `json:"challenge_visibility,omitempty"`
	Paused                 *bool   `json:"paused,omitempty"`
	RegistrationVisibility *string `json:"registration_visibility,omitempty"`
	ScoreVisibility        *string `json:"score_visibility,omitempty"`

	// Security

	HTMLSanitization *bool   `json:"html_sanitization,omitempty"`
	RegistrationCode *string `json:"registration_code,omitempty"`

	// Email

	MailPassword                       *string `json:"mail_password,omitempty"`
	MailPort                           *string `json:"mail_port,omitempty"`
	MailServer                         *string `json:"mail_server,omitempty"`
	MailSSL                            *bool   `json:"mail_ssl,omitempty"`
	MailTLS                            *bool   `json:"mail_tls,omitempty"`
	MailUseAuth                        *bool   `json:"mail_useauth,omitempty"`
	MailUsername                       *string `json:"mail_username,omitempty"`
	PasswordChangeAlertBody            *string `json:"password_change_alert_body,omitempty"`
	PasswordChangeAlertSubject         *string `json:"password_change_alert_subject,omitempty"`
	PasswordResetBody                  *string `json:"password_reset_body,omitempty"`
	PasswordResetSubject               *string `json:"password_reset_subject,omitempty"`
	SuccessfulRegistrationEmailBody    *string `json:"successful_registration_email_body,omitempty"`
	SuccessfulRegistrationEmailSubject *string `json:"successful_registration_email_subject,omitempty"`
	UserCreationEmailBody              *string `json:"user_creation_email_body,omitempty"`
	UserCreationEmailSubject           *string `json:"user_creation_email_subject,omitempty"`
	VerificationEmailBody              *string `json:"verification_email_body,omitempty"`
	VerificationEmailSubject           *string `json:"verification_email_subject,omitempty"`

	// DEPRECATED
	MailFromAddr *string `json:"mailfrom_addr,omitempty"`
	// DEPRECATED
	MailGunAPIKey *string `json:"mailgun_api_key,omitempty"`
	// DEPRECATED
	MailGunBaseURL *string `json:"mailgun_base_url,omitempty"`

	// Time

	End          *string `json:"end,omitempty"`
	Freeze       *string `json:"freeze,omitempty"`
	Start        *string `json:"start,omitempty"`
	ViewAfterCTF *bool   `json:"view_after_ctf,omitempty"`

	// Social

	SocialShares *bool `json:"social_shares,omitempty"`

	// Legal

	PrivacyText *string `json:"privacy_text,omitempty"`
	PrivacyURL  *string `json:"privacy_url,omitempty"`
	TOSText     *string `json:"tos_text,omitempty"`
	TOSURL      *string `json:"tos_url,omitempty"`

	// User Mode

	UserMode *string `json:"user_mode,omitempty"`
}

func (client *Client) PatchConfigs(params *PatchConfigsParams, opts ...Option) error {
	return client.Patch("/configs", params, nil, opts...)
}

type PostConfigsParams struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (client *Client) PostConfigs(params *PostConfigsParams, opts ...Option) (*Config, error) {
	config := &Config{}
	if err := client.Post("/configs", params, &config, opts...); err != nil {
		return nil, err
	}
	return config, nil
}

type GetConfigsFieldsParams struct {
	Type  *string `schema:"type,omitempty"`
	Q     *string `schema:"q,omitempty"`
	Field *string `schema:"field,omitempty"`
}

func (client *Client) GetConfigsFields(params *GetConfigsParams, opts ...Option) ([]*ConfigField, error) {
	fields := []*ConfigField{}
	if err := client.Get("/configs/fields", params, &fields, opts...); err != nil {
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

func (client *Client) PostConfigFields(params *PostConfigFieldsParams, opts ...Option) (*ConfigField, error) {
	field := &ConfigField{}
	if err := client.Post("/configs/fields", params, &field, opts...); err != nil {
		return nil, err
	}
	return field, nil
}

func (client *Client) GetConfigsField(id string, opts ...Option) (*ConfigField, error) {
	field := &ConfigField{}
	if err := client.Get("/configs/fields/"+id, nil, &field, opts...); err != nil {
		return nil, err
	}
	return field, nil
}

func (client *Client) DeleteConfigsField(id string, opts ...Option) error {
	return client.Delete("/configs/fields/"+id, nil, nil, opts...)
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

func (client *Client) PatchConfigsField(id string, params *PatchConfigsFieldParams, opts ...Option) (*ConfigField, error) {
	field := &ConfigField{}
	if err := client.Patch("/configs/fields/"+id, params, &field, opts...); err != nil {
		return nil, err
	}
	return field, nil
}

// TODO find model
func (client *Client) GetConfigsByKey(key string, opts ...Option) (any, error) {
	var config any
	if err := client.Get("/configs/"+key, nil, &config, opts...); err != nil {
		return nil, err
	}
	return config, nil
}

// TODO confirm delete does not take parameters and returns anything
func (client *Client) DeleteConfigsByKey(key string, opts ...Option) error {
	return client.Delete("/configs/"+key, nil, nil, opts...)
}

// TODO find input model
func (client *Client) PatchConfigsByKey(key string, params any, opts ...Option) (any, error) {
	var config any
	if err := client.Patch("/configs/"+key, params, &config, opts...); err != nil {
		return nil, err
	}
	return config, nil
}

type PatchConfigsCTFLogo struct {
	Value *string `json:"value"`
}

func (client *Client) PatchConfigsCTFLogo(params *PatchConfigsCTFLogo, opts ...Option) (*ThemeImage, error) {
	var ti *ThemeImage
	if err := client.Patch("/configs/ctf_logo", params, ti, opts...); err != nil {
		return nil, err
	}
	return ti, nil
}

func (client *Client) PatchConfigsCTFSmallIcon(params *PatchConfigsCTFLogo, opts ...Option) (*ThemeImage, error) {
	var ti *ThemeImage
	if err := client.Patch("/configs/ctf_small_icon", params, ti, opts...); err != nil {
		return nil, err
	}
	return ti, nil
}
