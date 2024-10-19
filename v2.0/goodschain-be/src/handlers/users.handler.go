package handlers

import (
	//"encoding/json"
	//"log"
	"net/http"

	"github.com/MatBarons/supplychain_tracking_system/databases"
	//"github.com/MatBarons/supplychain_tracking_system/models"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	query := "SELECT password FROM Users WHERE email=?;"
	result := databases.ExecQuery(query, r.URL.Query().Get("email"))

}

func UserCreation(w http.ResponseWriter, r *http.Request) {

}

func UserUpdateInfo(w http.ResponseWriter, r *http.Request) {

}
