package webhooksite

// The default endpoint, but obviously you can use any self-hosted instance as
// well.
const Endpoint = "https://webhook.site"

type Client struct {
	endpoint string
}

func New(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}
