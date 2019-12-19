package main

import (
	"CampToolDevelop/pkg/db"
	"cloud.google.com/go/firestore"
	"net/http"
	"testing"

	"CampToolDevelop/internal/apps"
)

func TestViewKotsuhi(t *testing.T) {

	form, err := apps.ExeKotsuhi(createSession(t))

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	list := form["list"].([]*apps.Kotsuhi)

	for _, kotsuhi := range list {
		t.Log(kotsuhi)
	}

}

func createSession(t *testing.T) (*http.Request, *firestore.Client) {

	client, err := db.OpenFirebase()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	var req *http.Request

	// req.URL := http.Request.URL{
	// 	Path:"/kotsuhi",
	// }

	return req, client

}
