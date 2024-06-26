package middleware

import (
	"fmt"
	"net/http"
	"time"

	"kgpsc-backend/utils"
)

func Logger(handler http.Handler, name string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		handler.ServeHTTP(w, r)

		utils.LogInfo(
			r,
			fmt.Sprintf("%s %s", name, time.Since(start)),
		)
	})
}
