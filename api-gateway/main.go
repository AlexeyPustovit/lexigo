package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const serviceName = "api-gateway"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("'PORT' env var is not set")
		return
	}

	srv := gin.Default()
	srv.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, serviceName+" is alive")
	})

	err := srv.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
