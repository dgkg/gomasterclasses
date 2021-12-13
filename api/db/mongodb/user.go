package mongodb

import (
	"context"
	"time"

	dbi "github.com/dgkg/gomasterclasses/api/db"
	"github.com/dgkg/gomasterclasses/api/model"
	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *MongoDB) CreateUser(u *model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	u.UUID = uuid.NewString()
	_, err := db.colUser.InsertOne(ctx, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (db *MongoDB) Delete(uuid string) error {
	return nil
}
func (db *MongoDB) Update(uuid string, data map[string]interface{}) (*model.User, error) {
	return nil, nil
}
func (db *MongoDB) Get(uuid string) (*model.User, error) {
	filter := bson.D{{"uuid", uuid}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var u model.User
	err := db.colUser.FindOne(ctx, filter).Decode(&u)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		return nil, dbi.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	return &u, nil
}
func (db *MongoDB) GetAll() (map[string]*model.User, error) {
	return nil, nil
}
func (db *MongoDB) GetByEmail(email string) (*model.User, error) {
	filter := bson.D{{"email", email}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var u model.User
	err := db.colUser.FindOne(ctx, filter).Decode(&u)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		return nil, dbi.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	return &u, nil
}
