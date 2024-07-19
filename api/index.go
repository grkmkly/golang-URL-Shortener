package index

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
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
	<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet"
		integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH" crossorigin="anonymous">
</head>
<body class="bg-black h-100">
	<div class="container-fluid text-center " style="height:100%%;">
		<div class="row d-flex align-items-center h-100">
			<div class="m-3 h-100">
				<form  action="/action-url" method="post">
					<div class="row d-flex mt-3 ">
						<p class="fs-2 text-light">--- enter url and click submit ---</p>
					</div>
					<div class="row d-flex m-3">
						<span class="col">
							<input type="text" id="url" name="url" class="fs-5" style="width: 50%%;" >
						</span>
					</div>
					<span class="row pt-2">
						<span class="col">
						<input style="background-color:#59CE8F;width: 50%%;height: 4rem;" type="submit" name="submit" value="Submit" >
					</span>
					<span class="col">
						<input style="background-color:#FF1E00; color:#E8F9FD; width: 50%%;height: 4rem;;" type="reset" name="Reset" value="Reset" >
					</span>
					</span>
			</div>
			<div class="row d-flex mt-3 ">
				<p class="text-light fs-5 m-3">link</p>
			</div>
			<div class="row d-flex mt-3 ">
				<p class="text-primary m-3" href="%v">%v</p>
			</div>
		</div>
	</div>
	<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js"
		integrity="sha384-YvpcrYf0tY3lHB60NNkmXc5s9fDVZLESaAA55NDzOxhy9GkcIdslK1eN7N6jIeHz"
		crossorigin="anonymous"></script>
</body>
</html>`, urlMap[newUrl], newUrl)
	os.WriteFile("api/newUrl.html", []byte(html), 0755)
}

// var urlMap map[string]string
var urlMap = make(map[string]string)
var port string

// Key üreteci
func generateKey() string {
	key := ""
	alphabet := "ABCDEFGHIJKLMNOPRSTUVYZabcdefghijklmnoprstuvyz"
	alphabe := strings.Split(alphabet, "")
	for i := 0; i < 10; i++ {
		rand := rand.Intn(46)
		key += alphabe[rand]
	}
	// Eğer linkler aynı olursa tekrar key üretme
	formatKey := fmt.Sprintf("localhost:%v/linkpage/%v", port, key)
	for key := range urlMap {
		if key == formatKey {
			generateKey()
		}
	}
	return key
}

// HomePage oluşturucu
func homePage(w http.ResponseWriter, r *http.Request) {
	htmlByte, err := os.ReadFile("api/homepage.html")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(htmlByte)
}

// Url kısaltıcı
func urlShorter(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		key := generateKey()
		newUrl := fmt.Sprintf("localhost:%v/linkpage/%v", port, key) // yeni url oluşturuldu
		oldUrl := r.FormValue("url")                                 //  gelen url alındı
		urlMap[newUrl] = oldUrl                                      // oldUrl = Gelen Url newUrl yarattığımız Url
		writeHtml(newUrl)
		logging(newUrl, oldUrl)
		//logging(newUrl, oldUrl)
		http.Redirect(w, r, "/linkpage", http.StatusSeeOther)
	}
}

// Url'yi yönlendirici
func redirectUrl(w http.ResponseWriter, r *http.Request) {
	hosts := fmt.Sprintf("localhost:%v%v", port, r.URL.Path)
	for key, value := range urlMap {
		fmt.Printf("Key : %v\n", key)
		if key == hosts {
			http.Redirect(w, r, value, http.StatusSeeOther)
			return
		}
	}
}

// Link sayfası oluşturucu
func linkPage(w http.ResponseWriter, r *http.Request) {
	htmlByte, err := os.ReadFile("api/newUrl.html")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(htmlByte)
}

// Ip adres getirici
func getIpAdrs() (net.IP, error) {
	var ipv4 net.IP

	host, err := os.Hostname() // PC'nin hostuna bakılıyor
	if err != nil {
		log.Fatal(err)
	}

	adrrs, err := net.LookupIP(host) // PC'nin ip değerlerine bakılıyor
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range adrrs {
		ipv4 = value.To4() // Pcnin ip değerlerindeki ipv4 değerini alıyor
		if ipv4 != nil {
			return ipv4, nil
		}
	}
	return nil, errors.New("IpNotFound")
}

// LOG dosyasına kayıt
func logging(newUrl string, oldUrl string) {
	ipv4, err := getIpAdrs() // İP adresini aldık
	if err != nil {
		log.Fatal(err)
	}
	// Dosya okuma ve yazma işlemi
	logByte, err := os.ReadFile("logging.txt")
	if err != nil {
		log.Fatal(err)
	}
	logText := string(logByte)

	if logText == "" {
		logText += fmt.Sprintf("Ip :%v,NewUrl :%v,OldUrl :%v", ipv4, newUrl, oldUrl)
		os.WriteFile("logging.txt", []byte(logText), 0755)
		return
	}
	logText += fmt.Sprintf("\nIP :%v,newUrl :%v,OldUrl :%v", ipv4, newUrl, oldUrl)
	os.WriteFile("logging.txt", []byte(logText), 0755)
}

func redirectHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/homepage", http.StatusSeeOther)
}

// MAİN
func main() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", redirectHome)
	http.HandleFunc("/linkpage", linkPage)
	http.HandleFunc("/homepage", homePage)
	http.HandleFunc("/action-url", urlShorter)
	http.HandleFunc("/linkpage/{key}", redirectUrl)
	http.ListenAndServe(":"+port, nil)
}
