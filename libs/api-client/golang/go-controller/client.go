package gocontroller

import (
	"context"
	"fmt"
	controllerInputDTO "libs/dtos/golang/dto-controller/input"
     controllerSharedDTO "libs/dtos/golang/dto-controller/shared"
	controllerOutputDTO "libs/dtos/golang/dto-controller/output"
	gorequest "libs/golang/go-request"
	"net/http"
)

type Client struct {
	ctx     context.Context
	baseURL string
}

func NewClient() *Client {
	return &Client{
		ctx:     context.Background(),
		baseURL: "http://lake-controller:8000",
	}
}

func (c *Client) CreateConfig(configInput controllerInputDTO.ConfigDTO) (controllerOutputDTO.ConfigDTO, error) {
	url := fmt.Sprintf("%s/configs", c.baseURL)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodPost, url, configInput)
	if err != nil {
		return controllerOutputDTO.ConfigDTO{}, err
	}

	var configOutput controllerOutputDTO.ConfigDTO
	err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &configOutput)
	if err != nil {
		return controllerOutputDTO.ConfigDTO{}, err
	}

	return configOutput, nil
}

func (c *Client) ListAllConfigs() ([]controllerOutputDTO.ConfigDTO, error) {
	url := fmt.Sprintf("%s/configs", c.baseURL)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var configList []controllerOutputDTO.ConfigDTO
	err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &configList)
	if err != nil {
		return nil, err
	}

	return configList, nil
}

func (c *Client) ListOneConfigById(id string) (controllerOutputDTO.ConfigDTO, error) {
	url := fmt.Sprintf("%s/configs/%s", c.baseURL, id)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return controllerOutputDTO.ConfigDTO{}, err
	}

	var configOutput controllerOutputDTO.ConfigDTO
	err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &configOutput)
	if err != nil {
		return controllerOutputDTO.ConfigDTO{}, err
	}

	return configOutput, nil
}

func (c *Client) ListAllConfigsByService(service string) ([]controllerOutputDTO.ConfigDTO, error) {
	url := fmt.Sprintf("%s/configs/service/%s", c.baseURL, service)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var configList []controllerOutputDTO.ConfigDTO
	err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &configList)
	if err != nil {
		return nil, err
	}

	return configList, nil
}

func (c *Client) ListAllConfigsByDependentJob(service string, source string) ([]controllerOutputDTO.ConfigDTO, error) {
	url := fmt.Sprintf("%s/configs/service/%s/source/%s", c.baseURL, service, source)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var configList []controllerOutputDTO.ConfigDTO
	err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &configList)
	if err != nil {
		return nil, err
	}

	return configList, nil
}

func (c *Client) CreateProcessingJobDependencies(jobInput controllerInputDTO.ProcessingJobDependenciesDTO) (controllerOutputDTO.ProcessingJobDependenciesDTO, error) {
     url := fmt.Sprintf("%s/jobs-dependencies", c.baseURL)
     req, err := gorequest.CreateRequest(c.ctx, http.MethodPost, url, jobInput)
     if err != nil {
          return controllerOutputDTO.ProcessingJobDependenciesDTO{}, err
     }

     var dependenciesOutput controllerOutputDTO.ProcessingJobDependenciesDTO
     err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &dependenciesOutput)
     if err != nil {
          return controllerOutputDTO.ProcessingJobDependenciesDTO{}, err
     }

     return dependenciesOutput, nil
}

func (c *Client) ListOneProcessingJobDependenciesById(id string) (controllerOutputDTO.ProcessingJobDependenciesDTO, error) {
     url := fmt.Sprintf("%s/jobs-dependencies/%s", c.baseURL, id)
     req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
     if err != nil {
          return controllerOutputDTO.ProcessingJobDependenciesDTO{}, err
     }

     var dependenciesOutput controllerOutputDTO.ProcessingJobDependenciesDTO
     err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &dependenciesOutput)
     if err != nil {
          return controllerOutputDTO.ProcessingJobDependenciesDTO{}, err
     }

     return dependenciesOutput, nil
}

func (c *Client) RemoveProcessingJobDependencies(id string) (controllerOutputDTO.ProcessingJobDependenciesDTO, error) {
     url := fmt.Sprintf("%s/jobs-dependencies/%s", c.baseURL, id)
     req, err := gorequest.CreateRequest(c.ctx, http.MethodDelete, url, nil)
     if err != nil {
          return controllerOutputDTO.ProcessingJobDependenciesDTO{}, err
     }

     var dependenciesOutput controllerOutputDTO.ProcessingJobDependenciesDTO
     err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &dependenciesOutput)
     if err != nil {
          return controllerOutputDTO.ProcessingJobDependenciesDTO{}, err
     }

     return dependenciesOutput, nil
}

func (c *Client) UpdateProcessingJobDependencies(id string, jobDep controllerSharedDTO.ProcessingJobDependencies) (controllerOutputDTO.ProcessingJobDependenciesDTO, error) {
     url := fmt.Sprintf("%s/jobs-dependencies/%s", c.baseURL, id)
     req, err := gorequest.CreateRequest(c.ctx, http.MethodPost, url, jobDep)
     if err != nil {
          return controllerOutputDTO.ProcessingJobDependenciesDTO{}, err
     }

     var dependenciesOutput controllerOutputDTO.ProcessingJobDependenciesDTO
     err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &dependenciesOutput)
     if err != nil {
          return controllerOutputDTO.ProcessingJobDependenciesDTO{}, err
     }

     return dependenciesOutput, nil
}
