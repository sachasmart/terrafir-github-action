package types

type Post struct {
	APIKey string `json:"api_key"`
	Email  string `json:"email"`
	Input  string `json:"input"`
}

const URL string = "https://api.terrafir.com/"
