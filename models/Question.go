package models

type Question struct {
	Id        interface{} `json:"-"`
	PublicId  string      `json:"id"`
	Title     string      `json:"title"`
	Responses []Response  `json:"responses"`
}

type Response struct {
	PublicId string `json:"id"`
	Answer   string `json:"answer"`
	Votes    uint   `json:"votes"`
}
