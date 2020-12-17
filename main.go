package main

import (
	"cms/bootstrap"
	_ "cms/config"
	"cms/pkg/config"
	"fmt"

	//"cms/pkg/encrypt"
	"cms/pkg/logger"
	//"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var Router *mux.Router

func main()  {

	fmt.Println("开始连接数据库...")
	bootstrap.SetupDB()
	fmt.Println("数据库连接成功，开始注册路由...")

	//password, encryptCode := encrypt.CreatePassword("123456")
	//fmt.Println("password=", password)
	//fmt.Println("encrypt=", encryptCode)
	router := bootstrap.SetupRoute()

	fmt.Println("路由注册成功，启动服务器套接字...")
	http.ListenAndServe(":" + config.GetString("app.port"), router)
}
