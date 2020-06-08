package cmd

import (
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configFileName = "govan"
	configFileType = "yaml"
	configFile     = "govan.yaml"
)

var (
	// v is the local instance of viper.
	v = viper.New()

	// output is the global flag determining output directory for govan.
	output string
)

// NewCommand generates a new instance of govan command.
func NewCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "govan",
		Short: "Go vanity import static generator",
		Long: `Govan generates static HTML files which serves as vanity imports
for Go packages. The static HTML files can be self-hosted with any
web-server or can be hosted on services like Github Pages, Netlify,
Surge.sh, etc.`,
		PersistentPreRunE: func(*cobra.Command, []string) (err error) {
			output, err = filepath.Abs(output)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cobra.OnInitialize(initConfig)

	root.PersistentFlags().StringVar(&output, "output", "./site", "Output directory for static HTML")

	root.AddCommand(initCommand())
	root.AddCommand(addCommand())
	root.AddCommand(removeCommand())
	root.AddCommand(generateCommand())
	return root
}

// Execute runs the govan command. This can be called in main.main() if govan is
// the root command.
func Execute() {
	if err := NewCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}

// initConfig loads and initialises configuration for govan.
func initConfig() {
	v.SetConfigType(configFileType)
	v.SetConfigName(configFileName)
	v.AddConfigPath(".")

	// Read the configuration file from current directory. Ignore the error if
	// the configuration file is not found.
	err := v.ReadInConfig()
	if err != nil && !isErrNotFound(err) {
		log.Fatal(err)
	}
}
