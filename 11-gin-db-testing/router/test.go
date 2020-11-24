package router

import (
	"gin-http-server/model"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func userCheck() gin.HandlerFunc {
	log.Debug().Msg("init userCheck middleware")
	return func(c *gin.Context) {
		u := model.User{
			Email: "bar@gmail.com",
			ID:    100,
			Name:  "bar sir",
		}
		c.Set("user", u)
		c.Next()
	}
}
