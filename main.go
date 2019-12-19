package main

import (
	"cloud.google.com/go/firestore"
	"html/template"

	"CampToolDevelop/internal/apps"
	"CampToolDevelop/pkg/db"
	"CampToolDevelop/pkg/util"

	"log"

	"github.com/gin-gonic/gin"
	"io/ioutil"
)

var templatePathMap map[string]string

func init() {

	templatePathMap = make(map[string]string)
	templateDir := "templates/"

	files, err := ioutil.ReadDir(templateDir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		templatePathMap[util.SubstrBefore(file.Name(), ".")] = templateDir + file.Name()
	}
}

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

	router.GET("/", htmlForward(router, client, templatePathMap["index"]))
	router.POST("/", htmlForward(router, client, templatePathMap["index"]))

	router.GET("/index", htmlForward(router, client, templatePathMap["index"]))
	router.POST("/index", htmlForward(router, client, templatePathMap["index"]))

	router.GET("/kotsuhi", htmlForward(router, client, templatePathMap["kotsuhi"]))
	router.POST("/kotsuhi", htmlForward(router, client, templatePathMap["kotsuhi"]))
	router.POST("kotsuhi:regist", htmlForward(router, client, templatePathMap["kotsuhi"]))

	router.Run()
}

func htmlForward(router *gin.Engine, client *firestore.Client, templatePath string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		html := template.Must(template.ParseFiles(templatePath, templatePathMap["base"]))
		router.SetHTMLTemplate(html)

		var err error
		form := gin.H{}

		actionPath := util.SubstrBefore(util.SubstrAfter(ctx.Request.URL.Path, "/"), ":")

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
