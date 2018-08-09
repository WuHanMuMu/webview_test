package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/zserge/webview"
	"html/template"
	"io"
	"net/http"
	"strconv"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func startServer() string {
	port := 8888
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}
	e.Renderer = t
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", "")
	})

	e.POST("/search", func(context echo.Context) error {
		return context.JSON(http.StatusOK, map[string]interface{}{
			"data": "fuck",
		})
	})
	go e.Start(":" + strconv.Itoa(port))
	return "http://127.0.0.1:" + strconv.Itoa(port)
}

func main() {
	localS := startServer()
	fmt.Println(localS)
	webview.Open("hello you",
		localS, 800, 600, true)
}
