package goctfd

type (
	Challenge struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Description    string `json:"description"`
		ConnectionInfo string `json:"connection_info"`
		NextID         int    `json:"next_id"`
		MaxAttempts    int    `json:"max_attempts"`
		Value          int    `json:"value"`
		Initial        *int   `json:"initial,omitempty"`
		Decay          *int   `json:"decay,omitempty"`
		Minimum        *int   `json:"minimum,omitempty"`
		Category       string `json:"category"`
		Type           string `json:"type"`
		TypeDate       *Type  `json:"type_data,omitempty"`
		State          string `json:"state"`
		Requirements   any    `json:"requirements"` // TODO find model
		Solves         int    `json:"solves"`
		SolvedByMe     bool   `json:"solved_by_me"`
	}

	Type struct {
		ID        *string `json:"id,omitempty"`
		Name      string  `json:"name"`
		Templates CUV     `json:"templates"`
		Scripts   *CUV    `json:"scripts,omitempty"`
		Create    *string `json:"create,omitempty"`
	}

	CUV struct {
		Create string  `json:"create"`
		Update string  `json:"update"`
		View   *string `json:"view,omitempty"`
	}

	File struct {
		ID       int    `json:"id"`
		Type     string `json:"type"`
		Location string `json:"location"`
	}

	Flag struct {
		ID          int    `json:"id"`
		ChallengeID int    `json:"challenge_id"`
		Content     string `json:"content"`
		Data        any    `json:"data"` // TODO find an example
		Type        string `json:"type"`
		Challenge   int    `json:"challenge"` // XXX may be duplicated with ChallengeID ?
	}

	Hint struct {
		ID           int           `json:"id"`
		HTML         *string       `json:"html,omitempty"`
		Content      *string       `json:"content,omitempty"`
		Cost         int           `json:"cost"`
		Type         string        `json:"type"`
		ChallengeID  int           `json:"challenge_id"`
		Challenge    int           `json:"challenge"` // XXX may be duplicated with ChallengeID ?
		Requirements *Requirements `json:"requirements,omitempty"`
	}

	Requirements struct {
		Prerequisites []int `json:"prerequisites"`
	}

	Tag struct {
		ID          int    `json:"id"`
		Challenge   *int   `json:"challenge,omitempty"` // XXX This may be duplicated with ChallengeID ?
		ChallengeID int    `json:"challenge_id"`
		Value       string `json:"value"`
	}

	Topic struct {
		ID          int    `json:"id"`
		ChallengeID *int   `json:"challenge_id,omitempty"`
		Challenge   *int   `json:"challenge,omitempty"` // XXX may be duplicated with ChallengeID ?
		TopicID     *int   `json:"topic_id,omitempty"`  // XXX may be duplicated with ID ?
		Topic       *int   `json:"topic,omitempty"`     // XXX may be duplicated with ID ?
		Value       string `json:"value"`
	}

	Award struct {
		ID           int           `json:"id"`
		TeamID       int           `json:"team_id"`
		Category     *string       `json:"category"`
		UserID       int           `json:"user_id"`
		Team         int           `json:"team"` // XXX may be duplicated with ID ?
		Date         string        `json:"date"`
		Description  *string       `json:"description"`
		User         int           `json:"user"` // XXX may be duplicated with UserID ?
		Type         string        `json:"type"`
		Value        int           `json:"value"`
		Requirements *Requirements `json:"requirements"`
		Name         string        `json:"name"`
		Icon         string        `json:"icon"`
	}

	Submission struct {
		ID          int    `json:"id"`
		TeamID      int    `json:"team_id"` // XXX may be duplicated with team.id ?
		IP          string `json:"ip"`
		ChallengeID int    `json:"challenge_id"`
		UserID      int    `json:"user_id"` // XXX may be duplicated with user.id ?
		Team        struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"team"`
		Date string `json:"date"`
		User struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"user"`
		Type      string `json:"type"`
		Challenge struct {
			Value    int    `json:"value"`
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Category string `json:"category"`
		} `json:"challenge"`
		Provided string `json:"provided"`
	}

	Scoreboard struct {
		Pos         int    `json:"pos"`
		AccountId   int    `json:"account_id"`
		AccountURL  string `json:"account_url"`
		AccountType string `json:"account_type"`
		OauthID     any    `json:"oauth_id"`
		Name        string `json:"name"`
		Score       int    `json:"score"`
		Members     []struct {
			ID      int    `json:"id"`
			OauthID any    `json:"oauth_id"`
			Name    string `json:"name"`
			Score   int    `json:"score"`
		} `json:"members"`
	}

	Team struct {
		Bracket     *string `json:"bracket"`
		Members     []int   `json:"members,omitempty"`
		ID          int     `json:"id"`
		Created     string  `json:"created"`
		Country     *string `json:"country"`
		Email       *string `json:"email"`
		Affiliation *string `json:"affiliation"`
		CaptainID   *int    `json:"captain_id"`
		Fields      []Field `json:"fields"`
		Banned      bool    `json:"banned"`
		Website     *string `json:"website"`
		Hidden      bool    `json:"hidden"`
		Secret      *bool   `json:"secret"`
		Name        string  `json:"name"`
		OauthID     *string `json:"oauth_id"`
		Place       *string `json:"place,omitempty"`
		Score       *int    `json:"score,omitempty"`
	}

	User struct {
		Bracket     *string `json:"bracket"`
		ID          int     `json:"id"`
		TeamID      *int    `json:"team_id"`
		Country     *string `json:"country"`
		Affiliation *string `json:"affiliation"`
		Fields      []Field `json:"fields"`
		Website     *string `json:"website"`
		Name        string  `json:"name"`
		Email       *string `json:"email,omitempty"`
		OauthID     *string `json:"oauth_id"`
		Verified    *bool   `json:"verified,omitempty"`
		Banned      *bool   `json:"banned,omitempty"`
		Hidden      *bool   `json:"hidden,omitempty"`
		Type        *string `json:"type,omitempty"`
		Created     *string `json:"created,omitempty"`
		Secret      *string `json:"secret,omitempty"`
	}

	StatChallSubmission struct {
		ID         int      `json:"id"`
		Name       string   `json:"name"`
		Solves     *int     `json:"solves,omitempty"`
		Percentage *float64 `json:"percentage,omitempty"`
	}

	Distribution struct {
		Brackets map[string]int `json:"brackets"`
	}

	StatTeams struct {
		Registered int `json:"registered"`
	}

	StatUsers struct {
		Registered int `json:"registered"`
		Confirmed  int `json:"confirmed"`
	}

	Notification struct {
		ID      int     `json:"id"`
		TeamID  *int    `json:"team_id"` // XXX may be duplicated with Team ?
		HTML    string  `json:"html"`
		UserID  *int    `json:"user_id"` // XXX may be duplicated with User ?
		Team    *int    `json:"team"`
		Content string  `json:"content"`
		Date    string  `json:"date"`
		Title   string  `json:"title"`
		User    *int    `json:"user"`
		Type    *string `json:"type,omitempty"`
		Sound   *bool   `json:"sound,omitempty"`
	}

	Config struct {
		ID    int    `json:"id"`
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	ConfigField struct {
		FieldType   any     `json:"field_type"`
		Editable    bool    `json:"editable"`
		Type        string  `json:"type"`
		Required    bool    `json:"required"`
		Public      bool    `json:"public"`
		Name        *string `json:"name"`
		Description *string `json:"description"`
		ID          int     `json:"id"`
	}

	Field struct {
		FieldID int    `json:"field_id"`
		Value   string `json:"value"` // seems could be int/bool/string, let CTFd interfer type
	}

	Page struct {
		Files        []any   `json:"files"` // XXX find model
		ID           int     `json:"id"`
		Draft        bool    `json:"draft"`
		Route        string  `json:"route"`
		Title        string  `json:"title"`
		Hidden       bool    `json:"hidden"`
		Format       string  `json:"format"`
		AuthRequired bool    `json:"auth_required"`
		Content      *string `json:"content,omitempty"`
	}

	Unlock struct {
		Type   string `json:"type"`
		TeamID int    `json:"team_id"`
		Target int    `json:"target"`
		Date   string `json:"date"`
		UserID int    `json:"user_id"`
		ID     int    `json:"id"`
	}

	Token struct {
		Expiration string  `json:"expiration"`
		ID         int     `json:"id"`
		Type       string  `json:"type"`
		Value      *string `json:"value,omitempty"`
		Created    *string `json:"created,omitempty"`
		UserID     *int    `json:"user_id,omitempty"`
	}

	Comment struct {
		ID       int     `json:"id"`
		AuthorID int     `json:"author_id"`
		Content  *string `json:"content"`
		Date     string  `json:"date"`
		HTML     *string `json:"html,omitempty"`
		Author   struct {
			Name string `json:"name"`
		} `json:"author"`
		Type string `json:"type"`
	}

	Attempt struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
)
