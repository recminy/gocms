package middlewares

import (
	"cms/pkg/session"
	"net/http"
)

func StartSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//start
		session.StartSession(w, r)

		//下一步请求
		next.ServeHTTP(w, r)
	})
}
