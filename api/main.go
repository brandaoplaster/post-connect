package main

import (
	"api/api/src/router"
	"log"
	"net/http"
)

func main()  {
	router := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", router))
}
