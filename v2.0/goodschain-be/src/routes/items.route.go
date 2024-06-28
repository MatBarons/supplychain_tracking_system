package routes

import (
	"net/http"

	"github.com/MatBarons/supplychain_tracking_system/handlers"
)

func CreateItemsRoute() *http.ServeMux {
	itemsRouter := http.NewServeMux()
	itemsRouter.HandleFunc("POST items/list", handlers.HandleItemsList)
	itemsRouter.HandleFunc("GET item/{id}/history", handlers.HandleItemHistory)
	return itemsRouter
}
