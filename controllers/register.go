package controllers

import (
	"kgpsc-backend/middleware"
	"kgpsc-backend/models"
	"kgpsc-backend/utils"
	"net/http"
)

type RegisterUserReqFields struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	app := r.Context().Value(middleware.APP_CTX_KEY).(*middleware.App)
	db := app.DB
	reqFields := RegisterUserReqFields{}

	err := utils.DecodeJSON(r, &reqFields)
	if err != nil {
		utils.LogErrAndRespond(r, w, err, "Invalid request Payload.", http.StatusBadRequest)
		return
	}

	user := models.User{}
	tx := db.
		Table("users").
		Where("email = ?", reqFields.Email).
		Or("username = ?", reqFields.Username).
		First(&user)

	if tx.Error == nil {
		utils.RespondWithHTTPMessage(r, w, http.StatusConflict, "User already exists.")
		return
	}

	hashedPassword, err := utils.HashPassword(reqFields.Password)
	if err != nil {
		utils.LogErrAndRespond(r, w, err, "Error hashing password.", http.StatusInternalServerError)
		return
	}

	tx = db.
		Create(&models.User{
			Username: reqFields.Username,
			Email:    reqFields.Email,
			Password: hashedPassword,
		})

	if tx.Error != nil {
		utils.LogErrAndRespond(r, w, tx.Error, "Error creating user.", http.StatusInternalServerError)
		return
	}

	utils.RespondWithHTTPMessage(r, w, http.StatusCreated, "User created successfully.")
}
