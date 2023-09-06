package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bepped/hello-api/handlers/rest"
)

func main() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if addr == ":" {
		addr = ":8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", rest.TranslateHandler)

	srv := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           mux,
	}
	log.Printf("listening on %s\n", addr)

	log.Fatal(srv.ListenAndServe())
}
