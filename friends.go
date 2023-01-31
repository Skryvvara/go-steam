package steam

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SteamFriend struct {
	UserID       string `json:"steamid"`
	Relationship string
	FriendSince  int64 `json:"friend_since"`
}

type playerFriendsListJSON struct {
	Friendslist *struct {
		Friends []SteamFriend
	}
}

func (a SteamApiClient) GetFriendsList(userID string) ([]SteamFriend, error) {
	url := fmt.Sprintf("%s/ISteamUser/GetFriendList/v0001/?key=%s&steamid=%s&relationship=friend",
		a.BaseURL,
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
	err = json.NewDecoder(resp.Body).Decode(&friends)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return friends.Friendslist.Friends, nil
}
