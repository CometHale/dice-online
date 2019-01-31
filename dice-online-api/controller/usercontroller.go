package controller

//controller controls the traffic, so it takes a route, makes the appropriate repo
//the repo returns an object to the controller and the controller passes that to the browser
// in some kind of response

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"dice-online-api/model/domain"
	"dice-online-api/shared/repositories/usermanagement/repository"
	"dice-online-api/shared/session"
	"dice-online-api/shared/utils"

	"dice-online-api/shared/database"
)

// UserGet takes a GET request and returns a user from the database
func UserGet(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		repo := repository.NewUserRepository(database.POSTGRESQL)

		err := r.ParseForm()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		id, err := strconv.Atoi(r.FormValue("id"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		user, err := repo.Get(int64(id))

		json, err := json.Marshal(user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)

	default:
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// UserCreate takes a POST request and creates a user if the required information is provided
func UserCreate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		repo := repository.NewUserRepository(database.POSTGRESQL)

		err := r.ParseForm()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		email := r.PostFormValue("email")
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		// not a switch because that would be unnecessarily long
		// not indiviual if-statements because all args are required
		if email == "" || username == "" || password == "" {
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}

		user, err := repo.Create(email, password, username)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json, err := json.Marshal(user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)

	default:
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

// UserLogin logs in the given user
func UserLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		repo := repository.NewUserRepository(database.POSTGRESQL)

		err := r.ParseForm()

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		email := r.PostFormValue("email")
		password := r.PostFormValue("password")

		if email == "" || password == "" {
			log.Println("400 Bad Request")
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}

		DBPassword, id, err := repo.GetPassword(email)

		// create session
		session := session.Instance(r, email)

		loggedIn := utils.AuthVerifyPassword(password, DBPassword)

		if loggedIn != nil {
			log.Println(loggedIn)
			http.Error(w, loggedIn.Error(), http.StatusUnauthorized)
			return
		}

		// Set user as authenticated
		session.Values["authenticated"] = true
		session.Save(r, w)

		user := &domain.User{ID: id, Email: email}
		json, err := json.Marshal(user)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)

		return
	default:
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

// UserLogout logs out the user and returns the index page
func UserLogout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		err := r.ParseForm()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		email := r.FormValue("email")
		session := session.Instance(r, email)

		// Revoke users authentication
		session.Values["authenticated"] = false
		session.Save(r, w)
		return

	default:
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
