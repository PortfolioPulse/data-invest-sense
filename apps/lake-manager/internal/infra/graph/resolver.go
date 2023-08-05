package graph

import (
     usecaseSchemas "apps/lake-manager/internal/usecase/schemas"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.


type Resolver struct{
     CreateSchemaInputUseCase usecaseSchemas.CreateSchemaInputUseCase
}
