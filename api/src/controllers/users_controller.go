package controllers

import (
	"errors"
	"net/http"

	"api/api/src/responses"
	"api/api/src/services"
)

type UsersController struct {
	service services.UserService
}

func NewUsersController(service services.UserService) *UsersController {
	return &UsersController{service: service}
}

func (c *UsersController) Index(w http.ResponseWriter, r *http.Request) {
	users, err := c.service.List()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func (c *UsersController) Create(w http.ResponseWriter, r *http.Request) {
	user, err := responses.DecodeUser(r)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err = c.service.Create(user)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func (c *UsersController) Show(w http.ResponseWriter, r *http.Request) {
	id, err := responses.ParseUserID(r)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := c.service.Find(id)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func (c *UsersController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := responses.ParseUserID(r)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := responses.DecodeUser(r)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err = c.service.Update(id, user)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func (c *UsersController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := responses.ParseUserID(r)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.service.Delete(id); err != nil {
		handleServiceError(w, err)
		return
	}

	responses.NoContent(w)
}

func handleServiceError(w http.ResponseWriter, err error) {
	if errors.Is(err, services.ErrUserNotFound) {
		responses.Error(w, http.StatusNotFound, err.Error())
		return
	}
	responses.Error(w, http.StatusInternalServerError, err.Error())
}
