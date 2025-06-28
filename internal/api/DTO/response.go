package DTO

type ResponseHTTP struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Err     string      `json:"err"`
}
