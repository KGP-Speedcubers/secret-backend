package server

import (
	"errors"
	"kgpsc-backend/middleware"
	"kgpsc-backend/utils"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var ErrRouteNotFound = errors.New("route not found")
var ErrMethodNotAllowed = errors.New("method not allowed")

func NewRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.LogErrAndRespond(r, w, ErrRouteNotFound, "404 Not Found.", http.StatusNotFound)
	})

	router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.LogErrAndRespond(r, w, ErrMethodNotAllowed, "405 Method Not Allowed.", http.StatusMethodNotAllowed)
	})

	app := &middleware.App{DB: db}
	routes := getRoutes(app)

	for _, route := range routes {
		if route.disabled {
			continue
		}

		handler := middleware.Logger(route.HandlerFunc, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
