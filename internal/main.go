package main

import (
	"CampToolDevelop/internal/apps"
	"CampToolDevelop/pkg/db"
	"CampToolDevelop/pkg/util"
	"cloud.google.com/go/firestore"
	"encoding/json"
	"firebase.google.com/go/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
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

	// setting session
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))

	// setting template
	router.Static("../web", ".././web")
	router.LoadHTMLGlob("../templates/*.html")

	// firestore
	client, err := db.OpenFirestore()
	if err != nil {
		log.Fatalf("erro in new db client. reason : %v\n", err)
	}
	defer client.Close()

	// ログイン画面
	router.GET("/", viewLogin(router, client, templatePathMap["login"]))
	router.POST("/", viewLogin(router, client, templatePathMap["login"]))
	router.GET("/login", viewLogin(router, client, templatePathMap["login"]))
	router.POST("/login", viewLogin(router, client, templatePathMap["login"]))

	// ログイン処理
	router.POST("/login:cmd/login", login)

	router.GET("/carfare", htmlForward(router, client, templatePathMap["carfare"]))
	router.POST("/carfare", htmlForward(router, client, templatePathMap["carfare"]))
	router.POST("/carfare:cmd/insert", jsonForward(client))
	router.POST("/carfare:cmd/update", jsonForward(client))
	router.POST("/carfare:cmd/delete", jsonForward(client))

	router.Run()
}

func viewLogin(router *gin.Engine, client *firestore.Client, templatePath string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		html := template.Must(template.ParseFiles(templatePath, templatePathMap["base"]))
		router.SetHTMLTemplate(html)

		ctx.HTML(http.StatusOK, "base.html", gin.H{})
	}
}

func login(ctx *gin.Context) {
	err := apps.Login(ctx)
	if err != nil {
		log.Println(err)
		ctx.Redirect(http.StatusMovedPermanently, "/index")
	} else {
		ctx.Redirect(http.StatusMovedPermanently, "/carfare")
	}
}

func htmlForward(router *gin.Engine, client *firestore.Client, templatePath string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		var err error
		form := map[string]interface{}{}

		actionPath := util.SubstrBefore(util.SubstrAfter(ctx.Request.URL.Path, "/"), ":")

		session := sessions.Default(ctx)
		userId := session.Get("userId")
		log.Println(userId)

		authClient, err := db.OpenAuth()
		if err != nil {
			log.Println(err)
		}

		var userRec *auth.UserRecord
		if userId != nil {
			userRec, _ = db.GetUserRecord(authClient, userId.(string))
		} else {
			log.Println("session time out")
		}

		switch actionPath {
		case "index":
			//form, err = apps.ViewIndex(client)
		case "carfare":
			form, err = apps.ExeCarfare(ctx.Request, client, userRec.UserInfo)
		default:
		}

		if err != nil {
			log.Printf("error : %v\n", err)
			code, errForm := errorHandring(err)

			html := template.Must(template.ParseFiles(templatePath, templatePathMap["500"]))
			router.SetHTMLTemplate(html)

			ctx.HTML(code, "base.html", errForm)
		} else {

			if userRec.UserInfo.DisplayName == "" {
				userRec.UserInfo.DisplayName = "未設定"
			}

			form["userInfo"] = userRec.UserInfo

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
		case "index":
		case "carfare":
			form, err = apps.ExeCarfare(ctx.Request, client, nil)
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
