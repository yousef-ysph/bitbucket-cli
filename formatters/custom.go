package formatters

import (
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

func CustomFormat(cmd *cobra.Command, res any) (bool, error) {

	customFormat, err := cmd.Flags().GetString("format")
	if err != nil {
		return false, err
	}

	if customFormat == "" {
		return false, err
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
