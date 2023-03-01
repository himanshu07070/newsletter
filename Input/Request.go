package input

type User struct {
	Email string `json:"email"`
	Tag   string `json:"tag"`
}

type MetaData struct {
	Link        string `json:"link"`
	Tag         string `json:"tag"`
	ArticleName string `json:"articleName"`
	Subject     string `json:"subject"`
}

type UserEmail struct {
	Email string `json:"email"`
}
