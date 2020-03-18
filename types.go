package webhooksite

type TokenOptions struct {
	DefaultStatus      int    `json:"default_status,omitempty"`
	DefaultContent     string `json:"default_content,omitempty"`
	DefaultContentType string `json:"default_content_type,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
}

type Token struct {
	Redirect           bool        `json:"redirect"`
	Alias              interface{} `json:"alias"`
	Timeout            int         `json:"timeout"`
	Premium            bool        `json:"premium"`
	UUID               string      `json:"uuid"`
	IP                 string      `json:"ip"`
	UserAgent          string      `json:"user_agent"`
	DefaultContent     string      `json:"default_content"`
	DefaultStatus      int         `json:"default_status"`
	DefaultContentType string      `json:"default_content_type"`
	PremiumExpiresAt   string      `json:"premium_expires_at"`
	CreatedAt          string      `json:"created_at"`
	UpdatedAt          string      `json:"updated_at"`

	// Filled by Client, can be used to call the webhook
	URL string `json:"-"`
}

type Requests struct {
	Data        []*Request `json:"data"`
	Total       int        `json:"total"`
	PerPage     int        `json:"per_page"`
	CurrentPage int        `json:"current_page"`
	IsLastPage  bool       `json:"is_last_page"`
	From        int        `json:"from"`
	To          int        `json:"to"`
}

type Request struct {
	UUID      string  `json:"uuid"`
	TokenID   string  `json:"token_id"`
	IP        string  `json:"ip"`
	Hostname  string  `json:"hostname"`
	Method    string  `json:"method"`
	UserAgent string  `json:"user_agent"`
	Content   string  `json:"content"`
	Query     Query   `json:"query"`
	Headers   Headers `json:"headers"`
	URL       string  `json:"url"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type Query struct {
	Action string `json:"action"`
}

type Headers struct {
	ContentLength []string `json:"content-length"`
	UserAgent     []string `json:"user-agent"`
}
