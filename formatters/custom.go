package formatters

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

func CustomFormat(cmd *cobra.Command, res any, rangeTemplate string) (bool, error) {

	isJsonFormat, err := cmd.Flags().GetBool("json")
	if isJsonFormat {
		jsonOutput, err := json.Marshal(res)
		fmt.Println(string(jsonOutput))
		return true, err
	}

	customFormat, err := cmd.Flags().GetString("format")
	if err != nil {
		return false, err
	}

	if customFormat == "" && !isJsonFormat {
		return false, err
	}

	if rangeTemplate != "" {
		customFormat = "{{ range ." + rangeTemplate + "}}" + customFormat + "{{end}}"
	}
	customOutputTemplate, err := template.New("customFormat").Parse(customFormat)
	if err != nil {
		return true, err
	}
	err = customOutputTemplate.Execute(os.Stdout, res)
	if err != nil {
		return true, err
	}
	return true, err

}
