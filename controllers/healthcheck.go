package controllers

import (
	"kgpsc-backend/middleware"
	"kgpsc-backend/utils"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	app := r.Context().Value(middleware.APP_CTX_KEY).(*middleware.App)
	db := app.DB

	err := db.Exec("SELECT 1").Error
	if err != nil {
		utils.LogErrAndRespond(r, w, err, "Could not ping database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("OK"))
	if err != nil {
		utils.LogErr(r, err, "Could not respond to HealthCheck")
		return
	}
}
