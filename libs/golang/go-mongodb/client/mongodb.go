package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	DBDriver string          `json:"db_driver"`
	User     string          `json:"user"`
	Password string          `json:"password"`
	Host     string          `json:"host"`
	Port     string          `json:"port"`
	DBName   string          `json:"db_name"`
	Dns      string          `json:"dns"`
	Ctx      context.Context `json:"ctx"`
	Client   *mongo.Client   `json:"client"`
}

func NewMongoDB(driver, user, password, host, port, dbName string, ctx context.Context) *MongoDB {
	mongoDB := MongoDB{
		DBDriver: driver,
		User:     user,
		Password: password,
		Host:     host,
		Port:     port,
		DBName:   dbName,
		Ctx:      ctx,
	}

	return &mongoDB
}

func (m *MongoDB) getMongoDBURI() string {
	return fmt.Sprintf(
		"%s://%s:%s", m.DBDriver, m.Host, m.Port)
}

func (m *MongoDB) Connect() (*mongo.Client, error) {
	m.Dns = m.getMongoDBURI()
	for attempt := 1; attempt <= 5; attempt++ {
		client, err := m.connect()
		if err == nil {
			return client, nil
		}
		time.Sleep(5 * time.Second)
	}
	return nil, fmt.Errorf("failed to connect to MongoDB after multiple attempts")
}

func (m *MongoDB) connect() (*mongo.Client, error) {
	var err error
	m.Client, err = mongo.Connect(m.Ctx, options.Client().ApplyURI(m.Dns))
	failOnError(err, "Failed to connect to MongoDB")
	if err != nil {
		return nil, err
	}
	return m.Client, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func (m *MongoDB) Disconnect(client *mongo.Client) error {
	err := client.Disconnect(m.Ctx)
	if err != nil {
		return err
	}
	return nil
}
