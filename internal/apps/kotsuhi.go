package apps

import (
	"CampToolDevelop/pkg/util"
	"cloud.google.com/go/firestore"
	"context"
	"github.com/gin-gonic/gin"
	"log"

	"errors"
	"google.golang.org/api/iterator"
	"net/http"
)

type Kotsuhi struct {
	End          string
	Start        string
	RoundTripFlg int64
	Price        int64
}

func ExeKotsuhi(req *http.Request, client *firestore.Client) (gin.H, error) {

	cmd := util.SubstrAfter(req.URL.Path, ":")

	log.Println(cmd)

	switch cmd {
	case "/kotsuhi":
		return view(req, client)
	case "regist":
		return regist(req, client)
	default:
		return nil, errors.New("invalid command")
	}
}

func view(req *http.Request, client *firestore.Client) (gin.H, error) {

	iter := client.Collection("tomoki").Documents(context.Background())

	// array for return
	list := make([]*Kotsuhi, 0, 10)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		data := doc.Data()

		kotsuhi := &Kotsuhi{
			data["end"].(string),
			data["start"].(string),
			data["round_trip_flg"].(int64),
			data["price"].(int64),
		}

		list = append(list, kotsuhi)
	}

	return gin.H{
		"title": "KOTSUHI",
		"list":  list,
	}, nil
}

func regist(req *http.Request, client *firestore.Client) (gin.H, error) {

	return view(req, client)
}
