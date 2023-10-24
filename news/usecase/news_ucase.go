package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/mayapada/models"
	news "github.com/mayapada/news"
)

type NewsUsecase struct {
	newsRepo       news.NewsRepositoryI
	contextTimeout time.Duration
}

func NewNewsUsecase(a news.NewsRepositoryI, timeout time.Duration) news.NewsUsecaseI {
	return &NewsUsecase{
		newsRepo:       a,
		contextTimeout: timeout,
	}
}

func formatDate(value string) (a time.Time, err error) {
	layoutFormat := "2006-01-02 15:04:05"
	date, err := time.Parse(layoutFormat, value)
	if err != nil {
		return date, err
	}

	return date, nil
}

func (a *NewsUsecase) GetNews(c context.Context) (result []*models.GetNews, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	resNews, err := a.newsRepo.GetNews(ctx)
	if err != nil {
		return nil, err
	}

	if len(resNews) > 0 {
		for _, v := range resNews {
			var new models.GetNews
			new.AccountNumber = v.AccountNumber
			new.Amount = v.Amount
			new.TransactionDate = v.TransactionDate.Format("2006-01-02 15:04:05")
			new.FraudCountTrxL7d = v.FraudCountTrxL7d
			new.FraudCountTrxL30d = v.FraudCountTrxL30d

			result = append(result, &new)
		}
	}

	return result, nil
}

func (a *NewsUsecase) AddNews(c context.Context, req models.AddNews) error {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	date, err := formatDate(req.TransactionDate)
	fmt.Println("err: ", err)
	if err != nil {
		return err
	}

	// add to news table
	reqAddNews := models.News{
		AccountNumber:     req.AccountNumber,
		Amount:            req.Amount,
		TransactionDate:   date,
		FraudCountTrxL7d:  req.FraudCountTrxL7d,
		FraudCountTrxL30d: req.FraudCountTrxL30d,
	}

	_, err = a.newsRepo.AddNews(ctx, &reqAddNews)
	fmt.Println("err2: ", err)
	if err != nil {
		return err
	}

	return nil
}
