/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bitbucket/bitbucketapi"
	"bitbucket/cliformat"
	"bitbucket/formatters"
	githelper "bitbucket/git"
	"encoding/json"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

type CreatePullRequestData struct {
	Title  string `json:"title"`
	Source struct {
		Branch struct {
			Name string `json:"name"`
		} `json:"branch"`
	} `json:"source"`
}

type CreatePullRequestDataWithDestenation struct {
	Title  string `json:"title"`
	Source struct {
		Branch struct {
			Name string `json:"name"`
		} `json:"branch"`
	} `json:"source"`
	Dest struct {
		Branch struct {
			Name string `json:"name"`
		} `json:"branch"`
	} `json:"destination"`
}
type CreatePullRequestsResponse struct {
	Links struct {
		Html struct {
			Href string `json:"href"`
		} `json:"html"`
	} `json:"links"`
}

var pullRequestCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create pull request",
	Long:  `Created pull reqeust`,

	Run: func(cmd *cobra.Command, args []string) {

		repo, err := githelper.GetCurrentRepo(cmd)
		if err != nil {
			cliformat.Error(err.Error())
		}
		url := repo + "/pullrequests"
		if err != nil {
			fmt.Println(cliformat.Error(err.Error()))
			return
		}
		title, err := cmd.Flags().GetString("title")
		if err != nil {
			fmt.Println(cliformat.Error(err.Error()))
			return
		}
		source, err := cmd.Flags().GetString("source")
		if err != nil {
			fmt.Println(cliformat.Error(err.Error()))
			return
		}
		dest, err := cmd.Flags().GetString("dest")
		if err != nil {
			fmt.Println(cliformat.Error(err.Error()))
			return
		}
		var requestDataJson []byte
		if dest == "" {
			createPullRequestData := &CreatePullRequestData{}
			createPullRequestData.Title = title
			createPullRequestData.Source.Branch.Name = source

			requestDataJson, err = json.Marshal(createPullRequestData)
		} else {

			createPullRequestData := &CreatePullRequestDataWithDestenation{}
			createPullRequestData.Title = title
			createPullRequestData.Source.Branch.Name = source
			createPullRequestData.Dest.Branch.Name = dest

			requestDataJson, err = json.Marshal(createPullRequestData)

		}
		resp, err := bitbucketapi.HttpRequestWithBitbucketAuthJson("POST", url, requestDataJson)

		defer resp.Body.Close()

		if err != nil {
			fmt.Println(cliformat.Error(err.Error()))
		}

		if resp.StatusCode != 201 {
			bodyText, _ := io.ReadAll(resp.Body)
			fmt.Println(cliformat.Error(string(bodyText)))
			return
		}
		var pullrequest CreatePullRequestsResponse
		json.NewDecoder(resp.Body).Decode(&pullrequest)
		fmt.Println("PR created\n" + pullrequest.Links.Html.Href)

	},
}

func getPullRequestDetails(cmd *cobra.Command, args []string) {
	repo, err := githelper.GetCurrentRepo(cmd)
	if err != nil {
		fmt.Println(cliformat.Error("No repo provided and current directory doesn't have a git remote repo"))
		return
	}
	url := ""
	if err != nil {
		return
	}
	if len(args) > 0 {
		url = repo + "/pullrequests/" + args[0] + "/activity?pagelen=50&sort=-created_on"
		resp, err := bitbucketapi.HttpRequestWithBitbucketAuthJson("GET", url, map[string]string{})
		defer resp.Body.Close()
		if err != nil {
			fmt.Println(cliformat.Error(err.Error()))
		}
		if resp.StatusCode != 200 {
			bodyText, _ := io.ReadAll(resp.Body)
			fmt.Println(cliformat.Error(string(bodyText)))
			return
		}
		var prActivitites formatters.PullRequestActivities
		json.NewDecoder(resp.Body).Decode(&prActivitites)
		formatters.FormatPullrequestActivitites(prActivitites)
	}

}

func listPullRequests(cmd *cobra.Command) {
	repo, err := githelper.GetCurrentRepo(cmd)
	if err != nil {
		fmt.Println(cliformat.Error("No repo provided and current directory doesn't have a git remote repo"))
		return
	}
	url := repo + "/pullrequests?pagelen=20&sort=-created_on"

	state, err := cmd.Flags().GetString("state")
	if err != nil || state != "" {
		url = url + "&state=" + state

	}
	page, err := cmd.Flags().GetString("page")
	if err != nil || page != "" {
		url = url + "&page=" + page

	}
	resp, err := bitbucketapi.HttpRequestWithBitbucketAuthJson("GET", url, map[string]string{})

	defer resp.Body.Close()

	if err != nil {
		fmt.Println(cliformat.Error(err.Error()))
	}

	if resp.StatusCode != 200 {
		bodyText, _ := io.ReadAll(resp.Body)
		fmt.Println(cliformat.Error(string(bodyText)))
		return
	}
	var pullrequests formatters.PullRequestsResponse
	json.NewDecoder(resp.Body).Decode(&pullrequests)
	formatters.FormatPullrequest(pullrequests)

}

// pipelinesCmd represents the pipelines command
var pullRequestsCmd = &cobra.Command{
	Use:   "pr",
	Short: "List pull-requests",
	Long:  `List pull-requests`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			listPullRequests(cmd)
			return
		}
		getPullRequestDetails(cmd, args)

	},
}

func init() {
	pullRequestsCmd.PersistentFlags().StringP("repo", "r", "", "Repo remote url")
	pullRequestsCmd.Flags().StringP("page", "p", "", "Page number for pullreuest pagination")
	pullRequestsCmd.Flags().StringP("state", "s", "", "Pull request state")
	pullRequestCreateCmd.Flags().StringP("source", "s", "", "Pull request source branch")
	pullRequestCreateCmd.Flags().StringP("title", "t", "", "Pull request title")
	pullRequestCreateCmd.Flags().StringP("dest", "d", "", "Pull request destenation branch")
	pullRequestCreateCmd.MarkFlagRequired("source")
	pullRequestCreateCmd.MarkFlagRequired("title")
	pullRequestCreateCmd.RegisterFlagCompletionFunc("dest", githelper.GetBranchSuggestions)
	pullRequestCreateCmd.RegisterFlagCompletionFunc("source", githelper.GetBranchSuggestions)
	pullRequestsCmd.AddCommand(pullRequestCreateCmd)
	rootCmd.AddCommand(pullRequestsCmd)
}
