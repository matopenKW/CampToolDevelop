package apps

import (
	"cloud.google.com/go/firestore"
	"context"

	"github.com/gin-gonic/gin"

	"google.golang.org/api/iterator"
)

func ViewKotsuhi(client *firestore.Client) (gin.H, error) {

	ctx := context.Background()

	list := make([]*Kotsuhi, 0, 10)

	iter := client.Collection("tomoki").Documents(ctx)
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

		// data := doc.Data()
		// for key, value := range data {
		// 	fmt.Printf("key: %v, value: %v\n", key, value)
		// }
	}
	defer client.Close()

	return gin.H{
		"title": "KOTSUHI",
		"list":  list,
	}, nil
}

type Kotsuhi struct {
	End          string
	Start        string
	RoundTripFlg int64
	Price        int64
}
