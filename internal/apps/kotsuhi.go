package apps

import "github.com/gin-gonic/gin"

func ViewKotsuhi() gin.H {
	return gin.H{
		"title": "KOTSUHI",
	}
}
