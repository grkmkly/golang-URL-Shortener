package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"api/main.go/api/model"
	"api/main.go/api/utils"

	"github.com/gorilla/mux"
)

func homepage() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		htmlFile, err := os.ReadFile("src/index.html") // Adjust the path to your HTML file
		if err != nil {
			http.Error(w, "Could not read HTML file", http.StatusInternalServerError)
			return
		}
		w.Write(htmlFile)
	}
}
func getLink(modURL *model.URLModel) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// Gelen json verisini modeldeki longlink'e ata
		var longLink model.Longlink
		x, _ := io.ReadAll(r.Body)
		err1 := json.Unmarshal(x, &longLink) // Gelen json verisini longlink modeline çevir
		if err1 != nil {
			log.Fatal(err1)
		}
		var splitString []string = strings.Split(longLink.LongLink, "://") // Gelen linki parçala
		fmt.Println(splitString[0])
		modURL.LongLink = splitString[1] // Parçalanan linki modeldeki longlink'e ata
		var ipv4, err = utils.GetIpAdrs()
		modURL.Ipv4 = ipv4.String() // Ip adresini modeldeki ipv4'e ata
		if err != nil {
			log.Fatal(err)
		}

		isTrue := utils.GenerateKey(modURL)
		if !isTrue {
			log.Fatal("Key oluşturulamadı")
		}

		var resp model.Resp
		resp.Status = true
		resp.ShortLink = modURL.ShortLink
		resp.LongLink = modURL.LongLink
		writeJson, err := json.Marshal(resp)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(w, string(writeJson))
	}
}
func redirect() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]
		for i := range utils.ModelArray {
			if utils.ModelArray[i].Key == key {
				http.Redirect(w, r, "https://"+utils.ModelArray[i].LongLink, http.StatusMovedPermanently)
			}
		}
	}

}

func MainHandler(r *mux.Router, model *model.URLModel) {
	r.HandleFunc("/", homepage()).Methods("GET")
	r.HandleFunc("/homepage", homepage()).Methods("GET")
	r.HandleFunc("/getlink", getLink(model)).Methods("POST")
	r.HandleFunc("/link/{key}", redirect()).Methods("GET")
}
