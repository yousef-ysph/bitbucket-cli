/*
Copyright © 2024 yousef@ysph.tech
*/

package cmd

import (
	"bitbucket/bitbucketapi"
	"bitbucket/cliformat"
	"bitbucket/constants"
	"bitbucket/formatters"
	githelper "bitbucket/git"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
)

// pipelinesCmd represents the pipelines command
func listPipelines(cmd *cobra.Command) {

	repo, err := githelper.GetCurrentRepo(cmd)
	if err != nil {
		fmt.Println(cliformat.Error("No repo provided and current directory doesn't have a git remote repo"))
		return
	}
	page, err := cmd.Flags().GetString("page")
	if err != nil {
		page = "1"
	}
	resp, err := bitbucketapi.HttpRequestWithBitbucketAuthJson("GET", repo+"/pipelines?pagelen=20&sort=-created_on&page="+page, map[string]string{})
	defer resp.Body.Close()

	if err != nil {
		cliformat.Error(err.Error())
	}

	if resp.StatusCode != 200 {
		bodyText, _ := io.ReadAll(resp.Body)
		fmt.Println(cliformat.Error(string(bodyText)))
		return
	}

	var res formatters.PipelinesResponse
	json.NewDecoder(resp.Body).Decode(&res)

	isCustomFormat, err := formatters.CustomFormat(cmd, res, "Values")
	if err != nil {
		cliformat.Error(err.Error())
	}
	if !isCustomFormat {
		formatters.FormatPipelines(res)
	}
}
func getStepLogById(repo string, pipelineId string, stepUUID string) {
	detailsRes, err := bitbucketapi.HttpRequestWithBitbucketAuth("GET", repo+"/pipelines/"+pipelineId+"/steps/"+stepUUID+"/log", map[string]string{}, "")
	defer detailsRes.Body.Close()
	if err != nil {
		cliformat.Error(err.Error())
	}
	if detailsRes.StatusCode != 200 {
		bodyText, _ := io.ReadAll(detailsRes.Body)
		fmt.Println(cliformat.Error(string(bodyText)))
		return
	}

	bodyText, _ := io.ReadAll(detailsRes.Body)
	fmt.Println(string(bodyText))

}

var getPipelineStep = &cobra.Command{
	Use:   "step",
	Short: "Pipeline step log",
	Long:  `Pipeline step log`,

	Run: func(cmd *cobra.Command, args []string) {
		repo, err := githelper.GetCurrentRepo(cmd)
		if err != nil {
			fmt.Println(cliformat.Error("No repo provided and current directory doesn't have a git remote repo"))
			return
		}

		pipelineId, err := cmd.Flags().GetString("pipelineId")
		if err != nil {
			cliformat.Error(err.Error())
		}

		stepUUID, err := cmd.Flags().GetString("step")
		if err != nil {
			cliformat.Error(err.Error())
		}

		getStepLogById(repo, pipelineId, stepUUID)
	},
}

func getStepLogByStatus(repo string, pipelineId string, steps formatters.PipelineStepsResponse, state string) bool {
	for stepIndex := 0; stepIndex < len(steps.Values); stepIndex++ {
		if steps.Values[stepIndex].State.Result.Name == state {
			getStepLogById(repo, pipelineId, steps.Values[stepIndex].UUID)
			return true
		}
	}
	return false

}

func getPipeline(cmd *cobra.Command, args []string) {
	repo, err := githelper.GetCurrentRepo(cmd)
	if err != nil {
		fmt.Println(cliformat.Error("No repo provided and current directory doesn't have a git remote repo"))
		return
	}
	pipelineId := args[0]
	stepsRes, err := bitbucketapi.HttpRequestWithBitbucketAuthJson("GET", repo+"/pipelines/"+pipelineId+"/steps", map[string]string{})
	var steps formatters.PipelineStepsResponse
	json.NewDecoder(stepsRes.Body).Decode(&steps)
	state, err := cmd.Flags().GetString("state")
	if state != "" {
		if getStepLogByStatus(repo, pipelineId, steps, state) {
			return
		}
	}

	detailsRes, err := bitbucketapi.HttpRequestWithBitbucketAuthJson("GET", repo+"/pipelines/"+pipelineId, map[string]string{})
	defer stepsRes.Body.Close()
	defer detailsRes.Body.Close()
	var details formatters.PipelineDetailsResponse
	json.NewDecoder(detailsRes.Body).Decode(&details)
	if err != nil {
		cliformat.Error(err.Error())
	}
	isCustomFormat, err := formatters.CustomFormat(cmd, details, "")
	if err != nil {
		cliformat.Error(err.Error())
	}
	if !isCustomFormat {
		formatters.FormatPipelineDetailsWithSteps(details, steps, cmd.Flag("detailed").Changed)
	}

}

var pipelinesCmd = &cobra.Command{
	Use:   "pipelines",
	Short: "Show pipelines",
	Long:  `Show pipelines`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			getPipeline(cmd, args)
			return
		}
		listPipelines(cmd)
	},
}

type PipelineVars []struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Secured string `json:"secured"`
}

type RunPipelineDataForCommit struct {
	Target struct {
		Commit struct {
			Hash string `json:"hash"`
			Type string `json:"type"`
		} `json:"commit"`
		Selector struct {
			Type    string `json:"type"`
			Pattern string `json:"pattern"`
		} `json:"selector"`
		Type string `json:"type"`
	} `json:"target"`
	Variables PipelineVars `json:"variables"`
}
type RunPipelineDataForBranch struct {
	Target struct {
		RefType  string `json:"ref_type"`
		RefName  string `json:"ref_name"`
		Type     string `json:"type"`
		Selector struct {
			Type    string `json:"type"`
			Pattern string `json:"pattern"`
		} `json:"selector"`
	} `json:"target"`
	Variables PipelineVars `json:"variables"`
}

type RunPipelineDataForBranchCommit struct {
	Target struct {
		RefType string `json:"ref_type"`
		RefName string `json:"ref_name"`
		Type    string `json:"type"`
		Commit  struct {
			Hash string `json:"hash"`
			Type string `json:"type"`
		} `json:"commit"`
		Selector struct {
			Type    string `json:"type"`
			Pattern string `json:"pattern"`
		} `json:"selector"`
	} `json:"target"`
	Variables PipelineVars `json:"variables"`
}

var runPipelineCmd = &cobra.Command{
	Use:   "run",
	Short: "Run pipelines",
	Long:  `Run pipelines`,

	Run: func(cmd *cobra.Command, args []string) {
		repo, err := githelper.GetCurrentRepo(cmd)
		if err != nil {
			cliformat.Error(err.Error())
		}
		pipelineInput, err := cmd.Flags().GetString("pipeline")
		commit, err := cmd.Flags().GetString("commit")
		branch, err := cmd.Flags().GetString("branch")
		variables, err := cmd.Flags().GetString("variables")
		if commit == "" && branch == "" {
			fmt.Println(cliformat.Error("--branch and --commit are required"))
			return
		}
		pipelineInputArr := strings.Split(pipelineInput, ":")
		pipelineType := "custom"
		pipelinePattern := pipelineInputArr[0]
		if len(pipelineInputArr) > 1 {
			pipelineType = pipelineInputArr[0]
			pipelinePattern = pipelineInputArr[1]
		}
		var reqDataJson []byte
		var pipelineVars PipelineVars
		fmt.Println(variables)
		if variables != "" {
			err := json.Unmarshal([]byte(variables), &pipelineVars)
			if err != nil {
				fmt.Println(cliformat.Error("JSON variables decoding error. " + err.Error()))
				return
			}
		} else {
			pipelineVars = PipelineVars{}
		}
		if commit != "" {
			pipelineReqData := &RunPipelineDataForCommit{}
			pipelineReqData.Target.Commit.Hash = commit
			pipelineReqData.Target.Commit.Type = "commit"
			pipelineReqData.Target.Type = "pipeline_commit_target"
			pipelineReqData.Target.Selector.Type = pipelineType
			pipelineReqData.Target.Selector.Pattern = pipelinePattern
			pipelineReqData.Variables = pipelineVars
			reqDataJson, err = json.Marshal(pipelineReqData)

		}
		if branch != "" {
			pipelineReqData := &RunPipelineDataForBranch{}
			pipelineReqData.Target.RefName = branch
			pipelineReqData.Target.RefType = "branch"
			pipelineReqData.Target.Type = "pipeline_ref_target"
			pipelineReqData.Target.Selector.Type = pipelineType
			pipelineReqData.Target.Selector.Pattern = pipelinePattern
			pipelineReqData.Variables = pipelineVars
			reqDataJson, err = json.Marshal(pipelineReqData)
		}
		if branch != "" && commit != "" {
			pipelineReqData := &RunPipelineDataForBranchCommit{}
			pipelineReqData.Target.Commit.Hash = commit
			pipelineReqData.Target.Commit.Type = "commit"
			pipelineReqData.Target.RefName = branch
			pipelineReqData.Target.RefType = "branch"
			pipelineReqData.Target.Type = "pipeline_commit_target"
			pipelineReqData.Target.Selector.Type = pipelineType
			pipelineReqData.Target.Selector.Pattern = pipelinePattern
			pipelineReqData.Variables = pipelineVars
			reqDataJson, err = json.Marshal(pipelineReqData)

		}
		if err != nil {
			cliformat.Error(err.Error())
			return
		}

		pipelineDetailsRes, err := bitbucketapi.HttpRequestWithBitbucketAuthJson("POST", repo+"/pipelines", reqDataJson)
		defer pipelineDetailsRes.Body.Close()
		if err != nil {
			fmt.Println(cliformat.Error(err.Error()))
		}
		if pipelineDetailsRes.StatusCode != 201 {
			bodyText, _ := io.ReadAll(pipelineDetailsRes.Body)
			fmt.Println(cliformat.Error(string(bodyText)))
			return
		}

		var details formatters.PipelineDetailsResponse
		json.NewDecoder(pipelineDetailsRes.Body).Decode(&details)
		fmt.Println("Pipeline running : ", details.BuildNumber)
		fmt.Println("UUID: ", details.UUID)

	},
}

var stopPipelineCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a running pipelines",
	Long:  `Stops a running pipeline given the pipeline id`,

	Run: func(cmd *cobra.Command, args []string) {
		repo, err := githelper.GetCurrentRepo(cmd)
		if err != nil {
			cliformat.Error(err.Error())
		}

		if len(args) == 0 {
			fmt.Println(cliformat.Error("Pipeline id required\n Example : bitbucket pipeline stop {id}"))
			return
		}
		pipelineId := args[0]
		pipelineDetailsRes, err := bitbucketapi.HttpRequestWithBitbucketAuthJson("POST", repo+"/pipelines/"+pipelineId+"/stopPipeline", map[string]string{})
		defer pipelineDetailsRes.Body.Close()
		if err != nil {
			fmt.Println(cliformat.Error(err.Error()))
		}
		if pipelineDetailsRes.StatusCode != 204 {
			bodyText, _ := io.ReadAll(pipelineDetailsRes.Body)
			fmt.Println(cliformat.Error(string(bodyText)))
			return
		}

		fmt.Println("Stopped Pipeline with the id: ", pipelineId)

	},
}

func getStepStates(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {

	strategies := []string{constants.PIPELINE_FAILED, constants.PIPELINE_COMPLETED, constants.PIPELINE_IN_PROGRESS, constants.PIPELINE_SUCCESSFUL}
	return strategies, cobra.ShellCompDirectiveDefault

}
func init() {
	rootCmd.AddCommand(pipelinesCmd)
	pipelinesCmd.Flags().StringP("page", "p", "", "Page number for pipelines pagination")
	pipelinesCmd.Flags().StringP("state", "s", "", "Stop by step with state {FAILED|IN_PROGRESS|SUCCESSFUL|COMPLETED}")
	pipelinesCmd.Flags().StringP("format", "f", "", "Output template format")
	pipelinesCmd.Flags().BoolP("json", "j", false, "Output as json")
	pipelinesCmd.PersistentFlags().StringP("repo", "r", "", "Repo remote url")
	pipelinesCmd.Flags()
	pipelinesCmd.Flags().BoolP("detailed", "d", false, "Detailed pipeline steps with commands")
	runPipelineCmd.Flags().StringP("pipeline", "p", "", "Pipeline name. Example type:pipelinename, default type is custom")
	runPipelineCmd.Flags().StringP("commit", "c", "", "Commit run specfic pipeline commit")
	runPipelineCmd.Flags().StringP("branch", "b", "", "Branch run specfic pipeline commit")
	runPipelineCmd.MarkFlagRequired("pipeline")

	runPipelineCmd.RegisterFlagCompletionFunc("branch", githelper.GetBranchSuggestions)
	runPipelineCmd.RegisterFlagCompletionFunc("pipeline", bitbucketapi.GetPipelineNames)
	runPipelineCmd.RegisterFlagCompletionFunc("state", getStepStates)
	runPipelineCmd.Flags().StringP("variables", "v", "", `Pipeline Variables [{ "key": "var1key",  "value": "var1value", "secured": true}]`)

	getPipelineStep.Flags().StringP("pipelineId", "p", "", "Pipeline build number or id")
	getPipelineStep.Flags().StringP("step", "s", "", "Step UUID")
	getPipelineStep.MarkFlagRequired("pipelineId")
	getPipelineStep.MarkFlagRequired("step")
	pipelinesCmd.AddCommand(runPipelineCmd)
	pipelinesCmd.AddCommand(getPipelineStep)
	pipelinesCmd.AddCommand(stopPipelineCmd)

}
