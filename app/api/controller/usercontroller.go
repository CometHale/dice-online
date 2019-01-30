package controller

//controller controls the traffic, so it takes a route, makes the appropriate repo
//the repo returns an object to the controller and the controller passes that to the browser
// in some kind of response

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/comethale/dice-online/app/api/shared/repositories/usermanagement/repository"

	"github.com/comethale/dice-online/app/api/shared/database"
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

		email := r.FormValue("email")
		username := r.FormValue("username")
		password := r.FormValue("password")

		// not a switch because that would be unnecessarily long
		// not indiviual if-statements because all args are required
		if email == "" || username == "" || password == "" {
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}

		user, err := repo.Create(email, password, username)

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

// // UpdateUser takes a POST request and updates a user
// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodPost:
// 		repo := repository.NewUserRepository(database.POSTGRESQL)

// 		err := r.ParseForm()

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}

// 		id, err := strconv.Atoi(r.FormValue("id"))

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}

// 		username := r.FormValue("username")
// 		email := r.FormValue("email")
// 		password := r.FormValue("password")

// 		user, err := repo.Update(email, password, username, 0, int64(id))

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}

// 		json, err := json.Marshal(user)
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(json)

// 	default:
// 		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
// 		return
// 	}
// }

// // UserDelete takes a DELETE request and deletes the given user from the database
// func UserDelete(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodDelete:
// 		repo := repository.NewUserRepository(database.POSTGRESQL)

// 		err := r.ParseForm()

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}

// 		id, err := strconv.Atoi(r.FormValue("id"))

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}

// 		user, err := repo.Delete(int64(id))

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}

// 		json, err := json.Marshal(user)

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(json)
// 	default:
// 		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// }

// // UserLogin logs in the given user
// func UserLogin(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodPost:

// 		err := r.ParseForm()

// 		email := r.FormValue("email")
// 		password := r.FormValue("password")

// 		if email == "" || password == "" {
// 			http.Error(w, "400 Bad Request", http.StatusBadRequest)
// 			return
// 		}

// 		// create session

// 	default:
// 		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
// 		return
// 	}
// }

// // UserLogout logs out the user and returns the index page
// func UserLogout(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodPost:
// 		fmt.Println("logged out")

// 		// mark session as inactive

// 	default:
// 		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
// 	}
// }
