package apps

import (
	"CampToolDevelop/pkg/db"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

func Login(ctx *gin.Context) error {

	req := ctx.Request

	req.ParseForm()
	uid := req.Form["uid"]
	if isBlank(uid) {
		return errors.New("ログイン情報が不正です。")
	}

	auth, err := db.OpenAuth()
	if err != nil {
		return err
	}

	userRec, err := db.GetUserRecord(auth, uid[0])
	if err != nil {
		return err
	}

	userInfo := *userRec.UserInfo
	if &userInfo == nil {
		return errors.New("ユーザー情報が不正です")
	}

	session := sessions.Default(ctx)
	if session.Get("userId") != userInfo.UID {
		session.Set("userId", userInfo.UID)
		session.Save()
	}

	return nil
}

func isBlank(param []string) bool {
	return param == nil || len(param) != 1 || param[0] == ""
}
