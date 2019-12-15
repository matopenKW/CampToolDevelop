package main

import (
	"html/template"

	"CampToolDevelop/internal/apps"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Static("assets", "./assets")

	router.LoadHTMLGlob("templates/*.html")

	router.GET("/index", htmlForward(router))
	router.GET("/kotsuhi", htmlForward(router))

	router.Run()
}

func htmlForward(router *gin.Engine) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		html := template.Must(template.ParseFiles("templates"+path+".html", "templates/base.html"))
		router.SetHTMLTemplate(html)
		actionPath := path[1:len(path)]

		form := gin.H{}

		switch actionPath {
		case "", "index":
			form = apps.ViewIndex()
		case "kotsuhi":
			form = apps.ViewKotsuhi()
		default:
		}

		ctx.HTML(200, "base.html", form)
	}
}
