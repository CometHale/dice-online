package route

import (
	"goblr/app/api/controller"
	"net/http"
)

// LoadRoutes loads the routes for the application
func LoadRoutes() http.Handler {

	return routes()
}

func routes() *http.ServeMux {
	mux := http.NewServeMux()

	// User
	mux.HandleFunc("/user-create/", controller.UserCreate)
	// mux.HandleFunc("/user-delete/", controller.UserDelete)
	// mux.HandleFunc("/user-login/", controller.UserLogin)
	// mux.HandleFunc("/user-logout/", controller.UserLogout)

	return mux
}
