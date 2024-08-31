/*
Copyright © 2024 yousef@ysph.tech
*/
package cmd

import (
	"bitbucket/bitbucketapi"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"

	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{

	Use:       "setup [TYPE]",
	Short:     "Sets up the authentication file with either password or token",
	Long:      "",
	ValidArgs: []string{"token", "password"},
	Example:   "bitbucket setup password\nbitbucket setup token",
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)
		homeDirectory, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		filename := homeDirectory + "/.bitbucketcmd.json"
		_, err = os.Stat(filename)
		var file *os.File

		file, err = os.Create(filename)
		defer file.Close()
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		configuration := bitbucketapi.Config{}
		if args[0] == "password" {
			fmt.Print("Enter Username: ")
			username, err := reader.ReadString('\n')

			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			fmt.Print("Enter Password: ")
			bytePassword, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			configuration.Password = string(bytePassword)
			configuration.User = strings.ReplaceAll(username, "\n", "")
			configurationJson, err := json.Marshal(configuration)
			_, err = file.Write(configurationJson)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}

		}

		if args[0] == "token" {

			fmt.Print("Enter Token: ")
			byteToken, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			configuration.Token = string(byteToken)

			configurationJson, err := json.Marshal(configuration)
			_, err = file.Write(configurationJson)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
		}
		fmt.Println("\nFile successsfully setup")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
