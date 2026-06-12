package main

import (
	"api/api/src/router"
	"api/api/src/config"
	"api/api/src/database"
	"log"
	"net/http"
	"strconv"
)

func main()  {
	config.Load()

	database, erro := database.Connect()
	if erro != nil {
		log.Fatal(erro)
	}
	defer database.Close()

	handler := NewHandlers(database)
	router := router.Generate(handler)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.PORT), router))
}
