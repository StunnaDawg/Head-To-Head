package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"jonson" : {
		AuthToken: "123ABC",
		Username: "stunnacat",
	},
	"johnny" : {
		AuthToken: "1234ABC",
		Username: "stunnadawg",
	},
	"jinks" : {
		AuthToken: "ABC123",
		Username: "stunnaloser",
	},
}

var mockFavouriteAthlete = map[string]UsersFavouriteAthleteDetails{
	"jonson" : {
		AthleteName: "Shohei",
		Username: "stunnacat",
	},
	"johnny" : {
		AthleteName: "Loser",
		Username: "stunnadawg",
	},
	"jinks" : {
		AthleteName: "Marner",
		Username: "stunnaloser",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second *1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetFavouriteAthlete(AthleteName string) *UsersFavouriteAthleteDetails {
	time.Sleep(time.Second *1)

	var clientData = UsersFavouriteAthleteDetails{}
	clientData, ok := mockFavouriteAthlete[AthleteName]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetUpDatabase() error {
	return nil
}

