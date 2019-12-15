package apps

import "github.com/gin-gonic/gin"

func viewIndex() gin.H {
	return gin.H{
		"title": "INDEX",
	}
}
