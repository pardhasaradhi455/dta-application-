package main

import (
	"delivery_tracking_api/consumer1DB/router"
	"delivery_tracking_api/consumer1DB/logger"
)

func main() {
	logger.Init()
	logger.Infoln("logger initialized")

	logger.Infoln("initializing router")
	router.Init()
}