package apps

import (
	"CampToolDevelop/pkg/db"
	"CampToolDevelop/pkg/util"
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go/auth"

	"errors"
	"log"
	"net/http"
)

type Carfare struct {
	Date         string
	End          string
	Start        string
	RoundTripFlg string
	Price        string
	Bikou        string
	DocumentId   string
}

func ExeCarfare(req *http.Request, client *firestore.Client, userInfo *auth.UserInfo) (map[string]interface{}, error) {

	url := util.SubstrAfter(req.URL.Path, ":")

	if url == "/carfare" {
		return view(req, client, userInfo)
	} else {
		cmd := util.SubstrAfter(util.SubstrAfter(req.URL.Path, ":"), "/")
		log.Println(cmd)
		switch cmd {
		case "insert":
			return insert(req, client)
		case "update":
			return update(req, client)
		case "delete":
			return delete(req, client)
		default:
			return nil, errors.New("invalid command")
		}

	}
}

func view(req *http.Request, client *firestore.Client, userInfo *auth.UserInfo) (map[string]interface{}, error) {

	req.ParseForm()

	retList := make([]*Carfare, 0, 10)

	orderBy := func() (string, firestore.Direction) {
		return "Date", firestore.Desc
	}

	list, err := db.SelectDocuments(client, userInfo.UID, orderBy)
	if err != nil {
		return nil, err
	}

	for _, v := range list {
		carfare := &Carfare{
			Date:         v["Date"].(string),
			End:          v["End"].(string),
			Start:        v["Start"].(string),
			RoundTripFlg: v["RoundTripFlg"].(string),
			Price:        v["Price"].(string),
			Bikou:        v["Bikou"].(string),
			DocumentId:   v["DocumentId"].(string),
		}

		retList = append(retList, carfare)
	}

	return map[string]interface{}{
		"title":  "CARFACE",
		"userId": userInfo.UID,
		"list":   retList,
	}, nil
}

func insert(req *http.Request, client *firestore.Client) (map[string]interface{}, error) {

	req.ParseForm()

	log.Println(req.Form)

	userID := req.Form["userId"][0]

	carfare := map[string]interface{}{
		"Date":         req.Form["date"][0],
		"End":          req.Form["end"][0],
		"Start":        req.Form["start"][0],
		"RoundTripFlg": "0",
		"Price":        req.Form["price"][0],
		"Bikou":        req.Form["bikou"][0],
	}

	documentID, err := db.InsertDocument(context.Background(), client, userID, carfare)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"param":      "SUCCSES",
		"documentID": documentID,
	}, nil
}

func update(req *http.Request, client *firestore.Client) (map[string]interface{}, error) {

	req.ParseForm()

	userID := req.Form["userId"][0]
	documentID := req.Form["documentId"][0]

	carfare := map[string]interface{}{
		"Date":         req.Form["date"][0],
		"End":          req.Form["end"][0],
		"Start":        req.Form["start"][0],
		"RoundTripFlg": "0",
		"Price":        req.Form["price"][0],
		"Bikou":        req.Form["bikou"][0],
	}

	err := db.UpdateDocument(context.Background(), client, userID, documentID, carfare)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"param": "SUCCSES",
	}, nil
}

func delete(req *http.Request, client *firestore.Client) (map[string]interface{}, error) {
	req.ParseForm()

	userID := req.Form["userId"][0]
	documentID := req.Form["documentId"][0]

	err := db.DeleteDocument(context.Background(), client, userID, documentID)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"param": "SUCCSES",
	}, nil
}
