package routes

import (
	"cms/app/http/controllers/admin"
	"github.com/gorilla/mux"
)

func RegisterAdminRouter(r *mux.Router)  {
	router := r.PathPrefix("/admin").Subrouter()

	loginController := admin.AuthController{}
	router.HandleFunc("/login", loginController.View).Methods("GET").Name("admin.loginView")
	router.HandleFunc("/login", loginController.Login).Methods("PATCH").Name("admin.loginSubmit")

	homeController := admin.HomeController{}
	router.HandleFunc("/", homeController.Welcome).Methods("GET").Name("admin.home")
}
