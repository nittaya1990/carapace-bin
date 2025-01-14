package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var secret_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List secrets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	secret_listCmd.Flags().StringP("env", "e", "", "List secrets for an environment")
	secret_listCmd.Flags().StringP("org", "o", "", "List secrets for an organization")
	secretCmd.AddCommand(secret_listCmd)

	carapace.Gen(secret_listCmd).FlagCompletion(carapace.ActionMap{
		"env": action.ActionEnvironments(secret_listCmd),
		"org": action.ActionUsers(secret_listCmd, action.UserOpts{Organizations: true}),
	})
}
