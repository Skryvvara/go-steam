package steam

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Get the owned games by user ID.
func (a SteamApiClient) GetOwnedGames(userID string) ([]SteamGame, error) {
	url := fmt.Sprintf("%s/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json",
		a.ApiBaseUrl,
		a.ApiKey,
		userID,
	)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()
	var games gameListResponseJSON
	if err = json.NewDecoder(resp.Body).Decode(&games); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return games.Response.Games, nil
}

// Get a game's details by it's ID.
func (a SteamApiClient) GetGameDetails(gameID string) (*GameDetails, error) {
	url := fmt.Sprintf("%s/api/appdetails?appids=%s", a.StoreBaseUrl, gameID)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()
	var details map[string]Response
	if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
		fmt.Println(err)
		return nil, err
	}

	gameDetails := details[gameID].Data
	return &gameDetails, nil
}
