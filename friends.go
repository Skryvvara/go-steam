package steam

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Get the friend list of a user by his ID.
func (a SteamApiClient) GetFriendsList(userID string) ([]SteamFriend, error) {
	url := fmt.Sprintf("%s/ISteamUser/GetFriendList/v0001/?key=%s&steamid=%s&relationship=friend",
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
	var friends playerFriendsListJSON
	if err = json.NewDecoder(resp.Body).Decode(&friends); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return friends.FriendsList.Friends, nil
}
