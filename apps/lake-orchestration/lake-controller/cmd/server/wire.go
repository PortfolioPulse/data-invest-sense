// go:build wireinject
// +build wireinject

package main

import (
	"apps/lake-orchestration/lake-controller/internal/entity"
	"apps/lake-orchestration/lake-controller/internal/infra/database"
	"apps/lake-orchestration/lake-controller/internal/event"
	webHandler "apps/lake-orchestration/lake-controller/internal/infra/web/handlers"
	"apps/lake-orchestration/lake-controller/internal/usecase"
	"libs/golang/events"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var setConfigRepositoryDependency = wire.NewSet(
	database.NewConfigRepository,
	wire.Bind(
		new(entity.ConfigInterface),
		new(*database.ConfigRepository),
	),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewConfigCreated,
	wire.Bind(new(events.EventInterface), new(*event.ConfigCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setConfigCreatedEvent = wire.NewSet(
	event.NewConfigCreated,
	wire.Bind(new(events.EventInterface), new(*event.ConfigCreated)),
)

// [Use Case]
func NewCreateConfigUseCase(client *mongo.Client, eventDispatcher events.EventDispatcherInterface, database string) *usecase.CreateConfigUseCase {
	wire.Build(
		setConfigRepositoryDependency,
		setConfigCreatedEvent,
		usecase.NewCreateConfigUseCase,
	)
	return &usecase.CreateConfigUseCase{}
}

// [Web Handler]
func NewWebConfigHandler(client *mongo.Client, eventDispatcher events.EventDispatcherInterface, database string) *webHandler.WebConfigHandler {
	wire.Build(
		setConfigRepositoryDependency,
		setConfigCreatedEvent,
		webHandler.NewWebConfigHandler,
	)
	return &webHandler.WebConfigHandler{}
}

func NewHealthzHandler() *webHandler.WebHealthzHandler {
     wire.Build(
          webHandler.NewWebHealthzHandler,
     )
     return &webHandler.WebHealthzHandler{}
}
