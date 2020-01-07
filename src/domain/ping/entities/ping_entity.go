package entities

type PingEntity struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}