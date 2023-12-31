package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/StunnaDawg/Head-To-Head/api"
	"github.com/StunnaDawg/Head-To-Head/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetFavouriteAthlete(w http.ResponseWriter, r *http.Request) {
	var params = api.ChosenAthleteParam{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err!= nil {
		log.Error(err)
		api.InternalErrorHandler(w, err)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err!= nil {
		log.Error(err)
		api.InternalErrorHandler(w, err)
		return
	}

	var tokenDetails *tools.UsersFavouriteAthleteDetails
	tokenDetails = (*database).GetFavouriteAthlete(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w, err)
		return
	}

	var response = api.AtheleteNameResponse{
		Name: (*tokenDetails).AthleteName,
		Code: http.StatusOK,
	}
	
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w, err)
		return
	}
}