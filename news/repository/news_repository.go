package repository

import (
	"context"
	"database/sql"

	"github.com/labstack/gommon/log"
	"github.com/mayapada/models"
	news "github.com/mayapada/news"
)

type NewsRepository struct {
	Conn *sql.DB
}

func NewNewsRepository(Conn *sql.DB) news.NewsRepositoryI {
	return &NewsRepository{Conn}
}

func (m *NewsRepository) GetNews(ctx context.Context) (res []*models.News, err error) {
	var rows *sql.Rows
	query := `select account_number, trx_amount, trx_date, fraud_c_count_trx_l7d, fraud_c_count_trx_l30d from transaction `
	rows, err = m.Conn.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*models.News, 0)
	for rows.Next() {
		s := new(models.News)
		err = rows.Scan(
			&s.AccountNumber,
			&s.Amount,
			&s.TransactionDate,
			&s.FraudCountTrxL7d,
			&s.FraudCountTrxL30d,
		)
		if err != nil {
			log.Errorf("error %v", err)
			return nil, err
		}

		result = append(result, s)
	}

	return result, nil
}

func (m *NewsRepository) AddNews(ctx context.Context, req *models.News) (int64, error) {
	query := `INSERT transaction SET account_number=?, trx_amount=?, trx_date=?, fraud_c_count_trx_l7d=?, fraud_c_count_trx_l30d=?  `
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, req.AccountNumber, req.Amount, req.TransactionDate, req.FraudCountTrxL7d, req.FraudCountTrxL30d)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
