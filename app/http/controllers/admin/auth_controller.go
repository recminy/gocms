package admin

import (
	"cms/app/http/requests"
	"cms/pkg/assets"
	"cms/pkg/config"
	"encoding/json"
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
	data["Title"] = "登录-" + config.GetString("APP_NAME") + "-后台管理"
	data["Mix"] = assets.Manifest()
	tmpl.Execute(w, data)
}

func (a *AuthController) Login(w http.ResponseWriter, r *http.Request)  {
	//user clerk.Clerk
	params := requests.FormBody(r)
	fmt.Println(json.NewDecoder(r.Body))
	fmt.Println(params)
	//fmt.Println(user)

}

