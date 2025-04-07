package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/*
	This session will cover:

1. Logging using logrus with GIN Framework in go.
2. Default logging in logrus.
3. Change format of the general logs with logrus.
4. Write logs to file.
5. Logging in JSON format in Logrus.
*/
func GetLogrus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello i am logrus",
	})
}
func main() {
	logrus.SetLevel(logrus.TraceLevel) // set lo level

	f, _ := os.Create("logger.log")       //create a file
	multi := io.MultiWriter(f, os.Stdout) // to get logs to file and terminal
	logrus.SetOutput(multi)

	logrus.Traceln("trace it")

	router := gin.New()
	router.GET("/logrus", GetLogrus)
	router.Run()
}
