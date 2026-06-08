package main

import (
	"api/api/src/router"
	"api/api/src/config"
	"log"
	"net/http"
	"strconv"
)

func main()  {
	router := router.Generate()

	config.Load()
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.PORT), router))
}
