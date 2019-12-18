package main

import (
	"CampToolDevelop/pkg/db"
	"github.com/gin-gonic/gin"
	"testing"

	"CampToolDevelop/internal/apps"
)

func TestViewKotsuhi(t *testing.T) {

	client, err := db.OpenFirebase()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	_ = gin.H{}

	form, err := apps.ViewKotsuhi(client)

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	list := form["list"].([]*apps.Kotsuhi)

	for _, kotsuhi := range list {
		t.Log(kotsuhi)
	}

}
