package local

import (
	"cloud.google.com/go/firestore"
)

func SelectDocuments(client *firestore.Client, userID string) ([]map[string]interface{}, error) {

	retList := make([]map[string]interface{}, 0, 10)

	doc := map[string]interface{}{
		"Date":         "20200101",
		"End":          "新宿",
		"Start":        "町田",
		"RoundTripFlg": "1",
		"Price":        "300",
	}
	retList = append(retList, doc)
	retList = append(retList, doc)
	retList = append(retList, doc)

	return retList, nil
}
