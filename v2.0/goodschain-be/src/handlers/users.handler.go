package handlers

import (
	//"encoding/json"
	//"log"
	"net/http"

	"github.com/MatBarons/supplychain_tracking_system/databases"
	//"github.com/MatBarons/supplychain_tracking_system/models"
)

func HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	query := "SELECT password FROM Users WHERE email=?;"
	result := databases.ExecQuery(query, r.URL.Query().Get("email"))

}

func HandlerUserCreation(w http.ResponseWriter, r *http.Request) {

}

func HandleUserUpdateInfo(w http.ResponseWriter, r *http.Request) {

}
