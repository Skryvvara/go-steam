package gosteam

// https://developer.valvesoftware.com/wiki/Steam_Web_API

type SteamApi struct {
	ApiKey  string
	BaseURL string
}

func NewSteamApi(apiKey string) SteamApi {
	a := SteamApi{
		ApiKey:  apiKey,
		BaseURL: "https://api.steampowered.com",
	}
	return a
}
