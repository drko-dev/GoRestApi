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

	// rows, err := db.Query("SELECT * FROM user")
	rows, err := db.Query("SELECT usuario.id, aviso.precio_publicacion, usuario.email FROM usuario INNER JOIN aviso ON aviso.usuario_id = usuario.id LIMIT 1000")
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
