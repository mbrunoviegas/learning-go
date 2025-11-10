package api

import (
	"errors"
	"go-bid/internal/jsonutils"
	"go-bid/internal/services"
	"go-bid/internal/usecase/user"
	"net/http"
)

func (api *Api) handleSingupUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[user.CreateUserReq](r)

	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := api.UserService.CreateUser(
		r.Context(),
		data.UserName,
		data.Email,
		data.Password,
		data.Bio,
	)

	if err != nil {
		if errors.Is(err, services.ErrorDuplicatedEmailOrUsername) {
			_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{
				"error": "duplicated email ou username",
			})
			return
		}

		_ = jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "internal server error",
		})
		return
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{
		"id": id,
	})
}

func (api *Api) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[user.LoginUserReq](r)
	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := api.UserService.AuthUser(r.Context(), data.Email, data.Password)

	if err != nil {
		if errors.Is(err, services.ErrorInvalidCredentials) {
			jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{
				"error": "invalid email or password",
			})

			return
		}

		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "internal server error",
		})

		return
	}

	err = api.Sessions.RenewToken(r.Context())
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "internal server error",
		})

		return
	}

	api.Sessions.Put(r.Context(), "AuthenticatedUserId", id)

	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"message": "success",
	})
}

func (api *Api) handleLogoutUser(w http.ResponseWriter, r *http.Request) {
	err := api.Sessions.RenewToken(r.Context())
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"error": "internal server error",
		})

		return
	}

	api.Sessions.Remove(r.Context(), "AuthenticatedUserId")

	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"message": "success",
	})
}
