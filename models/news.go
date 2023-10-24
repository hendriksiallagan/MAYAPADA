package models

import "time"

type News struct {
	AccountNumber     string    `json:"account_number"`
	Amount            float64   `json:"trx_amount"`
	TransactionDate   time.Time `json:"trx_date"`
	FraudCountTrxL7d  float64   `json:"fraud_c_count_trx_l7d"`
	FraudCountTrxL30d float64   `json:"fraud_c_count_trx_l30d"`
}

type GetNews struct {
	AccountNumber     string  `json:"account_number"`
	Amount            float64 `json:"trx_amount"`
	TransactionDate   string  `json:"trx_date"`
	FraudCountTrxL7d  float64 `json:"fraud_c_count_trx_l7d"`
	FraudCountTrxL30d float64 `json:"fraud_c_count_trx_l30d"`
}

type AddNews struct {
	AccountNumber     string  `json:"account_number"`
	Amount            float64 `json:"trx_amount"`
	TransactionDate   string  `json:"trx_date"`
	FraudCountTrxL7d  float64 `json:"fraud_c_count_trx_l7d"`
	FraudCountTrxL30d float64 `json:"fraud_c_count_trx_l30d"`
}
