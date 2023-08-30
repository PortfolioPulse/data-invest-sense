package service

import (
	"apps/lake-orchestration/lake-repository/internal/infra/grpc/pb"
	"apps/lake-orchestration/lake-repository/internal/usecase"
	"context"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

type SchemaService struct {
     pb.UnimplementedSchemaServiceServer
     CreateSchemaUseCase usecase.CreateSchemaUseCase
}

func NewSchemaService(
     createSchemaUseCase usecase.CreateSchemaUseCase,
) *SchemaService {
     return &SchemaService{
          CreateSchemaUseCase: createSchemaUseCase,
     }
}

type PropertiesValue struct {
	Value interface{}
}

func toMapStringInterface(req map[string]*anypb.Any) map[string]interface{} {
	res := make(map[string]interface{})
	for key, value := range req {
		res[key] = value
	}
	return res
}

func (v *PropertiesValue) ProtoReflect() protoreflect.Message {
	return (*anypb.Any)(nil).ProtoReflect()
}

func toMapAny(req map[string]interface{}) map[string]*anypb.Any {
	res := make(map[string]*anypb.Any, len(req))
	for key, value := range req {
		anyValue, _ := anypb.New(&PropertiesValue{Value: value})
		res[key] = anyValue
	}
	return res
}

func (s *SchemaService) CreateSchema(ctx context.Context, in *pb.CreateSchemaRequest) (*pb.CreateSchemaResponse, error) {
     dto := usecase.SchemaInputDTO{
          JsonSchema: toMapStringInterface(in.JsonSchema),
          SchemaType: in.SchemaType,
          Service:    in.Service,
          Source:     in.Source,
     }

     output, err := s.CreateSchemaUseCase.Execute(dto)
     if err != nil {
          return nil, err
     }

     return &pb.CreateSchemaResponse{
          Id:         output.ID,
          SchemaType: output.SchemaType,
          Service:    output.Service,
          Source:     output.Source,
          JsonSchema: toMapAny(output.JsonSchema),
          SchemaId:   output.SchemaID,
          CreatedAt:  output.CreatedAt,
          UpdatedAt:  output.UpdatedAt,
     }, nil
}
