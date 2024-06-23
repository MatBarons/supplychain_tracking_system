package main

import (
	"net/http"

	"github.com/MatBarons/supplychain_tracking_system/databases"
	"github.com/MatBarons/supplychain_tracking_system/routes"
)

func main() {
	databases.SetupDB()
	router := http.NewServeMux()
	requestsRouter := routes.CreateItemsRoute()
	usersRouter := routes.CreateUsersRoute()
	router.Handle("/api/", http.StripPrefix("/api", requestsRouter))
	router.Handle("/api/", http.StripPrefix("/api", usersRouter))

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
