package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
)

func Test(t *testing.T) {
	e := httptest.New(t, IPFSWebPublisher())
	// e.GET("/").Expect().Status(httptest.StatusOK.
	// 	ContentType("text/html", "utf-8").Body().Equal("<h1>Welcome</h1>")

	// TODO: refer to https://github.com/iris-contrib/examples/blob/master/mvc/hello-world/main.go and make this work.
	e.GET("/ping").Expect().Status(httptest.StatusOK).
		Body().Equal("pong")
}
