package utils

import (
	"api/model"
	"fmt"
	"math/rand"
	"strings"
)

var ModelArray []Model.URLModel

func GenerateKey(modURL *model.URLModel) bool {
	key := ""
	alphabet := "ABCDEFGHIJKLMNOPRSTUVYZabcdefghijklmnoprstuvyz"
	alphabe := strings.Split(alphabet, "")
	for i := 0; i < 6; i++ {
		rand := rand.Intn(46)
		key += alphabe[rand]
	}
	formatKey := fmt.Sprintf("%v:%v/link/%v", modURL.Ipv4, modURL.Port, key)
	fmt.Println(len(ModelArray))
	if len(ModelArray) == 0 { // Eğer array boşsa
		modURL.ShortLink = formatKey
		modURL.Key = key
		ModelArray = append(ModelArray, *modURL)
		return true
	} else if len(ModelArray) > 10 { // Eğer array doluysa arrayin ilk elemanını sil
		UpdateMapArray()
		modURL.ShortLink = formatKey
		modURL.Key = key
		ModelArray = append(ModelArray, *modURL)
		for i := range ModelArray {
			fmt.Println("i : ", i, ModelArray[i].ShortLink)
		}
		return true
	} else {
		for i := range ModelArray {
			if ModelArray[i].ShortLink == formatKey {
				GenerateKey(modURL)
			}
		}
		modURL.ShortLink = formatKey
		modURL.Key = key
		ModelArray = append(ModelArray, *modURL)
		for i := range ModelArray {
			fmt.Println("i : ", i, ModelArray[i].ShortLink)
		}
		return true
	}
}

// UpdateArray fonksiyonu arrayin ilk elemanını siler ve arrayi günceller
func UpdateMapArray() {
	ModelArray = ModelArray[1:] // Arrayin ilk elemanını sil
}
