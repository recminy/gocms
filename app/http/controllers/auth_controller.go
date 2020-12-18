package controllers

import (
	"cms/app/http/requests"
	"cms/app/models/clerk"
	"cms/pkg/assets"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type AuthController struct {
	BaseController
}

func (a *AuthController) View(w http.ResponseWriter, r *http.Request)  {
	file := "./resources/views" + r.URL.Path + ".gohtml"
	tmpl, _ := template.ParseFiles(file)
	data := make(map[string]interface{})
	data["Title"] = "登录"
	data["Mix"] = assets.Manifest()
	tmpl.Execute(w, data)
}

func (a *AuthController) Register(w http.ResponseWriter, r *http.Request)  {
	
}

func (a *AuthController) Login(w http.ResponseWriter, r *http.Request)  {
	var user clerk.Clerk
	formBody := requests.FormByteBody(r)
	fmt.Println("参数=", string(formBody))
	err := json.Unmarshal(formBody, &user)
	if err != nil {
		//json参数错误
		fmt.Println("参数错误！")
	}

	var user2 clerk.Clerk
	err2 := json.NewDecoder(r.Body).Decode(user2)
	if err2 != nil {
		fmt.Println("Ergs")
	}

	fmt.Println(user2)

	fmt.Println(user.Username)
	fmt.Println(user.Password)
}

