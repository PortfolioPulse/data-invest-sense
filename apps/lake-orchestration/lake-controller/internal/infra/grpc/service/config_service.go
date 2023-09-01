package service

import (
	"apps/lake-orchestration/lake-controller/internal/infra/grpc/pb"
	"apps/lake-orchestration/lake-controller/internal/usecase"
	"context"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

type ConfigService struct {
	pb.UnimplementedConfigServiceServer
	CreateConfigUseCase usecase.CreateConfigUseCase
}

func NewConfigService(
	createConfigUseCase usecase.CreateConfigUseCase,
) *ConfigService {
	return &ConfigService{
		CreateConfigUseCase: createConfigUseCase,
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

func convertDependsOn(pbDependsOn []*pb.DependsOn) []usecase.JobDependencies {
	jobDependsOn := make([]usecase.JobDependencies, 0, len(pbDependsOn))
	for _, pbDep := range pbDependsOn {
		jobDep := usecase.JobDependencies{
			Service: pbDep.Service,
			Source:  pbDep.Source,
		}
		jobDependsOn = append(jobDependsOn, jobDep)
	}
	return jobDependsOn
}

func ConvertJobDependenciesToPbDependsOn(jobDeps []usecase.JobDependencies) []*pb.DependsOn {
	pbDeps := make([]*pb.DependsOn, len(jobDeps))
	for i, jobDep := range jobDeps {
		pbDeps[i] = &pb.DependsOn{
			Service: jobDep.Service,
			Source:  jobDep.Source,
		}
	}
	return pbDeps
}

func (s *ConfigService) CreateConfig(ctx context.Context, in *pb.CreateConfigRequest) (*pb.CreateConfigResponse, error) {
	dto := usecase.ConfigInputDTO{
		Name:              in.Name,
		Active:            in.Active,
          Frequency:         in.Frequency,
		Service:           in.Service,
		Source:            in.Source,
		Context:           in.Context,
		DependsOn:         convertDependsOn(in.DependsOn),
		ServiceParameters: toMapStringInterface(in.ServiceParameters),
		JobParameters:     toMapStringInterface(in.JobParameters),
	}
	output, err := s.CreateConfigUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}

	return &pb.CreateConfigResponse{
		Id:                output.ID,
		Active:            output.Active,
          Frequency:         output.Frequency,
		Service:           output.Service,
		Source:            output.Source,
		Context:           output.Context,
		DependsOn:         ConvertJobDependenciesToPbDependsOn(output.DependsOn),
		ServiceParameters: toMapAny(output.ServiceParameters),
		JobParameters:     toMapAny(output.JobParameters),
		CreatedAt:         output.CreatedAt,
		UpdatedAt:         output.UpdatedAt,
	}, nil
}
