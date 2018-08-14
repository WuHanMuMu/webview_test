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
	e.Use(middleware.Logger())
	e.Static("/", "view")
	e.POST("/search", func(context echo.Context) error {
		//var searchParam requests.SearchParam
		search := context.FormValue("words")
		page := context.FormValue("page")
		//if err := context.Bind(searchParam); err != nil {
		//	fmt.Println(searchParam, "========>", searchParam.Words)
		//}
		pageIndex, _ := strconv.Atoi(page)
		fmt.Println("serach", search, page)
		reply := requests.GetGoodsFromHtml(search, pageIndex)
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
