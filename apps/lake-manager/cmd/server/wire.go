//go:build wireinject
// +build wireinject

package main

import (
	entitySchemas "apps/lake-manager/internal/entity/schemas"
	eventSchemas "apps/lake-manager/internal/event/schemas"
	databaseSchemas "apps/lake-manager/internal/infra/database/schemas"
	webHandlerSchemas "apps/lake-manager/internal/infra/web/handlers/schemas"
	usecaseSchemas "apps/lake-manager/internal/usecase/schemas"
	"apps/lake-manager/pkg/events"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var setSchemaInputRepositoryDependency = wire.NewSet(
	databaseSchemas.NewSchemaInputRepository,
	wire.Bind(
		new(entitySchemas.SchemaInputInterface),
		new(*databaseSchemas.SchemaInputRepository),
	),
)

var setSchemasEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	eventSchemas.NewSchemaInputCreated,
	wire.Bind(new(events.EventInterface), new(*eventSchemas.SchemaInputCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setSchemaInputCreatedEvent = wire.NewSet(
	eventSchemas.NewSchemaInputCreated,
	wire.Bind(new(events.EventInterface), new(*eventSchemas.SchemaInputCreated)),
)

func NewCreateSchemaInputUseCase(client *mongo.Client, eventDispatcher events.EventDispatcherInterface) *usecaseSchemas.CreateSchemaInputUseCase {
	wire.Build(
		setSchemaInputRepositoryDependency,
		setSchemaInputCreatedEvent,
		usecaseSchemas.NewCreateSchemaInputUseCase,
	)
	return &usecaseSchemas.CreateSchemaInputUseCase{}
}

func NewWebSchemaInputHandler(client *mongo.Client, eventDispatcher events.EventDispatcherInterface) *webHandlerSchemas.WebSchemaInputHandler {
	wire.Build(
		setSchemaInputRepositoryDependency,
		setSchemaInputCreatedEvent,
		webHandlerSchemas.NewWebSchemaInputHandler,
	)
	return &webHandlerSchemas.WebSchemaInputHandler{}
}
