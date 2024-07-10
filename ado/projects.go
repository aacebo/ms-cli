package ado

import (
	"context"
	"strconv"

	"github.com/microsoft/azure-devops-go-api/azuredevops/v7"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/core"
)

type ProjectsClient struct {
	client core.Client
}

func NewProjectsClient(conn *azuredevops.Connection) (*ProjectsClient, error) {
	client, err := core.NewClient(context.Background(), conn)

	if err != nil {
		return nil, err
	}

	return &ProjectsClient{client: client}, nil
}

func (self ProjectsClient) List() ([]core.TeamProjectReference, error) {
	res, err := self.client.GetProjects(context.Background(), core.GetProjectsArgs{})

	if err != nil {
		return nil, err
	}

	projects := []core.TeamProjectReference{}

	for {
		projects = append(projects, res.Value...)

		if res.ContinuationToken == "" {
			break
		}

		continuationToken, err := strconv.Atoi(res.ContinuationToken)

		if err != nil {
			return nil, err
		}

		res, err = self.client.GetProjects(context.Background(), core.GetProjectsArgs{
			ContinuationToken: &continuationToken,
		})

		if err != nil {
			return nil, err
		}
	}

	return projects, nil
}
