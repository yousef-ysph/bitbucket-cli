/*
Copyright © 2024 yousef@ysph.tech
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
	activityUrl := repo + "/pullrequests/" + args[0] + "/activity?pagelen=50&sort=-created_on"
	detailsUrl := repo + "/pullrequests/" + args[0]
	activityFetchChan := make(chan bitbucketapi.FetchToChanResponse)
	detailsFetchChan := make(chan bitbucketapi.FetchToChanResponse)
	go bitbucketapi.FetchToChannelJson("GET", activityUrl, map[string]string{}, activityFetchChan)
	go bitbucketapi.FetchToChannelJson("GET", detailsUrl, map[string]string{}, detailsFetchChan)
	activityResp := <-activityFetchChan
	detailsResp := <-detailsFetchChan
	defer activityResp.Resp.Body.Close()
	defer detailsResp.Resp.Body.Close()
	if activityResp.Err != nil || detailsResp.Err != err {
		fmt.Println(cliformat.Error(activityResp.Err.Error()))
		fmt.Println(cliformat.Error(detailsResp.Err.Error()))
	}

	if detailsResp.Resp.StatusCode != 200 {
		bodyText, _ := io.ReadAll(detailsResp.Resp.Body)
		fmt.Println(cliformat.Error(string(bodyText)))
		return
	}

	if activityResp.Resp.StatusCode != 200 {
		bodyText, _ := io.ReadAll(activityResp.Resp.Body)
		fmt.Println(cliformat.Error(string(bodyText)))
		return
	}

	var prActivitites formatters.PullRequestActivities
	var prDetails formatters.PullRequest
	json.NewDecoder(activityResp.Resp.Body).Decode(&prActivitites)
	json.NewDecoder(detailsResp.Resp.Body).Decode(&prDetails)
	isCustomFormat, err := formatters.CustomFormat(cmd, prDetails, "")

	if err != nil {
		fmt.Println(cliformat.Error(err.Error()))
	}

	if !isCustomFormat {
		fmt.Println(formatters.FormatPullrequest(prDetails))
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
	isCustomFormat, err := formatters.CustomFormat(cmd, pullrequests, "Values")

	if err != nil {
		fmt.Println(cliformat.Error(err.Error()))
	}

	if !isCustomFormat {
		formatters.FormatPullrequestResponse(pullrequests)
	}

}

type MergePullRequestsData struct {
	Type              string `json:"type"`
	Message           string `json:"message"`
	CloseSourceBranch bool   `json:"close_source_branch"`
	MergeStrategy     string `json:"merge_strategy"`
}

// pipelinesCmd represents the pipelines command
var mergePullRequestsCmd = &cobra.Command{
	Use:   "merge",
	Short: "merge pull-request",
	Long:  `merge pull-request`,

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println(cliformat.Error("Pull request id is required"))
			return
		}
		repo, err := githelper.GetCurrentRepo(cmd)
		var mergePullRequestsCmd MergePullRequestsData
		mergePullRequestsCmd.CloseSourceBranch, err = cmd.Flags().GetBool("close-source")
		mergePullRequestsCmd.Message, err = cmd.Flags().GetString("message")
		mergePullRequestsCmd.MergeStrategy, err = cmd.Flags().GetString("strategy")
		mergePullRequestsCmd.Type, err = cmd.Flags().GetString("type")
		requestDataJson, err := json.Marshal(mergePullRequestsCmd)

		if err != nil {
			fmt.Println(cliformat.Error("No repo provided and current directory doesn't have a git remote repo"))
			return
		}
		url := repo + "/pullrequests/" + args[0] + "/merge"

		resp, err := bitbucketapi.HttpRequestWithBitbucketAuthJson("POST", url, requestDataJson)
		defer resp.Body.Close()
		if err != nil {
			fmt.Println(cliformat.Error(err.Error()))
		}

		if resp.StatusCode != 200 {
			bodyText, _ := io.ReadAll(resp.Body)
			fmt.Println(resp.StatusCode)
			fmt.Println(cliformat.Error(string(bodyText)))
			return
		}
		fmt.Println(cliformat.Success("PR Merged"))

	},
}

var declinePullRequestsCmd = &cobra.Command{
	Use:   "decline",
	Short: "decline pull-request",
	Long:  `decline pull-request`,

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println(cliformat.Error("Pull request id is required"))
			return
		}
		repo, err := githelper.GetCurrentRepo(cmd)

		if err != nil {
			fmt.Println(cliformat.Error("No repo provided and current directory doesn't have a git remote repo"))
			return
		}
		url := repo + "/pullrequests/" + args[0] + "/decline"

		resp, err := bitbucketapi.HttpRequestWithBitbucketAuthJson("POST", url, map[string]string{})
		defer resp.Body.Close()
		if err != nil {
			fmt.Println(cliformat.Error(err.Error()))
		}

		if resp.StatusCode != 200 {
			bodyText, _ := io.ReadAll(resp.Body)
			fmt.Println(resp.StatusCode)
			fmt.Println(cliformat.Error(string(bodyText)))
			return
		}
		fmt.Println(cliformat.Success("PR Declined"))

	},
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

func getStrategies(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {

	strategies := []string{"merge_commit", "squash", "fast_forward"}
	return strategies, cobra.ShellCompDirectiveDefault

}

func init() {
	pullRequestsCmd.PersistentFlags().StringP("repo", "r", "", "Repo remote url")
	pullRequestsCmd.Flags().StringP("page", "p", "", "Page number for pullreuest pagination")
	pullRequestsCmd.Flags().StringP("format", "f", "", "Output template format")
	pullRequestsCmd.Flags().BoolP("json", "j", false, "Output as json")
	pullRequestsCmd.Flags().StringP("state", "s", "", "Pull request state")
	pullRequestCreateCmd.Flags().StringP("source", "s", "", "Pull request source branch")
	pullRequestCreateCmd.Flags().StringP("title", "t", "", "Pull request title")
	pullRequestCreateCmd.Flags().StringP("dest", "d", "", "Pull request destenation branch")
	pullRequestCreateCmd.MarkFlagRequired("source")
	pullRequestCreateCmd.MarkFlagRequired("title")
	pullRequestCreateCmd.RegisterFlagCompletionFunc("dest", githelper.GetBranchSuggestions)
	pullRequestCreateCmd.RegisterFlagCompletionFunc("source", githelper.GetBranchSuggestions)
	pullRequestsCmd.AddCommand(pullRequestCreateCmd)
	mergePullRequestsCmd.Flags().StringP("message", "m", "", "Merge pull request message")
	mergePullRequestsCmd.Flags().StringP("strategy", "s", "", "Merge strategy")
	mergePullRequestsCmd.Flags().BoolP("close-source", "c", false, "Close source branch")
	mergePullRequestsCmd.Flags().StringP("type", "t", "type", "Close source branch")
	mergePullRequestsCmd.RegisterFlagCompletionFunc("strategy", getStrategies)
	pullRequestsCmd.AddCommand(mergePullRequestsCmd)
	pullRequestsCmd.AddCommand(declinePullRequestsCmd)
	rootCmd.AddCommand(pullRequestsCmd)
}
