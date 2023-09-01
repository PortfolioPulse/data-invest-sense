package database

import (
	"apps/lake-orchestration/lake-gateway/internal/entity"
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StagingJobRepository struct {
	log        *log.Logger
	Client     *mongo.Client
	Database   string
	Collection string
}

func NewStagingJobRepository(client *mongo.Client, database string) *StagingJobRepository {
	return &StagingJobRepository{
		log:        log.New(os.Stdout, "[STAGING-JOB-REPOSITORY] ", log.LstdFlags),
		Client:     client,
		Database:   database,
		Collection: "staging-jobs",
	}
}

func (sjr *StagingJobRepository) getOneById(id string) (*entity.StagingJob, error) {
	collection := sjr.Client.Database(sjr.Database).Collection(sjr.Collection)

	filter := bson.M{"id": id}
	existingDoc := collection.FindOne(context.Background(), filter)
	// Check if the document does not exist
	if existingDoc.Err() != nil {
		return nil, existingDoc.Err()
	}

	var result entity.StagingJob
	if err := existingDoc.Decode(&result); err != nil {
		return nil, err
	}
     log.Printf("result: %+v\n", result)

	return &result, nil
}

func (sjr *StagingJobRepository) FindOneById(id string) (*entity.StagingJob, error) {
	result, err := sjr.getOneById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (sjr *StagingJobRepository) SaveStagingJob(stagingJob *entity.StagingJob) error {
	collection := sjr.Client.Database(sjr.Database).Collection(sjr.Collection)
	// Check if the document already exists based on the ID
	_, err := sjr.getOneById(string(stagingJob.ID))
	if err != nil {
		// Insert new document
		_, err := collection.InsertOne(context.Background(), bson.M{
			"id":            stagingJob.ID,
			"input_id":      stagingJob.InputId,
			"input":         stagingJob.Input,
			"source":        stagingJob.Source,
			"service":       stagingJob.Service,
			"processing_id": stagingJob.ProcessingId,
		})
		if err != nil {
			return err
		}
		return nil
	}
	// Update existing document
	_, err = collection.UpdateOne(context.Background(), bson.M{"id": stagingJob.ID}, bson.M{"$set": bson.M{
		"input_id":      stagingJob.InputId,
		"input":         stagingJob.Input,
		"source":        stagingJob.Source,
		"service":       stagingJob.Service,
		"processing_id": stagingJob.ProcessingId,
	}})
	if err != nil {
		return err
	}
	return nil
}

func (sjr *StagingJobRepository) DeleteById(id string) error {
	collection := sjr.Client.Database(sjr.Database).Collection(sjr.Collection)
	_, err := collection.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}

func (sjr *StagingJobRepository) FindAll() ([]*entity.StagingJob, error) {
	collection := sjr.Client.Database(sjr.Database).Collection(sjr.Collection)
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []*entity.StagingJob
	for cursor.Next(context.Background()) {
		var result entity.StagingJob
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, &result)
	}
	return results, nil
}

func (sjr *StagingJobRepository) FindAllByService(service string) ([]*entity.StagingJob, error) {
	collection := sjr.Client.Database(sjr.Database).Collection(sjr.Collection)

	filter := bson.M{"service": service}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []*entity.StagingJob
	for cursor.Next(context.Background()) {
		var result entity.StagingJob
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, &result)
	}
	return results, nil
}

func (sjr *StagingJobRepository) FindAllByServiceAndSource(service string, source string) ([]*entity.StagingJob, error) {
	collection := sjr.Client.Database(sjr.Database).Collection(sjr.Collection)

	filter := bson.M{"service": service, "source": source}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []*entity.StagingJob
	for cursor.Next(context.Background()) {
		var result entity.StagingJob
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, &result)
	}
	return results, nil
}

func (sjr *StagingJobRepository) FindAllByInputId(inputId string) ([]*entity.StagingJob, error) {
	collection := sjr.Client.Database(sjr.Database).Collection(sjr.Collection)

	filter := bson.M{"input_id": inputId}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []*entity.StagingJob
	for cursor.Next(context.Background()) {
		var result entity.StagingJob
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, &result)
	}
	return results, nil
}

func (sjr *StagingJobRepository) FindAllByProcessingId(processingId string) ([]*entity.StagingJob, error) {
	collection := sjr.Client.Database(sjr.Database).Collection(sjr.Collection)

	filter := bson.M{"processing_id": processingId}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []*entity.StagingJob
	for cursor.Next(context.Background()) {
		var result entity.StagingJob
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, &result)
	}
	return results, nil
}


