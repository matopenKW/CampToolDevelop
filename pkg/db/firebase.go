package db

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

const keyjson = "pkg/conf/key.json"

func OpenFirebase() (*firestore.Client, error) {
	opt := option.WithCredentialsFile(keyjson)
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
