package apps

import (
	"cloud.google.com/go/firestore"

	"github.com/gin-gonic/gin"
)

func ViewKotsuhi(client *firestore.Client) (gin.H, error) {

	return gin.H{
		"title": "KOTSUHI",
	}, nil
}
