package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"apps/lake-orchestration/lake-controller/internal/entity"
)

type ProcessingJobDependenciesRepository struct {
	log        *log.Logger
	Client     *mongo.Client
	Database   string
	Collection *mongo.Collection
}

func NewProcessingJobDependenciesRepository(client *mongo.Client, database string) *ProcessingJobDependenciesRepository {
	return &ProcessingJobDependenciesRepository{
		log:        log.New(os.Stdout, "[PROCESSING-JOB-DEPENDENCIES-REPOSITORY] ", log.LstdFlags),
		Client:     client,
		Database:   database,
		Collection: client.Database(database).Collection("job-staging-dependencies"),
	}
}

func (pjd *ProcessingJobDependenciesRepository) getOneById(id string) (*entity.ProcessingJobDependencies, error) {
	filter := bson.M{"id": id}
	existingDoc := pjd.Collection.FindOne(context.Background(), filter)
	// Check if the document does not exist
	if existingDoc.Err() != nil {
		return nil, existingDoc.Err()
	}

	var result entity.ProcessingJobDependencies
	if err := existingDoc.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (pjd *ProcessingJobDependenciesRepository) SaveProcessingJobDependencies(processingJobDependencies *entity.ProcessingJobDependencies) error {
	_, err := pjd.getOneById(string(processingJobDependencies.ID))
	if err != nil {
		// Insert new document
		_, err := pjd.Collection.InsertOne(context.Background(), bson.M{
			"id":               processingJobDependencies.ID,
			"service":          processingJobDependencies.Service,
			"source":           processingJobDependencies.Source,
			"job_dependencies": processingJobDependencies.JobDependencies,
		})
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (pjd *ProcessingJobDependenciesRepository) UpdateProcessingJobDependencies(jobDep *entity.JobDependenciesWithProcessingData, id string) error {
	existingProcessingJobDependencies, err := pjd.getOneById(id)
     log.Printf("UpdateProcessingJobDependencies: existingProcessingJobDependencies=%v", existingProcessingJobDependencies)
	if err != nil {
		return err
	}
	for i, job := range existingProcessingJobDependencies.JobDependencies {
		if job.Service == jobDep.Service && job.Source == jobDep.Source {
			existingProcessingJobDependencies.JobDependencies[i] = *jobDep
			break
		}
	}

     _, err = pjd.Collection.UpdateOne(context.Background(), bson.M{"id": id}, bson.M{"$set": bson.M{
          "job_dependencies": existingProcessingJobDependencies.JobDependencies,
     }})
     if err != nil {
          return err
     }
	return nil
}

func (pjd *ProcessingJobDependenciesRepository) DeleteProcessingJobDependencies(id string) error {
	_, err := pjd.Collection.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}

func (pjd *ProcessingJobDependenciesRepository) FindOneById(id string) (*entity.ProcessingJobDependencies, error) {
	return pjd.getOneById(id)
}
