package repository

import (
	"context"
	"english_bot_admin/internal/models"
	"english_bot_admin/internal/user"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type UserRepository struct {
	userCollection  *mongo.Collection
	adminCollection *mongo.Collection
}

func NewUserRepository(userCollection *mongo.Collection, adminCollection *mongo.Collection) user.Repository {
	return &UserRepository{
		userCollection:  userCollection,
		adminCollection: adminCollection,
	}
}

func (r *UserRepository) Select(ctx context.Context) ([]models.User, error) {
	filter := bson.M{}
	var users []models.User
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

func (r *UserRepository) AdminVerification(ctx context.Context, params *models.AdminSignInParams) (bool, error) {
	filter := bson.M{
		"login":    params.Login,
		"password": params.Password,
	}
	var admin models.Admin
	cursor, err := r.adminCollection.Find(ctx, filter)
	if err != nil {
		return false, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			log.Fatalf("Error closing cursor")
		}
	}(cursor, ctx)

	if cursor.Next(ctx) {
		if err = cursor.Decode(&admin); err != nil {
			return false, err
		}
		return true, nil
	}

	return false, fmt.Errorf("error: no admin")
}

func (r *UserRepository) InsertAdmin(ctx context.Context, newAdmin *models.Admin) error {
	_, err := r.adminCollection.InsertOne(ctx, newAdmin)
	if err != nil {
		return err
	}
	return nil
}
