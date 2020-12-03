package routes

import (
	"cms/app/http/controllers"
	"cms/app/http/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterWebRoute(r *mux.Router) {

	//静态资源加载
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))

	//auth
	auth := &controllers.AuthController{}
	r.HandleFunc("/login", auth.View).Methods("Get").Name("login.view")
	r.HandleFunc("/login", auth.Login).Methods("PATCH").Name("login.submit")

	pc := &controllers.PagesController{}
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About)
	//开启session中间件
	r.Use(middlewares.StartSession)
}
