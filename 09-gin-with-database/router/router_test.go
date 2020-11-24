package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/gavv/httpexpect"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func TestPing(t *testing.T) {
	// create http.Handler
	handler := Handler()

	// run server using httptest
	server := httptest.NewServer(handler)
	defer server.Close()

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)

	// is it working?
	e.GET("/ping").
		Expect().
		Status(http.StatusOK)
}

func TestPingFromGofight(t *testing.T) {
	r := gofight.New()

	r.GET("/ping").
		// turn on the debug mode.
		SetDebug(true).
		Run(Handler(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestGetUser(t *testing.T) {
	// create http.Handler
	handler := Handler()

	// run server using httptest
	server := httptest.NewServer(handler)
	defer server.Close()

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)

	// is it working?
	obj := e.GET("/user/1").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("user", "id")
	obj.Value("user").Object().ValueEqual("email", "foo@gmail.com")
}
