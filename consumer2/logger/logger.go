package logger

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Init() {

	//setting log format to json
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint:     true,
		TimestampFormat: "01-02-2006 15:04:05",
	})

	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.InfoLevel)

	//creating a file and logging in that file
	file, err := os.Create("consumer.log")

	if err == nil {
		logrus.SetOutput(file)
	} else {
		logrus.Println("file not created")
	}
}

// adding http request log to file
func LogInfo(message string, c *gin.Context) {
	logrus.WithFields(logrus.Fields{
		"path": c.Request.RequestURI,
		"type": c.Request.Method,
		"host": c.Request.Host,
	}).Info(message)
}

// logging custom messages
func Infoln(message string) {
	logrus.Infoln(message)
}
