package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WangYihang/go-net-conn-logger/pkg/listener"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new listener
	l, err := listener.NewListener(":8080")
	if err != nil {
		panic(err)
	}

	// Start the log consumer
	go func() {
		for logEntry := range l.LogChan() {
			data, _ := json.Marshal(logEntry)
			fmt.Println(string(data))
		}
	}()

	// Start the HTTP server
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	http.Serve(l, router)
}
