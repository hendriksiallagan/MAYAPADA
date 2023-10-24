package models

type GetResponseNews struct {
	Code int64   `json:"code"`
	Data []*News `json:"data"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
