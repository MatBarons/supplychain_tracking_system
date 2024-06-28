package handlers

import (
	"net/http"

	"github.com/MatBarons/supplychain_tracking_system/databases"
	"github.com/MatBarons/supplychain_tracking_system/models"
)

func HandleRequestsList(w http.ResponseWriter, r *http.Request) {
	var query string = "SELECT * FROM Requests WHERE user=?"
	result := databases.ExecQuery(query, r.URL.Query().Get("id"))
	var requests []models.Request
	for result.Next() {
		var req models.Request
		if err := result.Scan(&req.Id, &req.User, &req.ItemID, &req.IsConfirmed, &req.Borrower); err != nil {
			requests = append(requests, req)
			w.Write()
		}
	}
}

func HandleCreateNewRequest(w http.ResponseWriter, r *http.Request) {
}
