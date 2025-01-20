package main

type URLModel struct {
	LongLink  string `json:"longLink"`
	ShortLink string `json:"shortLink"`
	Port      string `json:"port"`
	Ipv4      string `json:"ipv4"`
	Key       string `json:"key"`
}
