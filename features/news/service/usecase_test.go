package service

import (
	"errors"
	"os"
	"testing"
	"time"
	"workuo/features/news"
	"workuo/features/news/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	newService    news.Service
	newRepository mocks.Repository
	newsData      []news.NewsCore
)

func TestMain(m *testing.M) {
	newService = NewApiService(&newRepository)
	newsData = []news.NewsCore{
		{
			Author:      "ABMN Staff",
			Category:    "general",
			Country:     "us",
			Description: "Uniper (ETR:UN01) Reaches New 1-Year High at $34.54",
			Image:       "",
			Language:    "en",
			PublishedAt: time.Now(),
			Source:      "americanbankingnews",
			Title:       "Uniper (ETR:UN01) Reaches New 1-Year High at $34.54",
			Url:         "https://www.americanbankingnews.com/2021/08/31/uniper-etrun01-reaches-new-1-year-high-at-34-54.html",
		},
	}

	os.Exit(m.Run())
}

func TestGetNews(t *testing.T) {
	t.Run("Get api success", func(t *testing.T) {
		newRepository.On("GetData", mock.AnythingOfType("string")).Return(newsData, nil).Once()
		resp, err := newService.GetNews("hockey")
		assert.Nil(t, err)
		assert.Equal(t, len(newsData), len(resp))
		assert.Equal(t, newsData[0].Author, resp[0].Author)
	})

	t.Run("Get api error", func(t *testing.T) {
		newRepository.On("GetData", mock.AnythingOfType("string")).Return([]news.NewsCore{}, errors.New("error get data from api")).Once()
		resp, err := newService.GetNews("hockey")
		assert.NotNil(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, 0, len(resp))
		assert.Equal(t, "error get data from api", err.Error())
	})
}
