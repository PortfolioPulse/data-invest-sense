package gorepository

import (
	"context"
	"fmt"
	repositoryInputDTO "libs/dtos/golang/dto-repository/input"
	repositoryOutputDTO "libs/dtos/golang/dto-repository/output"
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
		baseURL: "http://lake-repository:8000",
	}
}

func (c *Client) CreateSchema(schemaInput repositoryInputDTO.SchemaDTO) (repositoryOutputDTO.SchemaDTO, error) {
	url := fmt.Sprintf("%s/schemas", c.baseURL)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodPost, url, schemaInput)
	if err != nil {
		return repositoryOutputDTO.SchemaDTO{}, err
	}

	var schemaOutput repositoryOutputDTO.SchemaDTO
	err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &schemaOutput)
	if err != nil {
		return repositoryOutputDTO.SchemaDTO{}, err
	}

	return schemaOutput, nil
}

func (c *Client) ListAllSchemas() ([]repositoryOutputDTO.SchemaDTO, error) {
	url := fmt.Sprintf("%s/schemas", c.baseURL)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var schemaList []repositoryOutputDTO.SchemaDTO
	err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &schemaList)
	if err != nil {
		return nil, err
	}

	return schemaList, nil
}

func (c *Client) ListOneSchemaById(id string) (repositoryOutputDTO.SchemaDTO, error) {
	url := fmt.Sprintf("%s/schemas/%s", c.baseURL, id)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return repositoryOutputDTO.SchemaDTO{}, err
	}

	var schemaOutput repositoryOutputDTO.SchemaDTO
	err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &schemaOutput)
	if err != nil {
		return repositoryOutputDTO.SchemaDTO{}, err
	}

	return schemaOutput, nil
}

func (c *Client) ListAllSchemasByService(service string) ([]repositoryOutputDTO.SchemaDTO, error) {
	url := fmt.Sprintf("%s/schemas/service/%s", c.baseURL, service)
	req, err := gorequest.CreateRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var schemaList []repositoryOutputDTO.SchemaDTO
	err = gorequest.SendRequest(req, gorequest.DefaultHTTPClient, &schemaList)
	if err != nil {
		return nil, err
	}

	return schemaList, nil
}
