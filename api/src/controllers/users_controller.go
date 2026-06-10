package controllers

import (
	"api/api/src/database"
	"api/api/src/models"
	"api/api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	db, erro := database.Connect()
	if erro != nil {
		http.Error(write, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	users, erro := repository.All()
	if erro != nil {
		http.Error(write, erro.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(write).Encode(users)
}

func Show(write http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)

	userId, error := strconv.ParseUint(parameters["userId"], 10, 64)

	if error != nil {
		http.Error(write, "ID inválido", http.StatusBadRequest)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		http.Error(write, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user, erro := repository.FindById(userId)
	if erro != nil {
		http.Error(write, erro.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(write).Encode(user)
}

func Update(write http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)

	userId, error := strconv.ParseUint(parameters["userId"], 10, 64)

	if error != nil {
		http.Error(write, "ID inválido", http.StatusBadRequest)
		return
	}

	body, erro := ioutil.ReadAll(request.Body)
	if erro != nil {
		http.Error(write, erro.Error(), http.StatusUnprocessableEntity)
		return
	}

	var user models.User
	if erro = json.Unmarshal(body, &user); erro != nil {
		http.Error(write, erro.Error(), http.StatusBadRequest)
		return
	}

	if erro = user.Prepare("update"); erro != nil {
		http.Error(write, erro.Error(), http.StatusBadRequest)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		http.Error(write, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user, erro = repository.Update(userId, user)
	if erro != nil {
		http.Error(write, erro.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(write).Encode(user)
}

func Delete(write http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)

	userId, error := strconv.ParseUint(parameters["userId"], 10, 64)

	if error != nil {
		http.Error(write, "ID inválido", http.StatusBadRequest)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		http.Error(write, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	erro = repository.Delete(userId)
	if erro != nil {
		if erro.Error() == "user not found" {
			http.Error(write, erro.Error(), http.StatusNotFound)
			return
		}
		http.Error(write, erro.Error(), http.StatusInternalServerError)
		return
	}

	write.WriteHeader(http.StatusNoContent)
}
