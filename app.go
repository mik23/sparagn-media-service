package main

import (
	"sparagn.com/sparagn-media-service/api"
)

func main() {
	router := api.SetupRouter()
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


