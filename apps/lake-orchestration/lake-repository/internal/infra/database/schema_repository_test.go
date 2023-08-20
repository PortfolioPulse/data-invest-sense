package database

import (
	"apps/lake-orchestration/lake-repository/internal/entity"
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SchemaRepositorySuite struct {
	suite.Suite
	log        *log.Logger
	Client     *mongo.Client
	Database   string
	Collection string
	repo       *SchemaRepository
	tokenAuth  *jwtauth.JWTAuth
}

func TestSchemaRepositorySuite(t *testing.T) {
	suite.Run(t, new(SchemaRepositorySuite))
}

func (suite *SchemaRepositorySuite) SetupTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	mongoURI := "mongodb://localhost:27017"
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	suite.Client = client
	suite.Database = "test-database"
	suite.Collection = "test-service"
	suite.log = log.New(os.Stdout, "[SCHEMA-REPOSITORY] ", log.LstdFlags)
	suite.repo = NewSchemaRepository(suite.Client, suite.Database)
	suite.tokenAuth = jwtauth.New("HS256", []byte("your-secret-key"), nil)
}

func (suite *SchemaRepositorySuite) TearDownTest() {
	suite.Client.Database(suite.Database).Drop(context.Background())
	err := suite.Client.Disconnect(context.Background())
	suite.NoError(err)
}

func (suite *SchemaRepositorySuite) TestSaveSchemaWhenSchemaDoesNotExist() {
	jsonSchema := map[string]interface{}{
		"test": "test",
	}
	schema, err := entity.NewSchema("test", "test", "test", jsonSchema, suite.tokenAuth)
	suite.NoError(err)
	err = suite.repo.SaveSchema(schema)
	suite.NoError(err)
}

func (suite *SchemaRepositorySuite) TestFindOneByIdWhenSchemaExists() {
	jsonSchema := map[string]interface{}{
		"test": "test",
	}
	schema, err := entity.NewSchema("test", "test", "test", jsonSchema, suite.tokenAuth)
	suite.NoError(err)

	err = suite.repo.SaveSchema(schema)
	suite.NoError(err)

	foundSchema, err := suite.repo.FindOneById(string(schema.ID))
	suite.NoError(err)
	suite.NotNil(foundSchema)
	suite.Equal(schema.ID, foundSchema.ID)
}

func (suite *SchemaRepositorySuite) TestFindOneByIdWhenSchemaDoesNotExist() {
	foundSchema, err := suite.repo.FindOneById("non-existing-id")
	suite.Error(err)
	suite.Nil(foundSchema)
}

func (suite *SchemaRepositorySuite) TestFindAll() {
	jsonSchema1 := map[string]interface{}{
		"field1":  "value1",
		"field2":  42,
		"service": "test", // Make sure each schema has the "service" key
	}
	jsonSchema2 := map[string]interface{}{
		"field1":  "value2",
		"field2":  43,
		"service": "test", // Make sure each schema has the "service" key
	}
	schema1, err := entity.NewSchema("test1", "test", "test1", jsonSchema1, suite.tokenAuth)
	suite.NoError(err)
	schema2, err := entity.NewSchema("test2", "test", "test2", jsonSchema2, suite.tokenAuth)
	suite.NoError(err)

	err = suite.repo.SaveSchema(schema1)
	suite.NoError(err)
	err = suite.repo.SaveSchema(schema2)
	suite.NoError(err)

	foundSchemas, err := suite.repo.FindAll()
	suite.NoError(err)
	suite.NotNil(foundSchemas)
	suite.Equal(2, len(foundSchemas))
}

func (suite *SchemaRepositorySuite) TestSaveSchemaWhenSchemaExist() {
	jsonSchema := map[string]interface{}{
		"field1": map[string]interface{}{
			"type": "string",
		},
		"field2": map[string]interface{}{
			"type": "string",
		},
	}
	schema, err := entity.NewSchema("test", "test", "test", jsonSchema, suite.tokenAuth)
	suite.NoError(err)

	err = suite.repo.SaveSchema(schema)
	schemaID1 := schema.SchemaID
	suite.NoError(err)

	newJsonSchema := map[string]interface{}{
		"field1": map[string]interface{}{
			"type": "string",
		},
		"field2": map[string]interface{}{
			"type": "string",
		},
		"field3": map[string]interface{}{
			"type": "string",
		},
	}

	schema, err = entity.NewSchema("test", "test", "test", newJsonSchema, suite.tokenAuth)
	suite.NoError(err)

	err = suite.repo.SaveSchema(schema)
	schemaID2 := schema.SchemaID
	suite.NoError(err)

	suite.NotEqual(schemaID1, schemaID2)

	updatedSchema, err := suite.repo.FindOneById(string(schema.ID))
	suite.NoError(err)
	suite.NotNil(updatedSchema)
	suite.Equal(schemaID2, updatedSchema.SchemaID)
	suite.Equal(newJsonSchema, updatedSchema.JsonSchema)
}

func (suite *SchemaRepositorySuite) TestFindAllByServiceWhenSchemasExist() {
	jsonSchema1 := map[string]interface{}{
		"field1":  "value1",
		"field2":  42,
		"service": "test", // Make sure each schema has the "service" key
	}
	jsonSchema2 := map[string]interface{}{
		"field1":  "value2",
		"field2":  43,
		"service": "test", // Make sure each schema has the "service" key
	}
	schema1, err := entity.NewSchema("test1", "test", "test1", jsonSchema1, suite.tokenAuth)
	suite.NoError(err)
	schema2, err := entity.NewSchema("test2", "test", "test2", jsonSchema2, suite.tokenAuth)
	suite.NoError(err)

	err = suite.repo.SaveSchema(schema1)
	suite.NoError(err)
	err = suite.repo.SaveSchema(schema2)
	suite.NoError(err)

	foundSchemas, err := suite.repo.FindAllByService("test")
	suite.NoError(err)
	suite.NotNil(foundSchemas)
	suite.Equal(2, len(foundSchemas))
}

func (suite *SchemaRepositorySuite) TestFindAllByServiceWhenNoSchemasExist() {
	// Find all schemas for a non-existing service
	foundSchemas, err := suite.repo.FindAllByService("non-existing-service")
	log.Println(foundSchemas)
	suite.NoError(err)
	suite.Nil(foundSchemas)
}

func (suite *SchemaRepositorySuite) TestFindAllByServiceWhenSchemasExistButNotForService() {
	jsonSchema1 := map[string]interface{}{
		"field1":  "value1",
		"field2":  42,
		"service": "test", // Make sure each schema has the "service" key
	}
	jsonSchema2 := map[string]interface{}{
		"field1":  "value2",
		"field2":  43,
		"service": "test", // Make sure each schema has the "service" key
	}
	schema1, err := entity.NewSchema("test1", "test", "test1", jsonSchema1, suite.tokenAuth)
	suite.NoError(err)
	schema2, err := entity.NewSchema("test2", "test", "test2", jsonSchema2, suite.tokenAuth)
	suite.NoError(err)

	// Save the schemas
	err = suite.repo.SaveSchema(schema1)
	suite.NoError(err)
	err = suite.repo.SaveSchema(schema2)
	suite.NoError(err)

	// Find all schemas for a non-existing service
	foundSchemas, err := suite.repo.FindAllByService("non-existing-service")
	suite.NoError(err)
	suite.Nil(foundSchemas)
}
