package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"workuo/features/news"
)

type NewsApi struct {
	URL    string
	Client http.Client
	Key    string
}

func NewNewsApiRepository(url string, key string) news.Repository {
	return &NewsApi{
		URL:    url,
		Client: http.Client{},
		Key:    key,
	}
}

func (nr *NewsApi) GetData(keyword string) ([]news.NewsCore, error) {
	url := fmt.Sprintf("%v?access_key=%v&languages=en&keyword=%v", nr.URL, nr.Key, keyword)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var newsData Data

	response, err := nr.Client.Do(request)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(response.Body).Decode(&newsData)
	if err != nil {
		return nil, err
	}

	return ToCoreList(newsData.Data), nil
}
