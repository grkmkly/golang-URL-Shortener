package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"main.go/api"
	"main.go/model"
)

var myModel *model.URLModel = new(model.URLModel)
var port = ""

func main() {
	godotenv.Load(".env")
	port = os.Getenv("PORT")
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Burada tüm domainlere izin veriyoruz, belirli bir domain yazabilirsiniz.
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})
	handlerWithCORS := c.Handler(r)
	myModel.Port = port
	api.MainHandler(r, myModel)
	fmt.Println("Server is running on port: " + myModel.Port)
	http.ListenAndServe(":"+myModel.Port, handlerWithCORS)
}
