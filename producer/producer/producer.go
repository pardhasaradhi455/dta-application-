package producer

import (
	"delivery_tracking_api/producer/logger"
	"delivery_tracking_api/producer/model"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	kafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

var topic = "orders"
var producer *kafka.Producer

func Init() {
	//var path, key, secret string
	// fmt.Println("Enter Bootstrap Server path : ")
	// fmt.Scanln(&path)
	// fmt.Println("Enter key : ")
	// fmt.Scanln(&key)
	// fmt.Println("Enter secret : ")
	// fmt.Scanln(&secret)
	var err error
	//initializing producer
	producer, err = kafka.NewProducer(&kafka.ConfigMap{
		// User-specific properties that you must set
		"bootstrap.servers": "dta-kafka-service.default.svc.cluster.local:9093",
		//"sasl.username":     key,
		//"sasl.password":     secret,

		// Fixed properties
		// "security.protocol": "SASL_SSL",
		// "sasl.mechanisms":   "PLAIN",
	})

	if err != nil {
		logger.Infoln("Failed to create producer: " + err.Error())
		os.Exit(1)
	}
}

func Log() {
	for e := range producer.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				logger.Infoln("Failed to produce event for order : " + ev.TopicPartition.String())
				fmt.Println("Failed to produce event for order : " + ev.TopicPartition.String())
			} else {
				fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
					*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				logger.Infoln("Produced event to topic : " + *ev.TopicPartition.Topic + " key : " + string(ev.Key) + " value : " + string(ev.Value))
				fmt.Println("Produced event to topic : " + *ev.TopicPartition.Topic + " key : " + string(ev.Key) + " value : " + string(ev.Value))
			}
		default:
			logger.Infoln(ev.String())
		}
	}
}

func ProduceEvent(order model.Order) {
	j, _ := json.Marshal(order)
	partition, _ := strconv.Atoi(order.OrderId)
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: int32(partition)%5},
		Key:            []byte(order.OrderId),
		Value:          j,
	}, nil)
}

func ProduceEvents(orders []model.Order) {
	for _, order := range orders {
		ProduceEvent(order)
	}
}
