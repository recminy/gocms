package exceptions

import "net/http"

func HttpOk(w http.ResponseWriter, message string)  {
	baseHttpException(w, true, 200, message)
}

func HttpAccepted(w http.ResponseWriter, message string)  {
	baseHttpException(w, true, 200, message)
}
