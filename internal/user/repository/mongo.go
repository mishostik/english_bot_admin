package repository

import (
	"context"
	"english_bot_admin/internal/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type UserRepository struct {
	userCollection *mongo.Collection
}

func NewUserRepository(userCollection *mongo.Collection) user.Repository {
	return &UserRepository{
		userCollection: userCollection,
	}
}

func (r *UserRepository) Select(ctx context.Context) ([]user.User, error) {
	filter := bson.M{}
	var users []user.User
	cursor, err := r.userCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			log.Fatalf("Error closing cursor")
		}
	}(cursor, ctx)
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}
