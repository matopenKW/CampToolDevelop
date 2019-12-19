package main

import (
	"CampToolDevelop/pkg/db"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestHtmlForward(t *testing.T) {

	router := gin.Default()
	client, err := db.OpenFirebase()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	htmlForward(router, client)

}
