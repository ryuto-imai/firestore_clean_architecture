package gateways

import (
	"context"
	"encoding/json"
	"errors"
	"firestore_clean/database"
	"firestore_clean/entities"
	"firestore_clean/usecases/ports"
	"fmt"
)

type UserGateway struct {
	clientFactory database.FirestoreClientFactory
}

func NewUserRepository(clientFactory database.FirestoreClientFactory) ports.UserRepository {
	return &UserGateway{
		clientFactory: clientFactory,
	}
}

func (gateway *UserGateway) AddUser(ctx context.Context, user *entities.User) ([]*entities.User, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}

	client, err := gateway.clientFactory.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed AddUser NewClient: %v", err)
	}
	defer client.Close()

	_, err = client.Collection("users").Doc(user.Name).Set(ctx, map[string]interface{}{
		"age":     user.Age,
		"address": user.Address,
	})

	if err != nil {
		return nil, fmt.Errorf("failed AddUser Set: %v", err)
	}

	return gateway.GetUsers(ctx)
}

func (gateway *UserGateway) GetUsers(ctx context.Context) ([]*entities.User, error) {
	client, err := gateway.clientFactory.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed GetUsers NewClient: %v", err)
	}
	defer client.Close()

	allData := client.Collection("users").Documents(ctx)

	docs, err := allData.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed GetUsers GetAll: %v", err)
	}

	users := make([]*entities.User, 0)
	for _, doc := range docs {
		u := new(entities.User)
		err = mapToStruct(doc.Data(), &u)
		if err != nil {
			return nil, fmt.Errorf("failed GetUsers mapToStruct: %v", err)
		}
		u.Name = doc.Ref.ID
		users = append(users, u)
	}

	return users, nil
}

// map -> 構造体の変換
func mapToStruct(m map[string]interface{}, val interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}
