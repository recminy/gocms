package exceptions

import "net/http"

func HttpConflictException(w http.ResponseWriter, message string)  {
	baseHttpException(w, true, 409 , message)
}

func HttpConflictExceptionView(w http.ResponseWriter, message string)  {
	baseHttpException(w, false, 409 , message)
}
