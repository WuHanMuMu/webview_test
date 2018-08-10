package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/zserge/webview"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"web_view/requsets"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func startServer() string {
	port := 8888
	serverUrl := "http://127.0.0.1:" + strconv.Itoa(port)
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}
	e.Renderer = t
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		//return c.String(http.StatusOK, "fffff")
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"serverUrl": serverUrl,
		})
	})

	e.POST("/search", func(context echo.Context) error {
		search := context.FormValue("words")
		reply := requsets.GetGoods(1, 10, search)
		return context.JSON(http.StatusOK, reply)
	})
	go e.Start(":" + strconv.Itoa(port))

	return serverUrl
}

func main() {
	//startServer()
	localS := startServer()
	fmt.Println(localS)
	webview.Open("hello you",
		localS, 800, 600, true)
}
