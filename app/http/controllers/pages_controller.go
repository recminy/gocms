package controllers

import (
	"cms/pkg/logger"
	"cms/pkg/session"
	"fmt"
	"net/http"
)

type PagesController struct {
}

func (p *PagesController) Home(w http.ResponseWriter, r *http.Request) {
	session.Put("UID", "11233")
	logger.Info("HomePageIn")
	fmt.Fprint(w, "<h1>Hello, 欢迎来到 go-blog！</h1>")
}

func (p *PagesController) About(w http.ResponseWriter, r *http.Request) {
	logger.Error("PageErrorhandler")
	fmt.Fprint(w, "笔记：好记性不如烂笔头")
}

func (p *PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
}

