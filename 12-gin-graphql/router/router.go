package router

import (
	"net/http"
	"strconv"

	"gin-http-server/model"
	"gin-http-server/router/graphql"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Handler init
func Handler() http.Handler {
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

	r.GET("/users", func(c *gin.Context) {

		log.Debug().Msg("create user")
		if err := model.CreateUser(&model.User{
			Name:  "foo",
			Email: "bar@gmail.com",
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"errors": err,
			})
			return
		}
		users, err := model.FindAllUsers()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	})

	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		i, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err,
			})
			return
		}

		user := model.FindUserByID(i)
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})

	r.POST("/user", func(c *gin.Context) {
		data := &model.User{
			Name:  "foo",
			Email: "foo@gmail.com",
		}
		c.JSON(http.StatusOK, gin.H{
			"user": data,
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
		i, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err,
			})
			return
		}
		user := model.DeleteUser(i)
		c.JSON(http.StatusOK, gin.H{
			"user":    user,
			"success": "ok",
			"id":      id,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	g := r.Group("/graphql")
	g.Use()
	{
		g.POST("", graphql.Handler())
		g.GET("", graphql.Handler())
	}

	return r
}
