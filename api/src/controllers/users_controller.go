package controllers

import (
	"api/api/src/database"
	"api/api/src/models"
	"api/api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Create(write http.ResponseWriter, request *http.Request) {
	response, erro := io.ReadAll(request.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User
	if erro = json.Unmarshal(response, &user); erro != nil {
		log.Fatal(erro)
		return
	}

	if erro = user.Prepare("Create"); erro != nil {
		log.Fatal(erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		log.Fatal(erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user.ID, erro = repository.Create(user)
	if erro != nil {
		log.Fatal(erro)
		return
	}

	write.Write([]byte(fmt.Sprintf("id insert %d", user.ID)))
}

func Index(write http.ResponseWriter, request *http.Request) {

}

func Show(write http.ResponseWriter, request *http.Request) {

}

func Update(write http.ResponseWriter, request *http.Request) {

}

func Delete(write http.ResponseWriter, request *http.Request) {

}
