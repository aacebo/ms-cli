package ado

import (
	"context"

	"github.com/microsoft/azure-devops-go-api/azuredevops/v7"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/core"
)

type Client struct {
	client   core.Client
	projects ProjectsClient
}

func (self Client) Projects() ProjectsClient {
	return self.projects
}

func New(url string, accessToken string) (*Client, error) {
	connection := azuredevops.NewPatConnection(url, accessToken)
	client, err := core.NewClient(context.Background(), connection)

	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
		projects: ProjectsClient{
			client: client,
		},
	}, nil
}
