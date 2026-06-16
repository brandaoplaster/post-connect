package responses

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"api/api/src/models"
	"github.com/gorilla/mux"
)

func ParseUserID(r *http.Request) (uint64, error) {
	id, err := strconv.ParseUint(mux.Vars(r)["userId"], 10, 64)
	if err != nil {
		return 0, errors.New("ID inválido")
	}
	return id, nil
}

func DecodeUser(r *http.Request) (models.User, error) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return user, errors.New("JSON inválido")
	}
	return user, nil
}
