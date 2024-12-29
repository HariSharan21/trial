package api

type Request struct {
	User    string `json:"user"`
	Numbers []int  `json:"numbers"`
}
