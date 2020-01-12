package db

import (
	"context"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"

	"errors"
	"google.golang.org/api/iterator"
)

const keyjson = "../pkg/conf/key.json"

func OpenAuth() (*auth.Client, error) {
	ctx := context.Background()

	opt := option.WithCredentialsFile(keyjson)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func OpenFirestore() (*firestore.Client, error) {
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

func GetUserRecord(client *auth.Client, uid string) (*auth.UserRecord, error) {
	user, err := client.GetUser(context.Background(), uid)
	if err != nil {
		return nil, err
	}
	return user, err
}

func UpdateUserInfo(client *auth.Client, uid string, userToUodate *auth.UserToUpdate) (*auth.UserRecord, error) {
	user, err := client.UpdateUser(context.Background(), uid, userToUodate)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func SelectDocuments(client *firestore.Client, userID string, orderBy func() (string, firestore.Direction)) ([]map[string]interface{}, error) {

	list := make([]map[string]interface{}, 0, 10)

	colle := client.Collection(userID)
	if colle == nil {
		return nil, errors.New("failed to connect")
	}

	iter := colle.OrderBy(orderBy()).Documents(context.Background())

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		data := doc.Data()
		data["DocumentId"] = doc.Ref.ID
		list = append(list, data)
	}

	return list, nil
}

func DeleteCollection(ctx context.Context, client *firestore.Client,
	ref *firestore.CollectionRef, batchSize int) error {

	for {
		// Get a batch of documents
		iter := ref.Limit(batchSize).Documents(ctx)
		numDeleted := 0

		// Iterate through the documents, adding
		// a delete operation for each one to a
		// WriteBatch.
		batch := client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			batch.Delete(doc.Ref)
			numDeleted++
		}

		// If there are no documents to delete,
		// the process is over.
		if numDeleted == 0 {
			return nil
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return err
		}
	}
}

func DeleteDocument(ctx context.Context, client *firestore.Client, userId string, documentId string) error {

	doc := client.Collection(userId).Doc(documentId)

	batch := client.Batch()

	batch.Delete(doc)

	_, err := batch.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func UpdateDocument(ctx context.Context, client *firestore.Client, userId string, documentId string, data map[string]interface{}) error {

	_, err := client.Collection(userId).Doc(documentId).Set(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func InsertDocument(ctx context.Context, client *firestore.Client, userId string, data map[string]interface{}) (string, error) {

	doc, _, err := client.Collection(userId).Add(ctx, data)
	if err != nil {
		return "", err
	}

	return doc.ID, nil
}
