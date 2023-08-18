//go:build wireinject
// +build wireinject

package main

import (
	"apps/lake-orchestation/lake-gateway/internal/entity"
	"apps/lake-orchestation/lake-gateway/internal/event"
	"apps/lake-orchestation/lake-gateway/internal/infra/database"
	webHandler "apps/lake-orchestation/lake-gateway/internal/infra/web/handlers"
	"apps/lake-orchestation/lake-gateway/internal/usecase"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"libs/golang/events"
)

var setInputRepositoryDependency = wire.NewSet(
	database.NewInputRepository,
	wire.Bind(
		new(entity.InputInterface),
		new(*database.InputRepository),
	),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewInputCreated,
	event.NewInputUpdated,
	wire.Bind(new(events.EventInterface), new(*event.InputCreated)),
	wire.Bind(new(events.EventInterface), new(*event.InputUpdated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setInputCreatedEvent = wire.NewSet(
	event.NewInputCreated,
	wire.Bind(new(events.EventInterface), new(*event.InputCreated)),
)

var setInputUpdatedEvent = wire.NewSet(
	event.NewInputUpdated,
	wire.Bind(new(events.EventInterface), new(*event.InputUpdated)),
)

// [Use Case]
func NewCreateInputUseCase(client *mongo.Client, eventDispatcher events.EventDispatcherInterface, database string) *usecase.CreateInputUseCase {
	wire.Build(
		setInputRepositoryDependency,
		setInputCreatedEvent,
		usecase.NewCreateInputUseCase,
	)
	return &usecase.CreateInputUseCase{}
}

func NewUpdateStatusInputUseCase(client *mongo.Client, eventDispatcher events.EventDispatcherInterface, database string) *usecase.UpdateStatusInputUseCase {
	wire.Build(
		setInputRepositoryDependency,
		setInputUpdatedEvent,
		usecase.NewUpdateStatusInputUseCase,
	)
	return &usecase.UpdateStatusInputUseCase{}
}

func NewListAllByServiceAndSourceUseCase(client *mongo.Client, database string) *usecase.ListAllByServiceAndSourceUseCase {
	wire.Build(
		setInputRepositoryDependency,
		usecase.NewListAllByServiceAndSourceUseCase,
	)
	return &usecase.ListAllByServiceAndSourceUseCase{}
}

func NewListAllByServiceUseCase(client *mongo.Client, database string) *usecase.ListAllByServiceUseCase {
	wire.Build(
		setInputRepositoryDependency,
		usecase.NewListAllByServiceUseCase,
	)
	return &usecase.ListAllByServiceUseCase{}
}

func NewListOneByIdAndServiceUseCase(client *mongo.Client, database string) *usecase.ListOneByIdAndServiceUseCase {
	wire.Build(
		setInputRepositoryDependency,
		usecase.NewListOneByIdAndServiceUseCase,
	)
	return &usecase.ListOneByIdAndServiceUseCase{}
}

// [Web Handler]
func NewWebInputStatusHandler(client *mongo.Client, eventDispatcher events.EventDispatcherInterface, database string) *webHandler.WebInputStatusHandler {
	wire.Build(
		setInputRepositoryDependency,
		setInputUpdatedEvent,
		webHandler.NewWebInputStatusHandler,
	)
	return &webHandler.WebInputStatusHandler{}
}

func NewWebInputHandler(client *mongo.Client, eventDispatcher events.EventDispatcherInterface, database string) *webHandler.WebInputHandler {
	wire.Build(
		setInputRepositoryDependency,
		setInputCreatedEvent,
		webHandler.NewWebInputHandler,
	)
	return &webHandler.WebInputHandler{}
}
