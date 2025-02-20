package formatters

import (
	"fmt"
	"sort"
)

type Environment struct {
	Type            string `json:"type"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Rank            int    `json:"rank"`
	EnvironmentType struct {
		Name string `json:"name"`
		Rank int    `json:"rank"`
		Type string `json:"type"`
	} `json:"environment_type"`
	DeploymentGateEnabled bool `json:"deployment_gate_enabled"`
	Lock                  struct {
		LockOpener struct {
			Type                string `json:"type"`
			PipelineUUID        string `json:"pipeline_uuid"`
			DeploymentGroupUUID string `json:"deployment_group_uuid"`
			StepUUID            string `json:"step_uuid"`
		} `json:"lock_opener"`
		Triggerer struct {
			Type         string `json:"type"`
			PipelineUUID string `json:"pipeline_uuid"`
			StepUUID     string `json:"step_uuid"`
		} `json:"triggerer"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"lock"`
	Restrictions struct {
		Type      string `json:"type"`
		AdminOnly bool   `json:"admin_only"`
	} `json:"restrictions"`
	Hidden                 bool   `json:"hidden"`
	EnvironmentLockEnabled bool   `json:"environment_lock_enabled"`
	UUID                   string `json:"uuid"`
	Category               struct {
		Name string `json:"name"`
	} `json:"category"`
	Pipeline PipelineDetailsResponse
}

type EnvResponse struct {
	Values []Environment `json:"values"`
}

func FormatEnv(env Environment) string {
	return fmt.Sprintf("%s [%s] %s\n%s\n", env.Name, env.EnvironmentType.Name, env.Lock.Triggerer.PipelineUUID, FormatPipelineDetails(env.Pipeline))
}
func FormatEnvs(envRes EnvResponse, grouped bool) string {
	output := make(map[string]string)
	sort.Slice(envRes.Values, func(i, j int) bool {
		return envRes.Values[i].Name < envRes.Values[j].Name
	})
	for i := 0; i < len(envRes.Values); i++ {
		env := envRes.Values[i]

		if envString, ok := output[env.EnvironmentType.Name]; ok {

			output[env.EnvironmentType.Name] = envString + "\n" + FormatEnv(env)
		} else {
			output[env.EnvironmentType.Name] = FormatEnv(env)
		}
	}
	outputString := ""
	for envType, envDetails := range output {
		outputString = outputString + envType + "\n===========\n" + envDetails + "\n"
	}
	return outputString

}
