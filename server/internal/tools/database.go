package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	AuthToken string
	Username string
}

type UsersFavouriteAthleteDetails struct {
	AthleteName string
	Username string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetFavouriteAthlete(username string) *UsersFavouriteAthleteDetails
	SetUpDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {

	var database DatabaseInterface = &mockDB{}

	var err error = database.SetUpDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}