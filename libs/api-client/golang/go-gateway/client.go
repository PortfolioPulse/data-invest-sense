package gogateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	inputDTO "libs/dtos/golang/dto-gateway/input"
	outputDTO "libs/dtos/golang/dto-gateway/output"
)

const (
	baseURL = "http://lake-gateway:8000"
)

var httpClient = &http.Client{
	Timeout: time.Second * 10,
}

type Client struct {
     ctx context.Context
}

func NewClient() *Client {
     return &Client{
          ctx: context.Background(),
     }
}

func (c *Client) createRequest(method, url string, body interface{}) (*http.Request, error) {
	inputJSON, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(c.ctx, method, url, bytes.NewBuffer(inputJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func sendRequest(req *http.Request, result interface{}) error {
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("failed to decode response JSON: %w", err)
	}

	return nil
}

func (c *Client) CreateInput(service, source string, input inputDTO.InputDTO) (outputDTO.InputDTO, error) {
	url := fmt.Sprintf("%s/service/%s/source/%s", baseURL, service, source)
	req, err := c.createRequest(http.MethodPost, url, input)
	if err != nil {
		return outputDTO.InputDTO{}, err
	}

	var output outputDTO.InputDTO
	if err := sendRequest(req, &output); err != nil {
		return outputDTO.InputDTO{}, err
	}

	return output, nil
}

func (c *Client) ListAllInputsByServiceAndSource(service, source string) ([]outputDTO.InputDTO, error) {
	url := fmt.Sprintf("%s/service/%s/source/%s", baseURL, service, source)
	req, err := c.createRequest(http.MethodGet, url, nil)
     if err != nil {
          return nil, fmt.Errorf("failed to create HTTP request: %w", err)
     }

     var output []outputDTO.InputDTO
     if err := sendRequest(req, &output); err != nil {
          return nil, err
     }

     return output, nil
}

func (c *Client) ListAllInputsByService(service string) ([]outputDTO.InputDTO, error) {
	url := fmt.Sprintf("%s/service/%s", baseURL, service)
	req, err := c.createRequest(http.MethodGet, url, nil)
     if err != nil {
          return nil, fmt.Errorf("failed to create HTTP request: %w", err)
     }

	var output []outputDTO.InputDTO
	if err := sendRequest(req, &output); err != nil {
		return nil, err
	}

	return output, nil
}

func (c *Client) ListOneInputByIdAndService(id, service, source string) (outputDTO.InputDTO, error) {
	url := fmt.Sprintf("%s/service/%s/source/%s/%s", baseURL, service, source, id)
	req, err := c.createRequest(http.MethodGet, url, nil)
     if err != nil {
          return outputDTO.InputDTO{}, fmt.Errorf("failed to create HTTP request: %w", err)
     }

     var output outputDTO.InputDTO
     if err := sendRequest(req, &output); err != nil {
          return outputDTO.InputDTO{}, err
     }

     return output, nil
}

func (c *Client) UpdateInputStatus(inputStatus inputDTO.InputStatusDTO, service string, source string) (outputDTO.InputDTO, error) {
     // id := inputStatus.ID
	url := fmt.Sprintf("%s/service/%s/source/%s", baseURL, service, source)
	req, err := c.createRequest(http.MethodPost, url, inputStatus)
     if err != nil {
          return outputDTO.InputDTO{}, fmt.Errorf("failed to create HTTP request: %w", err)
     }

	var output outputDTO.InputDTO
	if err := sendRequest(req, &output); err != nil {
		return outputDTO.InputDTO{}, err
	}

	return output, nil
}

func (c *Client) CreateStagingJob(stagingJob inputDTO.StagingJobDTO) (outputDTO.StagingJobDTO, error) {
	url := fmt.Sprintf("%s/staging-jobs", baseURL)
	req, err := c.createRequest(http.MethodPost, url, stagingJob)
	if err != nil {
		return outputDTO.StagingJobDTO{}, err
	}

	var output outputDTO.StagingJobDTO
	if err := sendRequest(req, &output); err != nil {
		return outputDTO.StagingJobDTO{}, err
	}

	return output, nil
}

func (c *Client) RemoveStagingJob(id string) (outputDTO.StagingJobDTO, error) {
	url := fmt.Sprintf("%s/staging-jobs/%s", baseURL, id)
	req, err := c.createRequest(http.MethodDelete, url, nil)
	if err != nil {
		return outputDTO.StagingJobDTO{}, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	var output outputDTO.StagingJobDTO
	if err := sendRequest(req, &output); err != nil {
		return outputDTO.StagingJobDTO{}, err
	}

	return output, nil
}

func (c *Client) ListOneStagingJobUsingServiceSourceInputId(service, source, inputId string) (outputDTO.StagingJobDTO, error) {
     url := fmt.Sprintf("%s/staging-jobs/service/%s/source/%s/id/%s", baseURL, service, source, inputId)
     req, err := c.createRequest(http.MethodGet, url, nil)
     if err != nil {
          return outputDTO.StagingJobDTO{}, fmt.Errorf("failed to create HTTP request: %w", err)
     }

     var output outputDTO.StagingJobDTO
     if err := sendRequest(req, &output); err != nil {
          return outputDTO.StagingJobDTO{}, err
     }

     return output, nil
}
