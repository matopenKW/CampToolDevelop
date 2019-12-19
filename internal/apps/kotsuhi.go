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
	Price        string
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
		log.Println(data)

		kotsuhi := &Kotsuhi{
			isNotNil(data["End"]),
			isNotNil(data["Start"]),
			isNotNilNum(data["roundTripFlg"]),
			isNotNil(data["Price"]),
		}

		list = append(list, kotsuhi)
	}

	return gin.H{
		"title": "KOTSUHI",
		"list":  list,
	}, nil
}

func regist(req *http.Request, client *firestore.Client) (gin.H, error) {

	req.ParseForm()

	rowSize := len(req.Form["start"])

	_, err := client.Collection("tomoki").Doc("DC").Delete(context.Background())
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	for i := 0; i < rowSize; i++ {
		log.Printf("start:%v %T/n", req.Form["start"][i], req.Form["start"][i])
		log.Printf("end:%v %T/n", req.Form["end"][i], req.Form["end"][i])
		log.Printf("plice:%v %T/n", req.Form["plice"][i], req.Form["plice"][i])
		log.Printf("bikou:%v %T/n", req.Form["bikou"][i], req.Form["bikou"][i])

		kotsuhi := &Kotsuhi{
			req.Form["end"][i],
			req.Form["start"][i],
			//req.Form["round_trip_flg"][i].(int64),
			0,
			req.Form["plice"][i],
		}

		ref := client.Collection("tomoki").NewDoc()
		_, err := ref.Set(context.Background(), kotsuhi)

		//		_, _, err := client.Collection("tomoki").Add(context.Background(), kotsuhi)

		if err != nil {
			return nil, err
		}

	}

	return view(req, client)
}

func isNotNil(obj interface{}) string {
	if obj != nil {
		return obj.(string)
	} else {
		return ""
	}
}

func isNotNilNum(obj interface{}) int64 {
	if obj != nil {
		return obj.(int64)
	} else {
		return 0
	}
}

// func deleteCollection(ctx context.Context, client *firestore.Client,
// 	ref *firestore.CollectionRef, batchSize int) error {

// 	for {
// 		// Get a batch of documents
// 		iter := ref.Limit(batchSize).Documents(ctx)
// 		numDeleted := 0

// 		// Iterate through the documents, adding
// 		// a delete operation for each one to a
// 		// WriteBatch.
// 		batch := client.Batch()
// 		for {
// 			doc, err := iter.Next()
// 			if err == iterator.Done {
// 				break
// 			}
// 			if err != nil {
// 				return err
// 			}

// 			batch.Delete(doc.Ref)
// 			numDeleted++
// 		}

// 		// If there are no documents to delete,
// 		// the process is over.
// 		if numDeleted == 0 {
// 			return nil
// 		}

// 		_, err := batch.Commit(ctx)
// 		if err != nil {
// 			return err
// 		}
// 	}
// }
