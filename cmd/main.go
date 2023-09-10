package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bepped/hello-api/handlers/rest"
	"github.com/bepped/hello-api/translation"
)

func main() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if addr == ":" {
		addr = ":8080"
	}

	mux := http.NewServeMux()

	translationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(translationService)
	mux.HandleFunc("/hello", translateHandler.TranslateHandler)

	srv := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           mux,
	}
	log.Printf("listening on %s\n", addr)

	log.Fatal(srv.ListenAndServe())
}
