package entities

// datamodelに直す（createdATTT)
type PingEntity struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	CreatedAt string `json:"created_at"`
}