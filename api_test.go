package steam_test

import (
	"os"
	"testing"

	"github.com/skryvvara/go-steam"
)

var apiClient = steam.NewApiClient(os.Getenv("API_KEY"))
var testUserID = "76561198272838960"
var testGame = struct {
	ID   string
	Name string
}{
	ID:   "1245620",
	Name: "ELDEN RING",
}

func TestGetFriendList(t *testing.T) {
	friends, err := apiClient.GetFriendsList(testUserID)
	if err != nil {
		t.Error(err)
	}

	if len(friends) == 0 {
		t.Error("Received 0 friends for test user")
	}
}

func TestGetOwnedGames(t *testing.T) {
	games, err := apiClient.GetOwnedGames(testUserID)
	if err != nil {
		t.Error(err)
	}

	if len(games) == 0 {
		t.Error("Received 0 games for test user")
	}
}

func TestGetGameDetails(t *testing.T) {
	details, err := apiClient.GetGameDetails(testGame.ID)
	if err != nil {
		t.Error(err)
	}

	if details.Name != testGame.Name {
		t.Error("Received incorrect details for ID " + testGame.ID)
	}
}
