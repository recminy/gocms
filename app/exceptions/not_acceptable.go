package exceptions

import "net/http"

func NotAcceptableException(w http.ResponseWriter, message string)  {
	baseHttpException(w, true, 406 , message)
}

func NotAcceptableExceptionView(w http.ResponseWriter, message string)  {
	baseHttpException(w, false, 406 , message)
}
