package database

import (
	"apps/lake-orchestation/lake-gateway/internal/entity"
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type InputRepository struct {
	log      *log.Logger
	Client   *mongo.Client
	Database string
}

func NewInputRepository(client *mongo.Client, database string) *InputRepository {
	return &InputRepository{
		log:      log.New(os.Stdout, "[INPUT-REPOSITORY] ", log.LstdFlags),
		Client:   client,
		Database: database,
	}
}

func (ir *InputRepository) getOneById(id string, service string) (*entity.Input, error) {
     collection := ir.Client.Database(ir.Database).Collection(service)

     filter := bson.M{"id": id}
     existingDoc := collection.FindOne(context.Background(), filter)
     // Check if the document does not exist
     if existingDoc.Err() != nil {
          return nil, existingDoc.Err()
     }

     var result entity.Input
     if err := existingDoc.Decode(&result); err != nil {
          return nil, err
     }

     return &result, nil
}

func (ir *InputRepository) FindOneByIdAndService(id string, service string) (*entity.Input, error) {
     result, err := ir.getOneById(id, service)
     if err != nil {
          return nil, err
     }
     return result, nil
}

func (ir *InputRepository) SaveInput(input *entity.Input, service string) error {
	collection := ir.Client.Database(ir.Database).Collection(service)
	// Check if the document already exists based on the ID
	_, err := ir.getOneById(string(input.ID), service)
	if err != nil {
		// Insert new document
		_, err := collection.InsertOne(context.Background(), bson.M{
			"id":       input.ID,
			"data":     input.Data,
			"metadata": input.Metadata,
			"status":   input.Status,
		})
		if err != nil {
			return err
		}
		log.Printf("[INSERT] - Service: %s | Source: %s | ID: %s", service, input.Metadata.Source, input.ID)
		return nil
	} else {
		// Update existing document
		_, err := collection.UpdateOne(
			context.Background(),
			bson.M{"id": input.ID},
			bson.M{
				"$set": bson.M{
					"data":     input.Data,
					"metadata": input.Metadata,
					"status":   input.Status,
				},
			},
		)
		if err != nil {
			return err
		}
		log.Printf("[UPDATE] - Service: %s | Source: %s | ID: %s", service, input.Metadata.Source, input.ID)
		return nil
	}
}

func (ir *InputRepository) FindAllByService(service string) ([]*entity.Input, error) {
	collection := ir.Client.Database(ir.Database).Collection(service)

	filter := bson.M{"metadata.service": service}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []*entity.Input
	for cursor.Next(context.Background()) {
		var result entity.Input
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

func (ir *InputRepository) FindAllByServiceAndSource(service string, source string) ([]*entity.Input, error) {
	collection := ir.Client.Database(ir.Database).Collection(service)

	filter := bson.M{"metadata.source": source}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []*entity.Input
	for cursor.Next(context.Background()) {
		var result entity.Input
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

func (ir *InputRepository) FindAllByServiceAndSourceAndStatus(service string, source string, status int) ([]*entity.Input, error) {
	collection := ir.Client.Database(ir.Database).Collection(service)

	filter := bson.M{"metadata.source": source, "status.code": status}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []*entity.Input
	for cursor.Next(context.Background()) {
		var result entity.Input
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
