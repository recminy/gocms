package exceptions

import (
	"net/http"
)

func HttpNotFoundException(w http.ResponseWriter)  {
	baseHttpException(w, true, 404, "页面不存在")
}

func HttpNotFoundExceptionView(w http.ResponseWriter)  {
	baseHttpException(w, false, 404, "页面不存在")
}

