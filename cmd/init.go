package cmd

import (
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// initCommand represents the init subcommand for govan. It initialises the
// configuration for new domain. This command expects exactly one argument:
// Domain.
func initCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "init DOMAIN",
		Short: "Initialise new vanity domain",
		Long: `The init sub-command initializes new vanity domain. It will
generate a new configuration file with boilerplate.`,
		Example:      "govan init example.dev",
		Args:         cobra.ExactArgs(1),
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			// Expects exactly one argument: Domain
			domain := strings.TrimSuffix(args[0], "/")

			_, err := os.Stat(configFile)
			if err != nil && os.IsNotExist(err) {
				// TODO: Find a better way to create/touch the file.
				// Alternatively find a way to directly create a file through
				// Viper.
				// Note: The SafeWriteConfig is supposed to do that but it
				// throws NotFound error instead.
				f, err := os.Create(configFile)
				if err != nil {
					return err
				}
				f.Close()

				v.Set("domain", domain)
				v.Set("packages", map[string]string{})
				err = v.WriteConfig()
				if err != nil {
					return err
				}
			} else if err != nil {
				return err
			}

			filename := filepath.Join(output, "index.html")
			writeHTML("pkg.go.dev/search?q="+url.QueryEscape(domain), "", "", filename)

			return nil
		},
	}
}
