package main

import (
	"net/http"

	"github.com/MatBarons/supplychain_tracking_system/routes"
)

func main() {
	router := http.NewServeMux()

	itemsRouter := routes.CreateItemsRoute()
	usersRouter := routes.CreateUsersRoute()
	router.Handle("/api/", http.StripPrefix("/api", itemsRouter))
	router.Handle("/api/", http.StripPrefix("/api", usersRouter))

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
