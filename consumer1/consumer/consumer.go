package consumer

import (
	"delivery_tracking_api/consumer1/controller"
	"delivery_tracking_api/consumer1/logger"
	"delivery_tracking_api/consumer1/model"
	"delivery_tracking_api/consumer1/repo"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var consumer *kafka.Consumer
var topic = "orders"

var Repo = &repo.Repo{}

func Init() {
	// var path, key, secret string
	// fmt.Println("Enter Bootstrap Server path : ")
	// fmt.Scanln(&path)
	// fmt.Println("Enter key : ")
	// fmt.Scanln(&key)
	// fmt.Println("Enter secret : ")
	// fmt.Scanln(&secret)
	var err error
	consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		// User-specific properties that you must set
		"bootstrap.servers": "dta-kafka-service.default.svc.cluster.local:9093",
		//"sasl.username":     key,
		//"sasl.password":     secret,

		"group.id":          "consumer-group-A",
		"auto.offset.reset": "latest",
	})

	if err != nil {
		logger.Infoln("Failed to create producer: " + err.Error())
		os.Exit(1)
	}
	consumer.SubscribeTopics([]string{topic}, func(c *kafka.Consumer, e kafka.Event) error {
		switch e := e.(type) {
		case kafka.AssignedPartitions:
			fmt.Println("Partitions Assigned: ", e.Partitions)
			logger.Infoln("Partitions Assigned")
			err := c.Assign(e.Partitions)
			if err != nil {
				fmt.Printf("Error assigning partitions: %s\n", err)
				logger.Infoln("Error assigning partitions: " + err.Error())
			}
		case kafka.RevokedPartitions:
			fmt.Println("Partitions revoked: ", e.Partitions)
			logger.Infoln("Partitions revoked")
			c.Unassign()
		default:
			fmt.Println("unhandled event ", e)
			logger.Infoln("unhandled event " + e.String())
		}
		return nil
	})
	go FetchAll()
}

func FetchAll() {

	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages
	run := true
	for run {
		select {
		case sig := <-sigchan:
			logger.Infoln("Caught signal " + sig.String() + ": terminating")
			fmt.Println("Caught signal " + sig.String() + ": terminating")
			run = false
			consumer.Close()
			os.Exit(15)
		default:
			ev, err := consumer.ReadMessage(100 * time.Millisecond)
			if err != nil {
				continue
			}
			var order model.Order
			json.Unmarshal([]byte(ev.Value), &order)
			controller.AddOrder(order)
			fmt.Println("Read order from topic " + *ev.TopicPartition.Topic + ": key = " + string(ev.Key) + " value = " + string(ev.Value))
			logger.Infoln("Read order from topic " + *ev.TopicPartition.Topic + ": key = " + string(ev.Key) + " value = " + string(ev.Value))
		}
	}

}
