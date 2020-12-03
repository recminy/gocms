package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

type AuthController struct {
}

func (a *AuthController) View(w http.ResponseWriter, r *http.Request)  {
	file := "./resources/views" + r.URL.Path + ".gohtml"
	tmpl, _ := template.ParseFiles(file)
	data := make(map[string]interface{})
	data["Title"] = "登录"
	tmpl.Execute(w, data)
}

func (a *AuthController) Register(w http.ResponseWriter, r *http.Request)  {
	
}

func (a *AuthController) Login(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, r.URL)
}

