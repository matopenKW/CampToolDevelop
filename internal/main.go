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

	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	router.Static("../web", ".././web")

	router.LoadHTMLGlob("../templates/*.html")

	// firestore
	client, err := db.OpenFirestore()
	if err != nil {
		log.Fatalf("erro in new db client. reason : %v\n", err)
	}
	defer client.Close()

	router.GET("/", htmlForward(router, client, templatePathMap["login"]))
	router.POST("/", htmlForward(router, client, templatePathMap["login"]))
	router.GET("/login", htmlForward(router, client, templatePathMap["login"]))
	router.POST("/login", htmlForward(router, client, templatePathMap["login"]))

	// Ajaxでチェックをかましてからサブミット→サブミット時にサイドチェックで　クッキー作成とHTMLフォワード
	router.POST("/login:cmd/login", login)

	router.GET("/carfare", htmlForward(router, client, templatePathMap["carfare"]))
	router.POST("/carfare", htmlForward(router, client, templatePathMap["carfare"]))
	router.POST("/carfare:cmd/insert", jsonForward(client))
	router.POST("/carfare:cmd/update", jsonForward(client))
	router.POST("/carfare:cmd/delete", jsonForward(client))

	router.Run()
}

func login(ctx *gin.Context) {

	apps.Login(ctx)

	ctx.Redirect(http.StatusMovedPermanently, "/carfare")

	// forward := htmlForward(router, client, templatePathMap["login"])
	// forward(ctx)

}

func htmlForward(router *gin.Engine, client *firestore.Client, templatePath string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		var err error
		form := map[string]interface{}{}

		actionPath := util.SubstrBefore(util.SubstrAfter(ctx.Request.URL.Path, "/"), ":")

		switch actionPath {
		case "", "login":
			form, err = apps.ExeLogin(ctx.Request, client)
		case "index":
			//form, err = apps.ViewIndex(client)
		case "carfare":
			form, err = apps.ExeCarfare(ctx.Request, client)
		default:
		}

		if err != nil {
			log.Printf("error : %v\n", err)
			code, errForm := errorHandring(err)

			html := template.Must(template.ParseFiles(templatePath, templatePathMap["500"]))
			router.SetHTMLTemplate(html)

			ctx.HTML(code, "base.html", errForm)
		} else {
			html := template.Must(template.ParseFiles(templatePath, templatePathMap["base"]))
			router.SetHTMLTemplate(html)

			ctx.HTML(http.StatusOK, "base.html", form)
		}
	}
}

func jsonForward(client *firestore.Client) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		var err error
		form := map[string]interface{}{}

		actionPath := util.SubstrBefore(util.SubstrAfter(ctx.Request.URL.Path, "/"), ":")

		switch actionPath {
		case "login":
			form, err = apps.ExeLogin(ctx.Request, client)
		case "index":

		case "carfare":
			form, err = apps.ExeCarfare(ctx.Request, client)
		default:
		}

		setJsonFunc := setJson(ctx)
		if err != nil {
			log.Printf("error : %v\n", err)
			setJsonFunc(errorHandring(err))
		} else {
			jsonForm, err := json.Marshal(form)
			if err != nil {
				log.Printf("failed　to　json　convert : %v\n", err)
				setJsonFunc(errorHandring(err))
			} else {
				log.Println(string(jsonForm))
				setJsonFunc(200, form)
			}
		}
	}
}

func errorHandring(err error) (code int, form map[string]interface{}) {
	return http.StatusInternalServerError, map[string]interface{}{
		"errMssage": "error : " + err.Error(),
	}
}

func setJson(ctx *gin.Context) func(code int, obj interface{}) {
	return func(code int, obj interface{}) {
		ctx.JSON(code, obj)
	}
}
