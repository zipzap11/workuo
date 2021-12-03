package data

import (
	"time"
	"workuo/features/news"
)

type Data struct {
	Data []New `json: "data"`
}

type New struct {
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

func (n *New) ToCore() news.NewsCore {
	return news.NewsCore{
		Author:      n.Author,
		Category:    n.Category,
		Country:     n.Country,
		Description: n.Description,
		Image:       n.Image,
		Language:    n.Language,
		PublishedAt: n.PublishedAt,
		Source:      n.Source,
		Title:       n.Title,
		Url:         n.Url,
	}
}

func ToCoreList(data []New) []news.NewsCore {
	converted := []news.NewsCore{}
	for _, new := range data {
		converted = append(converted, new.ToCore())
	}
	return converted
}
