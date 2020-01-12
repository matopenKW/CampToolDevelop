package apps

import (
	"CampToolDevelop/pkg/db"
	util "CampToolDevelop/pkg/util"
	"cloud.google.com/go/firestore"

	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Login(ctx *gin.Context) (map[string]interface{}, error) {

	req := ctx.Request

	req.ParseForm()

	uid := req.Form["uid"]

	if isBlank(uid) {
		return nil, errors.New("ログイン情報が不正です。")
	}

	auth, err := db.OpenAuth()
	if err != nil {
		return nil, err
	}

	userRec, err := db.GetUserRecord(auth, uid[0])

	if err != nil {
		return nil, err
	}

	userInfo := *userRec.UserInfo

	if &userInfo == nil {
		return nil, errors.New("ユーザー情報が不正です")
	}

	log.Println(userInfo)

	// code from README
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	session := sessions.Default(ctx)

	if session.Get("hello") != "world" {
		session.Set("hello", "world")
		session.Save()
	}

	return map[string]interface{}{
		"TITLE":    "LOGIN",
		"UserInfo": userInfo,
	}, nil
}

func ExeLogin(req *http.Request, client *firestore.Client) (map[string]interface{}, error) {

	cmd := util.SubstrAfter(req.URL.Path, ":")

	log.Println(cmd)

	return map[string]interface{}{}, nil

}

func isBlank(param []string) bool {
	return param == nil || len(param) != 1 || param[0] == ""
}
