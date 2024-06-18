package routes

import (
	"net/http"

	"github.com/MatBarons/supplychain_tracking_system/handlers"
)

func CreateUsersRoute() *http.ServeMux {
	usersRouter := http.NewServeMux()
	usersRouter.HandleFunc("POST users/login", handlers.HandleUserLogin)
	usersRouter.HandleFunc("POST users", handlers.HandlerUserCreation)
	return usersRouter
}
