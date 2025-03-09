package repo

import (
	"delivery_tracking_api/consumer2DB/model"
)

type RepoInterface interface {
	InserOrUpdateRecord(key string, value model.Order) string
	FetchAllRecords() map[string]model.Order
	FetchItemByKey() model.Order
}