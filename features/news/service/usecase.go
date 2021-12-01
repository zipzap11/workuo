package service

import "workuo/features/news"

type ApiService struct {
	newsRepository news.Repository
}

func NewApiService(nr news.Repository) news.Service {
	return &ApiService{nr}
}

func (as *ApiService) GetNews(keyword string) ([]news.NewsCore, error) {
	data, err := as.newsRepository.GetData(keyword)
	if err != nil {
		return nil, err
	}

	return data, nil
}
