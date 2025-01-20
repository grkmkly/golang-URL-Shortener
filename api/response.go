package model

type Resp struct {
	ShortLink string `json:"shortLink"`
	LongLink  string `json:"longLink"`
	Status    bool   `json:"status"` // Status 0 offline and 1 online
}
