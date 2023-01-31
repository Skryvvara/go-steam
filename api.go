package steam

// Returns a new SteamApiClient.
// To get an API-Key visit https://steamcommunity.com/dev/apikey.
func NewApiClient(apiKey string) SteamApiClient {
	apiClient := SteamApiClient{
		ApiKey:       apiKey,
		ApiBaseUrl:   "https://api.steampowered.com",
		StoreBaseUrl: "https://store.steampowered.com",
	}
	return apiClient
}
