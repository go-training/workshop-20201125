package main

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	var port = flag.String("port", "8080", "server port")
	var server = flag.Bool("server", false, "enable server")
	flag.Parse()

	r := gin.Default()
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"user": id,
		})
	})

	r.POST("/user", func(c *gin.Context) {
		data := &user{
			Name:  "foo",
			Email: "foo@gmail.com",
		}
		c.JSON(http.StatusOK, gin.H{
			"user": data,
		})
	})

	r.PUT("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		data := &user{
			Name:  "foo",
			Email: "foo@gmail.com",
		}
		c.JSON(http.StatusOK, gin.H{
			"user": data,
			"id":   id,
		})
	})

	r.DELETE("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		data := &user{
			Name:  "foo",
			Email: "foo@gmail.com",
		}
		c.JSON(http.StatusOK, gin.H{
			"user": data,
			"id":   id,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/healthz", func(c *gin.Context) {
		c.AbortWithStatus(200)
	})

	if *server {
		r.Run(":" + *port) // listen on 0.0.0.0:8080
	}
}
