package cmd

import (
	"fmt"
	"os"

	"github.com/ameen/gitdir/internal/git"
	"github.com/ameen/gitdir/internal/github"
	"github.com/spf13/cobra"
)

var destDir string

var rootCmd = &cobra.Command{
	Use:   "gitdir [github-url]",
	Short: "Download specific directories from GitHub repositories",
	Long: `gitdir is a CLI tool that lets you download individual directories
from GitHub repositories without cloning the entire repo.

Under the hood it uses Git's sparse-checkout to fetch only the
files you need — fast, efficient, and bandwidth-friendly.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rawURL := args[0]

		repoInfo, err := github.ParseURL(rawURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fatal: invalid URL: %v\n", err)
			os.Exit(1)
		}
		
		var targetPath string
		if repoInfo.Path != "" {
			targetPath = repoInfo.Path
		} else {
			targetPath = "entire repository"
		}
		
		fmt.Printf("Downloading '%s' from %s/%s...\n", targetPath, repoInfo.Owner, repoInfo.Repo)

		err = git.Download(repoInfo, destDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fatal: failed to download: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Done.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&destDir, "output", "o", "", "output directory (default: current directory)")
}
