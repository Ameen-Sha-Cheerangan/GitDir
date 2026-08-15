package cmd

import (
	"fmt"
	"os"

	"github.com/ameen/gitdir/internal/git"
	"github.com/ameen/gitdir/internal/github"
	"github.com/pterm/pterm"
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

		pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithTextStyle(pterm.NewStyle(pterm.FgBlack)).Println("GitDir")

		spinner, _ := pterm.DefaultSpinner.Start("Parsing GitHub URL...")
		repoInfo, err := github.ParseURL(rawURL)
		if err != nil {
			spinner.Fail(fmt.Sprintf("Invalid URL: %v", err))
			os.Exit(1)
		}
		
		var targetPath string
		if repoInfo.Path != "" {
			targetPath = repoInfo.Path
		} else {
			targetPath = "entire repository"
		}
		spinner.UpdateText(fmt.Sprintf("Downloading %s from %s/%s...", targetPath, repoInfo.Owner, repoInfo.Repo))

		err = git.Download(repoInfo, destDir)
		if err != nil {
			spinner.Fail(fmt.Sprintf("Failed to download: %v", err))
			os.Exit(1)
		}

		spinner.Success("Successfully downloaded to your current directory!")
		pterm.Success.Printf("Ready to use. Happy coding!\n")
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
