package api

import (
	"fmt"
	"net/http"
)

//Downloader is the function to Dowload files
func Downloader(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Downloading file!")
}
