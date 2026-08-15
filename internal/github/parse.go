package github

import (
	"errors"
	"net/url"
	"strings"
)

// RepoInfo holds the parsed components of a GitHub URL.
type RepoInfo struct {
	Owner  string
	Repo   string
	Branch string
	Path   string
}

// URL returns the base repository URL.
func (r *RepoInfo) URL() string {
	return "https://github.com/" + r.Owner + "/" + r.Repo
}

// ParseURL extracts repository information from a GitHub URL.
func ParseURL(rawURL string) (*RepoInfo, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	if u.Host != "github.com" {
		return nil, errors.New("only github.com URLs are supported")
	}

	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(parts) < 2 {
		return nil, errors.New("invalid GitHub repository URL")
	}

	info := &RepoInfo{
		Owner: parts[0],
		Repo:  parts[1],
	}

	// Example: /owner/repo/tree/main/path/to/dir
	if len(parts) >= 4 && parts[2] == "tree" {
		info.Branch = parts[3] // We assume the branch name doesn't contain slashes for simplicity
		if len(parts) > 4 {
			info.Path = strings.Join(parts[4:], "/")
		}
	} else if len(parts) > 2 {
		return nil, errors.New("URL does not point to a specific directory (expected /tree/<branch>/<path>)")
	} else {
		// Just the repo URL
		info.Path = ""
	}

	return info, nil
}
