package graph

import (
     "apps/lake-orchestration/lake-controller/internal/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
     CreateConfigUseCase usecase.CreateConfigUseCase
}
