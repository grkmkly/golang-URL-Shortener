package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

func writeHtml(newUrl string) {
	var html string = fmt.Sprintf(
		`<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<form action="/action-url" method="post">
			<label for="fname">First name:<label><br>
			<input type="text" id="fname" name="fname" value="John"><br>
			<label for="lname">Last name:</label><br>
			<input type="text" id="lname" name="lname" value="Doe"><br><br>
			<input type="submit" name ="submit" value="Submit">
			<input type="reset" name ="Reset" value="Reset">
		  </form>
		  <p>%v</p>
	</body>
	</html>`, newUrl)
	os.WriteFile("src/index.html", []byte(html), 0755)
}

// var urlMap map[string]string
var urlMap = make(map[string]string)

func generateKey() string {
	key := ""
	alphabet := "ABCDEFGHIJKLMNOPRSTUVYZabcdefghijklmnoprstuvyz"
	alphabe := strings.Split(alphabet, "")
	for i := 0; i < 10; i++ {
		rand := rand.Intn(46)
		key += alphabe[rand]
	}
	return key
}
func homePage(w http.ResponseWriter, r *http.Request) {
	htmlByte, err := os.ReadFile("src/index.html")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(htmlByte)
}
func urlShorter(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		key := generateKey()

		newUrl := fmt.Sprintf("https://mgk.com/%v", key) // yeni url oluşturuldu
		oldUrl := r.FormValue("fname")
		urlMap[newUrl] = oldUrl // oldUrl = Gelen Url newUrl yarattığımız Url

		fmt.Println(urlMap)
		writeHtml(newUrl)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/action-url", urlShorter)
	http.ListenAndServe(":5000", nil)
}
