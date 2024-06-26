package controllers

import (
	"kgpsc-backend/middleware"
	"kgpsc-backend/models"
	"kgpsc-backend/utils"
	"net/http"
	"sort"
)

type SubmitTimesReqFields struct {
	Username string    `json:"username"`
	Event    string    `json:"event"`
	CompID   uint      `json:"comp"`
	Times    []float32 `json:"times"`
}

func calculateAverage(times []float32) (float32, float32) {
	sort.Slice(times, func(i, j int) bool { return times[i] < times[j] })
	return (times[1] + times[2] + times[3]) / 3, times[0]
}

func SubmitTimes(w http.ResponseWriter, r *http.Request) {

	app := r.Context().Value(middleware.APP_CTX_KEY).(*middleware.App)
	db := app.DB

	reqFields := SubmitTimesReqFields{}
	err := utils.DecodeJSON(r, &reqFields)
	if err != nil {
		utils.LogErrAndRespond(r, w, err, "Invalid request payload.", http.StatusBadRequest)
		return
	}

	if len(reqFields.Times) != 5 {
		utils.RespondWithHTTPMessage(r, w, http.StatusBadRequest, "Invalid number of times submitted.")
		return
	}

	ao5, bestTime := calculateAverage(reqFields.Times)

	var user struct {
		ID       uint
		Username string
	}

	tx := db.Table("users").
		Select("id, username").
		Where("username = ?", reqFields.Username).
		First(&user)

	if tx.Error != nil {
		utils.LogErrAndRespond(r, w, tx.Error, "Error fetching user.", http.StatusInternalServerError)
		return
	}

	tx = db.
		Create(&models.Results{
			UserID:   user.ID,
			Username: user.Username,
			CompID:   reqFields.CompID,
			Event:    reqFields.Event,
			Times:    reqFields.Times,
			Ao5:      ao5,
			Best:     bestTime,
		})

	if tx.Error != nil {
		utils.LogErrAndRespond(r, w, tx.Error, "Error submitting times.", http.StatusInternalServerError)
		return
	}

	utils.RespondWithHTTPMessage(r, w, http.StatusOK, "Times submitted successfully.")
}
