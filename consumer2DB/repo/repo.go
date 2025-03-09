package repo

import (
	"delivery_tracking_api/consumer2DB/logger"
	"delivery_tracking_api/consumer2DB/model"

	badger "github.com/dgraph-io/badger/v4"
)

type Repo struct {
	RepoInterface
}

var db *badger.DB
var err error

func Init() {
	db, err = badger.Open(badger.DefaultOptions("tmp/badger"))
	if err != nil {
		logger.Infoln("unable to open db")
	}
	logger.Infoln("db opened successfully")

	//defer db.Close()
}

func (ctrl *Repo) InserOrUpdateRecord(key string, value model.Order) string {
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), value.ToBytes())
		return err
	})
	if err != nil {
		return err.Error()
	}
	logger.Infoln("Record Inserted Successfully")
	return "Record Inserted Successfully"
}

func (ctrl *Repo) FetchAllRecords() map[string]model.Order {
	orders :=  map[string]model.Order{}
	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				orders[string(k)] = model.ByteToOrder(v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	logger.Infoln("All records fetched successfully")
	return orders
}

func (ctrl *Repo) FetchItemByKey(key string) model.Order {
	var order model.Order
	db.View(func(txn *badger.Txn) error {
		item, _ := txn.Get([]byte(key))
		err := item.Value(func(v []byte) error {
			order = model.ByteToOrder(v)
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})
	logger.Infoln("Records fetched successfully with key : " + key)
	return order
}
