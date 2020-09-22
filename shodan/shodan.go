package shodan

// BaseURL is for Shodan API
const BaseURL = "https://api.shodan.io"

// Client is the API key for Shodan
type Client struct {
	apiKey string
}

// New creates new Client for Shodan API
func New(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}
