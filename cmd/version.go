package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	Build   string
	Version string
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info(fmt.Sprintf("OTH export version %s - buildtime %s\n", Version, Build))

	},
}
