package data

import (
	"context"

	"github.com/mayapada/models"
)

type NewsRepositoryI interface {
	GetNews(ctx context.Context) (res []*models.News, err error)
	AddNews(ctx context.Context, req *models.News) (int64, error)
}
