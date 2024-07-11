package ado

import (
	"context"
	"strings"

	"github.com/microsoft/azure-devops-go-api/azuredevops/v7"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/search"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/webapi"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/workitemtracking"
)

type CreateWorkItemArgs struct {
	Title       string
	Type        string
	Description string
	Tags        []string
}

type WorkItemsClient struct {
	project   string
	search    search.Client
	workItems workitemtracking.Client
}

func NewWorkItemsClient(conn *azuredevops.Connection) (*WorkItemsClient, error) {
	search, err := search.NewClient(context.Background(), conn)

	if err != nil {
		return nil, err
	}

	workItems, err := workitemtracking.NewClient(context.Background(), conn)

	if err != nil {
		return nil, err
	}

	return &WorkItemsClient{
		search:    search,
		workItems: workItems,
	}, nil
}

func (self WorkItemsClient) Search(text string) ([]search.WorkItemResult, error) {
	top := 20
	skip := 0
	results := []search.WorkItemResult{}

	for {
		if skip >= 1000 {
			break
		}

		res, err := self.search.FetchWorkItemSearchResults(context.Background(), search.FetchWorkItemSearchResultsArgs{
			Project: &self.project,
			Request: &search.WorkItemSearchRequest{
				Top:        &top,
				Skip:       &skip,
				SearchText: &text,
				Filters: &map[string][]string{
					"System.TeamProject": {self.project},
					"System.AreaPath":    {self.project},
					"System.State":       {"Proposed", "New", "Active"},
				},
			},
		})

		if err != nil {
			return nil, err
		}

		if *res.Count == 0 {
			break
		}

		results = append(results, *res.Results...)
		skip += top
	}

	return results, nil
}

func (self WorkItemsClient) Create(args CreateWorkItemArgs) (*workitemtracking.WorkItem, error) {
	titlePath := "/fields/System.Title"
	descPath := "/fields/System.Description"
	statePath := "/fields/System.State"
	tagsPath := "/fields/System.Tags"
	validateOnly := false

	return self.workItems.CreateWorkItem(context.Background(), workitemtracking.CreateWorkItemArgs{
		Project:      &self.project,
		Type:         &args.Type,
		ValidateOnly: &validateOnly,
		Document: &[]webapi.JsonPatchOperation{
			{
				Op:    &webapi.OperationValues.Add,
				Path:  &titlePath,
				Value: args.Title,
			},
			{
				Op:    &webapi.OperationValues.Add,
				Path:  &descPath,
				Value: args.Description,
			},
			{
				Op:    &webapi.OperationValues.Add,
				Path:  &statePath,
				Value: "Proposed",
			},
			{
				Op:    &webapi.OperationValues.Add,
				Path:  &tagsPath,
				Value: strings.Join(args.Tags, "; "),
			},
		},
	})
}
