package formatters

import (
	"bitbucket/cliformat"
	"bitbucket/constants"
	"fmt"
	"time"
)

type Target struct {
	RefType  string `json:"ref_type"`
	RefName  string `json:"ref_name"`
	Selector struct {
		RefType string `json:"ref_type"`
		RefName string `json:"ref_name"`
		Type    string `json:"type"`
		Pattern string `json:"pattern"`
	} `json:"selector"`
	Commit struct {
		Hash  string `json:"hash"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
		} `json:"links"`
		Type string `json:"type"`
	} `json:"commit"`
	Type string `json:"type"`
}

type State struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Result struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"result"`
	Stage struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"stage"`
}

type Creator struct {
	DisplayName string `json:"display_name"`
	Links       struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Avatar struct {
			Href string `json:"href"`
		} `json:"avatar"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
	} `json:"links"`
	Type      string `json:"type"`
	UUID      string `json:"uuid"`
	AccountID string `json:"account_id"`
	Nickname  string `json:"nickname"`
}

type PipelineDetailsResponse struct {
	UUID        string    `json:"uuid"`
	State       State     `json:"state"`
	BuildNumber int       `json:"build_number"`
	Creator     Creator   `json:"creator"`
	CreatedOn   time.Time `json:"created_on"`
	CompletedOn time.Time `json:"completed_on"`
	Target      Target
	Trigger     struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"trigger"`
	RunNumber         int  `json:"run_number"`
	DurationInSeconds int  `json:"duration_in_seconds"`
	BuildSecondsUsed  int  `json:"build_seconds_used"`
	FirstSuccessful   bool `json:"first_successful"`
	Expired           bool `json:"expired"`
	Links             struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Steps struct {
			Href string `json:"href"`
		} `json:"steps"`
	} `json:"links"`
	HasVariables bool `json:"has_variables"`
	Labels       struct {
	} `json:"labels"`
	ConfigurationSources []struct {
		Source string `json:"source"`
		URI    string `json:"uri"`
	} `json:"configuration_sources"`
	Type string `json:"type"`
}

type Command struct {
	CommandType string `json:"commandType"`
	Name        string `json:"name"`
	Command     string `json:"command"`
}
type TeardownCommand struct {
	CommandType string `json:"commandType"`
	Action      string `json:"action"`
	Name        string `json:"name"`
	Command     string `json:"command"`
}

type Step struct {
	Pipeline struct {
		Type string `json:"type"`
		UUID string `json:"uuid"`
	} `json:"pipeline"`
	SetupCommands    []Command         `json:"setup_commands"`
	ScriptCommands   []Command         `json:"script_commands"`
	TeardownCommands []TeardownCommand `json:"teardown_commands"`
	UUID             string            `json:"uuid"`
	Name             string            `json:"name"`
	Trigger          struct {
		Type string `json:"type"`
	} `json:"trigger"`
	State State `json:"state"`
	Image struct {
		Name string `json:"name"`
	} `json:"image"`
	MaxTime           int       `json:"maxTime"`
	StartedOn         time.Time `json:"started_on"`
	CompletedOn       time.Time `json:"completed_on"`
	DurationInSeconds int       `json:"duration_in_seconds"`
	BuildSecondsUsed  int       `json:"build_seconds_used"`
	RunNumber         int       `json:"run_number"`
	Type              string    `json:"type"`
}

type PipelineStepsResponse struct {
	Page    int    `json:"page"`
	Values  []Step `json:"values"`
	Size    int    `json:"size"`
	Pagelen int    `json:"pagelen"`
}

type Pipeline struct {
	UUID string `json:"uuid"`

	State       State `json:"state"`
	BuildNumber int   `json:"build_number"`
	Creator     Creator
	CreatedOn   time.Time `json:"created_on"`
	CompletedOn time.Time `json:"completed_on"`
	Target      Target    `json:"target"`
	Trigger     struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"trigger"`
	RunNumber         int  `json:"run_number"`
	DurationInSeconds int  `json:"duration_in_seconds"`
	BuildSecondsUsed  int  `json:"build_seconds_used"`
	FirstSuccessful   bool `json:"first_successful"`
	Expired           bool `json:"expired"`
	Links             struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Steps struct {
			Href string `json:"href"`
		} `json:"steps"`
	} `json:"links"`
	HasVariables bool `json:"has_variables"`
	Labels       struct {
	} `json:"labels"`
	Type string `json:"type"`
}

type PipelinesResponse struct {
	Values []Pipeline
}

func formatPipeline(pipeline Pipeline) string {

	state := pipeline.State.Name
	if state == constants.PIPELINE_COMPLETED {
		state = pipeline.State.Result.Name
		if state == constants.PIPELINE_SUCCESSFUL {
			state = cliformat.Success(state)
		} else if state == constants.PIPELINE_FAILED {
			state = cliformat.Error(state)
		} else {
			state = cliformat.Info(state)
		}
	}
	if state == constants.PIPELINE_IN_PROGRESS {
		state = cliformat.Info(pipeline.State.Stage.Name)
	}
	commit := pipeline.Target.Commit

	location := time.Now().Location()
	createdOn := pipeline.CreatedOn.In(location).Format("2006-01-02 15:04")

	duration := time.Duration(pipeline.DurationInSeconds * 1000_000_000).String()
	createdBy := pipeline.Creator.DisplayName
	pipelineType := pipeline.Target.Selector.Type + ":" + pipeline.Target.Selector.Pattern
	buildNumber := fmt.Sprintf("%v", pipeline.BuildNumber)
	return fmt.Sprintf("%s %s %s %s %s (%s)\n%s %s\n",
		cliformat.RightPad(state, 37, " "),
		cliformat.RightPad(buildNumber, 13, " "),
		cliformat.RightPad(commit.Hash, 44, " "),
		cliformat.RightPad(createdBy, 24, " "),
		createdOn,
		duration,
		cliformat.RightPad(pipelineType, 9, " "),
		pipeline.Target.RefName,
	)
}

func FormatPipelines(pipelineResponse PipelinesResponse) {

	for i := 0; i < len(pipelineResponse.Values); i++ {
		fmt.Println(formatPipeline(pipelineResponse.Values[i]))
	}
}
func formatPipelineStepCommand(commands []Command) string {
	text := ""
	for i := 0; i < len(commands); i++ {

		text = text + fmt.Sprintf("    - %s", commands[i].Command)
	}
	if len(text) > 0 {

		return "\n---------\n\n" + text
	}
	return text
}
func formatPipelineStep(pipelineStep Step, showScripts bool) string {
	state := pipelineStep.State.Name
	if state == constants.PIPELINE_COMPLETED {
		state = pipelineStep.State.Result.Name
		if state == constants.PIPELINE_SUCCESSFUL {
			state = cliformat.Success(state)
		} else if state == constants.PIPELINE_FAILED {
			state = cliformat.Error(state)
		} else {
			state = cliformat.Info(state)
		}
	} else {
		if pipelineStep.State.Stage.Name != "" {
			state = cliformat.Info(pipelineStep.State.Stage.Name)
		} else {
			state = cliformat.Info(state)

		}
	}

	duration := time.Duration(pipelineStep.DurationInSeconds * 1000_000_000).String()
	commandsText := ""
	if showScripts {
		commandsText = formatPipelineStepCommand(pipelineStep.ScriptCommands)
	}
	return fmt.Sprintf("\n- %s %s %s\n  Duration: %s%s", state, pipelineStep.Name, pipelineStep.UUID, duration, commandsText)
}
func FormatPipelineDetails(pipelineDetails PipelineDetailsResponse) string {
	state := pipelineDetails.State.Name
	if state == constants.PIPELINE_COMPLETED {
		state = pipelineDetails.State.Result.Name
		if state == constants.PIPELINE_SUCCESSFUL {
			state = cliformat.Success(state)
		} else if state == constants.PIPELINE_FAILED {
			state = cliformat.Error(state)
		} else {
			state = cliformat.Info(state)
		}
	} else {
		state = cliformat.Info(pipelineDetails.State.Stage.Name)
	}
	commit := pipelineDetails.Target.Commit.Hash

	location := time.Now().Location()
	createdOn := pipelineDetails.CreatedOn.In(location).Format("2006-01-02 15:04")

	duration := time.Duration(pipelineDetails.DurationInSeconds * 1000_000_000).String()
	pipelineType := pipelineDetails.Target.Selector.Type + ":" + pipelineDetails.Target.Selector.Pattern
	pipelineRef := pipelineDetails.Target.RefType + " : " + pipelineDetails.Target.RefName
	buildNumber := fmt.Sprintf("%v", pipelineDetails.BuildNumber)

	creator := pipelineDetails.Creator.DisplayName
	return fmt.Sprintf("%s %s \n%s\n%s\nBuild number:%s\nCreated by : %s\nCreated on : %s (%s)", state, pipelineType, pipelineRef, commit, buildNumber, creator, createdOn, duration)
}

func FormatPipelineDetailsWithSteps(pipelineDetailsRes PipelineDetailsResponse, pipelineStepsRes PipelineStepsResponse, showScripts bool) {
	fmt.Println(FormatPipelineDetails(pipelineDetailsRes))
	for i := 0; i < len(pipelineStepsRes.Values); i++ {
		fmt.Println(formatPipelineStep(pipelineStepsRes.Values[i], showScripts))
	}

}
