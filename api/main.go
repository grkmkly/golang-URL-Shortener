package main

import (
	"fmt"
	"net/http"
	"os"

	"api/main.go/api/handler"
	"api/main.go/api/model"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var myModel *model.URLModel = new(model.URLModel)
var port = ""

func main() {
	godotenv.Load(".env")
	port = os.Getenv("PORT")
	r := mux.NewRouter()
	myModel.Port = port
	handler.MainHandler(r, myModel)
	fmt.Println("Server is running on port: " + myModel.Port)
	http.ListenAndServe(":"+myModel.Port, r)
}
