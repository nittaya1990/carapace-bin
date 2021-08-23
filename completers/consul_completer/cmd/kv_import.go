package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kv_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Imports a tree stored as JSON to the KV store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_importCmd).Standalone()
	addClientFlags(kv_importCmd)
	addServerFlags(kv_importCmd)
	kv_importCmd.Flags().String("namespace", "", "Specifies the namespace to query.")

	kvCmd.AddCommand(kv_importCmd)

	// TODO namespace completion

	carapace.Gen(kv_importCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if strings.HasPrefix(c.CallbackValue, "@") {
				c.CallbackValue = strings.TrimPrefix(c.CallbackValue, "@")
				return carapace.ActionFiles(".json").Invoke(c).Prefix("@").ToA()
			}
			return carapace.ActionValues()
		}),
	)
}
