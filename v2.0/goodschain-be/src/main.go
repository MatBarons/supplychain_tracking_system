package main

import (
	"net/http"

	"github.com/MatBarons/supplychain_tracking_system/blockchains"
	"github.com/MatBarons/supplychain_tracking_system/databases"
	"github.com/MatBarons/supplychain_tracking_system/routes"
)

func main() {
	databases.SetupDB()
	blockchains.SetupBlockchain()
	router := http.NewServeMux()
	requestsRouter := routes.CreateRequestsRoute()
	usersRouter := routes.CreateUsersRoute()
	itemsRouter := routes.CreateItemsRoute()
	router.Handle("/api/", http.StripPrefix("/api", requestsRouter))
	router.Handle("/api/", http.StripPrefix("/api", usersRouter))
	router.Handle("/api/", http.StripPrefix("/api", itemsRouter))
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
