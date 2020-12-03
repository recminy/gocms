package session

import (
	"cms/pkg/config"
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

var Store = sessions.NewCookieStore([]byte(config.GetString("APP_KEY")))

var Session *sessions.Session

var Request *http.Request

var Response http.ResponseWriter

func StartSession(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("启动session...")
	var err error
	Session, err = Store.Get(r, config.GetString("APP_NAME", "goApp"))
	//记录错误日志
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Session启动成功!")
	}
	//TODO
	Request = r
	Response = w
}

func Put(key string, value interface{}) {
	Session.Values[key] = value
	Save()
}

func Get(name string) interface{}  {
	return Session.Values[name]
}

func Forget(key string) {
	delete(Session.Values, key)
	Save()
}

func Flush()  {
	Session.Options.MaxAge = -1
	Save()
}

func Save()  {
	err := Session.Save(Request, Response)
	if err != nil {
		fmt.Println("SessionSavedError:", err)
	}
	//TODO
}