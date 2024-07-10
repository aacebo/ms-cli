package ado

import (
	"context"

	"github.com/microsoft/azure-devops-go-api/azuredevops/v7"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/workitemtracking"
)

type WorkItemsClient struct {
	client workitemtracking.Client
}

func NewWorkItemsClient(conn *azuredevops.Connection) (*WorkItemsClient, error) {
	client, err := workitemtracking.NewClient(context.Background(), conn)

	if err != nil {
		return nil, err
	}

	return &WorkItemsClient{client: client}, nil
}
