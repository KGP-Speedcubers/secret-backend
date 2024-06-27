package controllers

import (
	"encoding/json"
	"kgpsc-backend/middleware"
	"kgpsc-backend/models"
	"kgpsc-backend/utils"

	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	app := r.Context().Value(middleware.APP_CTX_KEY).(*middleware.App)
	db := app.DB

	users := []models.Users{}

	tx := db.
		Table("users").
		Find(&users)

	if tx.Error != nil {
		utils.LogErrAndRespond(r, w, tx.Error, "Error fetching users.", http.StatusInternalServerError)
		return
	}

	if len(users) == 0 {
		utils.RespondWithHTTPMessage(r, w, http.StatusOK, "No users exist.")
		return
	}

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		utils.LogErrAndRespond(r, w, err, "Error encoding users.", http.StatusInternalServerError)
		return
	}

	utils.RespondWithHTTPMessage(r, w, http.StatusOK, "Users fetched successfully.")
}
