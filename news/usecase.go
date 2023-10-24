package data

import (
	"context"

	"github.com/mayapada/models"
)

type NewsUsecaseI interface {
	GetNews(ctx context.Context) (res []*models.GetNews, err error)
	AddNews(c context.Context, req models.AddNews) error
}
