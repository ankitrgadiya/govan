package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// removeCommand represents the remove subcommand for govan. It removes the
// package from configuration and deletes HTML page.
func removeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "remove PATH",
		Short: "Removes package under the vanity domain",
		Long: `The remove sub-command removes existing package under the vanity domain. It
automatically removes the configuration and deletes static HTML file under the
output directory.`,
		Example:      "govan remove example",
		Args:         cobra.ExactArgs(1),
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			name := args[0]
			pkgs := v.GetStringMapString("packages")
			delete(pkgs, name)
			v.Set("packages", pkgs)
			err := v.WriteConfig()
			if err != nil {
				return err
			}

			return os.RemoveAll(filepath.Join(output, name))
		},
	}
}
