package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"main.go/handler"
	"main.go/model"
)

var myModel *model.URLModel = new(model.URLModel)
var port = "8080"

func main() {
	r := mux.NewRouter()
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Herhangi bir kaynaktan isteklere izin verir
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})
	myModel.Port = port
	handler.MainHandler(r, myModel)
	http.ListenAndServe(":"+myModel.Port, corsHandler.Handler(r))
}
