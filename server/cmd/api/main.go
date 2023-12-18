package main

import (
	"fmt"
	"net/http"
	"github.com/StunnaDawg/Head-To-Head/internal/handlers"
    "github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

func main() {

	var r * chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting Server")

	err := http.ListenAndServe("localhost:4000", r)

	if err != nil {
		log.Error(err)
	}
}