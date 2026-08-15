package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Ameen-Sha-Cheerangan/GitDir/internal/github"
)

// Download handles the sparse checkout of a specific directory.
func Download(repo *github.RepoInfo, destDir string) error {
	// If destDir is not provided, use the base name of the repo path
	if destDir == "" {
		if repo.Path != "" {
			destDir = filepath.Base(repo.Path)
		} else {
			destDir = repo.Repo
		}
	}

	// Check if destination already exists BEFORE doing any work
	if _, err := os.Stat(destDir); !os.IsNotExist(err) {
		return fmt.Errorf("destination '%s' already exists.\n💡 Tip: Provide a different name using the -o flag (e.g., gitdir <url> -o new_folder)\n   or remove the existing '%s' directory first", destDir, destDir)
	}

	// Create a temporary directory for the sparse clone
	tmpDir, err := os.MkdirTemp("", "gitdir-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	// Step 1: Clone with sparse checkout
	args := []string{"clone", "--progress", "--depth", "1", "--filter=blob:none", "--sparse"}
	if repo.Branch != "" {
		args = append(args, "--branch", repo.Branch)
	}
	args = append(args, repo.URL(), tmpDir)

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git clone failed: %w", err)
	}

	// Step 2: Set sparse checkout path
	if repo.Path != "" {
		cmd = exec.Command("git", "-C", tmpDir, "sparse-checkout", "set", repo.Path)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("git sparse-checkout failed: %w", err)
		}
	}

	// Step 3: Move the downloaded directory to the destination
	srcPath := filepath.Join(tmpDir, repo.Path)
	
	// Check if the source path exists
	if _, err := os.Stat(srcPath); os.IsNotExist(err) {
		return fmt.Errorf("path '%s' does not exist in the repository", repo.Path)
	}

	// Move the directory (we use rename)
	err = os.Rename(srcPath, destDir)
	if err != nil {
		// os.Rename can fail across different filesystems (e.g., /tmp to /home).
		// For a robust CLI, we might need a deep copy here, but for now we'll 
		// fallback to standard cp or an error if rename fails.
		
		// Fallback to cp -r
		cpCmd := exec.Command("cp", "-r", srcPath, destDir)
		if out, cpErr := cpCmd.CombinedOutput(); cpErr != nil {
			return fmt.Errorf("failed to move directory (rename error: %v, cp error: %v)\n%s", err, cpErr, out)
		}
	}
	
	// Remove the .git directory if it was copied (it would be in the root of temp dir, 
	// but if we copied the root repo, it might be in destDir)
	os.RemoveAll(filepath.Join(destDir, ".git"))

	return nil
}
