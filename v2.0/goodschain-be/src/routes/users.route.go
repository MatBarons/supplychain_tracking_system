package routes

import (
	"net/http"

	"github.com/MatBarons/supplychain_tracking_system/handlers"
)

func CreateUsersRoute() *http.ServeMux {
	usersRouter := http.NewServeMux()
	usersRouter.HandleFunc("POST users/{email}", handlers.HandleUserLogin)
	usersRouter.HandleFunc("POST users", handlers.HandlerUserCreation)
	usersRouter.HandleFunc("PUT users/{email}", handlers.HandleUserUpdateInfo)
	return usersRouter
}
