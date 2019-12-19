package apps

import (
	"CampToolDevelop/pkg/db"
	"CampToolDevelop/pkg/util"
	"cloud.google.com/go/firestore"
	"context"
	"github.com/gin-gonic/gin"

	"errors"
	"google.golang.org/api/iterator"
	"net/http"
)

type Kotsuhi struct {
	End          string
	Start        string
	RoundTripFlg int64
	Price        string
	Bikou        string
}

func ExeKotsuhi(req *http.Request, client *firestore.Client) (gin.H, error) {

	cmd := util.SubstrAfter(req.URL.Path, ":")

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

	req.ParseForm()

	var userId string
	if req.Form["userId"] != nil {
		userId = req.Form["userId"][0]
	} else {
		userId = ""
	}

	iter := client.Collection(userId).Documents(context.Background())

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
			data["End"].(string),
			data["Start"].(string),
			nonNilInt(data["roundTripFlg"]),
			data["Price"].(string),
			"",
		}

		list = append(list, kotsuhi)
	}

	return gin.H{
		"title":  "KOTSUHI",
		"userId": userId,
		"list":   list,
	}, nil
}

func regist(req *http.Request, client *firestore.Client) (gin.H, error) {

	ctx := context.Background()

	ref := client.Collection("tomoki")
	db.DeleteCollection(ctx, client, ref, 10)

	req.ParseForm()
	rowSize := len(req.Form["start"])
	for i := 0; i < rowSize; i++ {

		kotsuhi := &Kotsuhi{
			req.Form["end"][i],
			req.Form["start"][i],
			//req.Form["round_trip_flg"][i].(int64),
			0,
			// int に置換する事
			req.Form["plice"][i],
			req.Form["bikou"][i],
		}

		ref := client.Collection("tomoki").NewDoc()
		_, err := ref.Set(ctx, kotsuhi)

		//		_, _, err := client.Collection("tomoki").Add(context.Background(), kotsuhi)

		if err != nil {
			return nil, err
		}

	}

	return view(req, client)
}

func nonNilInt(t interface{}) int64 {

	if t != nil {
		return t.(int64)
	} else {
		return 0
	}
}
