package bootstrap

import (
	"cms/routes"
	"github.com/gorilla/mux"
)

func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoute(router)
	routes.RegisterAdminRouter(router)
	return router
}
