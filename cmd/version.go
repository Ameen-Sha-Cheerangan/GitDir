package cmd

import (
	"fmt"

	"github.com/Ameen-Sha-Cheerangan/GitDir/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of gitdir",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gitdir %s\n", version.String())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
