package handlers

import (
	"encoding/json"
	"net/http"
	"simple-crud/api/handlers/utils"
	"simple-crud/api/handlers/validators"
	"simple-crud/models"
	"simple-crud/services"
)

type UserHandler struct {
	service   *services.UserService
	validator *validators.UserValidator
}

func NewUserHandler(service *services.UserService, validator *validators.UserValidator) *UserHandler {
	return &UserHandler{service: service, validator: validator}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendJson(
			w,
			utils.Response{Error: "Invalid Body"},
			http.StatusUnprocessableEntity,
		)

		return
	}

	messages, isValid := h.validator.ValidateRequiredField(user)
	if !isValid {
		utils.SendJson(
			w,
			utils.Response{Error: messages},
			http.StatusBadRequest,
		)
		return
	}

	id := h.service.CreateUser(user)
	w.Header().Set("Content-Type", "application/json")
	utils.SendJson(
		w,
		utils.Response{Data: map[string]string{"id": string(id)}},
		http.StatusCreated,
	)
}
