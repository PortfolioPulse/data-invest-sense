package graph

import (
     "apps/lake-orchestation/lake-gateway/internal/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.


type Resolver struct{
     CreateInputUseCase usecase.CreateInputUseCase
}
