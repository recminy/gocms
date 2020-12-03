package main

import (
	"cms/bootstrap"
	_ "cms/config"
	"cms/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
)

var Router *mux.Router

func main()  {
	router := bootstrap.SetupRoute()
	logger.Info("服务器启动成功...")
	http.ListenAndServe(":8080", router)
}
