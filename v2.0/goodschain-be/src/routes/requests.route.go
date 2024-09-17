package routes

import (
	"net/http"

	"github.com/MatBarons/supplychain_tracking_system/handlers"
)

func CreateRequestsRoute() *http.ServeMux {
	requestsRouter := http.NewServeMux()
	requestsRouter.HandleFunc("POST requests", handlers.CreateNewRequest)
	requestsRouter.HandleFunc("POST requests/list", handlers.RequestsList)
	requestsRouter.HandleFunc("DELETE requests/{id}", handlers.DeleteRequest)
	return requestsRouter
}
