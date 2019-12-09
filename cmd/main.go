package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Static("../assets", "./assets")

	router.LoadHTMLGlob("../templates/*.html")

	router.GET("/index", func(ctx *gin.Context) {
		html := template.Must(template.ParseFiles("../templates/index.html", "../templates/base.html"))
		router.SetHTMLTemplate(html)
		ctx.HTML(200, "base.html", gin.H{})
	})

	router.GET("/kotsuhi", func(ctx *gin.Context) {
		html := template.Must(template.ParseFiles("../templates/kotsuhi.html", "../templates/base.html"))
		router.SetHTMLTemplate(html)
		ctx.HTML(200, "base.html", gin.H{})
	})

	router.Run()
}
