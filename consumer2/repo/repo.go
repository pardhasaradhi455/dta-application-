package repo

import (
	"bytes"
	"delivery_tracking_api/consumer2/logger"
	"delivery_tracking_api/consumer2/model"
	"io"
	"net/http"
)

type Repo struct {
	RepoInterface
}

func (ctrl *Repo) InserOrUpdateRecord(key string, value model.Order){
	jsonBody := value.ToBytes()
	bodyReader := bytes.NewReader(jsonBody)
	url := "http://localhost:7072/db/insert"
	response, err := http.Post(url, "application/json", bodyReader)
	if err!=nil{
		logger.Infoln("unable to insert record in : " + url)
	}
	defer response.Body.Close()
	logger.Infoln("Record Inserted Successfully into : " + url)
}

func (ctrl *Repo) FetchAllRecords() map[string]model.Order {
	orders := map[string]model.Order{}
	url :=  "http://localhost:7072/db/fetch"
	response, err := http.Get(url)
	if err != nil {
		logger.Infoln("unable to fetch records from : " + url)
		return orders
	}
	defer response.Body.Close()
	logger.Infoln("Records fetched Successfully from : " + url)
	content, _ := io.ReadAll(response.Body)
	return model.ByteToOrders(content)
}

func (ctrl *Repo) FetchItemByKey(key string) model.Order {
	order := model.Order{}
	url := "http://localhost:7072/db/fetch/" + key
	response, err := http.Get(url)
	if err != nil {
		logger.Infoln("unable to fetch record from : " + url)
		return order
	}
	defer response.Body.Close()
	logger.Infoln("Record fetched Successfully from : " + url)
	content, _ := io.ReadAll(response.Body)
	return model.ByteToOrder(content)
}
