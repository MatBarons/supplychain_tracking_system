package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MatBarons/supplychain_tracking_system/databases"
	"github.com/MatBarons/supplychain_tracking_system/models"
)

func HandleRequestsList(w http.ResponseWriter, r *http.Request) {
	query := "SELECT * FROM Requests WHERE user=?;"
	result := databases.ExecQuery(query, r.URL.Query().Get("id"))
	var requests []models.Request
	for result.Next() {
		var req models.Request
		if err := result.Scan(&req.Id, &req.User, &req.ItemID, &req.IsConfirmed, &req.Borrower); err != nil {
			requests = append(requests, req)

		}
	}
	buf, err := json.Marshal(requests)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(buf)
}

func HandleCreateNewRequest(w http.ResponseWriter, r *http.Request) {
	query := "INSERT INTO Request (UserID, itemID, isApproved) VALUES (?, ?, FALSE);"
	var request models.Request
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&request)
	if err != nil {
		log.Fatal(err)
	}
	databases.ExecUpdate(query, request.User, request.ItemID)
}
