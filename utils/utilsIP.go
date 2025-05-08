package utils

import (
	"errors"
	"log"
	"net"
	"os"
)

func GetIpAdrs() (net.IP, error) {
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
