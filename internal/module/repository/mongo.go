package repository

import (
	"context"
	"english_bot_admin/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type ModuleRepository struct {
	moduleCollection *mongo.Collection
}

func NewModuleRepository(moduleCollection *mongo.Collection) *ModuleRepository {
	return &ModuleRepository{
		moduleCollection: moduleCollection,
	}
}

func (r *ModuleRepository) NewModule(ctx context.Context, module *models.Module) error {
	_, err := r.moduleCollection.InsertOne(ctx, module)
	if err != nil {
		return err
	}
	return nil
}

func (r *ModuleRepository) SelectModules(ctx context.Context) ([]models.Module, error) {
	filter := bson.M{}
	var modules []models.Module
	cursor, err := r.moduleCollection.Find(ctx, filter)
	if err != nil {
		log.Println("error while getting modules:", err)
		return []models.Module{}, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			log.Println("error closing cursor")
		}
	}(cursor, ctx)
	if err = cursor.All(ctx, &modules); err != nil {
		log.Println("error while decoding modules:", err)
		return []models.Module{}, err
	}
	return modules, nil
}

func (r *ModuleRepository) InsertTask(ctx context.Context, params *models.TaskToModule) error {
	filter := bson.M{"module_id": params.ModuleId}
	update := bson.M{"$addToSet": bson.M{"task": params.Task}}

	_, err := r.moduleCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
