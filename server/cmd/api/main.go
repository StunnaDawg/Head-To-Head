package main

import (
	"fmt"
	"net/http"
	"github.com/StunnaDawg/Head-To-Head/internal/handlers"
    "github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReporterCaller(true)
	var r * chi.Mux = chi.newRouter()
	handler.Handlers(r)

	fmt.Println("Starting Server")

	err := http.ListenAndServe("localhost:4000")

	if err != nil {
		log.Error(err)
	}
}