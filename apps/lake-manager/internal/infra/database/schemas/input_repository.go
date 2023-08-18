package schemas

import (
	"apps/lake-manager/internal/entity/schemas"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SchemaInputRepository struct {
	client         *mongo.Client
	database       string
	collectionName string
}

func NewSchemaInputRepository(client *mongo.Client) *SchemaInputRepository {
	return &SchemaInputRepository{
		client:         client,
		database:       "schemas",
		collectionName: "input",
	}
}

func (si *SchemaInputRepository) Save(input *schemas.SchemaInput) error {
	collection := si.client.Database(si.database).Collection(si.collectionName)
	// Check if the document already exists based on the ID
	filter := bson.M{"id": input.ID}
	existingDoc := collection.FindOne(context.Background(), filter)

	if existingDoc.Err() == nil {
		// Document already exists, you can update it here if needed
		fmt.Println("Document already exists:", existingDoc)
		return nil
	} else if existingDoc.Err() != mongo.ErrNoDocuments {
		// Some error occurred while querying
		return existingDoc.Err()
	}

	fmt.Println("Input Document: ", input)

	res, err := collection.InsertOne(context.Background(), bson.M{
		"id":         input.ID,
		"required":   input.Required,
		"properties": input.Properties,
		"schema_id":  input.SchemaId,
	})

	if err != nil {
		return err
	}
	fmt.Println("Result: ", res)

	return nil
}

func (si *SchemaInputRepository) GetTotal() (int64, error) {
	collection := si.client.Database(si.database).Collection(si.collectionName)

	count, err := collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return 0, err
	}

	return count, nil
}
