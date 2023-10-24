package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/mayapada/models"
	news "github.com/mayapada/news"
)

type HttpNewsHandler struct {
	NUsecase news.NewsUsecaseI
}

func NewNewsHttpHandler(e *echo.Echo, as news.NewsUsecaseI) {
	handler := &HttpNewsHandler{
		NUsecase: as,
	}
	e.GET("/", handler.GetNews)
	e.POST("/", handler.AddNews)
}

func (a *HttpNewsHandler) GetNews(c echo.Context) error {

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	data, err := a.NUsecase.GetNews(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, data)
}

func (a *HttpNewsHandler) AddNews(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	fmt.Println("mashook")
	var req models.AddNews
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{Code: http.StatusUnprocessableEntity, Message: http.StatusText(http.StatusUnprocessableEntity)})
	}

	fmt.Println("mashook2")
	err = a.NUsecase.AddNews(ctx, req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, models.Response{Code: http.StatusUnprocessableEntity, Message: http.StatusText(http.StatusUnprocessableEntity)})
	}

	fmt.Println("mashook3")
	return c.JSON(http.StatusOK, models.Response{Code: http.StatusOK, Message: http.StatusText(http.StatusOK)})
}
