package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BaseController struct {

}

func (b *BaseController) Json(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	jsonRes, _ := json.Marshal(data)
	fmt.Fprint(w, string(jsonRes))
}