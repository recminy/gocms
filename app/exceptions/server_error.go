package exceptions

import "net/http"

func ServerErrorExceptionView(w http.ResponseWriter, message string)  {
	baseHttpException(w, false, 500 , message)
}
