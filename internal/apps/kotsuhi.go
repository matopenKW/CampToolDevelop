package apps

import (
	"cloud.google.com/go/firestore"
	"context"

	"github.com/gin-gonic/gin"

	"google.golang.org/api/iterator"
)

func ViewKotsuhi(client *firestore.Client) (gin.H, error) {

	ctx := context.Background()

	//	list := make([]*Kotsuhi, 1)
	list := make([]map[string]interface{}, 1, 10)

	iter := client.Collection("tomoki").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		list = append(list, doc.Data())

		// data := doc.Data()
		// for key, value := range data {
		// 	fmt.Printf("key: %v, value: %v\n", key, value)
		// }
	}

	return gin.H{
		"title": "KOTSUHI",
		"list":  list,
	}, nil
}

type Kotsuhi struct {
	end, start          string
	roundTripFlg, price int
}
