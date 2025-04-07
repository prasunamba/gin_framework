package main

import (
	"examplevideo4/logger"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

/*
	This session will cover:

1. Logging with GIN Framework.
2. Default logging in GIN.
3. Change formatting for the log of routes in GIN.
4. Change format of the general logs with GIN.
5. Write logs to file.
6. Controlling log output colouring in console with GIN.
7. Logging in JSON format in GIN.
*/
func getlog(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "for logger",
	})
}
func main() {
	gin.DebugPrintRouteFunc = func(httpMethod string, absolutePath string, handlerName string, nuHandlers int) {
		log.Printf("formatted log %v,%v,%v %v ", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	f, _ := os.Create("logfile.log")                 // create a file
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout) //copy the logs to  file and terminal
	log.SetOutput(gin.DefaultWriter)                 // to log, response codes to the file
	router := gin.Default()
	// router.Use(gin.LoggerWithFormatter(logger.Formatlogs))
	router.Use(gin.LoggerWithFormatter(logger.Jsonformatlogs)) //json formatted logs has advantages

	router.GET("/log", getlog)
	router.Run()
}
