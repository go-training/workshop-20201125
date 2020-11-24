package router

import (
	"net/http"

	"gin-http-server/model"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Handler init
func Handler(db *gorm.DB) http.Handler {
	log.Debug().Msg("load handler")
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	r.Use(test01())
	r.Use(test02())

	r.GET("/me", userCheck(), func(c *gin.Context) {
		if u, ok := c.Get("user"); ok {
			c.JSON(http.StatusOK, gin.H{
				"user": u,
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "not found",
		})
		return
	})

	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")

		user := &model.User{}
		if err := db.First(user, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})

	r.POST("/user", func(c *gin.Context) {
		user := &model.User{
			Name:  "test",
			Email: "test@gmail.com",
		}

		if err := db.Create(user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})

	r.PUT("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		data := &model.User{
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
		data := &model.User{
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

	return r
}
