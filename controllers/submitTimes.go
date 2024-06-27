package controllers

import (
	"kgpsc-backend/middleware"
	"kgpsc-backend/models"
	"kgpsc-backend/utils"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type SubmitTimesReqFields struct {
	Username string `json:"username"`
	Event    string `json:"event"`
	CompID   uint   `json:"comp"`
	Times    string `json:"times"` // Comma separated string of 5 times
}

func calculateAverage(times []float32) (float32, float32) {
	sort.Slice(times, func(i, j int) bool { return times[i] < times[j] })
	return (times[1] + times[2] + times[3]) / 3, times[0]
}

func StringToFloat32Slice(s string) []float32 {
	var floatSlice []float32
	for _, v := range strings.Split(s, ",") {
		fv, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return nil
		}
		floatSlice = append(floatSlice, float32(fv))
	}
	return floatSlice
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

	times := StringToFloat32Slice(reqFields.Times)

	ao5, bestTime := calculateAverage(times)

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
			Ao5:      strconv.FormatFloat(float64(ao5), 'f', 2, 64),
			Best:     strconv.FormatFloat(float64(bestTime), 'f', 2, 64),
		})

	if tx.Error != nil {
		utils.LogErrAndRespond(r, w, tx.Error, "Error submitting times.", http.StatusInternalServerError)
		return
	}

	utils.RespondWithHTTPMessage(r, w, http.StatusOK, "Times submitted successfully.")
}
