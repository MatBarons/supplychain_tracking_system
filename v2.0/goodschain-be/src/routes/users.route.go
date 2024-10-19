package routes

import (
	"net/http"

	"github.com/MatBarons/supplychain_tracking_system/handlers"
)

func CreateUsersRoute() *http.ServeMux {
	usersRouter := http.NewServeMux()
	usersRouter.HandleFunc("POST users/{email}", handlers.UserLogin)
	usersRouter.HandleFunc("POST users", handlers.UserCreation)
	usersRouter.HandleFunc("PUT users/{email}", handlers.UserUpdateInfo)
	return usersRouter
}
