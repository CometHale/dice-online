package route

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/comethale/dice-online/app/api/controller"
	"github.com/comethale/dice-online/app/api/route/middleware/cors"
	"github.com/comethale/dice-online/app/api/route/middleware/logrequest"
	"github.com/comethale/dice-online/app/api/shared/database"
	"github.com/comethale/dice-online/app/api/shared/repositories/usermanagement/repository"
)

// LoadRoutes loads the routes for the application
func LoadRoutes() http.Handler {
	return middleware(routes())
}

func routes() *http.ServeMux {
	mux := http.NewServeMux()

	// User
	mux.HandleFunc("/user-create/", controller.UserCreate)
	mux.HandleFunc("/user-login/", controller.UserLogin)
	mux.HandleFunc("/user-logout/", controller.UserLogout)

	// View User
	mux.HandleFunc("/view-high-score/", ViewUser)

	// View All Users
	mux.HandleFunc("/view-all/", ViewAll)

	// Game
	mux.HandleFunc("/start-game/", controller.StartGame)
	mux.HandleFunc("/roll-dice/", controller.RollDice)

	return mux
}

// ViewAll returns a JSON of all users
func ViewAll(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		repo := repository.NewUserRepository(database.POSTGRESQL)

		users, err := repo.GetAll()

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json, err := json.Marshal(users)

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
		return
	default:
		log.Println("405 Method Not Allowed")
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

// ViewUser returns a JSON of the requested user
func ViewUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		userID, err := strconv.Atoi(r.URL.Query().Get("userid"))
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		repo := repository.NewUserRepository(database.POSTGRESQL)
		user, err := repo.Get(int64(userID))

		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

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
		log.Println("405 Method Not Allowed")
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

}

// *******************************
// Middleware
// *******************************

func middleware(h http.Handler) http.Handler {
	// Log every request
	h = logrequest.Handler(h)

	// Cors for swagger-ui
	h = cors.Handler(h)

	// Clear handler for Gorilla Context
	// h = context.ClearHandler(h)

	return h
}
