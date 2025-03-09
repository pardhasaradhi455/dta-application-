package main

import (
	"delivery_tracking_api/consumer2DB/router"
	"delivery_tracking_api/consumer2DB/logger"
)

func main() {
	logger.Init()
	logger.Infoln("logger initialized")

	logger.Infoln("initializing router")
	router.Init()
}