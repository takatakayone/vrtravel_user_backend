package datamodels

import "github.com/takatakayone/vrtravel_user_backend/src/domain/ping/entities"

type Ping struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	CreatedAt string `json:"created_at"`
}

func (p *Ping) Mapping() entities.PingEntity {
	return entities.PingEntity{
		Id:    p.Id,
		Title: p.Title,
		Body:  p.Body,
	}
}