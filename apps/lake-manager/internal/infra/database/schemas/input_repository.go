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
		client:   client,
		database: "schemas",
          collectionName: "input",
	}
}

func (si *SchemaInputRepository) Save(input *schemas.SchemaInput) error {
	collection := si.client.Database(si.database).Collection(si.collectionName)

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
	// defer collection.Drop(ctx)

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
