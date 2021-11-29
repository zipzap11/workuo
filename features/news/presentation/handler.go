package presentation

import (
	"net/http"
	"workuo/features/news"

	"github.com/labstack/echo/v4"
)

type NewsHandler struct {
	newsService news.Service
}

func NewNewsHandler(ns news.Service) *NewsHandler {
	return &NewsHandler{ns}
}

func (ns *NewsHandler) GetNewsHandler(e echo.Context) error {
	data, err := ns.newsService.GetNews()
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    data,
	})
}
