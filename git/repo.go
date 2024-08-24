package githelper

import (
	"fmt"
	"os"
	"regexp"
	"slices"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
)

func getRepoPathFromRemote(remote string) string {
	re1, _ := regexp.Compile(`^.*bitbucket\.org.`)
	re2, _ := regexp.Compile(`\.git$`)
	repopath := re2.ReplaceAllString(re1.ReplaceAllString(remote, ""), "")

	return repopath
}
func getBranchName(branch string) string {
	re1, _ := regexp.Compile(`^refs\/heads\/`)
	return re1.ReplaceAllString(branch, "")

}

func GetCurrentRepoBranches() ([]string, error) {
	currentdir, err := os.Getwd()

	branches := &[]string{}
	if err != nil {
		return *branches, err
	}

	repo, err := git.PlainOpen(currentdir + "/.git")

	if err != nil {

		return *branches, err
	}
	b, err := repo.Branches()
	if err != nil {
		return *branches, err
	}
	b.ForEach(func(r *plumbing.Reference) error {
		*branches = slices.Insert(*branches, len(*branches), getBranchName(r.Name().String()))
		return nil
	})

	return *branches, nil
}

func GetBranchSuggestions(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	branches, err := GetCurrentRepoBranches()
	if err != nil {
		return branches, cobra.ShellCompDirectiveError
	}

	return branches, cobra.ShellCompDirectiveDefault

}

func GetCurrentRepoFromFile() (string, error) {
	currentdir, err := os.Getwd()

	if err != nil {

		return "", err
	}

	repo, err := git.PlainOpen(currentdir + "/.git")

	if err != nil {
		fmt.Println(currentdir + "/.git")

		return "", err
	}
	remote, err := repo.Remote("origin")
	if err != nil {
		return "", err
	}
	urls := remote.Config().URLs
	if len(urls) > 0 {
		return getRepoPathFromRemote(urls[0]), nil

	}
	return "", err

}

func GetCurrentRepo(cmd *cobra.Command) (string, error) {
	repo, err := cmd.Flags().GetString("repo")
	if err != nil || repo != "" {
		return repo, err
	}
	repo, err = GetCurrentRepoFromFile()
	if repo == "" && err == nil {
		return "", fmt.Errorf("No remote repo found")
	}
	return repo,err
}
