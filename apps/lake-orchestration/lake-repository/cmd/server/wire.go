// go:build wireinject
// +build wireinject

package main

import (
     "apps/lake-orchestration/lake-repository/internal/entity"
     "apps/lake-orchestration/lake-repository/internal/infra/database"
     "apps/lake-orchestration/lake-repository/internal/event"
     webHandler "apps/lake-orchestration/lake-repository/internal/infra/web/handlers"
     "apps/lake-orchestration/lake-repository/internal/usecase"

     "libs/golang/events"

     "github.com/google/wire"
     "go.mongodb.org/mongo-driver/mongo"
     "github.com/go-chi/jwtauth"
)


var setSchemaRepositoryDependency = wire.NewSet(
     database.NewSchemaRepository,
     wire.Bind(
          new(entity.SchemaInterface),
          new(*database.SchemaRepository),
     ),
)

var setEventDispatcherDependency = wire.NewSet(
     events.NewEventDispatcher,
     event.NewSchemaCreated,
     wire.Bind(new(events.EventInterface), new(*event.SchemaCreated)),
     wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setSchemaCreatedEvent = wire.NewSet(
     event.NewSchemaCreated,
     wire.Bind(new(events.EventInterface), new(*event.SchemaCreated)),
)

// [Use Case]
func NewCreateSchemaUseCase(client *mongo.Client, eventDispatcher events.EventDispatcherInterface, database string, tokenAuth *jwtauth.JWTAuth) *usecase.CreateSchemaUseCase {
     wire.Build(
          setSchemaRepositoryDependency,
          setSchemaCreatedEvent,
          usecase.NewCreateSchemaUseCase,
     )
     return &usecase.CreateSchemaUseCase{}
}

// [Web Handler]
func NewWebSchemaHandler(client *mongo.Client, eventDispatcher events.EventDispatcherInterface, database string, tokenAuth *jwtauth.JWTAuth) *webHandler.WebSchemaHandler {
     wire.Build(
          setSchemaRepositoryDependency,
          setSchemaCreatedEvent,
          webHandler.NewWebSchemaHandler,
     )
     return &webHandler.WebSchemaHandler{}
}
