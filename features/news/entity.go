package news

import "time"

type NewsCore struct {
	Author      string    `json: "author"`
	Category    string    `json: "category"`
	Country     string    `json: "country"`
	Description string    `json: "description"`
	Image       string    `json: "image"`
	Language    string    `json: "language"`
	PublishedAt time.Time `json: "published_at"`
	Source      string    `json: "source"`
	Title       string    `json: "title"`
	Url         string    `json: "url"`
}

type Service interface {
	GetNews() ([]NewsCore, error)
}

type Repository interface {
	GetData() ([]NewsCore, error)
}
