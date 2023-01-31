package steam

type SteamApiClient struct {
	ApiKey  string
	BaseURL string
}

// Returns a new SteamApiClient.
// To get an API-Key visit https://steamcommunity.com/dev/apikey.
func NewApiClient(apiKey string) SteamApiClient {
	apiClient := SteamApiClient{
		ApiKey:  apiKey,
		BaseURL: "https://api.steampowered.com",
	}
	return apiClient
}
