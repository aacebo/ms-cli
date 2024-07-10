package ado

import (
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7"
)

type Client struct {
	projects  *ProjectsClient
	workItems *WorkItemsClient
}

func New(url string, accessToken string) (*Client, error) {
	conn := azuredevops.NewPatConnection(url, accessToken)
	projects, err := NewProjectsClient(conn)

	if err != nil {
		return nil, err
	}

	workItems, err := NewWorkItemsClient(conn)

	if err != nil {
		return nil, err
	}

	return &Client{
		projects:  projects,
		workItems: workItems,
	}, nil
}

func (self Client) Projects() ProjectsClient {
	return *self.projects
}

func (self Client) WorkItems(project string) WorkItemsClient {
	self.workItems.project = project
	return *self.workItems
}
