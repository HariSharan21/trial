package api

type SumRequest struct {
	User    string  `json:"user"`
	Numbers []int   `json:"numbers"`
}

type SumResponse struct {
	User string `json:"user"`
	Sum  int64  `json:"sum"`
}

type DiffRequest struct {
	User      string  `json:"user"`
	Numbers   []int64   `json:"numbers"`
	MaxNumber int64   `json:"max_number"`
}
