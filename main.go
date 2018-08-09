package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/zserge/webview"
	"net/http"
	"strconv"
)

func startServer() string {
	port := 8888
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	go e.Start(":" + strconv.Itoa(port))
	return "http://127.0.0.1:8888"
}

func main() {
	localS := startServer()
	fmt.Println(localS)
	webview.Open("hello you",
		localS, 800, 600, true)
}
