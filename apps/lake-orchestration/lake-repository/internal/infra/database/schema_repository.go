package database

import (
	"apps/lake-orchestration/lake-repository/internal/entity"
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SchemaRepository struct {
	log        *log.Logger
	Client     *mongo.Client
	Database   string
	Collection *mongo.Collection
}

func NewSchemaRepository(client *mongo.Client, database string) *SchemaRepository {
	return &SchemaRepository{
		log:        log.New(os.Stdout, "[SCHEMA-REPOSITORY] ", log.LstdFlags),
		Client:     client,
		Database:   database,
		Collection: client.Database(database).Collection("schemas"),
	}
}

func (sr *SchemaRepository) getOneById(id string) (*entity.Schema, error) {
	filter := bson.M{"id": id}
	existingDoc := sr.Collection.FindOne(context.Background(), filter)
	// Check if the document does not exist
	if existingDoc.Err() != nil {
		return nil, existingDoc.Err()
	}

	var result entity.Schema
	if err := existingDoc.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (sr *SchemaRepository) SaveSchema(schema *entity.Schema) error {
	// Check if the document already exists based on the ID
	_, err := sr.getOneById(string(schema.ID))
	if err != nil {
		// Insert new document
          schema.CreatedAt = time.Now().Format(time.RFC3339)
		_, err := sr.Collection.InsertOne(context.Background(), bson.M{
			"id":          schema.ID,
			"schema_type": schema.SchemaType,
			"json_schema": schema.JsonSchema,
               "service":     schema.Service,
			"schema_id":   schema.SchemaID,
			"created_at":  schema.CreatedAt,
			"updated_at":  schema.UpdatedAt,
		})
		if err != nil {
			return err
		}
	} else {
		// Update existing document
		filter := bson.M{"id": string(schema.ID)}
		update := bson.M{
			"$set": bson.M{
				"schema_type": schema.SchemaType,
				"json_schema": schema.JsonSchema,
                    "service":     schema.Service,
				"schema_id":   schema.SchemaID,
				"created_at":  schema.CreatedAt,
				"updated_at":  schema.UpdatedAt,
			},
		}
		_, err := sr.Collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return err
		}
	}
	return nil
}

func (sr *SchemaRepository) FindOneById(id string) (*entity.Schema, error) {
	result, err := sr.getOneById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (sr *SchemaRepository) FindAll() ([]*entity.Schema, error) {
     cursor, err := sr.Collection.Find(context.Background(), bson.M{})
     if err != nil {
          return nil, err
     }
     defer cursor.Close(context.Background())

     var results []*entity.Schema
     for cursor.Next(context.Background()) {
          var result entity.Schema
          if err := cursor.Decode(&result); err != nil {
               return nil, err
          }
          results = append(results, &result)
     }
     if err := cursor.Err(); err != nil {
          return nil, err
     }
     return results, nil
}

func (sr *SchemaRepository) FindAllByService(service string) ([]*entity.Schema, error) {
	filter := bson.M{"service": service}
	cursor, err := sr.Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []*entity.Schema
	for cursor.Next(context.Background()) {
		var result entity.Schema
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, &result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return results, nil
}






