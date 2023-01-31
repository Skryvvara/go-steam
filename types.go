package steam

import "encoding/json"

type SteamApiClient struct {
	ApiKey       string
	ApiBaseUrl   string
	StoreBaseUrl string
}

type SteamFriend struct {
	UserID       string `json:"steamid"`
	Relationship string
	FriendSince  int64 `json:"friend_since"`
}

type playerFriendsListJSON struct {
	FriendsList *struct {
		Friends []SteamFriend
	}
}

type SteamGame struct {
	GameID          json.Number `json:"appid"`
	Playtime2Weeks  int64       `json:"playtime_2weeks"`
	PlaytimeTotal   int64       `json:"playtime_forever"`
	PlaytimeWindows int64       `json:"playtime_windows_forever"`
	PlaytimeMac     int64       `json:"playtime_mac_forever"`
	PlaytimeLinux   int64       `json:"playtime_linux_forever"`
}

type gameListResponseJSON struct {
	Response *struct {
		Count int `json:"game_count"`
		Games []SteamGame
	}
}

type GameDetails struct {
	Id                 json.Number `json:"steam_appid"`
	Name               string      `json:"name"`
	Type               string      `json:"type"`
	IsFree             bool        `json:"is_free"`
	ShortDescription   string      `json:"short_description"`
	HeaderImageUrl     string      `json:"header_image"`
	Website            string      `json:"website"`
	Developers         []string    `json:"developers"`
	Publishers         []string    `json:"publishers"`
	Platforms          Platforms   `json:"platforms"`
	Rating             Metacritic  `json:"metacritic"`
	Categories         []Tag       `json:"categories"`
	Genres             []Tag       `json:"genres"`
	BackgroundImageUrl string      `json:"background"`
	StoreUrl           string      `default:""`
	IsHidden           bool        `default:"false"`
}

type Platforms struct {
	Windows bool `json:"windows"`
	Mac     bool `json:"mac"`
	Linux   bool `json:"linux"`
}

type Metacritic struct {
	Score float32 `json:"score"`
	Url   string  `json:"url"`
}

type Tag struct {
	Id   json.Number `json:"id"`
	Name string      `json:"description"`
}

type Response struct {
	Success bool        `json:"success"`
	Data    GameDetails `json:"data"`
}
