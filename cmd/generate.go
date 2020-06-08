package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
)

// generateCommand represents the generate subcommand for govan. This generates
// the HTML pages for all the packages in configuration file.
func generateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "generate",
		Short: "Generates static HTML",
		Long: `The generate sub-command generates HTML pages for all the
packages in configuration. This can be used with manual modification to
configuration or in non-interactive environments`,
		Example:      "govan generate",
		Args:         cobra.NoArgs,
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			domain := v.GetString("domain")
			pkgs := v.GetStringMapString("packages")

			for name, source := range pkgs {
				filename := filepath.Join(output, name, "index.html")
				err := writeHTML(domain, name, source, filename)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}
}
