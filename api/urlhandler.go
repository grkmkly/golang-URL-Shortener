package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"main.go/model"
	"main.go/utils"
)

func GetLink(modURL *model.URLModel) func(w http.ResponseWriter, r *http.Request) {
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
		modURL.LongLink = splitString[1]                                   // Parçalanan linki modeldeki longlink'e ata
		//var ipv4, err = utils.GetIpAdrs()
		modURL.Ipv4 = os.Getenv("HOSTIP")

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
func Redirect() func(w http.ResponseWriter, r *http.Request) {
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
	r.HandleFunc("/getlink", GetLink(model)).Methods("POST")
	r.HandleFunc("/link/{key}", Redirect()).Methods("GET")
}
