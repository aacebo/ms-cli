package workItems

import (
	"context"
	"fmt"
	"slices"

	"github.com/aacebo/ms-cli/ado"
	"github.com/aacebo/ms-cli/cli"
	"github.com/google/go-github/v62/github"

	adoSearch "github.com/microsoft/azure-devops-go-api/azuredevops/v7/search"
)

func syncGithub(client *ado.Client) func(cli.Args) error {
	gh := github.NewClient(nil)
	defaultTags := []string{
		"AISDK",
		"beacon_board",
		"Plat_AIextensibility",
		"Plat_Ailib",
	}

	return func(args cli.Args) error {
		issues, _, err := gh.Issues.ListByRepo(
			context.Background(),
			args["gh-org"].String(),
			args["gh-repo"].String(),
			&github.IssueListByRepoOptions{
				State:       "open",
				Labels:      []string{"epic"},
				ListOptions: github.ListOptions{PerPage: 1000},
			},
		)

		if err != nil {
			return err
		}

		workItems, err := client.WorkItems(args["project"].String()).Search("Plat_Ailib")

		for j, issue := range issues {
			if j > 0 {
				break
			}

			i := slices.IndexFunc(workItems, func(item adoSearch.WorkItemResult) bool {
				title := (*item.Fields)["system.title"]
				return title == *issue.Title
			})

			fmt.Printf(
				"- [%s (%d)]:\n\t- Type => %s\n\t- State => %s\n\t- Tags => %s\n\t- ADO => %v\n\n",
				*issue.Title,
				issue.ID,
				"Epic",
				*issue.State,
				issue.Labels,
				i > -1,
			)

			if i > -1 {

			} else {
				_, err := client.WorkItems(args["project"].String()).Create(ado.CreateWorkItemArgs{
					Title:       *issue.Title,
					Type:        "Epic",
					Description: *issue.Body,
					Tags:        defaultTags,
				})

				if err != nil {
					return err
				}
			}
		}

		fmt.Println(len(issues))

		return nil
	}
}
