package apps

import "github.com/gin-gonic/gin"

func ViewIndex() gin.H {
	return gin.H{
		"title": "INDEX",
	}
}
