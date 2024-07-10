package ado

import (
	"context"

	"github.com/microsoft/azure-devops-go-api/azuredevops/v7"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/search"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/workitemtracking"
)

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
					"System.TeamProject":  {self.project},
					"System.AreaPath":     {self.project},
					"System.WorkItemType": {"Bug", "User Story", "Feature", "Task"},
					"System.State":        {"New", "Active", "Closed"},
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
