package admin

import (
	"cms/app/exceptions"
	"cms/app/http/controllers"
	"cms/app/models/clerk"
	"cms/pkg/assets"
	"cms/pkg/config"
	"cms/pkg/encrypt"
	"cms/pkg/model"
	"cms/pkg/session"
	"encoding/json"
	"html/template"
	"net/http"
)

type AuthController struct {
	controllers.BaseController
}

func (a *AuthController) View(w http.ResponseWriter, r *http.Request)  {
	file := "./resources/views" + r.URL.Path + ".gohtml"
	tmpl, _ := template.ParseFiles(file)
	data := make(map[string]interface{})
	data["Title"] = "登录-" + config.GetString("APP_NAME") + "-后台管理"
	data["Mix"] = assets.Manifest()
	tmpl.Execute(w, data)
}

func (a *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var user clerk.Clerk
	json.NewDecoder(r.Body).Decode(&user)
	password := user.Password
	model.DB.Where("username=?", user.Username).First(&user)
	if user.ID < 1 {
		exceptions.NotAcceptableException(w, "用户不存在")
		return
	}
	if user.IsActive < 1 {
		exceptions.NotAcceptableException(w, "登录失败：用户被禁用")
		return
	}
	if !encrypt.VerifyPassword(password, user.Encrypt, user.Password) {
		exceptions.NotAcceptableException(w, "用户名密码不正确")
		return
	}
	session.Put("UID", user.ID)
	exceptions.HttpOk(w, "登录成功")
}

