package formatters

import (
	"github.com/spf13/cobra"
	"os"
	"text/template"
)

func CustomFormat(cmd *cobra.Command, res any,rangeTemplate string) (bool, error) {

	customFormat, err := cmd.Flags().GetString("format")
	if err != nil {
		return false, err
	}

	if customFormat == "" {
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
