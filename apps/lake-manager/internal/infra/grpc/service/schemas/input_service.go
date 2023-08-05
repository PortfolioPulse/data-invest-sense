package schemas

import (
	"apps/lake-manager/internal/infra/grpc/pb"
	usecaseSchemas "apps/lake-manager/internal/usecase/schemas"
	"context"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

type SchemaInputService struct {
	pb.UnimplementedSchemaInputServiceServer
	CreateSchemaInputUseCase usecaseSchemas.CreateSchemaInputUseCase
}

func NewSchemaInputService(
	createSchemaInputUseCase usecaseSchemas.CreateSchemaInputUseCase,
) *SchemaInputService {
	return &SchemaInputService{
		CreateSchemaInputUseCase: createSchemaInputUseCase,
	}
}

func toMapStringInterface(req map[string]*anypb.Any) map[string]interface{} {
	res := make(map[string]interface{})
	for key, value := range req {
		res[key] = value
	}
	return res
}

type PropertiesValue struct {
	Value interface{}
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

func (s *SchemaInputService) CreateSchemaInput(ctx context.Context, in *pb.CreateSchemaInputRequest) (*pb.CreateSchemaInputResponse, error) {
	dto := usecaseSchemas.SchemaInputInputDTO{
		ID:         in.Id,
		Required:   in.Required,
		Properties: toMapStringInterface(in.Properties),
	}
	output, err := s.CreateSchemaInputUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateSchemaInputResponse{
		Id:         output.ID,
		Required:   output.Required,
		Properties: toMapAny(output.Properties),
		SchemaId:   output.SchemaId,
	}, nil
}
