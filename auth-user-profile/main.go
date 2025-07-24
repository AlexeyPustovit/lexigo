package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

const serviceName = "auth-user-profile"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("'PORT' env var is not set")
		return
	}

	srv := gin.Default()
	srv.GET("/ping", func(c *gin.Context) {
		c.String(200, fmt.Sprintf("%s is alive", serviceName))
	})

	srv.Run(fmt.Sprintf(":%s", port))
}
