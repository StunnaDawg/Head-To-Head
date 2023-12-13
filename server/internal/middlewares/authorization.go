package middlewares

import (
	"errors"
	"net/http"
	"fmt"

	"github.com/StunnaDawg/Head-To-Head/api"
	"github.com/StunnaDawg/Head-To-Head/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New(fmt.Sprintf("Invalid Credentials"))

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {	

		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" {
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w, UnAuthorizedError)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)

	})
}