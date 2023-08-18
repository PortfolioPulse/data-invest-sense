package database

import (
	"apps/lake-orchestration/lake-controller/internal/entity"
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ConfigRepository struct {
	log        *log.Logger
	Client     *mongo.Client
	Database   string
	Collection *mongo.Collection
}

func NewConfigRepository(client *mongo.Client, database string) *ConfigRepository {
	return &ConfigRepository{
		log:        log.New(os.Stdout, "[CONFIG-REPOSITORY] ", log.LstdFlags),
		Client:     client,
		Database:   database,
		Collection: client.Database(database).Collection("config"),
	}
}

func (cr *ConfigRepository) getOneById(id string) (*entity.Config, error) {
	filter := bson.M{"id": id}
	existingDoc := cr.Collection.FindOne(context.Background(), filter)
	// Check if the document does not exist
	if existingDoc.Err() != nil {
		return nil, existingDoc.Err()
	}

	var result entity.Config
	if err := existingDoc.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (cr *ConfigRepository) SaveConfig(config *entity.Config) error {
	// Check if the document already exists based on the ID
	_, err := cr.getOneById(string(config.ID))
	if err != nil {
		// Insert new document
		_, err := cr.Collection.InsertOne(context.Background(), bson.M{
			"id":            config.ID,
			"name":          config.Name,
			"active":        config.Active,
			"service":       config.Service,
			"source":        config.Source,
			"context":       config.Context,
			"dependsOn":     config.DependsOn,
			"jobParameters": config.JobParameters,
		})
		if err != nil {
			return err
		}
		return nil
	}
	// Update existing document
	_, err = cr.Collection.UpdateOne(context.Background(), bson.M{"id": config.ID}, bson.M{"$set": bson.M{
		"name":          config.Name,
		"active":        config.Active,
		"service":       config.Service,
		"source":        config.Source,
		"context":       config.Context,
		"dependsOn":     config.DependsOn,
		"jobParameters": config.JobParameters,
	}})
	if err != nil {
		return err
	}
	return nil
}


func (cr *ConfigRepository) FindAllByService(service string) ([]*entity.Config, error) {
     filter := bson.M{"service": service}
     cursor, err := cr.Collection.Find(context.Background(), filter)
     if err != nil {
          return nil, err
     }
     var results []*entity.Config
     if err := cursor.All(context.Background(), &results); err != nil {
          return nil, err
     }
     return results, nil
}

func (cr *ConfigRepository) FindOneById(id string) (*entity.Config, error) {
     result, err := cr.getOneById(id)
     if err != nil {
          return nil, err
     }
     return result, nil
}
