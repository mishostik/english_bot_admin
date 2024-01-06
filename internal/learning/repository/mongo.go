package repository

import (
	"context"
	"english_bot_admin/internal/learning"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type LearnRepository struct {
	learnCollection *mongo.Collection
}

func NewLearnRepository(learnCollection *mongo.Collection) learning.Repository {
	return &LearnRepository{
		learnCollection: learnCollection,
	}
}

func (r *LearnRepository) InsertRule(ctx context.Context, rule *learning.Rule) error {
	_, err := r.learnCollection.InsertOne(ctx, rule)
	if err != nil {
		return err
	}
	return nil
}

func (r *LearnRepository) SelectRules(ctx context.Context) ([]learning.Rule, error) {
	filter := bson.M{}
	var rules []learning.Rule
	cursor, err := r.learnCollection.Find(ctx, filter)
	if err != nil {
		log.Println("error while getting rules:", err)
		return []learning.Rule{}, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			log.Println("error closing cursor")
		}
	}(cursor, ctx)
	if err = cursor.All(ctx, &rules); err != nil {
		log.Println("error while decoding rules:", err)
		return []learning.Rule{}, err
	}
	return rules, nil
}
