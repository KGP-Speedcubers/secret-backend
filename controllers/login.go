package controllers

import (
	"kgpsc-backend/middleware"
	"kgpsc-backend/models"
	"kgpsc-backend/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type userLoginReqFields struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	app := r.Context().Value(middleware.APP_CTX_KEY).(*middleware.App)
	db := app.DB

	reqFields := userLoginReqFields{}
	err := utils.DecodeJSON(r, &reqFields)
	if err != nil {
		utils.LogErrAndRespond(r, w, err, "Invalid request payload.", http.StatusBadRequest)
		return
	}

	user := models.User{}
	tx := db.Table("users").
		Where("username = ?", reqFields.Username).
		First(&user)

	if tx.Error != nil {
		utils.LogErrAndRespond(r, w, tx.Error, "Error fetching user.", http.StatusInternalServerError)
		return
	}

	if tx.RowsAffected == 0 {
		utils.RespondWithHTTPMessage(r, w, http.StatusUnauthorized, "Invalid credentials.")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqFields.Password))
	if err != nil {
		utils.RespondWithHTTPMessage(r, w, http.StatusUnauthorized, "Invalid credentials.")
		return
	}

	utils.RespondWithHTTPMessage(r, w, http.StatusOK, "Logged in successfully.")
}
