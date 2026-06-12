package controllers

import (
	"api/api/src/models"
	"api/api/src/repositories"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UsersController struct {
	repo *repositories.Users
}

func NewUsersController(repo *repositories.Users) *UsersController {
	return &UsersController{repo: repo}
}

func (userRepo *UsersController) Index(write http.ResponseWriter, request *http.Request) {
	users, erro := userRepo.repo.All()
	if erro != nil {
		http.Error(write, erro.Error(), http.StatusInternalServerError)
		return
	}
	write.WriteHeader(http.StatusOK)
    json.NewEncoder(write).Encode(users)
}

func (userRepo *UsersController) Create(write http.ResponseWriter, request *http.Request) {
	response, erro := io.ReadAll(request.Body)
	if erro != nil {
		http.Error(write, erro.Error(), http.StatusUnprocessableEntity)
		return
	}

	var user models.User
	if erro = json.Unmarshal(response, &user); erro != nil {
		http.Error(write, erro.Error(), http.StatusBadRequest)
		return
	}

	if erro = user.Prepare("Create"); erro != nil {
		http.Error(write, erro.Error(), http.StatusBadRequest)
		return
	}

	user, erro = userRepo.repo.Create(user)
	if erro != nil {
		http.Error(write, erro.Error(), http.StatusInternalServerError)
		return
	}

	write.WriteHeader(http.StatusCreated)
	json.NewEncoder(write).Encode(user)
}

func (userRepo *UsersController) Show(write http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)

	userId, error := strconv.ParseUint(parameters["userId"], 10, 64)

	if error != nil {
		http.Error(write, "ID inválido", http.StatusBadRequest)
		return
	}

	user, erro := userRepo.repo.FindById(userId)
	if erro != nil {
		http.Error(write, erro.Error(), http.StatusInternalServerError)
		return
	}

	write.WriteHeader(http.StatusOK)
	json.NewEncoder(write).Encode(user)
}

func (userRepo *UsersController) Update(write http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)

	userId, error := strconv.ParseUint(parameters["userId"], 10, 64)

	if error != nil {
		http.Error(write, "ID inválido", http.StatusBadRequest)
		return
	}

	body, erro := io.ReadAll(request.Body)
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

	user, erro = userRepo.repo.Update(userId, user)
	if erro != nil {
		http.Error(write, erro.Error(), http.StatusInternalServerError)
		return
	}

	write.WriteHeader(http.StatusOK)
	json.NewEncoder(write).Encode(user)
}

func (userRepo *UsersController) Delete(write http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)

	userId, error := strconv.ParseUint(parameters["userId"], 10, 64)

	if error != nil {
		http.Error(write, "ID inválido", http.StatusBadRequest)
		return
	}

	erro := userRepo.repo.Delete(userId)
	if erro != nil {
		http.Error(write, erro.Error(), http.StatusInternalServerError)
		return
	}

	write.WriteHeader(http.StatusNoContent)
}
