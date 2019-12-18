package apps

import (
	"cloud.google.com/go/firestore"
	_ "firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

func ViewIndex(_ *firestore.Client) (gin.H, error) {
	return gin.H{
		"title": "INDEX",
	}, nil
}
