package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MatBarons/supplychain_tracking_system/databases"
	"github.com/MatBarons/supplychain_tracking_system/models"
)

func RequestsList(w http.ResponseWriter, r *http.Request) {
	query := "SELECT * FROM Requests WHERE user=?;"
	result := databases.ExecQuery(query, r.URL.Query().Get("id"))
	var requests []models.Request
	for result.Next() {
		var req models.Request
		if err := result.Scan(&req.Id, &req.UserID, &req.ItemID, &req.IsConfirmed, &req.Requester); err != nil {
			requests = append(requests, req)
		}
	}
	buf, err := json.Marshal(requests)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Write(buf)
}

func CreateNewRequest(w http.ResponseWriter, r *http.Request) {
	query := "INSERT INTO Request (Borrower, itemID, isApproved) VALUES (?, ?, FALSE);"
	var request models.Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if request.ItemID == "" || request.Requester == "" {
		http.Error(w, "request data is missing", http.StatusNotFound)
	}
	databases.ExecUpdate(query, request.UserID, request.ItemID)
}

func DeleteRequest(w http.ResponseWriter, r *http.Request) {
	query := "DELETE FROM Requests WHERE Id=?;"
	databases.ExecUpdate(query, r.URL.Query().Get("id"))
}
