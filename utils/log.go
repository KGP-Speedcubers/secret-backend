package utils

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func LogErrAndRespond(r *http.Request, w http.ResponseWriter, err error, msg string, code int) {
	LogErr(r, err, msg)
	RespondWithHTTPMessage(r, w, code, msg)
}

func LogErr(r *http.Request, err error, errMsg string) {
	log.Err(err).Msgf(
		"%s %s: %s",
		r.Method,
		r.RequestURI,
		errMsg,
	)
}

func LogInfo(r *http.Request, msg string) {
	log.Info().Msgf(
		"%s %s: %s",
		r.Method,
		r.RequestURI,
		msg,
	)
}
