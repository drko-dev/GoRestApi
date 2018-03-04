package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-api/models"
	// mysql import
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func dbConn() (db *sql.DB) {

	db, err := sql.Open("mysql", "user_proyect:monaFU69ma@/symfony")
	if err != nil {
		panic(err.Error())
	}

	return db
}

// Index pagina principal
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

// GetUsers listar usuarios
func GetUsers(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	var scanField models.Users
	result := []models.Users{}

	rows, err := db.Query("SELECT id, username, email FROM user")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&scanField.ID, &scanField.Username, &scanField.Email)
		if err != nil {
			panic(err.Error())
		}

		result = append(result, scanField)
	}

	defer db.Close()

	if len(result) == 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

// GetUser Listar usuario unico
func GetUser(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	var scanField models.Users
	result := []models.Users{}
	vars := mux.Vars(r)

	rows, err := db.Query("SELECT id, username, email FROM user WHERE id = ?", vars["id"])
	if err != nil && err != sql.ErrNoRows {
		log.Panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&scanField.ID, &scanField.Username, &scanField.Email)
		if err != nil && err != sql.ErrNoRows {
			log.Panic(err)
		}
		result = append(result, scanField)
	}
	defer db.Close()

	if len(result) == 0 {
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

// UserDelete Function Delete destroys a row based on ID
func UserDelete(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	vars := mux.Vars(r)
	// Prepare the SQL Delete
	delForm, err := db.Prepare("DELETE FROM names WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	// Execute the Delete SQL
	delForm.Exec(vars["id"])

	// Show on console the action
	log.Println("DELETE")

	defer db.Close()

	// Redirect a HOME
	http.Redirect(w, r, "/", 301)
}

// UserUpdate It's the same as Insert and New
func UserUpdate(w http.ResponseWriter, r *http.Request) {

	db := dbConn()

	if r.Method == "POST" {

		// Get the values from form
		name := r.FormValue("username")
		email := r.FormValue("email")
		id := r.FormValue("id") // This line is a hidden field on form (View the file: `tmpl/Edit`)

		// Prepare the SQL Update
		insForm, err := db.Prepare("UPDATE names SET username=?, email=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}

		// Update row based on hidden form field ID
		insForm.Exec(name, email, id)

		// Show on console the action
		log.Println("UPDATE: Name: " + name + " | E-mail: " + email)
	}

	defer db.Close()

	// Redirect to Home
	http.Redirect(w, r, "/", 301)
}

// UserInsert Function Insert puts data into the database
func UserInsert(w http.ResponseWriter, r *http.Request) {

	// Open database connection
	db := dbConn()

	// Check the request form METHOD
	if r.Method == "POST" {

		// Get the values from Form
		name := r.FormValue("username")
		email := r.FormValue("email")

		// Prepare a SQL INSERT and check for errors
		insForm, err := db.Prepare("INSERT INTO user(username, email) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}

		// Execute the prepared SQL, getting the form fields
		insForm.Exec(name, email)

		// Show on console the action
		log.Println("INSERT: Name: " + name + " | E-mail: " + email)
	}

	// Close database connection
	defer db.Close()

	// Redirect to HOME
	http.Redirect(w, r, "/", 301)
}

/* func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
*/
