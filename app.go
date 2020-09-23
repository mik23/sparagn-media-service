package main

import (
	"log"
	"net/http"

	"sparagn.com/sparagn-media-service/api"
)

func handleRequests() {
	http.HandleFunc("/upload", api.Upload)
	http.HandleFunc("/downloader", api.Downloader)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
