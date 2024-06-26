package server

import (
	"kgpsc-backend/controllers"
	"kgpsc-backend/middleware"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	disabled    bool
}

func getRoutes(app *middleware.App) []Route {
	return []Route{
		{
			"HealthCheck",
			"GET",
			"/healthcheck/",
			middleware.WrapApp(app, controllers.HealthCheck),
			false,
		},
		{
			"Register",
			"POST",
			"/register/",
			middleware.WrapApp(app, controllers.Register),
			false,
		},
		{
			"Login",
			"POST",
			"/login/",
			middleware.WrapApp(app, controllers.Login),
			false,
		},
		{
			"Users",
			"Get",
			"/users/",
			middleware.WrapApp(app, controllers.GetAllUsers),
			false,
		},
	}
}
