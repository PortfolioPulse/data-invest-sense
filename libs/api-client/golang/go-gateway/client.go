package gogateway

import (
	"context"
	"fmt"
	"net/http"

	inputDTO "libs/dtos/golang/dto-gateway/input"
	outputDTO "libs/dtos/golang/dto-gateway/output"
	gorequest "libs/golang/go-request"
)

type Client struct {
	ctx     context.Context
	baseURL string
}

func NewClient() *Client {
	return &Client{
		ctx:     context.Background(),
		baseURL: "http://lake-gateway:8000",
	}
}

func (c *Client) CreateInput(service, source string, input inputDTO.InputDTO) (outputDTO.InputDTO, error) {
	url := fmt.Sprintf("%s/service/%s/source/%s", c.baseURL, service, source)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodPost, url, input)
	if err != nil {
		return outputDTO.InputDTO{}, err
	}

	var output outputDTO.InputDTO
	if err := gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &output); err != nil {
		return outputDTO.InputDTO{}, err
	}

	return output, nil
}

func (c *Client) ListAllInputsByServiceAndSource(service, source string) ([]outputDTO.InputDTO, error) {
	url := fmt.Sprintf("%s/service/%s/source/%s", c.baseURL, service, source)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	var output []outputDTO.InputDTO
	if err := gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &output); err != nil {
		return nil, err
	}

	return output, nil
}

func (c *Client) ListAllInputsByService(service string) ([]outputDTO.InputDTO, error) {
	url := fmt.Sprintf("%s/service/%s", c.baseURL, service)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	var output []outputDTO.InputDTO
	if err := gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &output); err != nil {
		return nil, err
	}

	return output, nil
}

func (c *Client) ListOneInputByIdAndService(id, service, source string) (outputDTO.InputDTO, error) {
	url := fmt.Sprintf("%s/service/%s/source/%s/%s", c.baseURL, service, source, id)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return outputDTO.InputDTO{}, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	var output outputDTO.InputDTO
	if err := gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &output); err != nil {
		return outputDTO.InputDTO{}, err
	}

	return output, nil
}

func (c *Client) UpdateInputStatus(inputStatus inputDTO.InputStatusDTO, service string, source string, id string) (outputDTO.InputDTO, error) {
	// id := inputStatus.ID
	url := fmt.Sprintf("%s/service/%s/source/%s/%s", c.baseURL, service, source, id)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodPost, url, inputStatus)
	if err != nil {
		return outputDTO.InputDTO{}, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	var output outputDTO.InputDTO
	if err := gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &output); err != nil {
		return outputDTO.InputDTO{}, err
	}

	return output, nil
}

func (c *Client) CreateStagingJob(stagingJob inputDTO.StagingJobDTO) (outputDTO.StagingJobDTO, error) {
	url := fmt.Sprintf("%s/staging-jobs", c.baseURL)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodPost, url, stagingJob)
	if err != nil {
		return outputDTO.StagingJobDTO{}, err
	}

	var output outputDTO.StagingJobDTO
	if err := gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &output); err != nil {
		return outputDTO.StagingJobDTO{}, err
	}

	return output, nil
}

func (c *Client) RemoveStagingJob(id string) (outputDTO.StagingJobDTO, error) {
	url := fmt.Sprintf("%s/staging-jobs/%s", c.baseURL, id)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodDelete, url, nil)
	if err != nil {
		return outputDTO.StagingJobDTO{}, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	var output outputDTO.StagingJobDTO
	if err := gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &output); err != nil {
		return outputDTO.StagingJobDTO{}, err
	}

	return output, nil
}

func (c *Client) ListOneStagingJobUsingServiceSourceInputId(service, source, inputId string) (outputDTO.StagingJobDTO, error) {
	url := fmt.Sprintf("%s/staging-jobs/service/%s/source/%s/%s", c.baseURL, service, source, inputId)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return outputDTO.StagingJobDTO{}, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	var output outputDTO.StagingJobDTO
	if err := gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &output); err != nil {
		return outputDTO.StagingJobDTO{}, err
	}

	return output, nil
}
