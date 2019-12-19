package main

import (
	"cloud.google.com/go/firestore"
	"html/template"

	"CampToolDevelop/internal/apps"
	"CampToolDevelop/pkg/db"
	"CampToolDevelop/pkg/util"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Static("assets", "./assets")

	router.LoadHTMLGlob("templates/*.html")

	// firebase接続
	client, err := db.OpenFirebase()
	if err != nil {
		log.Fatalf("erro in new db client. reason : %v\n", err)
	}
	defer client.Close()

	router.POST("/", htmlForward(router, client))
	router.GET("/", htmlForward(router, client))

	router.POST("/index", htmlForward(router, client))
	router.GET("/index", htmlForward(router, client))

	router.POST("kotsuhi:regist", htmlForward(router, client))
	router.GET("/kotsuhi", htmlForward(router, client))

	router.Run()
}

func htmlForward(router *gin.Engine, client *firestore.Client) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path

		htmlPath := util.SubstrBefore(path, ":")
		log.Println(htmlPath)
		if htmlPath == "/" {
			htmlPath = htmlPath + "index"
		}

		actionPath := util.SubstrAfter(path, "/")

		html := template.Must(template.ParseFiles("templates"+htmlPath+".html", "templates/base.html"))
		router.SetHTMLTemplate(html)

		var err error
		form := gin.H{}

		switch actionPath {
		case "", "index":
			//form, err = apps.ViewIndex(client)
		case "kotsuhi":
			form, err = apps.ExeKotsuhi(ctx.Request, client)
		default:
		}

		if err != nil {
			log.Fatalf("erro in new db client. reason : %v\n", err)
		}

		ctx.HTML(200, "base.html", form)
	}
}
