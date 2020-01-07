package apps

import (
	_ "CampToolDevelop/pkg/db"
	"CampToolDevelop/pkg/util"
	"cloud.google.com/go/firestore"
	_ "context"

	"errors"
	"log"
	"net/http"
)

func ExeLogin(req *http.Request, client *firestore.Client) (map[string]interface{}, error) {

	url := util.SubstrAfter(req.URL.Path, ":")

	log.Println(url)

	if url == "login" {
		req.ParseForm()

		email := req.Form["email"][0]
		password := req.Form["password"][0]

		log.Println(email)
		log.Println(password)

		// ログイン時にエラーの場合
		if true {
			return nil, errors.New("login error")
		}

		return map[string]interface{}{
			"title": "LOGIN",
		}, errors.New("invalid command")

	} else {
		return map[string]interface{}{}, nil
	}
}

func login(req *http.Request, client *firestore.Client) (map[string]interface{}, error) {
	return map[string]interface{}{
		"title":  "LOGIN",
		"userId": "userID",
		"list":   "retList",
	}, nil
}
