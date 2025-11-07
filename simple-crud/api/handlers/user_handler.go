package handlers

import (
	"encoding/json"
	"net/http"
	"simple-crud/api/handlers/utils"
	"simple-crud/api/handlers/validators"
	"simple-crud/models"
	"simple-crud/services"

	"github.com/go-chi/chi/v5"
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

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	user, exists := h.service.GetUserById(models.ID(userId))

	if !exists {
		utils.SendJson(
			w,
			utils.Response{Error: "user not found"},
			http.StatusNotFound,
		)

		return
	}

	utils.SendJson(
		w,
		utils.Response{Data: user},
		http.StatusOK,
	)
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users := h.service.ListUsers()

	w.Header().Set("Content-Type", "application/json")
	utils.SendJson(
		w,
		utils.Response{Data: users},
		http.StatusOK,
	)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateData models.User
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		utils.SendJson(
			w,
			utils.Response{Error: "Invalid Body"},
			http.StatusUnprocessableEntity,
		)

		return
	}

	messages, isValid := h.validator.ValidateRequiredField(updateData)
	if !isValid {
		utils.SendJson(
			w,
			utils.Response{Error: messages},
			http.StatusBadRequest,
		)
		return
	}

	id := chi.URLParam(r, "id")

	updatedUser, exists := h.service.UpdateUser(models.ID(id), updateData)

	if !exists {
		utils.SendJson(
			w,
			utils.Response{Error: "user not found"},
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.SendJson(
		w,
		utils.Response{Data: updatedUser},
		http.StatusCreated,
	)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	deletedUser, exists := h.service.DeleteUser(models.ID(userId))

	if !exists {
		utils.SendJson(
			w,
			utils.Response{Error: "user not found"},
			http.StatusNotFound,
		)

		return
	}

	utils.SendJson(
		w,
		utils.Response{Data: deletedUser},
		http.StatusOK,
	)
}
