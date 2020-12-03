package routes

import (
	"cms/app/http/controllers/admin"
	"github.com/gorilla/mux"
)

func RegisterAdminRouter(r *mux.Router)  {
	router := r.PathPrefix("/admin").Subrouter()

	homeController := admin.HomeController{}
	router.HandleFunc("/", homeController.Welcome).Methods("GET").Name("admin.home")
}
