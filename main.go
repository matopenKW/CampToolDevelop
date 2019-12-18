package main

import (
	"html/template"

	"CampToolDevelop/internal/apps"
	"CampToolDevelop/pkg/db"
	"log"

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

		// firebase接続
		client, err := db.OpenFirebase()
		if err != nil {
			log.Fatalf("erro in new db client. reason : %v\n", err)
		}
		form := gin.H{}

		switch actionPath {
		case "", "index":
			form, err = apps.ViewIndex(client)
		case "kotsuhi":
			form, err = apps.ViewKotsuhi(client)
		default:
		}

		if err != nil {
			log.Fatalf("erro in new db client. reason : %v\n", err)
		}

		ctx.HTML(200, "base.html", form)
	}
}
