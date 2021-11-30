package presentation

import (
	"net/http"
	"workuo/features/news"
	"workuo/helper"

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
		return helper.ErrorResponse(e, http.StatusInternalServerError, "something went wrong", nil)
	}

	return helper.SuccessResponse(e, data)
}
