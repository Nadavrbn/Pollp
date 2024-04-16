package models

type Question struct {
	Id        interface{} `json:"-"`
	PublicId  uint32      `json:"id"`
	Title     string      `json:"title"`
	Responses Response    `json:"responses"`
}

type Response struct {
	Answer string `json:"answer"`
	Votes  uint   `json:"votes"`
}
