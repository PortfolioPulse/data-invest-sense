package service

import (
	"apps/lake-orchestration/lake-gateway/internal/infra/grpc/pb"
	"apps/lake-orchestration/lake-gateway/internal/usecase"
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

type InputService struct {
	pb.UnimplementedInputServiceServer
	CreateInputUseCase usecase.CreateInputUseCase
}

func NewInputService(
	createInputUseCase usecase.CreateInputUseCase,
) *InputService {
	return &InputService{
		CreateInputUseCase: createInputUseCase,
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

func (s *InputService) CreateInput(ctx context.Context, in *pb.CreateInputRequest) (*pb.CreateInputResponse, error) {
     md, ok := metadata.FromIncomingContext(ctx)
     if !ok {
          return nil, nil
     }
     service := md.Get("service")[0]
     source := md.Get("source")[0]

     dto := usecase.InputInputDTO{
		Data: toMapStringInterface(in.Data),
	}
	output, err := s.CreateInputUseCase.Execute(dto, service, source)
	if err != nil {
		return nil, err
	}
	return &pb.CreateInputResponse{
		Id:   output.ID,
		Data: toMapAny(output.Data),
		Metadata: &pb.Metadata{
			ProcessingId:        output.Metadata.ProcessingId,
			ProcessingTimestamp: output.Metadata.ProcessingTimestamp,
			Source:              output.Metadata.Source,
			Service:             output.Metadata.Service,
		},
		Status: &pb.Status{
			Code:   int32(output.Status.Code),
			Detail: output.Status.Detail,
		},
	}, nil
}



// import (
//      "context"
//      "google.golang.org/grpc/metadata"
//  )

//  service := "your-service-value"
//  source := "your-source-value"

//  ctx := metadata.NewOutgoingContext(
//      context.Background(),
//      metadata.Pairs(
//          "service", service,
//          "source", source,
//      ),
//  )

//  response, err := client.CreateInput(ctx, &pb.InputRequest{ /* ... */ })
