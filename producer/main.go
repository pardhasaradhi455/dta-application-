package main

import (
	"delivery_tracking_api/producer/logger"
	"delivery_tracking_api/producer/router"
)

func main(){

	logger.Init()
	logger.Infoln("logger initialized")

	logger.Infoln("router initializing")
	router.Init()
}