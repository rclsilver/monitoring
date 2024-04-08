package cmd

import (
	"os"

	"github.com/ovh/configstore"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	verbose    bool
	configFile string
)

var rootCmd = &cobra.Command{
	Use:   "monitoring-daemon",
	Short: "monitoring-daemon server",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose {
			logrus.SetLevel(logrus.DebugLevel)
		}

		configstore.InitFromEnvironment()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable to verbose mode")
}
