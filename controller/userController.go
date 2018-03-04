package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-api/models"

	"github.com/gorilla/mux"
)

// Index pagina principal
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsonSuccess{Code: http.StatusOK, Text: "Success"})
}

// GetUsers listar usuarios
func GetUsers(w http.ResponseWriter, r *http.Request) {

	db := dbConn()

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()

	result, err := models.ScanUsers(rows)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	if len(result) == 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			log.Panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Panic(err)
	}
}

// GetUser Listar usuario unico
func GetUser(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	vars := mux.Vars(r)

	row := db.QueryRow("SELECT * FROM user WHERE id = ?", vars["id"])

	result, err := models.ScanUser(row)
	defer db.Close()

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}
