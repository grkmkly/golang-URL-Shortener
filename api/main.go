package main

import (
	"api/main.go/api/model"
	"net/http"

	"github.com/gorilla/mux"
)

var myModel *model.URLModel = new(model.URLModel)
var port = "8080"

func main() {
	r := mux.NewRouter()
	myModel.Port = port
	MainHandler(r, myModel)
	http.ListenAndServe(":"+myModel.Port, r)
}
