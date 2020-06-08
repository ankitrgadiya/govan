package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
)

// addCommand represents the add subcommand for govan. It adds new package under
// the vanity domain. This command expects exactly two arguments: Path and
// Source.
func addCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "add PATH SOURCE",
		Short: "Adds new package under the vanity domain",
		Long: `The add sub-command adds new package under the vanity domain. It
automatically modified the configuration and generates new static HTML
file under the output directory.`,
		Example:      "govan add example github.com/username/example",
		Args:         cobra.ExactArgs(2),
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			name, source := args[0], args[1]
			pkgs := v.GetStringMapString("packages")
			pkgs[name] = source
			v.Set("packages", pkgs)
			err := v.WriteConfig()
			if err != nil {
				return err
			}

			filename := filepath.Join(output, name, "index.html")
			return writeHTML(v.GetString("domain"), name, source, filename)
		},
	}
}
