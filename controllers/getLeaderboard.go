package controllers

import (
	"encoding/json"
	"kgpsc-backend/middleware"
	"kgpsc-backend/models"
	"kgpsc-backend/utils"
	"net/http"
)

func GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	app := r.Context().Value(middleware.APP_CTX_KEY).(*middleware.App)
	db := app.DB

	leaderboard := []models.Results{}

	tx := db.
		Table("results").
		Order("ao5 desc").
		Find(&leaderboard)

	if tx.Error != nil {
		utils.LogErrAndRespond(r, w, tx.Error, "Error fetching leaderboard.", http.StatusInternalServerError)
		return
	}

	if len(leaderboard) == 0 {
		utils.RespondWithHTTPMessage(r, w, http.StatusOK, "No results exist.")
		return
	}

	err := json.NewEncoder(w).Encode(leaderboard)
	if err != nil {
		utils.LogErrAndRespond(r, w, err, "Error encoding leaderboard.", http.StatusInternalServerError)
		return
	}

	utils.RespondWithHTTPMessage(r, w, http.StatusOK, "Leaderboard fetched successfully.")
}
