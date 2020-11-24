package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User data
type User struct {
	ID    int64  `json:"user_id"`
	Name  string `json:"name222"`
	Email string `json:"email333"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/user", func(c *gin.Context) {
		user := &User{
			ID:    100,
			Name:  "appleboy",
			Email: "test@gmail.com",
		}
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})

	r.GET("/healthz", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})

	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}
