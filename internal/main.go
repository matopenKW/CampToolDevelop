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

	"encoding/json"
)

var templatePathMap map[string]string

func init() {

	templatePathMap = make(map[string]string)
	templateDir := "../templates/"

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
	router.Static("../web", ".././web")

	router.LoadHTMLGlob("../templates/*.html")

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

	router.GET("/carfare", htmlForward(router, client, templatePathMap["carfare"]))
	router.POST("/carfare", htmlForward(router, client, templatePathMap["carfare"]))
	router.POST("/carfare:cmd/insert", jsonForward(client))
	router.POST("/carfare:cmd/update", jsonForward(client))
	router.POST("/carfare:cmd/delete", jsonForward(client))

	router.Run()
}

func htmlForward(router *gin.Engine, client *firestore.Client, templatePath string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		html := template.Must(template.ParseFiles(templatePath, templatePathMap["base"]))
		router.SetHTMLTemplate(html)

		var err error
		form := map[string]interface{}{}

		actionPath := util.SubstrBefore(util.SubstrAfter(ctx.Request.URL.Path, "/"), ":")

		switch actionPath {
		case "", "index":
			//form, err = apps.ViewIndex(client)
		case "carfare":
			form, err = apps.ExeCarfare(ctx.Request, client)
		default:
		}

		if err != nil {
			log.Fatalf("erro in new db client. reason : %v\n", err)
		}
		ctx.HTML(200, "base.html", form)
	}
}

func jsonForward(client *firestore.Client) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		form, err := apps.ExeCarfare(ctx.Request, client)

		if err != nil {
			log.Fatalf("erro in new db client. reason : %v\n", err)
		}

		jsonForm, err := json.Marshal(form)
		if err != nil {
			log.Fatalf("erro in new db client. reason : %v\n", err)
		}

		log.Println(string(jsonForm))
		ctx.JSON(200, string(jsonForm))
	}
}
