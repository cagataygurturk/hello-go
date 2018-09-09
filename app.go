package main

import (
	"github.com/cagataygurturk/hello-go/controllers"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	r := gin.New()
	r.Use(
		gin.LoggerWithWriter(getLogWriter()),
		gin.Recovery())

	r.NoRoute(getErrorResponse(http.StatusNotFound, "Not found"));
	r.NoMethod(getErrorResponse(http.StatusMethodNotAllowed, "Method not allowed"))

	r.GET("/ping", controllers.Ping)
	r.Run()
}

func getErrorResponse(statusCode int, message string) func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(statusCode, gin.H{"code": statusCode, "message": message})
	}
}

// Gets Log file reader
func getLogWriter() io.Writer {
	/*
	if gin.Mode() == gin.ReleaseMode {
		logPath := "logs"
		_ = os.Mkdir(logPath, 0777)

		logFile, err := os.Create(logPath + "/server.log")

		if err != nil {
			panic(err)
		}

		return logFile
	}
	*/
	return os.Stdout
}
