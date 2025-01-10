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

func getEnvPipeline(repo string, id string) formatters.PipelineDetailsResponse {
	var details formatters.PipelineDetailsResponse
	detailsRes, err := bitbucketapi.HttpRequestWithBitbucketAuthJson("GET", repo+"/pipelines/"+id, map[string]string{})
	defer detailsRes.Body.Close()
	if err != nil {
		return details
	}
	json.NewDecoder(detailsRes.Body).Decode(&details)
	return details

}

var environmentCmd = &cobra.Command{
	Use:   "envs",
	Short: "List environments",
	Long:  `List envs`,

	Run: func(cmd *cobra.Command, args []string) {
		page, err := cmd.Flags().GetString("page")
		if err != nil || page == "" {
			page = "1"
		}
		repo, err := githelper.GetCurrentRepo(cmd)
		if err != nil {
			fmt.Println(cliformat.Error("No repo provided and current directory doesn't have a git remote repo"))
			return
		}
		envsRes, err := bitbucketapi.HttpRequestWithBitbucketAuthJson("GET", repo+"/environments?page="+page, map[string]string{})
		defer envsRes.Body.Close()
		if err != nil {

			fmt.Println(cliformat.Error(err.Error()))
			return
		}
		var envResJson formatters.EnvResponse
		json.NewDecoder(envsRes.Body).Decode(&envResJson)
		grouped, err := cmd.Flags().GetBool("grouped")
		if err != nil {
			grouped = false
		}
		if envsRes.StatusCode != 200 {
			bodyText, _ := io.ReadAll(envsRes.Body)
			fmt.Println(cliformat.Error(string(bodyText)))
			return
		}
		ch := make(chan formatters.PipelineDetailsResponse, len(envResJson.Values))
		for i := 0; i < len(envResJson.Values); i++ {
			go func(index int) {

				ch <- getEnvPipeline(repo, envResJson.Values[index].Lock.Triggerer.PipelineUUID)
			}(i)
		}
		pipelineMap := map[string]formatters.PipelineDetailsResponse{}
		for range envResJson.Values {
			pipeline := <-ch
			pipelineMap[pipeline.UUID] = pipeline

		}
		for j := 0; j < len(envResJson.Values); j++ {
			envResJson.Values[j].Pipeline = pipelineMap[envResJson.Values[j].Lock.Triggerer.PipelineUUID]
		}

		isCustomFormat, err := formatters.CustomFormat(cmd, envResJson, "Values")

		if err != nil {
			fmt.Println(cliformat.Error(err.Error()))
		}

		if !isCustomFormat {
			fmt.Println(formatters.FormatEnvs(envResJson, grouped))
		}

	},
}

func init() {
	environmentCmd.PersistentFlags().StringP("repo", "r", "", "Repo remote url")
	environmentCmd.Flags().BoolP("json", "j", false, "Output as json")
	environmentCmd.Flags().BoolP("grouped", "g", true, "Group enviroments by type")
	environmentCmd.Flags().StringP("page", "p", "", "Page number for environments pagination")
	environmentCmd.Flags().StringP("format", "f", "", "Output template format")

	rootCmd.AddCommand(environmentCmd)

}
