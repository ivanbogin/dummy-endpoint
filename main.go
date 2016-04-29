package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	uri := r.URL.Path

	switch uri {
	case "/500":
		status = http.StatusInternalServerError
	case "/wait3":
		time.Sleep(3 * time.Second)
	case "/wait5":
		time.Sleep(5 * time.Second)
	case "/wait60":
		time.Sleep(time.Minute)
	}

	w.WriteHeader(status)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", HttpHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
