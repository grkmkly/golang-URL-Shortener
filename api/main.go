package main

import (
	"net/http"

	"api/main.go/api/handler"
	"api/main.go/api/model"

	"github.com/gorilla/mux"
)

var myModel *model.URLModel = new(model.URLModel)
var port = "8080"

func main() {
	r := mux.NewRouter()
	myModel.Port = port
	handler.MainHandler(r, myModel)
	http.ListenAndServe(":"+myModel.Port, r)
}
