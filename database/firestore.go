package database

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type MyFirestoreClientFactory struct{}

func (f *MyFirestoreClientFactory) NewClient(ctx context.Context) (*firestore.Client, error) {
	sa := option.WithCredentialsFile("atomic-key-339415-1d169f4a8048.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}
