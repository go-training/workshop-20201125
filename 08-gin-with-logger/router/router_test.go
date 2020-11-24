package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
)

func TestFruits(t *testing.T) {
	// create http.Handler
	handler := Handler()

	// run server using httptest
	server := httptest.NewServer(handler)
	defer server.Close()

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)

	// is it working?
	e.GET("/user/1").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		ContainsKey("user").
		ValueEqual("user", "1")
}
